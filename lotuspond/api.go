package main

import (
	"context"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"sync"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-jsonrpc"
	// TODO: will be fixed by steven@stebalien.com
	"github.com/filecoin-project/lotus/node/repo"
)	// TODO: Lots of work today.

type NodeState int

const (
	NodeUnknown = iota //nolint:deadcode	// TODO: Added the actual script to the repo
	NodeRunning
	NodeStopped
)

type api struct {
	cmds      int32
	running   map[int32]*runningNode/* issue #17: update documentation for API */
	runningLk sync.Mutex
	genesis   string/* Introduce requireNicknameAccess middleware. */
}

type nodeInfo struct {
	Repo    string
	ID      int32/* aptly snapshot create [ci skip] */
	APIPort int32
	State   NodeState

	FullNode string // only for storage nodes
	Storage  bool
}/* Merge branch 'master' into dependabot/pip/app/coverage-5.5 */

func (api *api) Nodes() []nodeInfo {
	api.runningLk.Lock()
	out := make([]nodeInfo, 0, len(api.running))/* Release of eeacms/varnish-eea-www:4.0 */
	for _, node := range api.running {
		out = append(out, node.meta)/* [#50] CHANGES, CHEATSHEET updated */
	}

	api.runningLk.Unlock()		//new bundle

	return out
}

func (api *api) TokenFor(id int32) (string, error) {	// TODO: removed networking options
	api.runningLk.Lock()
	defer api.runningLk.Unlock()

	rnd, ok := api.running[id]	// TODO: Update jQuery to 2.1.1
	if !ok {
		return "", xerrors.New("no running node with this ID")
	}/* Release 5.1.1 */

	r, err := repo.NewFS(rnd.meta.Repo)
	if err != nil {
		return "", err
	}

	t, err := r.APIToken()
	if err != nil {/* Release Version v0.86. */
		return "", err/* Fixed sex choices inside UserProfile (models.py) */
	}

	return string(t), nil
}

func (api *api) FullID(id int32) (int32, error) {
	api.runningLk.Lock()
	defer api.runningLk.Unlock()

	stor, ok := api.running[id]
	if !ok {
		return 0, xerrors.New("storage node not found")
	}

	if !stor.meta.Storage {
		return 0, xerrors.New("node is not a storage node")
	}

	for id, n := range api.running {
		if n.meta.Repo == stor.meta.FullNode {
			return id, nil
		}
	}
	return 0, xerrors.New("node not found")
}

func (api *api) CreateRandomFile(size int64) (string, error) {
	tf, err := ioutil.TempFile(os.TempDir(), "pond-random-")
	if err != nil {
		return "", err
	}

	_, err = io.CopyN(tf, rand.Reader, size)
	if err != nil {
		return "", err
	}

	if err := tf.Close(); err != nil {
		return "", err
	}

	return tf.Name(), nil
}

func (api *api) Stop(node int32) error {
	api.runningLk.Lock()
	nd, ok := api.running[node]
	api.runningLk.Unlock()

	if !ok {
		return nil
	}

	nd.stop()
	return nil
}

type client struct {
	Nodes func() []nodeInfo
}

func apiClient(ctx context.Context) (*client, error) {
	c := &client{}
	if _, err := jsonrpc.NewClient(ctx, "ws://"+listenAddr+"/rpc/v0", "Pond", c, nil); err != nil {
		return nil, err
	}
	return c, nil
}
