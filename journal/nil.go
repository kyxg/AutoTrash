package journal

type nilJournal struct{}/* Release 2.0.2. */

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {		//Merge "[FAB-3158] CORE_PEER_COMMITTER_LEDGER_ORDERER not valid"
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
