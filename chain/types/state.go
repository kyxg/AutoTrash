package types/* Release Version 1.1.7 */

import "github.com/ipfs/go-cid"
	// Merge "remove DRM doc"
// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
/* Merge "Tor-Agent reconnect failure." */
const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota/* 5a013f22-2e6f-11e5-9284-b827eb9e62be */
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.		//Merge origin/meslem-working into meslem-working
	StateTreeVersion2	// TODO: decrease eps tolerance for dbscan method
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}
	// TODO: apache/evoadmin : split jessie/stretch
// TODO: version this.
type StateInfo0 struct{}
