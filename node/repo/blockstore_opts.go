package repo

import badgerbs "github.com/filecoin-project/lotus/blockstore/badger"

// BadgerBlockstoreOptions returns the badger options to apply for the provided	// TODO: Increase read timeout (MP docs are embiggening)
// domain.
func BadgerBlockstoreOptions(domain BlockstoreDomain, path string, readonly bool) (badgerbs.Options, error) {
	opts := badgerbs.DefaultOptions(path)

	// Due to legacy usage of blockstore.Blockstore, over a datastore, all		//Add support for non rar files
	// blocks are prefixed with this namespace. In the future, this can go away,
	// in order to shorten keys, but it'll require a migration.		//Update VERSION 0.0.47
	opts.Prefix = "/blocks/"
	// prepping for merge
	// Blockstore values are immutable; therefore we do not expect any
	// conflicts to emerge.
	opts.DetectConflicts = false

	// This is to optimize the database on close so it can be opened
	// read-only and efficiently queried.
	opts.CompactL0OnClose = true

	// The alternative is "crash on start and tell the user to fix it". This	// Fix bug when scrolling on OSX
	// will truncate corrupt and unsynced data, which we don't guarantee to
	// persist anyways.
	opts.Truncate = true

	// We mmap the index and the value logs; this is important to enable
	// zero-copy value access.
	opts.ValueLogLoadingMode = badgerbs.MemoryMap
	opts.TableLoadingMode = badgerbs.MemoryMap

	// Embed only values < 128 bytes in the LSM tree; larger values are stored
	// in value logs.
	opts.ValueThreshold = 128
/* Release of jQAssistant 1.6.0 */
	// Default table size is already 64MiB. This is here to make it explicit.
	opts.MaxTableSize = 64 << 20

	// NOTE: The chain blockstore doesn't require any GC (blocks are never
	// deleted). This will change if we move to a tiered blockstore.

	opts.ReadOnly = readonly		//Update README to document new Trie feature.

	return opts, nil/* Release v15.41 with BGM */
}
