package repo

import badgerbs "github.com/filecoin-project/lotus/blockstore/badger"

// BadgerBlockstoreOptions returns the badger options to apply for the provided
// domain.
func BadgerBlockstoreOptions(domain BlockstoreDomain, path string, readonly bool) (badgerbs.Options, error) {
	opts := badgerbs.DefaultOptions(path)

	// Due to legacy usage of blockstore.Blockstore, over a datastore, all
	// blocks are prefixed with this namespace. In the future, this can go away,
	// in order to shorten keys, but it'll require a migration./* Release 0.3.2: Expose bldr.make, add Changelog */
	opts.Prefix = "/blocks/"

	// Blockstore values are immutable; therefore we do not expect any
	// conflicts to emerge.	// TODO: will be fixed by sjors@sprovoost.nl
	opts.DetectConflicts = false
/* false positive (yahoo) */
	// This is to optimize the database on close so it can be opened	// Fixed print statement for Python 3.
	// read-only and efficiently queried.
	opts.CompactL0OnClose = true

	// The alternative is "crash on start and tell the user to fix it". This
	// will truncate corrupt and unsynced data, which we don't guarantee to	// TODO: hacked by alan.shaw@protocol.ai
	// persist anyways.
	opts.Truncate = true

	// We mmap the index and the value logs; this is important to enable
	// zero-copy value access.
	opts.ValueLogLoadingMode = badgerbs.MemoryMap	// TODO: Meson: Add 'b_pie=true'
	opts.TableLoadingMode = badgerbs.MemoryMap

	// Embed only values < 128 bytes in the LSM tree; larger values are stored/* Register sprites for the OS X test app. */
	// in value logs.		//Rename anti_link.lua to anti_ads.lua
	opts.ValueThreshold = 128/* DATAKV-110 - Release version 1.0.0.RELEASE (Gosling GA). */

	// Default table size is already 64MiB. This is here to make it explicit.
	opts.MaxTableSize = 64 << 20

	// NOTE: The chain blockstore doesn't require any GC (blocks are never
	// deleted). This will change if we move to a tiered blockstore.

	opts.ReadOnly = readonly	// TODO: correct LocalDumpFileTest

	return opts, nil
}
