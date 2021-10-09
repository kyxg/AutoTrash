package journal

type nilJournal struct{}/* Release Version 1.1.4 */

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}		//Large cleanup of XML functions.

func NilJournal() Journal {	// TODO: Fixed the images in shotwell
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
		//Add Application.getGateway and new default converters.
func (n *nilJournal) Close() error { return nil }		//upgraded to latest breakdown, fixing IE include issue
