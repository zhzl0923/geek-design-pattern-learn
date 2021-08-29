package chain

import (
	"fmt"
)

type Handler interface {
	SetSuccessor(successor Handler)
	Handle()
}

type HandlerA struct {
	successor Handler
}

func (a *HandlerA) SetSuccessor(successor Handler) {
	a.successor = successor
}

func (a *HandlerA) doHandle() bool {
	fmt.Println("HandlerA.doHandle")
	return true
}

func (a *HandlerA) Handle() {
	handled := a.doHandle()
	if handled && a.successor != nil {
		a.successor.Handle()
	}
}

type HandlerB struct {
	successor Handler
}

func (a *HandlerB) SetSuccessor(successor Handler) {
	a.successor = successor
}

func (a *HandlerB) doHandle() bool {
	fmt.Println("HandlerB.doHandle")
	return true
}

func (a *HandlerB) Handle() {
	handled := a.doHandle()
	if handled && a.successor != nil {
		a.successor.Handle()
	}
}

type HandlerChain struct {
	head Handler
	tail Handler
}

func (c *HandlerChain) AddHandler(handler Handler) {
	if c.head == nil {
		c.head = handler
		c.tail = handler
		return
	}
	c.tail.SetSuccessor(handler)
	c.tail = handler
}

func (c *HandlerChain) Handle() {
	if c.head != nil {
		c.head.Handle()
	}
}
