package eventconsumer

import (
	"log"
	"time"

	"github.com/KingDaveII/darita-food-bot/events"
)

type Consumer struct {
	fetcher events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher: fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer : %s", err.Error())
			// TODO: add retry with backoff
			continue
		}
		
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}

	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("[ERR] consumer : can't process event: %s", err.Error())
			// TODO: add retry with backoff and saving failed events to DB
			// maybe we can use dead letter queue for this purpose (fallback queue)
			// somehow sync with Fetcher so we don't lose events and don't process them twice
			// maybe add parallel processing of events with worker pool and channels (waitGroup)
			continue
		}


	}

	return nil
}