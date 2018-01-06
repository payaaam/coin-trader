package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
)

type MarketStore struct {
	db *sql.DB
}

func NewMarketStore(db *sql.DB) *MarketStore {
	return &MarketStore{
		db: db,
	}
}

func (m *MarketStore) Save(ctx context.Context, market *models.Market) error {
	return market.Insert(m.db)
}

func (m *MarketStore) GetMarket(ctx context.Context, marketName string, marketKey string) (*models.Market, error) {
	market, err := models.Markets(m.db,
		qm.Where("market_name=?", marketName),
		qm.And("market_key=?", marketKey),
	).One()

	if err != nil {
		return nil, err
	}

	return market, nil
}
