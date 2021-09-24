package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj/* Docs(guide): Bad link for ShadyCSS Polyfill */
}
/* New version of Cherish - 0.8 */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
