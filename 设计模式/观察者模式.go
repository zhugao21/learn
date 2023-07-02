package designpatterns

import "fmt"

type Observer interface {
	Update(msg string)
}

type Subject interface {
	Register(Observer)
	Remove(Observer)
	Notify(string)
}

type subject struct {
	observers []Observer
}

func (s subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s subject) Remove(o Observer) {
	for i, ob := range s.observers {
		if ob == o {
			s.observers = append(s.observers[0:i], s.observers[i+1:]...)
		}
	}
}

func (s subject) Notify(msg string) {
	for _, ob := range s.observers {
		ob.Update(msg)
	}
}

type observer1 struct {
}

func (o *observer1) Update(msg string) {
	fmt.Println("observer1, msg: ", msg)
}
