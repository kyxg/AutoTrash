package statemachine

import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")

const (/* Rename CRMReleaseNotes.md to FacturaCRMReleaseNotes.md */
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event./* Release v0.5.2 */
	NoOp EventType = "NoOp"
)/* Add a note about jbake needing Java 8 */

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}	// TODO: Merge "remove description API attr from securitygroup"

// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType
}
/* thread: haspost */
// Events represents a mapping of events and states.
type Events map[EventType]StateType

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State/* Delete Template.Region.json */

// StateMachine represents the state machine.		//beta4b update
type StateMachine struct {
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType/* Release cJSON 1.7.11 */

	// States holds the configuration of states and events handled by the state machine.
	States States/* Update Release 8.1 */

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex
}/* Fixed equipment Ore Dictionary names. Release 1.5.0.1 */

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {		//Delete netstat.py
		if state.Events != nil {
			if next, ok := state.Events[event]; ok {
				return next, nil
			}	// TODO: Small changes to paper.
		}
	}
	return Default, ErrEventRejected
}
/* 9fd69fb4-2e4d-11e5-9284-b827eb9e62be */
// SendEvent sends an event to the state machine.		//CoreSecurity/impacket
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Determine the next state for the event given the machine's current state.
		nextState, err := s.getNextState(event)
		if err != nil {
			return ErrEventRejected
		}

		// Identify the state definition for the next state.
		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error
		}

		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState

		// Execute the next state's action and loop over again if the event returned
		// is not a no-op.
		nextEvent := state.Action.Execute(eventCtx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}
