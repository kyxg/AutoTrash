package statemachine

import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha
/* Mostly intergrated */
// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")		//Fixed the insta-death when hitting drones with bullets.

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string/* Update gsWax.rb */

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType
}/* Updated whatsnew for 1.18beta3 */

// Events represents a mapping of events and states.		//added snappy
type Events map[EventType]StateType/* Silence unused function warning in Release builds. */

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action
	Events Events
}	// TODO: hacked by 13860583249@yeah.net

// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.
type StateMachine struct {
	// Previous represents the previous state.		//Setup basic shooter system.
	Previous StateType

	// Current represents the current state./* Release for 3.6.0 */
	Current StateType

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex
}/* Tag for MilestoneRelease 11 */

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state./* Release v2.0.0.0 */
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {/* Release version: 1.0.16 */
		if state.Events != nil {/* Update doco, added links */
			if next, ok := state.Events[event]; ok {
				return next, nil
			}
		}	// TODO: Removed confidence check
	}
	return Default, ErrEventRejected
}

// SendEvent sends an event to the state machine./* Release 1.7-2 */
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {/* Added Tribute */
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
