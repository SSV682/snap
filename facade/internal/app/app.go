package app

import (
	"context"
	"facade/internal/handlers"
	v1 "facade/internal/handlers/v1"
	"facade/internal/infrastructure/analyzer"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"facade/internal/config"

	"github.com/go-playground/validator/v10"
	routing "github.com/qiangxue/fasthttp-routing"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Runner interface {
	Run() error
}

type App struct {
	cfg        *config.Config
	httpServer *fasthttp.Server
	runners    []Runner
	closers    []io.Closer
}

func NewApp(cfg config.Config) *App {
	router := routing.New()
	var runners []Runner
	var closers []io.Closer

	analyzerClient, err := analyzer.NewGRCPClient(context.Background(), &analyzer.GRCPConfig{
		Addr:    cfg.Clients.Solver.Address,
		Retries: cfg.Clients.Solver.Retries,
		Timeout: cfg.Clients.Solver.Timeout,
	})
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	log.Info("App created")

	handlers.Register(
		router,
		v1.NewHandlers(v1.Config{
			AnalyzerClient: analyzerClient,

			Validator: validator.New(),
		}),
	)

	return &App{
		cfg: &cfg,
		httpServer: &fasthttp.Server{
			Handler:      router.HandleRequest,
			ReadTimeout:  cfg.HTTPServer.ReadTimeout,
			WriteTimeout: cfg.HTTPServer.WriteTimeout,
		},
		runners: runners,
		closers: closers,
	}
}

func (a *App) Run() {
	for _, runner := range a.runners {
		if err := runner.Run(); err != nil {
			log.Fatalf("Failed to run runner: %v", err)
		}
	}

	go func() {
		if err := a.httpServer.ListenAndServe(a.cfg.HTTPServer.Listen); err != nil {
			log.Fatalf("Failed listen and serve http server: %v", err)
		}
	}()

	log.Info("App has been started")

	a.waitGracefulShutdown()
}

// waitGracefulShutdown waits for a graceful shutdown signal and then shuts down the application.
// It attempts to close the HTTP server connections and gracefully close any background processes.
func (a *App) waitGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGABRT, // syscall.SIGABRT: abort signal from the operating system
		syscall.SIGQUIT, // syscall.SIGQUIT: terminal interrupt signal
		syscall.SIGHUP,  // syscall.SIGHUP: terminal hangup signal
		syscall.SIGTERM, // syscall.SIGTERM: termination signal
		os.Interrupt,    // os.Interrupt: interrupt signal sent from the terminal
	)

	log.Infof("Caught signal %s. Shutting down...", <-quit)

	done := make(chan struct{})
	go func() {
		// try to close http server connections
		if err := a.httpServer.Shutdown(); err != nil {
			log.Errorf("Failed to shutdown http server: %v", err)
		}

		close(done)
	}()

	select {
	case <-time.After(a.cfg.GracefulTimeout):
	case <-done:
	}

	// try to close background processes
	for _, c := range a.closers {
		if err := c.Close(); err != nil {
			log.Errorf("Failed to close: %v", err)
		}
	}
}
