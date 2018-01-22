package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
	"golang.org/x/net/context"
)

type ChartStore struct {
	db *sql.DB
}

func NewChartStore(db *sql.DB) ChartStoreInterface {
	return &ChartStore{
		db: db,
	}
}

func (c *ChartStore) Upsert(ctx context.Context, chart *models.Chart) error {
	return chart.Upsert(c.db, true, []string{"market_id", "interval"}, []string{"market_id"})
}

func (c *ChartStore) Save(ctx context.Context, chart *models.Chart) error {
	return chart.Insert(c.db)
}
