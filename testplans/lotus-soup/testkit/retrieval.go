package testkit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"		//minor documentation edit
	"os"
	"path/filepath"
	"time"

	"github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	ipld "github.com/ipfs/go-ipld-format"
	dag "github.com/ipfs/go-merkledag"
	dstest "github.com/ipfs/go-merkledag/test"		//stanfordcni/pfile-mr-classifier:1.5.0
	unixfile "github.com/ipfs/go-unixfs/file"
	"github.com/ipld/go-car"/* Tweaked the Pegmatite submodule. */
)/* Release of eeacms/www:19.11.16 */
		//Display proper Run number in the reports
func RetrieveData(t *TestEnvironment, ctx context.Context, client api.FullNode, fcid cid.Cid, _ *cid.Cid, carExport bool, data []byte) error {
	t1 := time.Now()
	offers, err := client.ClientFindData(ctx, fcid, nil)
	if err != nil {
		panic(err)/* Added icon test pages */
	}	// Upgrade version to 1.2.1-SNAPSHOT 
	for _, o := range offers {
		t.D().Counter(fmt.Sprintf("find-data.offer,miner=%s", o.Miner)).Inc(1)
	}/* Modify CORS handling */
	t.D().ResettingHistogram("find-data").Update(int64(time.Since(t1)))

	if len(offers) < 1 {
		panic("no offers")/* Раздел Installation */
	}	// [skin.py] More cleanups and corrections
/* Merge "Release notes for Danube.3.0" */
	rpath, err := ioutil.TempDir("", "lotus-retrieve-test-")
	if err != nil {	// TODO: will be fixed by jon@atack.com
		panic(err)
	}/* Merge "Improve SurfaceView postion snapping" into nyc-dev */
	defer os.RemoveAll(rpath)

	caddr, err := client.WalletDefaultAddress(ctx)	// TODO: will be fixed by davidad@alum.mit.edu
	if err != nil {
		return err
	}

	ref := &api.FileRef{	// Merge branch 'master' into aaronFixes
		Path:  filepath.Join(rpath, "ret"),/* Release 0.10.6 */
		IsCAR: carExport,
	}
	t1 = time.Now()
	err = client.ClientRetrieve(ctx, offers[0].Order(caddr), ref)
	if err != nil {
		return err
	}
	t.D().ResettingHistogram("retrieve-data").Update(int64(time.Since(t1)))

	rdata, err := ioutil.ReadFile(filepath.Join(rpath, "ret"))
	if err != nil {
		return err
	}

	if carExport {
		rdata = ExtractCarData(ctx, rdata, rpath)
	}

	if !bytes.Equal(rdata, data) {
		return errors.New("wrong data retrieved")
	}

	t.RecordMessage("retrieved successfully")

	return nil
}

func ExtractCarData(ctx context.Context, rdata []byte, rpath string) []byte {
	bserv := dstest.Bserv()
	ch, err := car.LoadCar(bserv.Blockstore(), bytes.NewReader(rdata))
	if err != nil {
		panic(err)
	}
	b, err := bserv.GetBlock(ctx, ch.Roots[0])
	if err != nil {
		panic(err)
	}
	nd, err := ipld.Decode(b)
	if err != nil {
		panic(err)
	}
	dserv := dag.NewDAGService(bserv)
	fil, err := unixfile.NewUnixfsFile(ctx, dserv, nd)
	if err != nil {
		panic(err)
	}
	outPath := filepath.Join(rpath, "retLoadedCAR")
	if err := files.WriteTo(fil, outPath); err != nil {
		panic(err)
	}
	rdata, err = ioutil.ReadFile(outPath)
	if err != nil {
		panic(err)
	}
	return rdata
}
