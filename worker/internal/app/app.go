package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"worker/internal/config"
	"worker/internal/handlers"
	v1 "worker/internal/handlers/v1"
	"worker/internal/infrastructure/tinkoff"
	"worker/internal/service"

	validator "github.com/go-playground/validator/v10"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type RunAsService interface {
	Run()
}

type App struct {
	cfg        *config.Config
	httpServer *fasthttp.Server
	runners    []RunAsService
}

func NewApp(configPath string) *App {
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	router := routing.New()

	client, err := tinkoff.NewClient(
		context.Background(),
		investgo.Config{
			EndPoint:                      cfg.InvestConfig.EndPoint,
			Token:                         cfg.InvestConfig.Token,
			AppName:                       cfg.InvestConfig.AppName,
			AccountId:                     cfg.InvestConfig.AccountId,
			DisableResourceExhaustedRetry: cfg.InvestConfig.DisableResourceExhaustedRetry,
			DisableAllRetry:               cfg.InvestConfig.DisableAllRetry,
			MaxRetries:                    cfg.InvestConfig.MaxRetries,
		},
		nil,
	)
	if err != nil {
		return nil
	}

	signalCh := make(chan service.Event)

	//tradingService := service.NewTradingService(
	//	&service.Config{
	//		TradingInfoProvider: client,
	//		SignalCh:            signalCh,
	//	},
	//)

	backTestService := service.NewBackTestService(
		&service.BackTestConfig{
			TradingInfoProvider: client,
			BrokerProvider:      client,
		},
	)

	var runners []RunAsService

	manager := service.NewManager(signalCh)
	runners = append(runners, manager)

	handlers.Register(
		router,
		v1.NewInvestHandler(backTestService, validator.New()),
	)

	log.Info("App created")

	return &App{
		cfg: &cfg,
		httpServer: &fasthttp.Server{
			Handler:      router.HandleRequest,
			ReadTimeout:  cfg.HTTPServerConfig.ReadTimeout,
			WriteTimeout: cfg.HTTPServerConfig.WriteTimeout,
		},
		runners: runners,
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	for _, runner := range a.runners {
		runner.Run()
	}

	go func() {
		if err := a.httpServer.ListenAndServe(a.cfg.HTTPServerConfig.Listen); err != nil {
			log.Fatalf("Failed listen and serve http server: %v", err)
		}
	}()

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

	cancel()

	if err := a.httpServer.Shutdown(); err != nil {
		log.Errorf("Failed to shutdown http server: %v", err)
	}
}
