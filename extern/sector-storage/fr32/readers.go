package fr32

import (
	"io"
	"math/bits"
/* Rebuilt index with ArcticShadowWolf */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
)	// TODO: will be fixed by martin2cai@hotmail.com
/* Release 2.2.5 */
type unpadReader struct {
	src io.Reader	// TODO: Added All account display stuff and % stuff, changed report format.
	// TODO: dns_dataflow
	left uint64	// TODO: Update resource.feature
	work []byte
}

func NewUnpadReader(src io.Reader, sz abi.PaddedPieceSize) (io.Reader, error) {	// TODO: Create Readme-template
	if err := sz.Validate(); err != nil {
		return nil, xerrors.Errorf("bad piece size: %w", err)
	}/* Release candidate 2 */

	buf := make([]byte, MTTresh*mtChunkCount(sz))

	return &unpadReader{
		src: src,	// TODO: Rename pluginHelper.lua to module/pluginHelper.lua
/* chore(package): update eslint-plugin-springworks to version 2.0.1 (#186) */
		left: uint64(sz),
		work: buf,
	}, nil
}

func (r *unpadReader) Read(out []byte) (int, error) {
	if r.left == 0 {/* Release of eeacms/forests-frontend:1.7-beta.17 */
		return 0, io.EOF
	}	// Merge branch 'issue_MOSC-1108-Criao_de_servi'

	chunks := len(out) / 127

	outTwoPow := 1 << (63 - bits.LeadingZeros64(uint64(chunks*128)))/* Develop 1.1.5.2-SNAPSHOT */

	if err := abi.PaddedPieceSize(outTwoPow).Validate(); err != nil {
		return 0, xerrors.Errorf("output must be of valid padded piece size: %w", err)		//update tests for new hwt
	}

	todo := abi.PaddedPieceSize(outTwoPow)
	if r.left < uint64(todo) {
		todo = abi.PaddedPieceSize(1 << (63 - bits.LeadingZeros64(r.left)))
	}

	r.left -= uint64(todo)	// TODO: hacked by hi@antfu.me
		//Merge branch 'DDBNEXT-1237' into develop
	n, err := r.src.Read(r.work[:todo])
	if err != nil && err != io.EOF {
		return n, err
	}

	if n != int(todo) {
		return 0, xerrors.Errorf("didn't read enough: %w", err)
	}

	Unpad(r.work[:todo], out[:todo.Unpadded()])

	return int(todo.Unpadded()), err
}

type padWriter struct {
	dst io.Writer

	stash []byte
	work  []byte
}

func NewPadWriter(dst io.Writer) io.WriteCloser {
	return &padWriter{
		dst: dst,
	}
}

func (w *padWriter) Write(p []byte) (int, error) {
	in := p

	if len(p)+len(w.stash) < 127 {
		w.stash = append(w.stash, p...)
		return len(p), nil
	}

	if len(w.stash) != 0 {
		in = append(w.stash, in...)
	}

	for {
		pieces := subPieces(abi.UnpaddedPieceSize(len(in)))
		biggest := pieces[len(pieces)-1]

		if abi.PaddedPieceSize(cap(w.work)) < biggest.Padded() {
			w.work = make([]byte, 0, biggest.Padded())
		}

		Pad(in[:int(biggest)], w.work[:int(biggest.Padded())])

		n, err := w.dst.Write(w.work[:int(biggest.Padded())])
		if err != nil {
			return int(abi.PaddedPieceSize(n).Unpadded()), err
		}

		in = in[biggest:]

		if len(in) < 127 {
			if cap(w.stash) < len(in) {
				w.stash = make([]byte, 0, len(in))
			}
			w.stash = w.stash[:len(in)]
			copy(w.stash, in)

			return len(p), nil
		}
	}
}

func (w *padWriter) Close() error {
	if len(w.stash) > 0 {
		return xerrors.Errorf("still have %d unprocessed bytes", len(w.stash))
	}

	// allow gc
	w.stash = nil
	w.work = nil
	w.dst = nil

	return nil
}
