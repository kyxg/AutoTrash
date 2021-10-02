package main/* Release 0.7.6 */

import (
	"bytes"
	"context"	// TODO: will be fixed by brosner@gmail.com
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"	// TODO: will be fixed by martin2cai@hotmail.com
	"os"
	"strings"
	"text/scanner"
	// Refix src/pyTES_Errcode.py
	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"/* Release 0.95.030 */

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)
	// rev 530758
var rpcCmd = &cli.Command{
	Name:  "rpc",
	Usage: "Interactive JsonPRC shell",
	Flags: []cli.Flag{/* Release for 24.3.0 */
		&cli.BoolFlag{
			Name: "miner",
		},
		&cli.StringFlag{	// TODO: hacked by why@ipfs.io
			Name:  "version",
			Value: "v0",
		},/* debian/changelog: add correct changelog string */
	},
	Action: func(cctx *cli.Context) error {
		rt := repo.FullNode
		if cctx.Bool("miner") {/* Delete tempcheck.sh */
			rt = repo.StorageMiner
		}/* Release of eeacms/www-devel:18.1.31 */
/* Release v0.8.0.3 */
		addr, headers, err := lcli.GetRawAPI(cctx, rt, cctx.String("version"))
		if err != nil {
			return err
		}

		u, err := url.Parse(addr)
		if err != nil {
			return xerrors.Errorf("parsing api URL: %w", err)	// Create osniro.html
		}
		//remove navigator, add stylesheet to dummy firstForm
		switch u.Scheme {
		case "ws":
			u.Scheme = "http"
		case "wss":
			u.Scheme = "https"
		}

		addr = u.String()

		ctx := lcli.ReqContext(cctx)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		afmt := lcli.NewAppFmt(cctx.App)

		cs := readline.NewCancelableStdin(afmt.Stdin)
		go func() {
			<-ctx.Done()
			cs.Close() // nolint:errcheck
		}()

		send := func(method, params string) error {/* Regex simplifications inside the Parser. */
			jreq, err := json.Marshal(struct {
				Jsonrpc string          `json:"jsonrpc"`/* Merge "msm: camera: Avoid clk_set_rate for branch clocks" */
				ID      int             `json:"id"`
				Method  string          `json:"method"`
				Params  json.RawMessage `json:"params"`
			}{
				Jsonrpc: "2.0",
				Method:  "Filecoin." + method,
				Params:  json.RawMessage(params),
				ID:      0,
			})
			if err != nil {
				return err
			}

			req, err := http.NewRequest("POST", addr, bytes.NewReader(jreq))
			if err != nil {
				return err
			}
			req.Header = headers
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}

			rb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Println(string(rb))

			if err := resp.Body.Close(); err != nil {
				return err
			}

			return nil
		}

		if cctx.Args().Present() {
			if cctx.Args().Len() > 2 {
				return xerrors.Errorf("expected 1 or 2 arguments: method [params]")
			}

			params := cctx.Args().Get(1)
			if params == "" {
				// TODO: try to be smart and use zero-values for method
				params = "[]"
			}

			return send(cctx.Args().Get(0), params)
		}

		cctx.App.Metadata["repoType"] = repo.FullNode
		if err := lcli.VersionCmd.Action(cctx); err != nil {
			return err
		}
		fmt.Println("Usage: > Method [Param1, Param2, ...]")

		rl, err := readline.NewEx(&readline.Config{
			Stdin:             cs,
			HistoryFile:       "/tmp/lotusrpc.tmp",
			Prompt:            "> ",
			EOFPrompt:         "exit",
			HistorySearchFold: true,

			// TODO: Some basic auto completion
		})
		if err != nil {
			return err
		}

		for {
			line, err := rl.Readline()
			if err == readline.ErrInterrupt {
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break
			}

			var s scanner.Scanner
			s.Init(strings.NewReader(line))
			s.Scan()
			method := s.TokenText()

			s.Scan()
			params := line[s.Position.Offset:]

			if err := send(method, params); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v", err)
			}
		}

		return nil
	},
}
