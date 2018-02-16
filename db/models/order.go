// This file is generated by SQLBoiler (https://github.com/volatiletech/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// Order is an object representing the database table.
type Order struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	MarketID        int         `boil:"market_id" json:"market_id" toml:"market_id" yaml:"market_id"`
	Type            string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	ExchangeOrderID string      `boil:"exchange_order_id" json:"exchange_order_id" toml:"exchange_order_id" yaml:"exchange_order_id"`
	Limit           string      `boil:"limit" json:"limit" toml:"limit" yaml:"limit"`
	Quantity        string      `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	Status          string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	SellPrice       null.String `boil:"sell_price" json:"sell_price,omitempty" toml:"sell_price" yaml:"sell_price,omitempty"`
	OpenTime        int64       `boil:"open_time" json:"open_time" toml:"open_time" yaml:"open_time"`
	CloseTime       null.Int64  `boil:"close_time" json:"close_time,omitempty" toml:"close_time" yaml:"close_time,omitempty"`
	QuantityFilled  null.String `boil:"quantity_filled" json:"quantity_filled,omitempty" toml:"quantity_filled" yaml:"quantity_filled,omitempty"`

	R *orderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderColumns = struct {
	ID              string
	MarketID        string
	Type            string
	ExchangeOrderID string
	Limit           string
	Quantity        string
	Status          string
	SellPrice       string
	OpenTime        string
	CloseTime       string
	QuantityFilled  string
}{
	ID:              "id",
	MarketID:        "market_id",
	Type:            "type",
	ExchangeOrderID: "exchange_order_id",
	Limit:           "limit",
	Quantity:        "quantity",
	Status:          "status",
	SellPrice:       "sell_price",
	OpenTime:        "open_time",
	CloseTime:       "close_time",
	QuantityFilled:  "quantity_filled",
}

// orderR is where relationships are stored.
type orderR struct {
	Market *Market
}

// orderL is where Load methods for each relationship are stored.
type orderL struct{}

var (
	orderColumns               = []string{"id", "market_id", "type", "exchange_order_id", "limit", "quantity", "status", "sell_price", "open_time", "close_time", "quantity_filled"}
	orderColumnsWithoutDefault = []string{"market_id", "type", "exchange_order_id", "limit", "quantity", "status", "sell_price", "open_time", "close_time", "quantity_filled"}
	orderColumnsWithDefault    = []string{"id"}
	orderPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrderSlice is an alias for a slice of pointers to Order.
	// This should generally be used opposed to []Order.
	OrderSlice []*Order
	// OrderHook is the signature for custom Order hook methods
	OrderHook func(boil.Executor, *Order) error

	orderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderType                 = reflect.TypeOf(&Order{})
	orderMapping              = queries.MakeStructMapping(orderType)
	orderPrimaryKeyMapping, _ = queries.BindMapping(orderType, orderMapping, orderPrimaryKeyColumns)
	orderInsertCacheMut       sync.RWMutex
	orderInsertCache          = make(map[string]insertCache)
	orderUpdateCacheMut       sync.RWMutex
	orderUpdateCache          = make(map[string]updateCache)
	orderUpsertCacheMut       sync.RWMutex
	orderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var orderBeforeInsertHooks []OrderHook
var orderBeforeUpdateHooks []OrderHook
var orderBeforeDeleteHooks []OrderHook
var orderBeforeUpsertHooks []OrderHook

var orderAfterInsertHooks []OrderHook
var orderAfterSelectHooks []OrderHook
var orderAfterUpdateHooks []OrderHook
var orderAfterDeleteHooks []OrderHook
var orderAfterUpsertHooks []OrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Order) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range orderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Order) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range orderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Order) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range orderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Order) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range orderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Order) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range orderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Order) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range orderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Order) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range orderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Order) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range orderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Order) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range orderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderHook registers your hook function for all future operations.
func AddOrderHook(hookPoint boil.HookPoint, orderHook OrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		orderBeforeInsertHooks = append(orderBeforeInsertHooks, orderHook)
	case boil.BeforeUpdateHook:
		orderBeforeUpdateHooks = append(orderBeforeUpdateHooks, orderHook)
	case boil.BeforeDeleteHook:
		orderBeforeDeleteHooks = append(orderBeforeDeleteHooks, orderHook)
	case boil.BeforeUpsertHook:
		orderBeforeUpsertHooks = append(orderBeforeUpsertHooks, orderHook)
	case boil.AfterInsertHook:
		orderAfterInsertHooks = append(orderAfterInsertHooks, orderHook)
	case boil.AfterSelectHook:
		orderAfterSelectHooks = append(orderAfterSelectHooks, orderHook)
	case boil.AfterUpdateHook:
		orderAfterUpdateHooks = append(orderAfterUpdateHooks, orderHook)
	case boil.AfterDeleteHook:
		orderAfterDeleteHooks = append(orderAfterDeleteHooks, orderHook)
	case boil.AfterUpsertHook:
		orderAfterUpsertHooks = append(orderAfterUpsertHooks, orderHook)
	}
}

// OneP returns a single order record from the query, and panics on error.
func (q orderQuery) OneP() *Order {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single order record from the query.
func (q orderQuery) One() (*Order, error) {
	o := &Order{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for order")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Order records from the query, and panics on error.
func (q orderQuery) AllP() OrderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Order records from the query.
func (q orderQuery) All() (OrderSlice, error) {
	var o []*Order

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Order slice")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Order records in the query, and panics on error.
func (q orderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Order records in the query.
func (q orderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count order rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q orderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q orderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if order exists")
	}

	return count > 0, nil
}

// MarketG pointed to by the foreign key.
func (o *Order) MarketG(mods ...qm.QueryMod) marketQuery {
	return o.Market(boil.GetDB(), mods...)
}

// Market pointed to by the foreign key.
func (o *Order) Market(exec boil.Executor, mods ...qm.QueryMod) marketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MarketID),
	}

	queryMods = append(queryMods, mods...)

	query := Markets(exec, queryMods...)
	queries.SetFrom(query.Query, "\"market\"")

	return query
} // LoadMarket allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (orderL) LoadMarket(e boil.Executor, singular bool, maybeOrder interface{}) error {
	var slice []*Order
	var object *Order

	count := 1
	if singular {
		object = maybeOrder.(*Order)
	} else {
		slice = *maybeOrder.(*[]*Order)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &orderR{}
		}
		args[0] = object.MarketID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &orderR{}
			}
			args[i] = obj.MarketID
		}
	}

	query := fmt.Sprintf(
		"select * from \"market\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Market")
	}
	defer results.Close()

	var resultSlice []*Market
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Market")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Market = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MarketID == foreign.ID {
				local.R.Market = foreign
				break
			}
		}
	}

	return nil
}

// SetMarketG of the order to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Orders.
// Uses the global database handle.
func (o *Order) SetMarketG(insert bool, related *Market) error {
	return o.SetMarket(boil.GetDB(), insert, related)
}

// SetMarketP of the order to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Orders.
// Panics on error.
func (o *Order) SetMarketP(exec boil.Executor, insert bool, related *Market) {
	if err := o.SetMarket(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMarketGP of the order to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Orders.
// Uses the global database handle and panics on error.
func (o *Order) SetMarketGP(insert bool, related *Market) {
	if err := o.SetMarket(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMarket of the order to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Orders.
func (o *Order) SetMarket(exec boil.Executor, insert bool, related *Market) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"market_id"}),
		strmangle.WhereClause("\"", "\"", 2, orderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MarketID = related.ID

	if o.R == nil {
		o.R = &orderR{
			Market: related,
		}
	} else {
		o.R.Market = related
	}

	if related.R == nil {
		related.R = &marketR{
			Orders: OrderSlice{o},
		}
	} else {
		related.R.Orders = append(related.R.Orders, o)
	}

	return nil
}

// OrdersG retrieves all records.
func OrdersG(mods ...qm.QueryMod) orderQuery {
	return Orders(boil.GetDB(), mods...)
}

// Orders retrieves all the records using an executor.
func Orders(exec boil.Executor, mods ...qm.QueryMod) orderQuery {
	mods = append(mods, qm.From("\"order\""))
	return orderQuery{NewQuery(exec, mods...)}
}

// FindOrderG retrieves a single record by ID.
func FindOrderG(id int, selectCols ...string) (*Order, error) {
	return FindOrder(boil.GetDB(), id, selectCols...)
}

// FindOrderGP retrieves a single record by ID, and panics on error.
func FindOrderGP(id int, selectCols ...string) *Order {
	retobj, err := FindOrder(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrder(exec boil.Executor, id int, selectCols ...string) (*Order, error) {
	orderObj := &Order{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(orderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from order")
	}

	return orderObj, nil
}

// FindOrderP retrieves a single record by ID with an executor, and panics on error.
func FindOrderP(exec boil.Executor, id int, selectCols ...string) *Order {
	retobj, err := FindOrder(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Order) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Order) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Order) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Order) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no order provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	orderInsertCacheMut.RLock()
	cache, cached := orderInsertCache[key]
	orderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			orderColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderType, orderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into order")
	}

	if !cached {
		orderInsertCacheMut.Lock()
		orderInsertCache[key] = cache
		orderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Order record. See Update for
// whitelist behavior description.
func (o *Order) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Order record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Order) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Order, and panics on error.
// See Update for whitelist behavior description.
func (o *Order) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Order.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Order) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	orderUpdateCacheMut.RLock()
	cache, cached := orderUpdateCache[key]
	orderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			orderColumns,
			orderPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, append(wl, orderPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update order row")
	}

	if !cached {
		orderUpdateCacheMut.Lock()
		orderUpdateCache[key] = cache
		orderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q orderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q orderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for order")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OrderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OrderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in order slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Order) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Order) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Order) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Order) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no order provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderUpsertCacheMut.RLock()
	cache, cached := orderUpsertCache[key]
	orderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			orderColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			orderColumns,
			orderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert order, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderPrimaryKeyColumns))
			copy(conflict, orderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"order\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderType, orderMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert order")
	}

	if !cached {
		orderUpsertCacheMut.Lock()
		orderUpsertCache[key] = cache
		orderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Order record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Order) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Order record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Order) DeleteG() error {
	if o == nil {
		return errors.New("models: no Order provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Order record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Order) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Order record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Order) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Order provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderPrimaryKeyMapping)
	sql := "DELETE FROM \"order\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from order")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q orderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q orderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no orderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from order")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OrderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OrderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Order slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OrderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Order slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(orderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from order slice")
	}

	if len(orderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Order) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Order) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Order) ReloadG() error {
	if o == nil {
		return errors.New("models: no Order provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Order) Reload(exec boil.Executor) error {
	ret, err := FindOrder(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OrderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	orders := OrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order\".* FROM \"order\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&orders)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderSlice")
	}

	*o = orders

	return nil
}

// OrderExists checks if the Order row exists.
func OrderExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if order exists")
	}

	return exists, nil
}

// OrderExistsG checks if the Order row exists.
func OrderExistsG(id int) (bool, error) {
	return OrderExists(boil.GetDB(), id)
}

// OrderExistsGP checks if the Order row exists. Panics on error.
func OrderExistsGP(id int) bool {
	e, err := OrderExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OrderExistsP checks if the Order row exists. Panics on error.
func OrderExistsP(exec boil.Executor, id int) bool {
	e, err := OrderExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
