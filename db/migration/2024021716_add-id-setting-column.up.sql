BEGIN;

ALTER TABLE analyzer_service.strategy_settings
ADD COLUMN id SERIAL PRIMARY KEY;

COMMIT;
