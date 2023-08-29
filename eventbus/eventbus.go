package eventbus

import (
	"fmt"
)

type Event int64

type EventPayload struct {
	eventType Event
	data      int
}

type Dispatcher struct {
	bus        map[Event][]chan int
	eventQueue chan EventPayload
}

func NewDispatcher() Dispatcher {
	eb := Dispatcher{
		make(map[Event][]chan int),
		make(chan EventPayload, 100),
	}

	go dispatch(&eb)

	return eb
}

func (eb *Dispatcher) Subscribe(e Event) chan int {

	newChan := make(chan int)
	eb.bus[e] = append(eb.bus[e], newChan)

	fmt.Println("subscribed to ", e)

	return newChan
}

func (eb *Dispatcher) Publish(e Event, data int) {
	payload := EventPayload{
		e,
		data,
	}

	eb.eventQueue <- payload
}

func dispatch(d *Dispatcher) {
	var payload EventPayload
	for {
		payload = <-d.eventQueue
		for _, channel := range d.bus[payload.eventType] {
			channel <- payload.data
		}
	}
}
