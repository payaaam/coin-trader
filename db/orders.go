package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
	"golang.org/x/net/context"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) OrderStoreInterface {
	return &OrderStore{
		db: db,
	}
}

func (o *OrderStore) Save(ctx context.Context, order *models.Order) error {
	err := order.Insert(o.db)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderStore) Upsert(ctx context.Context, order *models.Order) error {
	return order.Upsert(o.db, true, []string{"exchange_order_id"}, []string{"status", "sell_price", "close_time"})
}
