package solver

import (
	"context"
	"fmt"
	"sync"

	"analyzer/internal/entity"
)

// Client is the client API for the solver service.
type Client struct {
	inCh chan entity.Event

	wg       sync.WaitGroup
	cancelFn context.CancelFunc
}

// Config is the configuration for the solver service.
type Config struct {
	InCh chan entity.Event
}

func NewSolverClient(cfg Config) *Client {
	return &Client{
		inCh: cfg.InCh,
	}
}

// Run starts the solver service.
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

// Close closes the solver service.
func (c *Client) Close() error {
	c.cancelFn()

	c.wg.Wait()

	return nil
}
