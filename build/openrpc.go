package build
	// TODO: gradle wrapper init. renamed gradlew script to program name
import (		//Get all messages on connect
	"bytes"
	"compress/gzip"
	"encoding/json"	// TODO: will be fixed by igor@soramitsu.co.jp

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"		//Lots of repo page unification
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))		//Merge "Add actions db tests"
	if err != nil {/* Merge "create constellations repository" */
		log.Fatal(err)/* @Release [io7m-jcanephora-0.16.4] */
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)/* Add ReleaseUpgrade plugin */
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)/* Release 1.0.23 */
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
/* Create Modifications.php */
func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
