BEGIN;

CREATE SCHEMA IF NOT EXISTS solver_service;

CREATE TABLE IF NOT EXISTS solver_service.limitations
(
    ticker                      VARCHAR(64)                       NOT NULL,
    limit_up                    INTEGER,
    limit_down                  INTEGER
);

COMMIT;