package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}/* removed logging from mix() */

func NilJournal() Journal {		//Description & data upload
	return nilj
}/* Update WebAppReleaseNotes.rst */

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }	// Delete analyze_trajectory.m

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}/* UP to Pre-Release or DOWN to Beta o_O */
/* :relieved: :relieved: */
func (n *nilJournal) Close() error { return nil }
