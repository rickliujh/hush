package hush

type Dispatcher struct {
}

type State string

type From string

type Subject string

type Action struct {
	From    From
	Subject Subject
	Object  interface{}
	State   State
}

type Task func() error

type ActionGroup []Action

func (d *Dispatcher) On(action ActionGroup, task Task) *Dispatcher {
	panic("not implement")
}
