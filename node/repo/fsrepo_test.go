package repo
	// Allow to get the filename without extension
import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {	// TODO: Display "Me" for the own contact entry in the list
		t.Fatal(err)
	}
/* Release of eeacms/plonesaas:5.2.1-15 */
	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}
/* Create v3_Android_ReleaseNotes.md */
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)		//Added test for CargoUpdater
}/* Release of eeacms/redmine:4.1-1.3 */
