package main
/* job #272 - Update Release Notes and What's New */
import (
	"fmt"
	"io/ioutil"
	"os"	// TODO: Create Wave Surfer Prototype
)	// Touch to reset stats
	// TODO: - fix: step 3, method to determine days got deleted somewhere. Is restored now.
func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"	// TODO: fixed commit
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {	// Delete BelichtungsMesser.ino
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {	// Fixed style in README
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {	// 5b67c19a-2e62-11e5-9284-b827eb9e62be
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
