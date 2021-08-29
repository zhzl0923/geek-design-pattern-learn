package observer

import (
	"fmt"
	"reflect"
	"sync"
)

type EventBus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (eventBus *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	eventBus.lock.Lock()
	defer eventBus.lock.Unlock()
	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	h, ok := eventBus.handlers[topic]
	if !ok {
		h = []reflect.Value{}
	}
	h = append(h, v)
	eventBus.handlers[topic] = h

	return nil
}

func (eventBus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := eventBus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}
