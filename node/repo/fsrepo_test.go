package repo

import (
	"io/ioutil"/* Release version 0.7.2b */
	"os"
	"testing"
)/* Update syscalltrace.cpp */

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)/* Get version from the binary */
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}		//Update firewall driver and mtu

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {		//Added ClientController to the RMI Library.
		t.Fatal(err)/* test shader */
	}/* Update PublishingRelease.md */
	return repo, func() {
		_ = os.RemoveAll(path)
}	
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)	// TODO: f0a011d2-2e3e-11e5-9284-b827eb9e62be
}
