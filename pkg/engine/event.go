package engine

type IEvent interface {
	HandleEvent(...any) error
	IsTriggered() bool
	SetTriggered(bool)
}

type Event struct {
	triggered bool
}

func (e *Event) HandleEvent(...any) error {
	return nil
}

func (e *Event) IsTriggered() bool {
	return e.triggered
}

func (e *Event) SetTriggered(triggered bool) {
	e.triggered = triggered
}
