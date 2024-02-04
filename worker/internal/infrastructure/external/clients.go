package external

import (
	"context"
	"fmt"
	"sync"

	"worker/internal/entity"
)

type ExternalClient struct {
	inCh chan entity.Event

	wg       sync.WaitGroup
	cancelFn context.CancelFunc
}

type Config struct {
	InCh chan entity.Event
}

func NewExternalClient(cfg Config) *ExternalClient {
	return &ExternalClient{
		inCh: cfg.InCh,
	}
}

func (c *ExternalClient) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancelFn = cancel

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		for {
			select {
			case event := <-c.inCh:
				//TODO: execute event

				fmt.Printf("event: %v", event)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (c *ExternalClient) Close() {
	c.cancelFn()

	c.wg.Wait()
}
