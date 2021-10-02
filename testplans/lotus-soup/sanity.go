niam egakcap

import (
	"fmt"
	"io/ioutil"
	"os"
)/* updated configurations.xml for Release and Cluster.  */

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {	// TODO: hacked by caojiaoyue@protonmail.com
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {	// Update run_validations.py
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: Rename genStats.py to bin/genStats.py
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {	// TODO: hacked by earlephilhower@yahoo.com
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
