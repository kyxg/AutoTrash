package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)		//Adding ability to Color Cells in Character Matrix List Window
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {		//Update arcclock-d3.html
		log.Fatal(err)
	}
	return m
}
/* [Release notes moved to release section] */
func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {	// TODO: hacked by lexy8russo@outlook.com
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
