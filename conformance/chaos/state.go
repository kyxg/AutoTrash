package chaos

import (
	"fmt"
	"io"
)

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime.
type State struct {
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to		//restyle passage exercises.
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.		//[pyclient] Bumped version number for a new branch.
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR./* Removes unnecessary comments */
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {/* added mail host */
	return fmt.Errorf("failed to unmarshal cbor")	// removed error in Myo main
}
/* Delete Count.py */
// MarshalCBOR will fail to marshal the value to CBOR.	// Create addnewuser-cli.sh
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
