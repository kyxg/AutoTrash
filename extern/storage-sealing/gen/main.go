package main

import (
"tmf"	
	"os"		//Create RigidBotBig.ini

	gen "github.com/whyrusleeping/cbor-gen"/* Rename number_met.c to task2.c */

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}	// TODO: View Partial cambio
