package statemachine

import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha	// TODO: Added information about folders

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in./* Update ngBits.Breeze.Directives.js */
var ErrEventRejected = errors.New("event rejected")	// TODO: will be fixed by mail@bitpshr.net

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// Action represents the action to be executed in a given state./* Fixed name and added aliases */
type Action interface {	// TODO: will be fixed by julia@jvns.ca
	Execute(eventCtx EventContext) EventType
}		//Delete learnings.md

// Events represents a mapping of events and states.
type Events map[EventType]StateType
	// TODO: hacked by alan.shaw@protocol.ai
// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action/* Release 1.0.3 */
	Events Events/* Rename formas-de-pagamentos to formas-de-pagamentos.md */
}

// States represents a mapping of states and their implementations.
type States map[StateType]State
		//Merge branch 'b/Reg-Test-Plots' into f/Linear
// StateMachine represents the state machine.
type StateMachine struct {/* Ensure we have better validation */
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType/* Create p_configure_multicast.me */

	// States holds the configuration of states and events handled by the state machine.	// TODO: :police_car: Castle license information
	States States
/* Release version: 2.0.0-alpha01 [ci skip] */
	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex		//Added MATLAB emulation functions and docstrings for Python.
}		//refcount now uses atomic operations if possible

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
