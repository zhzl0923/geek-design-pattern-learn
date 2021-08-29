package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &ConcreteSubject{}
	sub.RegisterObserver(&ConcreteObserverOne{})
	sub.RegisterObserver(&ConcreteObserverTwo{})
	sub.Notify("hi")
}
