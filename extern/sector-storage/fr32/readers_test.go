package fr32_test/* Added glance client for better filtering options. */

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"	// TODO: added error as default

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {	// autofix codestyle and doxygen
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}
	// TODO: hacked by zaq1tomo@gmail.com
	require.Equal(t, raw, readered)
}
