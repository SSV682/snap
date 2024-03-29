package solver

import (
	"analyzer/internal/entity"
	"context"
	"fmt"
	"time"

	solverv1 "github.com/SSV682/snap/protos/gen/go/solver"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type GRCPConfig struct {
	Addr    string
	Timeout time.Duration
	Retries int
}

type GRCPClient struct {
	client solverv1.SolverClient
}

// NewGRCPClient TODO: set security
// NewGRCPClient creates a new GRCPClient
func NewGRCPClient(ctx context.Context, cfg *GRCPConfig) (*GRCPClient, error) {
	retryOpts := []retry.CallOption{
		retry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		retry.WithMax(uint(cfg.Retries)),
		retry.WithPerRetryTimeout(cfg.Timeout),
	}

	cc, err := grpc.DialContext(
		ctx,
		cfg.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		return nil, fmt.Errorf("can't connect to %s: %v", cfg.Addr, err)
	}

	return &GRCPClient{
		client: solverv1.NewSolverClient(cc),
	}, nil
}

// MakeDecision sends an event to the gRPC server and returns an error if one occurs.
func (c *GRCPClient) MakeDecision(ctx context.Context, event entity.Event) error {
	var eventType solverv1.EventType

	switch event.Typ {
	case entity.Buy:
		eventType = solverv1.EventType_EVENT_TYPE_BUY
	case entity.Sell:
		eventType = solverv1.EventType_EVENT_TYPE_SELL
	}

	if _, err := c.client.MakeDecision(ctx, &solverv1.EventRequest{
		Ticker:    event.Ticker,
		EventType: eventType,
		Price:     float32(event.Price),
	}); err != nil {
		return err
	}

	return nil
}
