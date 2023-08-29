package playground

import (
	"fmt"
	"sync"

	"github.com/Chris2Schill/learngo/eventbus"
)

const (
	Event1 eventbus.Event = iota
	Event2
	Event3
	Event4
)

type Playground struct {
	*sync.WaitGroup
	eventBus     *eventbus.Dispatcher
	subbedEvents map[eventbus.Event](chan int)
}

func New(eb *eventbus.Dispatcher) *Playground {
	subbedEvents := make(map[eventbus.Event]chan int, 0)

	subbedEvents[Event1] = eb.Subscribe(Event1)
	subbedEvents[Event2] = eb.Subscribe(Event2)
	subbedEvents[Event3] = eb.Subscribe(Event3)

	pg := Playground{
		eventBus:     eb,
		subbedEvents: subbedEvents,
	}
	return &pg
}

func (pg *Playground) Process() {
	var val int
	for {
		select {
		case val = <-pg.subbedEvents[Event1]:
			fmt.Println("Recieved event1, val=", val)
		case val = <-pg.subbedEvents[Event2]:
			fmt.Println("Recieved event2, val=", val)
		case val = <-pg.subbedEvents[Event3]:
			fmt.Println("Recieved event3, val=", val)
		}
		pg.Done()
	}
}
