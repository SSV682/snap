package postgres

import (
	sq "github.com/Masterminds/squirrel"
)

func newQueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type settingsSelectBuilder struct {
	b sq.SelectBuilder
}

func newSettingsSelectBuilder() *settingsSelectBuilder {
	selectBuilder := newQueryBuilder().
		Select("ss.id",
			"ss.ticker",
			"ss.strategy",
			"ss.strategy_time_from",
			"ss.strategy_time_to",
			"ss.trading_time_from",
			"ss.trading_time_to",
		).
		From("analyzer_service.strategy_settings ss")

	return &settingsSelectBuilder{b: selectBuilder}
}

func (a *settingsSelectBuilder) ToSql() (query string, args []any, err error) {
	return a.b.ToSql()
}
