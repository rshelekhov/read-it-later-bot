package event_consumer

import (
	"github.com/rshelekhov/read-it-later-bot/events"
	"log"
	"sync"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)

		if err != nil {
			log.Printf("Error fetching events: %s\n", err)

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err = c.handleEvents(gotEvents); err != nil {
			log.Printf("Error handling events: %s\n", err)

			continue
		}
	}
}

func (c *Consumer) handleEvents(eventsList []events.Event) error {
	var wg sync.WaitGroup

	for _, event := range eventsList {
		wg.Add(1)

		go func(e events.Event) {
			defer wg.Done()
			log.Printf("Got new event: %v\n", e)

			if err := c.processor.Process(e); err != nil {
				log.Printf("Error processing event: %s\n", err)

			}
		}(event)
	}

	wg.Wait()

	return nil
}
