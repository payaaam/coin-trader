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

func (m *MarketStore) Upsert(ctx context.Context, market *models.Market) error {
	return market.Upsert(m.db, true, []string{"exchange_name", "market_key"}, []string{"market_key"})
}

func (m *MarketStore) Save(ctx context.Context, market *models.Market) error {
	return market.Insert(m.db)
}

func (m *MarketStore) GetMarkets(ctx context.Context, exchange string) ([]*models.Market, error) {
	return models.Markets(m.db,
		qm.Where("exchange_name=?", exchange),
	).All()
}

func (m *MarketStore) GetMarket(ctx context.Context, exchangeName string, marketKey string) (*models.Market, error) {
	market, err := models.Markets(m.db,
		qm.Where("exchange_name=?", exchangeName),
		qm.And("market_key=?", marketKey),
	).One()

	if err != nil {
		return nil, err
	}

	return market, nil
}
