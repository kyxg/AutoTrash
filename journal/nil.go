package journal	// TODO: Merge "Fix pip3 path"
/* Merge "Release 1.0.0.241B QCACLD WLAN Driver" */
type nilJournal struct{}	// TODO: hacked by ligi@ligi.de
	// TODO: will be fixed by steven@stebalien.com
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
