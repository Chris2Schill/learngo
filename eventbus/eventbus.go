package eventbus

import (
	"sync"

	"github.com/Chris2Schill/learngo/logger"
)

type Event int64

type Dispatcher interface {
	Subscribe(e Event) chan int
	Publish(e Event, data int)
	Dispatch()
}

type EventPayload struct {
	eventType Event
	data      int
}

type defaultDispatcher struct {
	busLock    sync.Mutex
	bus        map[Event][]chan int
	eventQueue chan EventPayload
}

func NewDispatcher() Dispatcher {
	d := defaultDispatcher{
		sync.Mutex{},
		make(map[Event][]chan int),
		make(chan EventPayload, 100),
	}

	return d
}

func (d defaultDispatcher) Subscribe(e Event) chan int {
	d.busLock.Lock()
	defer d.busLock.Unlock()

	newChan := make(chan int)
	d.bus[e] = append(d.bus[e], newChan)

	logger.Println("subscribed to ", e)

	return newChan
}

func (d defaultDispatcher) Publish(e Event, data int) {
	payload := EventPayload{
		e,
		data,
	}

	d.eventQueue <- payload
}

func (d defaultDispatcher) Dispatch() {
	var payload EventPayload
	for {
		payload = <-d.eventQueue
		for _, channel := range d.bus[payload.eventType] {
			channel <- payload.data
		}
	}
}
