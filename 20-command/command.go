package command

import "fmt"

type Command interface {
	Execute() error
}

type StartCommand struct{}

func (s *StartCommand) Execute() error {
	fmt.Println("start...")
	return nil
}

type EndCommand struct{}

func (e *EndCommand) Execute() error {
	fmt.Println("end...")
	return nil
}
