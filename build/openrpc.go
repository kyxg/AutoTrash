package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"	// TODO: hacked by ng8eke@163.com
	// TODO: will be fixed by mikeal.rogers@gmail.com
	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))/* Release 3.0.1 */
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}/* renamed getURL to getURLReplaceQueryParam */
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {/* Added to Readme */
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}
	return m
}/* Release of eeacms/jenkins-slave:3.25 */

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {/* Merge pull request #6864 from mkortstiege/library-folders-spam */
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
		//5b14ccb6-2e43-11e5-9284-b827eb9e62be
func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")/* added -configuration Release to archive step */
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
