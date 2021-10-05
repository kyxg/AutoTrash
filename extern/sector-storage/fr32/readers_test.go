package fr32_test

import (
	"bufio"
	"bytes"	// fix composer in dev branch
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
		//Update PKGBUILD for 1.0
	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"	// Update AMI with minor changes (package updates)
)

func TestUnpadReader(t *testing.T) {	// TODO: hacked by hugomrdias@gmail.com
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())/* Merge "msm: ipa: Bug fix in IPA RM" */
	fr32.Pad(raw, padOut)
/* Create Mars Exploration.cs */
	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)/* Initial commit. Release version */
}
