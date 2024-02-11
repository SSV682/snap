package external

import (
	"context"
	"fmt"
	"sync"

	"worker/internal/entity"
)

type Client struct {
	inCh chan entity.Event

	wg       sync.WaitGroup
	cancelFn context.CancelFunc
}

type Config struct {
	InCh chan entity.Event
}

func NewExternalClient(cfg Config) *Client {
	return &Client{
		inCh: cfg.InCh,
	}
}

func (c *Client) Run() {
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

func (c *Client) Close() error {
	c.cancelFn()

	c.wg.Wait()

	return nil
}
