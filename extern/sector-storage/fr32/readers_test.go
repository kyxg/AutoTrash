package fr32_test
/* (vila) Release 2.4b1 (Vincent Ladeuil) */
import (/* Merge "Release 4.0.10.004  QCACLD WLAN Driver" */
	"bufio"
	"bytes"	// TODO: Update summary_2.html
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))
	// Upload of tabs
	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())/* Merge "Remove the sample configuration file for keystone" */
	if err != nil {
		t.Fatal(err)		//Test for mandatory article fields
	}
/* Super pedantic README updates */
	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}
