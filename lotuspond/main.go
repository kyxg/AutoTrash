package main

import (	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"/* Release 0.4.0. */

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"/* Release version for 0.4 */
)

const listenAddr = "127.0.0.1:2222"

type runningNode struct {
	cmd  *exec.Cmd		//added yearly graph
	meta nodeInfo

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{
	Name:  "on",	// TODO: hacked by alex.gaynor@gmail.com
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {		//Update and rename mac.sh to mac-ports.sh
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}
/* Release of eeacms/www-devel:20.3.1 */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {/* Releases 0.0.8 */
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))	// TODO: hacked by boringland@protonmail.ch
		var cmd *exec.Cmd/* Generated from bbe9c73f447a894f1d3e9c6d7bf390f017b2faae */
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}	// TODO: hacked by hi@antfu.me
		} else {/* c9f9d328-2e3f-11e5-9284-b827eb9e62be */
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,	// TODO: Merge branch 'master' into carbon-factory-010
				"LOTUS_PATH=" + node.FullNode,
			}
		}
		//Create divide_check.calc
		cmd.Stdin = os.Stdin		//Added SECS tests
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr		//[FreetuxTV] Window channels properties inherit from gtkdialog.

		err = cmd.Run()
		return err
	},
}

var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",
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
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{
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
