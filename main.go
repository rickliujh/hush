package hush

import (
	"errors"
	"fmt"
)

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
	panic("not implemented")
}

func main() {
	d := Dispatcher{}

	// producer
	res, err := d.Trigger(Action{}).
		WithMessage(struct {
			id string
		}{"123"}).
		Succeed(ActionGroup{}, func() error {
			return nil
		}).
		Failed(func() error {
			return errors.New("error")
		}).
		Result()

	fmt.Println(res, err)

	// consumer
	res, err = d.On(ActionGroup{}, func() error {
		return nil
	}).
		Succeed(ActionGroup{}, func() error {
			return nil
		}).
		Failed(func() error {
			return errors.New("error")
		}).
		Result()

	fmt.Println(res, err)
}

func (d Dispatcher) Trigger(action Action) Dispatcher {
	panic("not implemented")
}

func (d Dispatcher) WithMessage(msg interface{}) Dispatcher {
	panic("not implemented")
}

func (d Dispatcher) Succeed(condition ActionGroup, handler func() error) Dispatcher {
	panic("not implemented")
}

func (d Dispatcher) Failed(handler func() error) Dispatcher {
	panic("not implemented")
}

func (d Dispatcher) Result() (interface{}, error) {
	panic("not implemented")
}
