package chain

import "testing"

func TestHandlerChain_Handle(t *testing.T) {
	chain := &HandlerChain{}
	chain.AddHandler(&HandlerA{})
	chain.AddHandler(&HandlerB{})
	chain.Handle()
}
