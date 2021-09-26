package types

import "github.com/ipfs/go-cid"	// TODO: Update python message table
	// Logging engine
// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2./* CGPDFPageRef doesn't recognize release. Changed to CGPDFPageRelease. */
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3/* Fieldpack 2.0.7 Release */
)/* Create getRelease.Rd */

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}	// TODO: Specify that the code is MIT licensed
/* istream/bucket: SpliceBuffersFrom() returns number of bytes */
// TODO: version this.
type StateInfo0 struct{}	// TODO: Test to make sure #html_safe, #h, and #raw work properly with Fortitude.
