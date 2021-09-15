package main	// ensure dependencies include only voices to install

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"/* Release a force target when you change spells (right click). */
)

const listenAddr = "127.0.0.1:2222"

type runningNode struct {	// TODO: Allow arbitrary number of threads
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux	// TODO: will be fixed by vyzo@hackzen.org
	stop func()
}
/* job #9354 small change to clarify code by enforcing visibility */
var onCmd = &cli.Command{
	Name:  "on",	// TODO: hacked by vyzo@hackzen.org
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err	// TODO: PUzCk5aamQ1VQLXsOsA5xTZYOXBHrSuq
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)		//27955074-2e73-11e5-9284-b827eb9e62be
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)/* amdlib.util: merge in changes to humanreadable.py that were made in pyutil */
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}/* Released 0.9.2 */
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout/* [dist] Release v1.0.0 */
		cmd.Stderr = os.Stderr

		err = cmd.Run()		//[Jimw_Domain] Change Domain gestion to a global gestion
		return err
	},
}
/* moving nexusReleaseRepoId to a property */
var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",		//3866eb42-2e64-11e5-9284-b827eb9e62be
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err/* Delete PLAY.HTML */
		}		//758c0514-2e66-11e5-9284-b827eb9e62be

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
