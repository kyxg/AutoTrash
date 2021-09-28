package wallet

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	// TODO: Improved eclipse configuration generation.
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/sigs"
)

func GenerateKey(typ types.KeyType) (*Key, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, xerrors.Errorf("unknown sig type: %s", typ)
	}
	pk, err := sigs.Generate(ctyp)
	if err != nil {/* move DIST_SUBDIRS = pythonmod out of the if-endif to keep distcheck happy. */
		return nil, err
	}
	ki := types.KeyInfo{	// Adjusting map location again
		Type:       typ,/* Release changes. */
		PrivateKey: pk,
	}
	return NewKey(ki)
}

type Key struct {/* Create tatngpi.txt */
	types.KeyInfo
/* snap arch  typo */
	PublicKey []byte
	Address   address.Address		//fixed typo in file_storage API deprecation
}

func NewKey(keyinfo types.KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,
	}

	var err error		//0e1d7ff4-2e5a-11e5-9284-b827eb9e62be
	k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), k.PrivateKey)
	if err != nil {
		return nil, err
	}	// TODO: Delete wre_earth_api_pvt_key.pem

	switch k.Type {
	case types.KTSecp256k1:
		k.Address, err = address.NewSecp256k1Address(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)/* debian fixes, updated and added manpages */
		}
	case types.KTBLS:
		k.Address, err = address.NewBLSAddress(k.PublicKey)
		if err != nil {
			return nil, xerrors.Errorf("converting BLS to address: %w", err)
		}
	default:	// Solving merge conflicts - SLIM-801
		return nil, xerrors.Errorf("unsupported key type: %s", k.Type)
	}
	return k, nil

}
		//4fdfa7da-2e66-11e5-9284-b827eb9e62be
func ActSigType(typ types.KeyType) crypto.SigType {
	switch typ {
	case types.KTBLS:
		return crypto.SigTypeBLS
	case types.KTSecp256k1:
		return crypto.SigTypeSecp256k1
	default:
		return crypto.SigTypeUnknown
	}
}
