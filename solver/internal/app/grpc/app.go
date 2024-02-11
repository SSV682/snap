package grpcapp

import (
	"fmt"
	"net"
	solvergrpc "solver/internal/grpc"
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type App struct {
	gRPCServer *grpc.Server
	port       int
}

// NewApp creates a new instance of the application.
func NewApp(cfg *Config) *App {
	// grpcServer is the gRPC server instance.
	grpcServer := grpc.NewServer()

	// RegisterServerAPI registers the gRPC server API.
	solvergrpc.RegisterServerAPI(grpcServer, NewManager(cfg)) //TODO: new manager

	// app is the application instance.
	app := &App{
		port:       cfg.Port,
		gRPCServer: grpcServer,
	}

	return app
}

// Run starts the gRPC server.
func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}

	return nil
}

// Close gracefully shuts down the gRPC server.
func (a *App) Close() error {
	a.gRPCServer.GracefulStop()

	return nil
}
