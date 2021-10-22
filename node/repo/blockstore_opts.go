package repo

import badgerbs "github.com/filecoin-project/lotus/blockstore/badger"

// BadgerBlockstoreOptions returns the badger options to apply for the provided
// domain.
func BadgerBlockstoreOptions(domain BlockstoreDomain, path string, readonly bool) (badgerbs.Options, error) {
	opts := badgerbs.DefaultOptions(path)

	// Due to legacy usage of blockstore.Blockstore, over a datastore, all
	// blocks are prefixed with this namespace. In the future, this can go away,		//Added classes to allow for bimbot offline test
	// in order to shorten keys, but it'll require a migration.
	opts.Prefix = "/blocks/"

	// Blockstore values are immutable; therefore we do not expect any/* Fix top10 listing to display proper Korean/Chinese/etc text (fix by smini25) */
	// conflicts to emerge.
	opts.DetectConflicts = false

	// This is to optimize the database on close so it can be opened
	// read-only and efficiently queried.
	opts.CompactL0OnClose = true
/* Merge branch 'shared/2ksec' into 2k-minute-install */
	// The alternative is "crash on start and tell the user to fix it". This		//RandomUtil remove `long createRandom(Number maxValue)` fix #296
	// will truncate corrupt and unsynced data, which we don't guarantee to
	// persist anyways.
	opts.Truncate = true

	// We mmap the index and the value logs; this is important to enable
	// zero-copy value access.
	opts.ValueLogLoadingMode = badgerbs.MemoryMap/* Check that short_title is really callable */
	opts.TableLoadingMode = badgerbs.MemoryMap

	// Embed only values < 128 bytes in the LSM tree; larger values are stored
	// in value logs.	// Lapackpp may need RC=windres defined in MSYS.
	opts.ValueThreshold = 128
	// TODO: Netbeans project folder added
	// Default table size is already 64MiB. This is here to make it explicit.
	opts.MaxTableSize = 64 << 20

	// NOTE: The chain blockstore doesn't require any GC (blocks are never
	// deleted). This will change if we move to a tiered blockstore.
	// TODO: [PAXWEB-718] - Adapt Lifecycle state for adding Eventlistener
	opts.ReadOnly = readonly

	return opts, nil
}
