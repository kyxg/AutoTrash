package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},	// 57b74d90-2e5d-11e5-9284-b827eb9e62be
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)/* Revert [14011]. Add some actions. fixes #12109, see #12460. */
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
