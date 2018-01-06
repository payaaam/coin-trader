package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
	"golang.org/x/net/context"
)

type ChartStore struct {
	db *sql.DB
}

func (c *ChartStore) Save(ctx context.Context, chart *models.Chart) error {
	return chart.Insert(c.db)
}
