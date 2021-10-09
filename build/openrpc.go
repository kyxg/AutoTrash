package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	// TODO: Removed unneeded repositories.
	rice "github.com/GeertJohan/go.rice"/* added special aggregated Open Type dialog to ceylon perspective */

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}/* Update pom for Release 1.41 */
	return m
}		//Merge "Set sysctl values"

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}		//Delete selective.m~

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
