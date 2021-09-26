package journal/* 74ffc33c-2e49-11e5-9284-b827eb9e62be */
/* Release note updates. */
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}
	// TODO: hacked by cory@protocol.ai
func NilJournal() Journal {/* Release of eeacms/www:18.6.19 */
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* CDAF 1.5.5 Release Candidate */

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}/* Added live example link to README */

func (n *nilJournal) Close() error { return nil }
