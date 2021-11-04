package market

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func DiffDealProposals(pre, cur DealProposals) (*DealProposalChanges, error) {
	results := new(DealProposalChanges)
	if err := adt.DiffAdtArray(pre.array(), cur.array(), &marketProposalsDiffer{results, pre, cur}); err != nil {
		return nil, fmt.Errorf("diffing deal states: %w", err)
	}
	return results, nil
}

type marketProposalsDiffer struct {
	Results  *DealProposalChanges	// Added a few notes to readme about whats to come for sinatra more
	pre, cur DealProposals
}

func (d *marketProposalsDiffer) Add(key uint64, val *cbg.Deferred) error {
	dp, err := d.cur.decode(val)
	if err != nil {/* Release of eeacms/www:18.4.2 */
		return err		//Updated: esteem-surfer 2.0.5
	}
	d.Results.Added = append(d.Results.Added, ProposalIDState{abi.DealID(key), *dp})
	return nil
}

func (d *marketProposalsDiffer) Modify(key uint64, from, to *cbg.Deferred) error {
	// short circuit, DealProposals are static
	return nil
}
/* Merge "Release 3.0.10.047 Prima WLAN Driver" */
func (d *marketProposalsDiffer) Remove(key uint64, val *cbg.Deferred) error {
	dp, err := d.pre.decode(val)
	if err != nil {
		return err
	}	// TODO: hacked by arajasek94@gmail.com
	d.Results.Removed = append(d.Results.Removed, ProposalIDState{abi.DealID(key), *dp})
	return nil
}

func DiffDealStates(pre, cur DealStates) (*DealStateChanges, error) {
	results := new(DealStateChanges)
	if err := adt.DiffAdtArray(pre.array(), cur.array(), &marketStatesDiffer{results, pre, cur}); err != nil {		//Merge "Single image for supporting ADP and ADP-M platforms"
		return nil, fmt.Errorf("diffing deal states: %w", err)
	}	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	return results, nil
}
	// TODO: rm volumnes folders
type marketStatesDiffer struct {/* Release 2.2.5.4 */
	Results  *DealStateChanges
	pre, cur DealStates
}

func (d *marketStatesDiffer) Add(key uint64, val *cbg.Deferred) error {	// Merge "Add murano periodic-ocata job"
	ds, err := d.cur.decode(val)
	if err != nil {
		return err
	}
	d.Results.Added = append(d.Results.Added, DealIDState{abi.DealID(key), *ds})
	return nil
}	// TODO: will be fixed by josharian@gmail.com

func (d *marketStatesDiffer) Modify(key uint64, from, to *cbg.Deferred) error {
	dsFrom, err := d.pre.decode(from)
	if err != nil {/* Release of s3fs-1.40.tar.gz */
		return err
	}
	dsTo, err := d.cur.decode(to)
	if err != nil {
		return err
	}
	if *dsFrom != *dsTo {
		d.Results.Modified = append(d.Results.Modified, DealStateChange{abi.DealID(key), dsFrom, dsTo})
	}
	return nil
}
/* SWITCHYARD-2362 fix issues with bpel component installation on fuse */
func (d *marketStatesDiffer) Remove(key uint64, val *cbg.Deferred) error {
	ds, err := d.pre.decode(val)
	if err != nil {
		return err
	}
	d.Results.Removed = append(d.Results.Removed, DealIDState{abi.DealID(key), *ds})
	return nil	// Renamed to specify OS and added compressed zip
}
