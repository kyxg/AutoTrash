package journal	// fixed seg-fault after read service with a still buggy mockup.

type nilJournal struct{}

// nilj is a singleton nil journal.	// TODO: will be fixed by davidad@alum.mit.edu
var nilj Journal = &nilJournal{}
	// TODO: hacked by ng8eke@163.com
func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }		//Delete Secret.java
/* de angular service maken (nog niet af) */
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}	// TODO: Found a legacy typo from skeleton and just fixed it

func (n *nilJournal) Close() error { return nil }
