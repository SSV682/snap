package telegram

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type command string

const (
	startCommand  command = "start"
	helpCommand   command = "help"
	statusCommand command = "status"
)

func toCommand(s string) (command, error) {
	switch s {
	case "start", "help", "status":
		return command(s), nil
	default:
		return "", fmt.Errorf("unknown command: %s", s)
	}
}

type ViewFunc func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error

type Bot struct {
	api          *tgbotapi.BotAPI
	commandViews map[command]ViewFunc

	cancel context.CancelFunc
	wg     sync.WaitGroup
}

type Config struct {
	Token string
}

func NewBot(c *Config) *Bot {
	api, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		return nil
	}

	return &Bot{
		api: api,
		commandViews: map[command]ViewFunc{
			statusCommand: viewCmdStatus(),
		},
	}
}

func (b *Bot) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	b.cancel = cancel

	b.wg.Add(1)
	go func() {
		defer b.wg.Done()

		b.run(ctx)
	}()
}

func (b *Bot) run(ctx context.Context) {
	updatesCh := b.api.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	})

	for {
		select {
		case update := <-updatesCh:
			updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Second)
			b.handleUpdate(updateCtx, update)
			updateCancel()
		case <-ctx.Done():
			return
		}
	}
}

func (b *Bot) Close() error {
	b.cancel()

	b.wg.Wait()

	return nil
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			// TODO: log error
			log.Printf("Recovered from panic: %v\n%s", err, string(debug.Stack()))
		}
	}()

	var view ViewFunc

	if update.Message == nil && !update.Message.IsCommand() {
		return
	}

	cmd, err := toCommand(update.Message.Command())
	if err != nil {
		log.Errorf("failed to parse command: %v", err)

		return
	}

	view, ok := b.commandViews[cmd]
	if !ok {
		return
	}

	if err = view(ctx, b.api, update); err != nil {
		log.Errorf("error handling update: %v", err)

		errMessage := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Error: %s", err.Error()))

		if _, err = b.api.Send(errMessage); err != nil {
			log.Errorf("error sending message: %v", err)
		}
	}

}

func viewCmdStatus() ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Status:OK"))); err != nil {
			return fmt.Errorf("error sending message to bot: %v", err)
		}

		return nil
	}
}
