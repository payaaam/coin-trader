package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

func (o *OrderStore) Save(order *models.Order) error {

	err := order.Insert(o.db)
	if err != nil {
		return err
	}

	return nil
}
