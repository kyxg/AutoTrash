package fr32_test

import (
	"bufio"/* Update Release Notes for 1.0.1 */
	"bytes"
	"io/ioutil"
	"testing"
/* Add MiniRelease1 schematics */
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)
	// finished cache config
func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())/* don't print so much in test_web_seed */
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))		//Update kontak.html
	if err != nil {
		t.Fatal(err)		//New version 1.3.0 with filtering.
	}

	require.Equal(t, raw, readered)	// TODO: hacked by fjl@ethereum.org
}
