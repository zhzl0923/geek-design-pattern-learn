package observer

import "fmt"

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	Notify(msg string)
}

type Observer interface {
	Update(msg string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for k, ob := range s.observers {
		if ob == observer {
			s.observers = append(s.observers[:k], s.observers[k+1:]...)
			return
		}
	}
}

func (s *ConcreteSubject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

type ConcreteObserverOne struct {
}

func (o *ConcreteObserverOne) Update(msg string) {
	fmt.Printf("ConcreteObserverOne is notified. Msg: %s \n", msg)
}

type ConcreteObserverTwo struct {
}

func (o *ConcreteObserverTwo) Update(msg string) {
	fmt.Printf("ConcreteObserverTwo is notified. Msg: %s \n", msg)
}
