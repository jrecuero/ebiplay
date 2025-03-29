package main

import "github.com/jrecuero/ebiplay/pkg/engine"

type State int

const (
	STATE_START State = iota
	STATE_WAIT_INPUT
	STATE_WAIT_ACTION
	STATE_RUN_OTHERS
	STATE_END
)

type StateMachine struct {
	*engine.Base
	state State
}

func NewStateMachine(name string) *StateMachine {
	return &StateMachine{
		Base:  engine.NewBase(name),
		state: STATE_START,
	}
}

func (s *StateMachine) GetState() State {
	return s.state
}

func (s *StateMachine) Run() {
	switch s.state {
	case STATE_START:
		// state machine and all actors setup.
	case STATE_WAIT_INPUT:
		// wait for user directional input for movement.
	case STATE_WAIT_ACTION:
		// wait for user action input.
	case STATE_RUN_OTHERS:
		// run any other actor in the board.
	case STATE_END:
		// end state machine and all actors.
	default:
	}
}

func (s *StateMachine) SetState(state State) {
	s.state = state
}
