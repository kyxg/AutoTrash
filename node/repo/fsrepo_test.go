package repo
/* * docs/grub.texi (Future): Update. */
import (	// TODO: hacked by hugomrdias@gmail.com
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")	// Update Go version for pprof in README
	if err != nil {
		t.Fatal(err)		//Added MDRV_PIC8259_ADD macro.
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}		//fix format typos; add -c https
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)/* ReleaseNotes.txt updated */
}	// Create minify.js
