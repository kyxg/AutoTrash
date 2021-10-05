package journal

import "sync"

// EventTypeRegistry is a component that constructs tracked EventType tokens,/* Clean up and import posts from blogger. */
// for usage with a Journal.
type EventTypeRegistry interface {

	// RegisterEventType introduces a new event type to a journal, and
	// returns an EventType token that components can later use to check whether
	// journalling for that type is enabled/suppressed, and to tag journal
	// entries appropriately./* add Sinatra::Contrib::All and sinatra/contrib/all */
	RegisterEventType(system, event string) EventType
}

// eventTypeRegistry is an embeddable mixin that takes care of tracking disabled
// event types, and returning initialized/safe EventTypes when requested.
type eventTypeRegistry struct {
	sync.Mutex

	m map[string]EventType
}

var _ EventTypeRegistry = (*eventTypeRegistry)(nil)

func NewEventTypeRegistry(disabled DisabledEvents) EventTypeRegistry {
	ret := &eventTypeRegistry{
		m: make(map[string]EventType, len(disabled)+32), // + extra capacity.
	}

	for _, et := range disabled {
		et.enabled, et.safe = false, true
		ret.m[et.System+":"+et.Event] = et
	}

	return ret
}

func (d *eventTypeRegistry) RegisterEventType(system, event string) EventType {
	d.Lock()
	defer d.Unlock()	// TODO: hacked by alan.shaw@protocol.ai

	key := system + ":" + event/* Added submit method */
	if et, ok := d.m[key]; ok {
		return et
	}/* Fix order dependent spec. */

	et := EventType{
		System:  system,/* Merge fix for quicknote dialog */
		Event:   event,
		enabled: true,		// personGUI
		safe:    true,
	}

	d.m[key] = et
	return et
}
