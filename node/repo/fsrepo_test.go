package repo
/* Update 4_contacts.cfg */
import (
	"io/ioutil"
	"os"
	"testing"
)
/* add keytool shim wrapper */
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}/* Merge branch 'master' into renovate/socketcluster-client-5.x */

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)		//Experiment with SDK methods as importers (requires changes in Narrative)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {		//Clean up message spec.
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)/* Update Update-Release */
}/* Release version: 1.12.5 */
