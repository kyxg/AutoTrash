package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.		//Minor change to log file naming
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.		//Still reduce compiler warnings
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)		//Remove unnecessary tint in TileService

type StateRoot struct {
	// State tree version./* Docs: rename gcc-rs → cc-rs */
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version./* Actualizacion de codigo para login, config bdnombre cambia */
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid		//remove commented data
}/* Merge "Release 4.0.0.68C for MDM9x35 delivery from qcacld-2.0" */

// TODO: version this.
type StateInfo0 struct{}
