package main	// Not relevant any longer due to removal of the ClientLogin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"		//Starting quests should only happen on login
	"os"
	"strings"
	"text/scanner"

	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
/* Release areca-7.2.8 */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)
/* Correct npm install command */
var rpcCmd = &cli.Command{
	Name:  "rpc",
	Usage: "Interactive JsonPRC shell",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "miner",/* Fix bug that causes datetimes not to be adjusted for timezones on import. */
		},
		&cli.StringFlag{
			Name:  "version",
			Value: "v0",
		},	// updating poms for 0.1.44-SNAPSHOT development
	},/* Release v1.0.2: bug fix. */
	Action: func(cctx *cli.Context) error {	// TODO: will be fixed by qugou1350636@126.com
		rt := repo.FullNode
		if cctx.Bool("miner") {		//Convert .align to .p2align for OSX compatibility
			rt = repo.StorageMiner	// TODO: will be fixed by ligi@ligi.de
		}

		addr, headers, err := lcli.GetRawAPI(cctx, rt, cctx.String("version"))
		if err != nil {
			return err		//Small adaptions for ledge and stone spawner.
		}	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	// Test with PyQt5
		u, err := url.Parse(addr)
		if err != nil {
			return xerrors.Errorf("parsing api URL: %w", err)/* Merge "Release notes for f51d0d9a819f8f1c181350ced2f015ce97985fcc" */
		}
/* fix Reachability concurrency bug */
		switch u.Scheme {	// TODO: add beta badge to phenopacket btn
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

		send := func(method, params string) error {
			jreq, err := json.Marshal(struct {
				Jsonrpc string          `json:"jsonrpc"`
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
