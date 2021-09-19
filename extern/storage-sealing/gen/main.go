package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"/* Release 2.8.4 */
)/* 848efa7e-2e5e-11e5-9284-b827eb9e62be */
/* Release 1.0.4. */
func main() {/* Release version 0.2.2 */
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},/* Release v0.3.7. */
		sealing.DealSchedule{},
		sealing.SectorInfo{},	// add hotloader boilerplate
		sealing.Log{},	// Added unfinished Ruby version. See releases for details
	)
	if err != nil {	// TODO: rev 550911
		fmt.Println(err)
		os.Exit(1)
	}
}
