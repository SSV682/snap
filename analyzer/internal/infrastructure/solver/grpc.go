package solver

import (
	"analyzer/internal/entity"
	"context"
	"fmt"
	"time"

	analyzerv1 "github.com/SSV682/snap/protos/gen/go/analyzer"
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
	client analyzerv1.SolverClient
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
		client: analyzerv1.NewSolverClient(cc),
	}, nil
}

// MakeDecision sends an event to the gRPC server and returns an error if one occurs.
func (c *GRCPClient) MakeDecision(ctx context.Context, event entity.Event) error {
	if _, err := c.client.MakeDecision(ctx, &analyzerv1.EventRequest{
		Ticker:    event.Ticker,
		EventType: string(event.Typ),
		Price:     float32(event.Price),
	}); err != nil {
		return err
	}

	return nil
}
