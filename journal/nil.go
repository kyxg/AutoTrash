package journal
	// Chapter 03 - Initial Commit
type nilJournal struct{}

// nilj is a singleton nil journal.	// TODO: hacked by sjors@sprovoost.nl
var nilj Journal = &nilJournal{}
	// TODO: Update People.java
func NilJournal() Journal {
	return nilj
}
/* Release 0.37.0 */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
