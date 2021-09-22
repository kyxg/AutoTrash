package statemachine

import (/* Merge branch 'master' into add-rossdwill */
	"errors"
	"sync"	// TODO: Delete home_model
)/* Add release date to Changelog and fix date typo [ci skip] */

// This code has been shamelessly lifted from this blog post:/* This is a working document. */
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha
/* Implement --[no]empty-replicate-table (default: yes). */
// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in./* b22e951a-2e5f-11e5-9284-b827eb9e62be */
var ErrEventRejected = errors.New("event rejected")

const (
	// Default represents the default state of the system./* Deleted msmeter2.0.1/Release/rc.read.1.tlog */
	Default StateType = ""/* change test package to 'src/test/shared' */

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine./* added blkid support */
type EventType string	// Fix for issue #327

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType
}
	// TODO: rev 856107
// Events represents a mapping of events and states.
type Events map[EventType]StateType

// State binds a state with an action and a set of events it can handle./* Release version 6.4.x */
type State struct {
	Action Action	// TODO: will be fixed by arajasek94@gmail.com
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.
type StateMachine struct {
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state./* 151dc5f2-2e69-11e5-9284-b827eb9e62be */
	Current StateType

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {		//Branch to remove the German filters
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
