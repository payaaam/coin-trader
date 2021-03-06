// This file is generated by SQLBoiler (https://github.com/volatiletech/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Charts", testCharts)
	t.Run("Exchanges", testExchanges)
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("Markets", testMarkets)
	t.Run("Orders", testOrders)
	t.Run("Ticks", testTicks)
}

func TestDelete(t *testing.T) {
	t.Run("Charts", testChartsDelete)
	t.Run("Exchanges", testExchangesDelete)
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("Markets", testMarketsDelete)
	t.Run("Orders", testOrdersDelete)
	t.Run("Ticks", testTicksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Charts", testChartsQueryDeleteAll)
	t.Run("Exchanges", testExchangesQueryDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("Markets", testMarketsQueryDeleteAll)
	t.Run("Orders", testOrdersQueryDeleteAll)
	t.Run("Ticks", testTicksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Charts", testChartsSliceDeleteAll)
	t.Run("Exchanges", testExchangesSliceDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("Markets", testMarketsSliceDeleteAll)
	t.Run("Orders", testOrdersSliceDeleteAll)
	t.Run("Ticks", testTicksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Charts", testChartsExists)
	t.Run("Exchanges", testExchangesExists)
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("Markets", testMarketsExists)
	t.Run("Orders", testOrdersExists)
	t.Run("Ticks", testTicksExists)
}

func TestFind(t *testing.T) {
	t.Run("Charts", testChartsFind)
	t.Run("Exchanges", testExchangesFind)
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("Markets", testMarketsFind)
	t.Run("Orders", testOrdersFind)
	t.Run("Ticks", testTicksFind)
}

func TestBind(t *testing.T) {
	t.Run("Charts", testChartsBind)
	t.Run("Exchanges", testExchangesBind)
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("Markets", testMarketsBind)
	t.Run("Orders", testOrdersBind)
	t.Run("Ticks", testTicksBind)
}

func TestOne(t *testing.T) {
	t.Run("Charts", testChartsOne)
	t.Run("Exchanges", testExchangesOne)
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("Markets", testMarketsOne)
	t.Run("Orders", testOrdersOne)
	t.Run("Ticks", testTicksOne)
}

func TestAll(t *testing.T) {
	t.Run("Charts", testChartsAll)
	t.Run("Exchanges", testExchangesAll)
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("Markets", testMarketsAll)
	t.Run("Orders", testOrdersAll)
	t.Run("Ticks", testTicksAll)
}

func TestCount(t *testing.T) {
	t.Run("Charts", testChartsCount)
	t.Run("Exchanges", testExchangesCount)
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("Markets", testMarketsCount)
	t.Run("Orders", testOrdersCount)
	t.Run("Ticks", testTicksCount)
}

func TestHooks(t *testing.T) {
	t.Run("Charts", testChartsHooks)
	t.Run("Exchanges", testExchangesHooks)
	t.Run("GorpMigrations", testGorpMigrationsHooks)
	t.Run("Markets", testMarketsHooks)
	t.Run("Orders", testOrdersHooks)
	t.Run("Ticks", testTicksHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Charts", testChartsInsert)
	t.Run("Charts", testChartsInsertWhitelist)
	t.Run("Exchanges", testExchangesInsert)
	t.Run("Exchanges", testExchangesInsertWhitelist)
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("Markets", testMarketsInsert)
	t.Run("Markets", testMarketsInsertWhitelist)
	t.Run("Orders", testOrdersInsert)
	t.Run("Orders", testOrdersInsertWhitelist)
	t.Run("Ticks", testTicksInsert)
	t.Run("Ticks", testTicksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChartToMarketUsingMarket", testChartToOneMarketUsingMarket)
	t.Run("OrderToMarketUsingMarket", testOrderToOneMarketUsingMarket)
	t.Run("TickToChartUsingChart", testTickToOneChartUsingChart)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ChartToTicks", testChartToManyTicks)
	t.Run("MarketToCharts", testMarketToManyCharts)
	t.Run("MarketToOrders", testMarketToManyOrders)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChartToMarketUsingMarket", testChartToOneSetOpMarketUsingMarket)
	t.Run("OrderToMarketUsingMarket", testOrderToOneSetOpMarketUsingMarket)
	t.Run("TickToChartUsingChart", testTickToOneSetOpChartUsingChart)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ChartToTicks", testChartToManyAddOpTicks)
	t.Run("MarketToCharts", testMarketToManyAddOpCharts)
	t.Run("MarketToOrders", testMarketToManyAddOpOrders)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Charts", testChartsReload)
	t.Run("Exchanges", testExchangesReload)
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("Markets", testMarketsReload)
	t.Run("Orders", testOrdersReload)
	t.Run("Ticks", testTicksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Charts", testChartsReloadAll)
	t.Run("Exchanges", testExchangesReloadAll)
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("Markets", testMarketsReloadAll)
	t.Run("Orders", testOrdersReloadAll)
	t.Run("Ticks", testTicksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Charts", testChartsSelect)
	t.Run("Exchanges", testExchangesSelect)
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("Markets", testMarketsSelect)
	t.Run("Orders", testOrdersSelect)
	t.Run("Ticks", testTicksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Charts", testChartsUpdate)
	t.Run("Exchanges", testExchangesUpdate)
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("Markets", testMarketsUpdate)
	t.Run("Orders", testOrdersUpdate)
	t.Run("Ticks", testTicksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Charts", testChartsSliceUpdateAll)
	t.Run("Exchanges", testExchangesSliceUpdateAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("Markets", testMarketsSliceUpdateAll)
	t.Run("Orders", testOrdersSliceUpdateAll)
	t.Run("Ticks", testTicksSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Charts", testChartsUpsert)
	t.Run("Exchanges", testExchangesUpsert)
	t.Run("GorpMigrations", testGorpMigrationsUpsert)
	t.Run("Markets", testMarketsUpsert)
	t.Run("Orders", testOrdersUpsert)
	t.Run("Ticks", testTicksUpsert)
}
