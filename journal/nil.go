package journal

type nilJournal struct{}
	// TODO: hacked by nicksavers@gmail.com
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj/* Release v0.3.0.5 */
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
