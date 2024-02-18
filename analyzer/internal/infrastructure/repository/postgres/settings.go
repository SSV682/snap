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
}

const (
	insertSettingsQuery = `INSERT INTO analyzer_service.strategy_settings (ticker, strategy, strategy_time_from, strategy_time_to, trading_time_from, trading_time_to)
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
)

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

func (r *SettingsRepository) Create(ctx context.Context, setting *entity.StrategySettings) (*entity.StrategySettings, error) {
	if err := r.db.QueryRowxContext(
		ctx,
		insertSettingsQuery,
		setting.Ticker,
		setting.Strategy,
		setting.StrategyTimeFrom,
		setting.StrategyTimeTo,
		setting.TradingTimeFrom,
	).Scan(&setting.ID); err != nil {
		return nil, fmt.Errorf("insert setting: %v", err)
	}

	return setting, nil
}

func (r *SettingsRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.db.ExecContext(ctx, `DELETE FROM analyzer_service.strategy_settings WHERE id = $1`, id); err != nil {
		return fmt.Errorf("delete setting: %v", err)
	}

	return nil
}
