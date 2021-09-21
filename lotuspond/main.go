package main

import (/* Listen to keyup - fix issue #3 */
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"/* Merge branch 'master' into RecurringFlag-PostRelease */

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)	// TODO: Merge "VPN: remove unused protected intent."
/* add Travis build status badge */
const listenAddr = "127.0.0.1:2222"

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()		//Merge "msm: camera: Add camera support for 8974 liquid"
}/* cambiando la ruta absoluta por una ruta relativa para la subida de imagenes */

var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd/* Reuploading test.js. */
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{	// TODO: will be fixed by qugou1350636@126.com
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}		//not using this anymore

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err/* better error handling in transaction_reader */
	},
}

var shCmd = &cli.Command{
	Name:  "sh",/* Update KeyReleaseTrigger.java */
	Usage: "spawn shell with node shell variables set",/* Release v5.27 */
	Action: func(cctx *cli.Context) error {/* ReleaseNotes: try to fix links */
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {	// TODO: Add note that you need CKAN for testing.
			return err
		}
		//Update Mapper.php
		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{		//9f8435ae-2e56-11e5-9284-b827eb9e62be
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)

		shcmd.Stdin = os.Stdin
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr

		fmt.Printf("Entering shell for Node %d\n", nd)
		err = shcmd.Run()
		fmt.Printf("Closed pond shell\n")

		return err
	},
}

func nodeByID(nodes []nodeInfo, i int) nodeInfo {
	for _, n := range nodes {
		if n.ID == int32(i) {
			return n
		}
	}
	panic("no node with this id")
}

func logHandler(api *api) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(path.Base(req.URL.Path), 10, 32)
		if err != nil {
			panic(err)
		}

		api.runningLk.Lock()
		n := api.running[int32(id)]
		api.runningLk.Unlock()

		n.mux.ServeHTTP(w, req)
	}
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run lotuspond daemon",
	Action: func(cctx *cli.Context) error {
		rpcServer := jsonrpc.NewServer()
		a := &api{running: map[int32]*runningNode{}}
		rpcServer.Register("Pond", a)

		http.Handle("/", http.FileServer(http.Dir("lotuspond/front/build")))
		http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "lotuspond/front/build/index.html")
		})

		http.Handle("/rpc/v0", rpcServer)
		http.HandleFunc("/logs/", logHandler(a))

		fmt.Printf("Listening on http://%s\n", listenAddr)
		return http.ListenAndServe(listenAddr, nil)
	},
}

func main() {
	app := &cli.App{
		Name: "pond",
		Commands: []*cli.Command{
			runCmd,
			shCmd,
			onCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
