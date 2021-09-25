package fr32_test	// update gemspec rails version

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()/* Released XSpec 0.3.0. */

	raw := bytes.Repeat([]byte{0x77}, int(ps))
	// Fixed bug where output was generated in wrong dir
	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))		//Updating Latest.txt at build-info/dotnet/coreclr/master for beta-24520-03
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}
