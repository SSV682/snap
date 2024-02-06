BEGIN;

CREATE SCHEMA IF NOT EXISTS analyzer_service;

CREATE TABLE IF NOT EXISTS analyzer_service.strategy_settings
(
    ticker                      VARCHAR(64)                       NOT NULL,
    strategy                    VARCHAR(64)                       NOT NULL,
    strategy_time_from          TIMESTAMP WITH TIME ZONE          NOT NULL DEFAULT NOW(),
    strategy_time_to            TIMESTAMP WITH TIME ZONE          NOT NULL,
    trading_time_from           TIME                              NOT NULL,
    trading_time_to             TIME                              NOT NULL,
    UNIQUE(ticker, strategy)
);

COMMIT;
