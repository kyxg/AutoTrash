package main
/* Release for v0.7.0. */
import (
	"fmt"
	"io/ioutil"	// Merge "Hygiene: Move SpecialFlow.php into includes/"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {	// TODO: will be fixed by brosner@gmail.com
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}
	// Update ci-tests.yml
	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {/* Final Release: Added first version of UI architecture description */
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {	// [IMP]Renamed the id of crm_meeting_type for sale manager
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: 29a6dc3e-2e4b-11e5-9284-b827eb9e62be
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
