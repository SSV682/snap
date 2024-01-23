package postgres

import (
	"context"
	"log"

	"snap/worker/internal/entity"

	"github.com/jmoiron/sqlx"
)

type InvestRepository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewActionRepository(db *sqlx.DB) *InvestRepository {
	return &InvestRepository{db: db}
}

func (r *InvestRepository) List(ctx context.Context) ([]entity.Candle, error) {
	panic("implement me")

	return nil, nil
}

func (r *InvestRepository) Create(ctx context.Context, candle *entity.Candle) error {
	//tx, err := r.db.BeginTxx(ctx, nil)
	//if err != nil {
	//	return fmt.Errorf("begin transaction: %v", err)
	//}
	//
	//defer func() {
	//	if txErr := tx.Rollback(); txErr != nil && !errors.Is(txErr, sql.ErrTxDone) {
	//		//r.log.GetLoggerFromContext(ctx).Errorf("Failed rollback transaction: %v", txErr)
	//	}
	//}()
	//
	//candleRow := NewCandleRow(candle)
	panic("implement me")

	return nil
}
