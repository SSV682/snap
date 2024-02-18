package analyzer

import (
	"context"
	"fmt"
	"time"

	"github.com/SSV682/snap/protos/gen/go/analyzer"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRCPConfig struct {
	Addr    string
	Timeout time.Duration
	Retries int
}

type GRCPClient struct {
	client analyzer_v1.AnalyzerClient
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
		client: analyzer_v1.NewAnalyzerClient(cc),
	}, nil
}

func (c *GRCPClient) ListActualSettings(ctx context.Context) (*analyzer_v1.ActualSettingsResponse, error) {
	return c.client.ListActualSettings(ctx, &emptypb.Empty{})
}

func (c *GRCPClient) CreateSetting(ctx context.Context, req *analyzer_v1.CreateSettingRequest) (*analyzer_v1.CreateSettingResponse, error) {
	return c.client.CreateSetting(ctx, req)
}

func (c *GRCPClient) DeleteSetting(ctx context.Context, req *analyzer_v1.DeleteSettingRequest) (*emptypb.Empty, error) {
	return c.client.DeleteSetting(ctx, req)
}
