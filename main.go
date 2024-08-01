package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Chris2Schill/learngo/eventbus"
	"github.com/Chris2Schill/learngo/logger"
	"github.com/Chris2Schill/learngo/playground"
	"github.com/Chris2Schill/learngo/writers"
	"github.com/Chris2Schill/learngo/chatroom"
)

func runWriterDecoratorTest() {
	boombabier := writers.BoomBabyWriter{
		os.Stdout,
	}

	logger.Println("derp69")

	someText := "derp "

	fmt.Fprintln(&boombabier, someText)
}

func runEventBusTest(ctx context.Context) {
	// Create an event bus
	eventBus := eventbus.NewDispatcher()

	// Create a playground which subscribes to some events
	pg := playground.New(eventBus)

	// Start the event dispatcher event-loop
	go eventBus.Dispatch()

	// Some time goes by
	time.Sleep(time.Second * 2)

	logger.Println("Publishing events...")
	eventBus.Publish(playground.Event1, 3)
	eventBus.Publish(playground.Event2, 5)
	eventBus.Publish(playground.Event3, 69)

	pg.Process(ctx)
}

func main() {
	logger.Default().SetOptions(logger.Timestamp | logger.CallerInfo)
	logger.Println("TESTING DEFAULT")

	defer logger.Flush()
	customLogger := logger.New(os.Stdout, logger.Timestamp)
	customLogger.Println("log without callerInfo")

	runWriterDecoratorTest()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go runEventBusTest(ctx)

	<-ctx.Done()

	logger.Println("Exiting...")
}
