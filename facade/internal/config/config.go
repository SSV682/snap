package config

import (
	"time"

	"facade/internal/bot/telegram"

	"github.com/ilyakaznacheev/cleanenv"
)

type BotsConfig struct {
	Telegram telegram.Config `yaml:"telegram"`
}

type GRPCClient struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
	Retries int           `yaml:"retries"`
}

type ClientsConfig struct {
	Solver   GRPCClient `yaml:"solver"`
	Analyzer GRPCClient `yaml:"analyzer"`
}

type HTTPServerConfig struct {
	Listen       string        `yaml:"listen"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Config struct {
	HTTPServer      HTTPServerConfig `yaml:"httpserver"`
	GracefulTimeout time.Duration    `yaml:"graceful_timeout"`
	Clients         ClientsConfig    `yaml:"clients"`
}

func ReadConfig(filePath string) (Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
