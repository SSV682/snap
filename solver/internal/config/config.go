package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
)

//type ConnConfig struct {
//	Network  string   `yaml:"network"`
//	Database string   `yaml:"database"`
//	Hosts    []string `yaml:"hosts"`
//	Ports    []string `yaml:"ports"`
//	Username string   `yaml:"username"`
//	Password string   `yaml:"password"`
//}
//
//type SQLConfig struct {
//	ConnConfig      `yaml:"conn_config"`
//	MaxOpenConns    int           `yaml:"max_open_conns"`
//	MaxIdleConns    int           `yaml:"max_idle_conns"`
//	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
//	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
//}

//type DatabaseConfig struct {
//	Postgres SQLConfig `yaml:"postgres"`
//}

// HTTPServerConfig represents the configuration of the HTTP server.
type HTTPServerConfig struct {
	Listen       string        `yaml:"listen"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Config represents	the configuration of the application.
type Config struct {
	Invest          investgo.Config  `yaml:"invest"`
	HTTPServer      HTTPServerConfig `yaml:"httpserver"`
	GRPC            GRPCConfig       `yaml:"grpc"`
	GracefulTimeout time.Duration    `yaml:"graceful_timeout"`
}

// ReadConfig reads the configuration from the given file path and
// unmarshal it into the given struct.
func ReadConfig(filePath string) (Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
