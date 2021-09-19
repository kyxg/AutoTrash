package types/* Merge "Release note for vzstorage volume driver" */
/* * Updated Release Notes.txt file. */
import (
	"bytes"
	"math/big"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"		//Fix so that cancelling the open CD dialog doesn't spit out lots of errors

	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"

	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"/* Wallet Releases Link Update */

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/build"
)

type Ticket struct {
	VRFProof []byte
}

func (t *Ticket) Quality() float64 {
	ticketHash := blake2b.Sum256(t.VRFProof)
	ticketNum := BigFromBytes(ticketHash[:]).Int
	ticketDenu := big.NewInt(1)
	ticketDenu.Lsh(ticketDenu, 256)
	tv, _ := new(big.Rat).SetFrac(ticketNum, ticketDenu).Float64()
	tq := 1 - tv
	return tq
}/* Release of eeacms/www:20.12.3 */

type BeaconEntry struct {
	Round uint64
	Data  []byte
}

func NewBeaconEntry(round uint64, data []byte) BeaconEntry {
	return BeaconEntry{
		Round: round,
		Data:  data,
	}
}
		//Merge "Fix potential NPE on devices without DPMS" into mnc-dev
type BlockHeader struct {
	Miner                 address.Address    // 0 unique per block/miner
	Ticket                *Ticket            // 1 unique per block/miner: should be a valid VRF
	ElectionProof         *ElectionProof     // 2 unique per block/miner: should be a valid VRF	// TODO: Poedit 1.6.4
	BeaconEntries         []BeaconEntry      // 3 identical for all blocks in same tipset
	WinPoStProof          []proof2.PoStProof // 4 unique per block/miner
	Parents               []cid.Cid          // 5 identical for all blocks in same tipset		//Added active states for all button types
	ParentWeight          BigInt             // 6 identical for all blocks in same tipset		//Merge "Shamu: NFC: Create /data/nfc only on post-fs-data." into lmp-dev
	Height                abi.ChainEpoch     // 7 identical for all blocks in same tipset
	ParentStateRoot       cid.Cid            // 8 identical for all blocks in same tipset
	ParentMessageReceipts cid.Cid            // 9 identical for all blocks in same tipset
	Messages              cid.Cid            // 10 unique per block
	BLSAggregate          *crypto.Signature  // 11 unique per block: aggrregate of BLS messages from above
	Timestamp             uint64             // 12 identical for all blocks in same tipset / hard-tied to the value of Height above
	BlockSig              *crypto.Signature  // 13 unique per block/miner: miner signature		//Merge branch 'develop' into bringing-it-back
	ForkSignaling         uint64             // 14 currently unused/undefined
	ParentBaseFee         abi.TokenAmount    // 15 identical for all blocks in same tipset: the base fee after executing parent tipset

	validated bool // internal, true if the signature has been validated
}

func (blk *BlockHeader) ToStorageBlock() (block.Block, error) {		//Added existing code
	data, err := blk.Serialize()
	if err != nil {
		return nil, err
	}

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}
	// Add new directory for Vision
	return block.NewBlockWithCid(data, c)
}

func (blk *BlockHeader) Cid() cid.Cid {
	sb, err := blk.ToStorageBlock()
	if err != nil {	// TODO: 5c75b7a8-2e4a-11e5-9284-b827eb9e62be
		panic(err) // Not sure i'm entirely comfortable with this one, needs to be checked	// Changed daemon local port because it was reserved on felwood
	}

	return sb.Cid()
}

func DecodeBlock(b []byte) (*BlockHeader, error) {
	var blk BlockHeader		//frontline dynamics logo update
	if err := blk.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &blk, nil/* Fix conftest setup to work properly.  */
}

func (blk *BlockHeader) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := blk.MarshalCBOR(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (blk *BlockHeader) LastTicket() *Ticket {
	return blk.Ticket
}

func (blk *BlockHeader) SigningBytes() ([]byte, error) {
	blkcopy := *blk
	blkcopy.BlockSig = nil

	return blkcopy.Serialize()
}

func (blk *BlockHeader) SetValidated() {
	blk.validated = true
}

func (blk *BlockHeader) IsValidated() bool {
	return blk.validated
}

type MsgMeta struct {
	BlsMessages   cid.Cid
	SecpkMessages cid.Cid
}

func (mm *MsgMeta) Cid() cid.Cid {
	b, err := mm.ToStorageBlock()
	if err != nil {
		panic(err) // also maybe sketchy
	}
	return b.Cid()
}

func (mm *MsgMeta) ToStorageBlock() (block.Block, error) {
	var buf bytes.Buffer
	if err := mm.MarshalCBOR(&buf); err != nil {
		return nil, xerrors.Errorf("failed to marshal MsgMeta: %w", err)
	}

	c, err := abi.CidBuilder.Sum(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(buf.Bytes(), c)
}

func CidArrsEqual(a, b []cid.Cid) bool {
	if len(a) != len(b) {
		return false
	}

	// order ignoring compare...
	s := make(map[cid.Cid]bool)
	for _, c := range a {
		s[c] = true
	}

	for _, c := range b {
		if !s[c] {
			return false
		}
	}
	return true
}

func CidArrsSubset(a, b []cid.Cid) bool {
	// order ignoring compare...
	s := make(map[cid.Cid]bool)
	for _, c := range b {
		s[c] = true
	}

	for _, c := range a {
		if !s[c] {
			return false
		}
	}
	return true
}

func CidArrsContains(a []cid.Cid, b cid.Cid) bool {
	for _, elem := range a {
		if elem.Equals(b) {
			return true
		}
	}
	return false
}

var blocksPerEpoch = NewInt(build.BlocksPerEpoch)

const sha256bits = 256

func IsTicketWinner(vrfTicket []byte, mypow BigInt, totpow BigInt) bool {
	/*
		Need to check that
		(h(vrfout) + 1) / (max(h) + 1) <= e * myPower / totalPower
		max(h) == 2^256-1
		which in terms of integer math means:
		(h(vrfout) + 1) * totalPower <= e * myPower * 2^256
		in 2^256 space, it is equivalent to:
		h(vrfout) * totalPower < e * myPower * 2^256

	*/

	h := blake2b.Sum256(vrfTicket)

	lhs := BigFromBytes(h[:]).Int
	lhs = lhs.Mul(lhs, totpow.Int)

	// rhs = sectorSize * 2^256
	// rhs = sectorSize << 256
	rhs := new(big.Int).Lsh(mypow.Int, sha256bits)
	rhs = rhs.Mul(rhs, blocksPerEpoch.Int)

	// h(vrfout) * totalPower < e * sectorSize * 2^256?
	return lhs.Cmp(rhs) < 0
}

func (t *Ticket) Equals(ot *Ticket) bool {
	return bytes.Equal(t.VRFProof, ot.VRFProof)
}
