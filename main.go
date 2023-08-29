package main

import (
	"fmt"
	"time"

	"github.com/Chris2Schill/learngo/eventbus"
	"github.com/Chris2Schill/learngo/playground"
)

func main() {
	eventBus := eventbus.NewDispatcher()
	pg := playground.New(&eventBus)

	time.Sleep(time.Second * 2)

	fmt.Println("Publishing events...")
	pg.Add(3)
	eventBus.Publish(playground.Event1, 3)
	eventBus.Publish(playground.Event2, 5)
	eventBus.Publish(playground.Event3, 69)

	go pg.Process()

	pg.Wait()
}
