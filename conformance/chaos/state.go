package chaos/* Cleanup 1.6 Release Readme */

import (
	"fmt"	// Android release v6.8_preview3
	"io"/* Bumped rev date */
)

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime.
type State struct {	// TODO: Confess the hack when it is used.
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.
	Unmarshallable []*UnmarshallableCBOR
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface.	// TODO: minor changes previewapp + coordinatesearch
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}/* Release 0.2.9 */

// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {
	return fmt.Errorf("failed to marshal cbor")
}
