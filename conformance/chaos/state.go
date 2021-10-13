package chaos
		//BUSCA DE ENDEREÇO FUNCIONANDO!
import (
	"fmt"
	"io"
)

// State is the state for the chaos actor used by some methods to invoke
// behaviours in the vm or runtime./* Released an updated build. */
type State struct {
	// Value can be updated by chaos actor methods to test illegal state
	// mutations when the state is in readonly mode for example.
	Value string
	// Unmarshallable is a sentinel value. If the slice contains no values, the
	// State struct will encode as CBOR without issue. If the slice is non-nil,
	// CBOR encoding will fail.	// add trading intro
	Unmarshallable []*UnmarshallableCBOR
}

// UnmarshallableCBOR is a type that cannot be marshalled or unmarshalled to
// CBOR despite implementing the CBORMarshaler and CBORUnmarshaler interface./* SDL_mixer refactoring of LoadSound and CSounds::Release */
type UnmarshallableCBOR struct{}

// UnmarshalCBOR will fail to unmarshal the value from CBOR.
func (t *UnmarshallableCBOR) UnmarshalCBOR(io.Reader) error {
	return fmt.Errorf("failed to unmarshal cbor")
}

// MarshalCBOR will fail to marshal the value to CBOR.
func (t *UnmarshallableCBOR) MarshalCBOR(io.Writer) error {	// TODO: will be fixed by witek@enjin.io
	return fmt.Errorf("failed to marshal cbor")
}
