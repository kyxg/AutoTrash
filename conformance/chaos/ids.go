package chaos
	// TODO: Create AssemblyAutoBoard.ino
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))/* Update positions.xml */
	if err != nil {
		panic(err)
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98	// Favicon refix
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton./* added epi_distcorr to init files */
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {
		panic(err)
	}
	return addr
}()
