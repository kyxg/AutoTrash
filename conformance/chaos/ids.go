package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)
	// TODO: will be fixed by lexy8russo@outlook.com
// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}	// So many slots! So many colors!
	return c
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds/* Documentation and website changes. Release 1.3.1. */
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)/* Release for 3.15.0 */
	addr, err := address.NewIDAddress(98)
	if err != nil {
		panic(err)
	}
	return addr
}()		//add synaptics touchscreen
