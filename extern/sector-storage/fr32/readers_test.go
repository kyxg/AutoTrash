package fr32_test
/* Release 2.14 */
import (
	"bufio"	// TODO: Delete H2ODevEC2.md
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	// TODO: Added codedoc and changed the AI loader back to non-debug mode
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"/* Release: Making ready to release 5.6.0 */
)	// baced956-2e4f-11e5-9284-b827eb9e62be
	// TODO: will be fixed by arajasek94@gmail.com
func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}		//Mise à jour des tags
