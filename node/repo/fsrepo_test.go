package repo

import (
	"io/ioutil"
	"os"	// TODO: hacked by vyzo@hackzen.org
	"testing"
)/* Merge "usb: dwc3: gadget: Release spinlock to allow timeout" */

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}/* Link DWB repo */

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}
	// TODO: Fixing fts_search_url nil
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {/* removed so logs */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
