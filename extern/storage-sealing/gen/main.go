package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {		//2 spaces, not 4
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},/* Update ppd_options.c */
	)/* Banner image started to change #118 */
	if err != nil {
		fmt.Println(err)
		os.Exit(1)	// TODO: Merge "Bug 5368 - NeutronL3Adapter ipv6 work around for mac address resolver"
	}
}
