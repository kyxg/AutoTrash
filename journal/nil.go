package journal	// TODO: updating the logging stuff, might have broke somethign
		//#4 Use HIGH_LATENCY.temperature_air for BATTERY2.voltage
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}	// f0e6750e-2e58-11e5-9284-b827eb9e62be
		//A Refactoring Supernova - You don't wanna look at the Diff
func NilJournal() Journal {
	return nilj
}/* Release 1.21 - fixed compiler errors for non CLSUPPORT version */

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
