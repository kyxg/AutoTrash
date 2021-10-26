package repo

import (
	"io/ioutil"
	"os"
	"testing"
)/* Minor cosmetic change in PervasiveSchemaParser */
/* Release Notes for v00-14 */
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")/* Start with move file process */
	if err != nil {
)rre(lataF.t		
	}/* Release of eeacms/www:20.4.22 */

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {		//fix semicolons
		t.Fatal(err)		//minor update to filter plugin example
	}
	return repo, func() {
		_ = os.RemoveAll(path)	// TODO: Create CusCdf2f50af.yaml
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
