package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)	// TODO: 7986fb21-2eae-11e5-9021-7831c1d44c14

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},	// ea2b08a2-2e57-11e5-9284-b827eb9e62be
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)		//Need to generate the receipt BEFORE it sends
		os.Exit(1)
	}
}
