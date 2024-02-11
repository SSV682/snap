package postgres

import (
	//"log"

	"analyzer/internal/entity"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SettingsRepository struct {
	db *sqlx.DB
	//log log.Logger
}

func NewSettingsRepository(db *sqlx.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (r *SettingsRepository) List(ctx context.Context) ([]*entity.StrategySettings, error) {
	builder := newSettingsSelectBuilder()

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build select query: %v", err)
	}

	return r.list(ctx, query, args...)
}

func (r *SettingsRepository) list(ctx context.Context, query string, args ...any) ([]*entity.StrategySettings, error) {
	var settingsRows []strategySettingsRow

	if err := r.db.SelectContext(ctx, &settingsRows, query, args...); err != nil {
		return nil, fmt.Errorf("select deployment action list: %v", err)
	}

	settings := make([]*entity.StrategySettings, len(settingsRows))
	for i := range settings {
		settings[i] = settingsRows[i].ToModel()
	}

	return settings, nil
}
