package app

import (
	"context"
	"io"
	"os"
	"os/signal"
	"solver/internal/service"
	"syscall"
	"time"

	grpcapp "solver/internal/app/grpc"
	"solver/internal/config"

	log "github.com/sirupsen/logrus"
)

type Runner interface {
	Run() error
}

type App struct {
	cfg     *config.Config
	runners []Runner
	closers []io.Closer
}

// NewApp creates a new instance of the application
func NewApp(cfg *config.Config) *App {
	var runners []Runner
	var closers []io.Closer

	solver := service.Solver{}

	// create a new gRPC server
	grpcServer := grpcapp.NewApp(&grpcapp.Config{
		Manager: &solver,
		Port:    cfg.GRPC.Port,
		Timeout: cfg.GRPC.Timeout,
	})

	// add the gRPC server to the list of runners
	runners = append(runners, grpcServer)
	closers = append(closers, grpcServer)

	// return a new instance of the application with the gRPC server
	return &App{
		runners: runners,
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	for _, runner := range a.runners {
		err := runner.Run()
		if err != nil {
			log.Errorf("Failed to run the runner: %v", err)
		}
	}

	log.Info("App has been started")
	a.waitGracefulShutdown(ctx, cancel)
}

func (a *App) waitGracefulShutdown(_ context.Context, cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt,
	)

	log.Infof("Caught signal %s. Shutting down...", <-quit)

	//done := make(chan struct{})
	//go func() {
	//	// try to close http server connections
	//	if err := a.httpServer.Shutdown(); err != nil {
	//		log.Errorf("Failed to shutdown http server: %v", err)
	//	}
	//
	//	close(done)
	//}()

	select {
	case <-time.After(a.cfg.GracefulTimeout):
		//case <-done:
	}

	// try to close background processes
	for _, c := range a.closers {
		if err := c.Close(); err != nil {
			log.Errorf("Failed to close: %v", err)
		}
	}
}
