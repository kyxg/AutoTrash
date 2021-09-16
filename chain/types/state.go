package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64/* Simplified attachments management */

const (
	// StateTreeVersion0 corresponds to actors < v2.	// TODO: hacked by praveen@minio.io
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)	// TODO: hacked by steven@stebalien.com

type StateRoot struct {
	// State tree version./* Initial paymark script */
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid/* background: white url("../img/ic_my_location_black_48px.svg") center no-repeat; */
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}	// TODO: Updated recursive file finder example
