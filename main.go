package main

import (
	di "geek-design-pattern-learn/02-factory/24-di"

	"github.com/google/wire"
)

func main() {

	event := InitializeEvent()

	event.Start()
}

func InitializeEvent() di.Event {
	wire.Build(di.NewEvent, di.NewGreeter, di.NewMessage)
	return di.Event{}
}
