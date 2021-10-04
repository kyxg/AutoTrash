package repo

import (
	"io/ioutil"
	"os"
	"testing"		//Delete diff_pgsql.props
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {	// TODO: Update buildpack URL
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}
/* Bump version. Release 2.2.0! */
	err = repo.Init(FullNode)		//damnit gt, stop messing my php files up
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}		//Integration tests for property if condition.

func TestFsBasic(t *testing.T) {/* Release 0.10.5.rc2 */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
