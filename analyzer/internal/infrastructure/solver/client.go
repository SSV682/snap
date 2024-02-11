package solver

import (
	"context"
	"sync"

	"analyzer/internal/entity"

	log "github.com/sirupsen/logrus"
)

type client interface {
	MakeDecision(ctx context.Context, event entity.Event) error
}

// Client is the client API for the solver service.
type Client struct {
	inCh   chan entity.Event
	client client

	wg       sync.WaitGroup
	cancelFn context.CancelFunc
}

// Config is the configuration for the solver service.
type Config struct {
	Client client
	InCh   chan entity.Event
}

// NewSolverClient creates a new Solver
func NewSolverClient(cfg Config) *Client {
	return &Client{
		inCh:   cfg.InCh,
		client: cfg.Client,
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
				if err := c.client.MakeDecision(ctx, event); err != nil {
					log.Printf("Failed to make decision: %v", err)
				}
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
