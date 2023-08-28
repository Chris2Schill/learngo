package main

import (
    "fmt"
    "sync"
    "time"
    "playground/eventbus"
)


const (
    Event1 eventbus.Event = 0
    Event2 eventbus.Event = 1
    Event3 eventbus.Event = 2
    Event4 eventbus.Event = 3
)

type Playground struct {
    eventBus *eventbus.Dispatcher
    wg *sync.WaitGroup
    subbedEvents [](chan int)
}

func NewPlayground(eb *eventbus.Dispatcher) Playground {
    subbedEvents := make([]chan int, 0)

    subbedEvents = append(subbedEvents,
                          eb.Subscribe(Event1))
    subbedEvents = append(subbedEvents,
                          eb.Subscribe(Event2))
    subbedEvents = append(subbedEvents,
                          eb.Subscribe(Event3))

    pg := Playground{
        eb,
        &sync.WaitGroup{},
        subbedEvents,
    }
    return pg
}

func process1(pg *Playground) {
    var val int
    for {
        select {
        case val = <- pg.subbedEvents[0]:
            fmt.Println("Recieved event1, val=", val)
        case val = <- pg.subbedEvents[1]:
            fmt.Println("Recieved event2, val=", val)
        case val = <- pg.subbedEvents[2]:
            fmt.Println("Recieved event3, val=", val)
        }
        pg.wg.Done()
    }
}

func main() {
    eventBus := eventbus.NewDispatcher()

    pg := NewPlayground(&eventBus)

    time.Sleep(2)

    fmt.Println("Publishing events...")
    pg.wg.Add(3)
    eventBus.Publish(Event1, 3)
    eventBus.Publish(Event2, 5)
    eventBus.Publish(Event3, 69)

    go process1(&pg)

    pg.wg.Wait()
}
