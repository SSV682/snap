package app

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
	"worker/internal/config"
	"worker/internal/entity"
	"worker/internal/handlers"
	v1 "worker/internal/handlers/v1"
	"worker/internal/infrastructure/broker"
	"worker/internal/infrastructure/external"
	"worker/internal/infrastructure/repository/postgres"
	"worker/internal/service"
	"worker/internal/service/backtest"
	"worker/internal/service/manager"

	"github.com/go-playground/validator/v10"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Runner interface {
	Run()
}

type App struct {
	cfg        *config.Config
	httpServer *fasthttp.Server
	runners    []Runner
	closers    []io.Closer
}

func NewApp(configPath string) *App {
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	router := routing.New()

	brokerClient, err := broker.NewClient(
		context.Background(),
		investgo.Config{
			EndPoint:                      cfg.Invest.EndPoint,
			Token:                         cfg.Invest.Token,
			AppName:                       cfg.Invest.AppName,
			AccountId:                     cfg.Invest.AccountId,
			DisableResourceExhaustedRetry: cfg.Invest.DisableResourceExhaustedRetry,
			DisableAllRetry:               cfg.Invest.DisableAllRetry,
			MaxRetries:                    cfg.Invest.MaxRetries,
		},
		nil,
	)
	if err != nil {
		return nil
	}

	db, err := initDB(cfg.Databases.Postgres)
	if err != nil {
		log.Fatalf("Failed to init db pool of connections: %v", err)
	}

	settingsRepo := postgres.NewSettingsRepository(db)

	managerService := manager.NewManager(manager.Config{
		SettingsRepository: settingsRepo,
		BrokerProvider:     brokerClient,
		InfoProvider:       brokerClient,
	})

	signalCh := make(chan entity.Event, 1)

	backTestService := backtest.NewBackTestService(
		&backtest.Config{
			TradingInfoProvider: brokerClient,
			BrokerProvider:      brokerClient,
		},
	)

	tradingService := service.NewTradingService(
		&service.TradingConfig{
			TradingInfoProvider: brokerClient,
		},
	)

	var runners []Runner
	var closers []io.Closer

	ec := external.NewExternalClient(external.Config{InCh: signalCh})
	runners = append(runners, ec)
	closers = append(closers, ec)

	runners = append(runners, managerService)
	closers = append(closers, managerService)

	handlers.Register(
		router,
		v1.NewInvestHandler(v1.Config{
			BackTestService: backTestService,
			TradingService:  tradingService,
			Validator:       validator.New(),
		}),
	)

	log.Info("App created")

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
		runner.Run()
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
