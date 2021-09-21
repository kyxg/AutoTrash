package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}/* Homiwpf: update Release with new compilation and dll */

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)		//eb874ee3-2ead-11e5-8a09-7831c1d44c14
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)		//Remove meta validation, not needed at this step anyway
	defer closer()
	basicTest(t, repo)
}		//Drop @openapitools/openapi-generator-cli
