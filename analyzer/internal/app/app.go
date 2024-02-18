package app

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcapp "analyzer/internal/app/grpc"
	"analyzer/internal/config"
	"analyzer/internal/infrastructure/broker"
	"analyzer/internal/infrastructure/repository/postgres"
	"analyzer/internal/infrastructure/solver"
	"analyzer/internal/service/manager"

	"github.com/russianinvestments/invest-api-go-sdk/investgo"
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

func NewApp(cfg config.Config) *App {
	var runners []Runner
	var closers []io.Closer

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
		log.Fatalf("Failed to: %v", err)

		return nil
	}

	db, err := initDB(cfg.Databases.Postgres)
	if err != nil {
		log.Fatalf("Failed to init db pool of connections: %v", err)
	}

	settingsRepo := postgres.NewSettingsRepository(db)

	solverClient, err := solver.NewGRCPClient(context.Background(), &solver.GRCPConfig{
		Addr:    cfg.Clients.Solver.Address,
		Retries: cfg.Clients.Solver.Retries,
		Timeout: cfg.Clients.Solver.Timeout,
	})
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	managerService := manager.NewManager(manager.Config{
		SettingsRepository: settingsRepo,
		BrokerProvider:     brokerClient,
		InfoProvider:       brokerClient,
		SolverClient:       solverClient,
	})
	runners = append(runners, managerService)
	closers = append(closers, managerService)

	// TODO: return BackTestService
	//backTestService := backtest.NewBackTestService(
	//	&backtest.Config{
	//		TradingInfoProvider: brokerClient,
	//		BrokerProvider:      brokerClient,
	//	},
	//)

	grpcServer := grpcapp.NewApp(&grpcapp.Config{
		Port:           cfg.GRPC.Port,
		Timeout:        cfg.GRPC.Timeout,
		ManagerService: managerService,
	})

	// add the gRPC server to the list of runners
	runners = append(runners, grpcServer)
	closers = append(closers, grpcServer)

	//TODO: return tradingService
	//tradingService := service.NewTradingService(
	//	&service.TradingConfig{
	//		TradingInfoProvider: brokerClient,
	//	},
	//)

	//handlers.Register(
	//	router,
	//	v1.NewInvestHandler(v1.Config{
	//		BackTestService: backTestService,
	//		TradingService:  tradingService,
	//		Validator:       validator.New(),
	//	}),
	//)

	log.Info("App created")

	return &App{
		cfg:     &cfg,
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
		// try to close background workers
		for _, closer := range a.closers {
			if err := closer.Close(); err != nil {
				log.Errorf("Failed to close closer: %v", err)
			}
		}

		// wait for all background workers to finish
		<-done
	}()

	select {
	case <-time.After(a.cfg.GracefulTimeout):
	case <-done:
	}

	// try to close background processes

}
