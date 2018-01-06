package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
	"golang.org/x/net/context"
)

type TickStore struct {
	db *sql.DB
}

func NewTickStore(db *sql.DB) *TickStore {
	return &TickStore{
		db: db,
	}
}

func (t *TickStore) Upsert(ctx context.Context, tick *models.Tick) error {
	return tick.Upsert(t.db, true, []string{"chart_id", "timestamp"}, []string{"open", "close", "high", "low"})
}

func (t *TickStore) Save(ctx context.Context, tick *models.Tick) error {
	return tick.Insert(t.db)
}
