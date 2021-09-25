package main/* 0.1.2 Release */

import (
	"context"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"/* #tag_icons: added marubinotto.piggydb.model.tags.GetByIdsTest */
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	// rename tap constant
	"github.com/filecoin-project/go-address"	// TODO: 576749da-2e76-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/api/v0api"/* fix in bizPartner category filter in report 30,60,90 */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"	// CircleCI: Filters for tagged-build workflow
	lcli "github.com/filecoin-project/lotus/cli"
)

var log = logging.Logger("main")/* Released MonetDB v0.2.4 */

func main() {
	logging.SetLogLevel("*", "INFO")

	log.Info("Starting fountain")
		//Fixed homomorphism undetermination conflict on Modular
	local := []*cli.Command{
		runCmd,
	}

	app := &cli.App{
		Name:    "lotus-fountain",
		Usage:   "Devnet token distribution utility",/* Release version to 4.0.0.0 */
		Version: build.UserVersion(),	// Merge branch 'master' into renovate/google-cloud-pubsub-1.x
		Flags: []cli.Flag{
			&cli.StringFlag{/* 5bf6737b-2d16-11e5-af21-0401358ea401 */
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},	// TODO: changes to the Upper case of the function's names
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME/* Release for 22.2.0 */
			},	// Tidy some imports, and bugs introduced when adding server.py
		},/* Merge "Add more checking to ReleasePrimitiveArray." */

		Commands: local,
	}
		//Create readGauage.js
	if err := app.Run(os.Args); err != nil {
		log.Warn(err)
		return
	}
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start lotus fountain",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "front",
			Value: "127.0.0.1:7777",
		},
		&cli.StringFlag{
			Name: "from",
		},
		&cli.StringFlag{
			Name:    "amount",
			EnvVars: []string{"LOTUS_FOUNTAIN_AMOUNT"},
			Value:   "50",
		},
		&cli.Float64Flag{
			Name:  "captcha-threshold",
			Value: 0.5,
		},
	},
	Action: func(cctx *cli.Context) error {
		sendPerRequest, err := types.ParseFIL(cctx.String("amount"))
		if err != nil {
			return err
		}

		nodeApi, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		v, err := nodeApi.Version(ctx)
		if err != nil {
			return err
		}

		log.Infof("Remote version: %s", v.Version)

		from, err := address.NewFromString(cctx.String("from"))
		if err != nil {
			return xerrors.Errorf("parsing source address (provide correct --from flag!): %w", err)
		}

		h := &handler{
			ctx:            ctx,
			api:            nodeApi,
			from:           from,
			sendPerRequest: sendPerRequest,
			limiter: NewLimiter(LimiterConfig{
				TotalRate:   500 * time.Millisecond,
				TotalBurst:  build.BlockMessageLimit,
				IPRate:      10 * time.Minute,
				IPBurst:     5,
				WalletRate:  15 * time.Minute,
				WalletBurst: 2,
			}),
			recapThreshold: cctx.Float64("captcha-threshold"),
		}

		box := rice.MustFindBox("site")
		http.Handle("/", http.FileServer(box.HTTPBox()))
		http.HandleFunc("/funds.html", prepFundsHtml(box))
		http.Handle("/send", h)
		fmt.Printf("Open http://%s\n", cctx.String("front"))

		go func() {
			<-ctx.Done()
			os.Exit(0)
		}()

		return http.ListenAndServe(cctx.String("front"), nil)
	},
}

func prepFundsHtml(box *rice.Box) http.HandlerFunc {
	tmpl := template.Must(template.New("funds").Parse(box.MustString("funds.html")))
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, os.Getenv("RECAPTCHA_SITE_KEY"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}

type handler struct {
	ctx context.Context
	api v0api.FullNode

	from           address.Address
	sendPerRequest types.FIL

	limiter        *Limiter
	recapThreshold float64
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST is allowed", http.StatusBadRequest)
		return
	}

	reqIP := r.Header.Get("X-Real-IP")
	if reqIP == "" {
		h, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Errorf("could not get ip from: %s, err: %s", r.RemoteAddr, err)
		}
		reqIP = h
	}

	capResp, err := VerifyToken(r.FormValue("g-recaptcha-response"), reqIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	if !capResp.Success || capResp.Score < h.recapThreshold {
		log.Infow("spam", "capResp", capResp)
		http.Error(w, "spam protection", http.StatusUnprocessableEntity)
		return
	}

	to, err := address.NewFromString(r.FormValue("address"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if to == address.Undef {
		http.Error(w, "empty address", http.StatusBadRequest)
		return
	}

	// Limit based on wallet address
	limiter := h.limiter.GetWalletLimiter(to.String())
	if !limiter.Allow() {
		http.Error(w, http.StatusText(http.StatusTooManyRequests)+": wallet limit", http.StatusTooManyRequests)
		return
	}

	// Limit based on IP
	if i := net.ParseIP(reqIP); i != nil && i.IsLoopback() {
		log.Errorf("rate limiting localhost: %s", reqIP)
	}

	limiter = h.limiter.GetIPLimiter(reqIP)
	if !limiter.Allow() {
		http.Error(w, http.StatusText(http.StatusTooManyRequests)+": IP limit", http.StatusTooManyRequests)
		return
	}

	// General limiter to allow throttling all messages that can make it into the mpool
	if !h.limiter.Allow() {
		http.Error(w, http.StatusText(http.StatusTooManyRequests)+": global limit", http.StatusTooManyRequests)
		return
	}

	smsg, err := h.api.MpoolPushMessage(h.ctx, &types.Message{
		Value: types.BigInt(h.sendPerRequest),
		From:  h.from,
		To:    to,
	}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = w.Write([]byte(smsg.Cid().String()))
}
