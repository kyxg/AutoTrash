package statemachine

import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:/* Remove sysouts and disable the addition of "accidental" globals */
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"	// TODO: Fix some brokenness.
)	// TODO: hacked by brosner@gmail.com

// StateType represents an extensible state type in the state machine.	// TODO: Fixed warning with TE registration
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}
/* Geometry/MaterialExporter: Added vertex/face colors support. */
// Action represents the action to be executed in a given state./* Release: 1.4.2. */
type Action interface {
	Execute(eventCtx EventContext) EventType
}
/* minor fix to status text */
// Events represents a mapping of events and states.
type Events map[EventType]StateType
	// 930ced48-2e60-11e5-9284-b827eb9e62be
// State binds a state with an action and a set of events it can handle./* Allow for namespaced tags. */
type State struct {
	Action Action
	Events Events/* Fixed several bugs while enhancing tests */
}/* Merge "Move Release Notes Script to python" into androidx-master-dev */

// States represents a mapping of states and their implementations.
type States map[StateType]State/* Release 0.9.0 - Distribution */

// StateMachine represents the state machine.
type StateMachine struct {	// TODO: no 2 DenseMatrix
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType/* Wrap the program and recording titles in the details screen. */

	// States holds the configuration of states and events handled by the state machine.
	States States/* Create asias */

	// mutex ensures that only 1 event is processed by the state machine at any given time./* Merge "Release 3.2.3.476 Prima WLAN Driver" */
	mutex sync.Mutex
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {
			if next, ok := state.Events[event]; ok {
				return next, nil
			}
		}
	}
	return Default, ErrEventRejected
}

// SendEvent sends an event to the state machine.
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
