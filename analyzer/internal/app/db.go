package app

import (
	"fmt"
	"strings"

	"worker/internal/config"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func initDB(cfg config.SQLConfig) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable target_session_attrs=read-write statement_cache_mode=describe",
		strings.Join(cfg.Hosts, ","),
		cfg.Username,
		cfg.Password,
		cfg.Database,
		strings.Join(cfg.Ports, ","),
	)

	db, err := sqlx.Open("pgx", dataSource)
	if err != nil {
		return nil, fmt.Errorf("create pool of connections to database: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return db, nil
}
