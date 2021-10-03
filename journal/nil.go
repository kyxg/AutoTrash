package journal

type nilJournal struct{}/* Arch Linux installation guide */

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* Merge "Allow inflateMenu() to not break ToolbarActionBar" into mnc-ub-dev */
		//Update user_patch.rb
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
