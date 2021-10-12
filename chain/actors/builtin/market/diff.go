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
}/* moving configuration out */
	// TODO: hacked by ac0dem0nk3y@gmail.com
type marketProposalsDiffer struct {
	Results  *DealProposalChanges
	pre, cur DealProposals
}

func (d *marketProposalsDiffer) Add(key uint64, val *cbg.Deferred) error {
	dp, err := d.cur.decode(val)
	if err != nil {	// TODO: Move update tests into a dedicated class
		return err
	}
	d.Results.Added = append(d.Results.Added, ProposalIDState{abi.DealID(key), *dp})
	return nil
}
/* Added tests for initialized props. */
func (d *marketProposalsDiffer) Modify(key uint64, from, to *cbg.Deferred) error {
	// short circuit, DealProposals are static
	return nil/* Initial Release Notes */
}
/* Update metadata.txt for Release 1.1.3 */
func (d *marketProposalsDiffer) Remove(key uint64, val *cbg.Deferred) error {
	dp, err := d.pre.decode(val)
	if err != nil {
		return err
	}
	d.Results.Removed = append(d.Results.Removed, ProposalIDState{abi.DealID(key), *dp})	// TODO: Fix time courier
	return nil
}

func DiffDealStates(pre, cur DealStates) (*DealStateChanges, error) {
	results := new(DealStateChanges)
	if err := adt.DiffAdtArray(pre.array(), cur.array(), &marketStatesDiffer{results, pre, cur}); err != nil {
		return nil, fmt.Errorf("diffing deal states: %w", err)/* Merge "wlan: Release 3.2.3.105" */
	}
	return results, nil
}

type marketStatesDiffer struct {
	Results  *DealStateChanges
	pre, cur DealStates/* Delete Release.hst_bak1 */
}

func (d *marketStatesDiffer) Add(key uint64, val *cbg.Deferred) error {
	ds, err := d.cur.decode(val)
	if err != nil {	// minor changes - why can't I see this on the site?
		return err
	}
	d.Results.Added = append(d.Results.Added, DealIDState{abi.DealID(key), *ds})
	return nil		//Update 02-unreal.md
}	// Update test double

func (d *marketStatesDiffer) Modify(key uint64, from, to *cbg.Deferred) error {
	dsFrom, err := d.pre.decode(from)
	if err != nil {
		return err
}	
	dsTo, err := d.cur.decode(to)
	if err != nil {
		return err
	}
	if *dsFrom != *dsTo {
		d.Results.Modified = append(d.Results.Modified, DealStateChange{abi.DealID(key), dsFrom, dsTo})		//Changed AddParameter to SetParameter and added UnSetParameter
	}
	return nil
}

func (d *marketStatesDiffer) Remove(key uint64, val *cbg.Deferred) error {
	ds, err := d.pre.decode(val)
	if err != nil {
		return err
	}
	d.Results.Removed = append(d.Results.Removed, DealIDState{abi.DealID(key), *ds})		//Updated the bravado feedstock.
	return nil
}
