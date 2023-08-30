package playground

import (
	"context"

	"github.com/Chris2Schill/learngo/eventbus"
	"github.com/Chris2Schill/learngo/logger"
)

const (
	Event1 eventbus.Event = iota
	Event2
	Event3
	Event4
)

type Playground struct {
	eventBus     eventbus.Dispatcher
	subbedEvents map[eventbus.Event](chan int)
}

func New(eb eventbus.Dispatcher) *Playground {
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

func (pg *Playground) Process(ctx context.Context) {
	var val int
	for {
		select {
		case val = <-pg.subbedEvents[Event1]:
			logger.Println("Recieved event1, val=", val)
		case val = <-pg.subbedEvents[Event2]:
			logger.Println("Recieved event2, val=", val)
		case val = <-pg.subbedEvents[Event3]:
			logger.Println("Recieved event3, val=", val)
		case <-ctx.Done():
			return
		}
	}
}
