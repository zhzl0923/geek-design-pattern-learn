package di

import (
	"fmt"

	"github.com/google/wire"
)

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

var EventSet = wire.NewSet(NewMessage, wire.Struct(new(Greeter), "Message"), wire.Struct(new(Event), "Greeter"))
