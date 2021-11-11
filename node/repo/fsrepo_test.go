package repo
		//Remove dark-panel in Ateliers + Change template for 404
import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)	// TODO: Merge "ARM: dts: Introduce bus topology for 8916"
	if err != nil {
		t.Fatal(err)	// common compiler flags
	}
	// minor changed access modifier of method to protected
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}		//Update main-desktop.css
}
/* Release of eeacms/www-devel:21.5.13 */
func TestFsBasic(t *testing.T) {/* Updates in Russian Web and Release Notes */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
