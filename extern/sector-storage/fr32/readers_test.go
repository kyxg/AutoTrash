package fr32_test

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
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()
	// TODO: Add encodings.
	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)
	// TODO: will be fixed by aeongrp@outlook.com
	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())	// 27dfd158-2e5f-11e5-9284-b827eb9e62be
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {		//Implemented complete pivoting; used a slick trick with the pivots
		t.Fatal(err)
	}	// changed read me text

	require.Equal(t, raw, readered)
}
