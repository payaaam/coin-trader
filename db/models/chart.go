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
)

// Chart is an object representing the database table.
type Chart struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	MarketID int    `boil:"market_id" json:"market_id" toml:"market_id" yaml:"market_id"`
	Interval string `boil:"interval" json:"interval" toml:"interval" yaml:"interval"`

	R *chartR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chartL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChartColumns = struct {
	ID       string
	MarketID string
	Interval string
}{
	ID:       "id",
	MarketID: "market_id",
	Interval: "interval",
}

// chartR is where relationships are stored.
type chartR struct {
	Market *Market
	Ticks  TickSlice
}

// chartL is where Load methods for each relationship are stored.
type chartL struct{}

var (
	chartColumns               = []string{"id", "market_id", "interval"}
	chartColumnsWithoutDefault = []string{"market_id", "interval"}
	chartColumnsWithDefault    = []string{"id"}
	chartPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChartSlice is an alias for a slice of pointers to Chart.
	// This should generally be used opposed to []Chart.
	ChartSlice []*Chart
	// ChartHook is the signature for custom Chart hook methods
	ChartHook func(boil.Executor, *Chart) error

	chartQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chartType                 = reflect.TypeOf(&Chart{})
	chartMapping              = queries.MakeStructMapping(chartType)
	chartPrimaryKeyMapping, _ = queries.BindMapping(chartType, chartMapping, chartPrimaryKeyColumns)
	chartInsertCacheMut       sync.RWMutex
	chartInsertCache          = make(map[string]insertCache)
	chartUpdateCacheMut       sync.RWMutex
	chartUpdateCache          = make(map[string]updateCache)
	chartUpsertCacheMut       sync.RWMutex
	chartUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var chartBeforeInsertHooks []ChartHook
var chartBeforeUpdateHooks []ChartHook
var chartBeforeDeleteHooks []ChartHook
var chartBeforeUpsertHooks []ChartHook

var chartAfterInsertHooks []ChartHook
var chartAfterSelectHooks []ChartHook
var chartAfterUpdateHooks []ChartHook
var chartAfterDeleteHooks []ChartHook
var chartAfterUpsertHooks []ChartHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Chart) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Chart) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chartBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Chart) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chartBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Chart) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Chart) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Chart) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range chartAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Chart) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chartAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Chart) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chartAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Chart) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chartAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChartHook registers your hook function for all future operations.
func AddChartHook(hookPoint boil.HookPoint, chartHook ChartHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chartBeforeInsertHooks = append(chartBeforeInsertHooks, chartHook)
	case boil.BeforeUpdateHook:
		chartBeforeUpdateHooks = append(chartBeforeUpdateHooks, chartHook)
	case boil.BeforeDeleteHook:
		chartBeforeDeleteHooks = append(chartBeforeDeleteHooks, chartHook)
	case boil.BeforeUpsertHook:
		chartBeforeUpsertHooks = append(chartBeforeUpsertHooks, chartHook)
	case boil.AfterInsertHook:
		chartAfterInsertHooks = append(chartAfterInsertHooks, chartHook)
	case boil.AfterSelectHook:
		chartAfterSelectHooks = append(chartAfterSelectHooks, chartHook)
	case boil.AfterUpdateHook:
		chartAfterUpdateHooks = append(chartAfterUpdateHooks, chartHook)
	case boil.AfterDeleteHook:
		chartAfterDeleteHooks = append(chartAfterDeleteHooks, chartHook)
	case boil.AfterUpsertHook:
		chartAfterUpsertHooks = append(chartAfterUpsertHooks, chartHook)
	}
}

// OneP returns a single chart record from the query, and panics on error.
func (q chartQuery) OneP() *Chart {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single chart record from the query.
func (q chartQuery) One() (*Chart, error) {
	o := &Chart{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chart")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Chart records from the query, and panics on error.
func (q chartQuery) AllP() ChartSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Chart records from the query.
func (q chartQuery) All() (ChartSlice, error) {
	var o []*Chart

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Chart slice")
	}

	if len(chartAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Chart records in the query, and panics on error.
func (q chartQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Chart records in the query.
func (q chartQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chart rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q chartQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q chartQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chart exists")
	}

	return count > 0, nil
}

// MarketG pointed to by the foreign key.
func (o *Chart) MarketG(mods ...qm.QueryMod) marketQuery {
	return o.Market(boil.GetDB(), mods...)
}

// Market pointed to by the foreign key.
func (o *Chart) Market(exec boil.Executor, mods ...qm.QueryMod) marketQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MarketID),
	}

	queryMods = append(queryMods, mods...)

	query := Markets(exec, queryMods...)
	queries.SetFrom(query.Query, "\"market\"")

	return query
}

// TicksG retrieves all the tick's tick.
func (o *Chart) TicksG(mods ...qm.QueryMod) tickQuery {
	return o.Ticks(boil.GetDB(), mods...)
}

// Ticks retrieves all the tick's tick with an executor.
func (o *Chart) Ticks(exec boil.Executor, mods ...qm.QueryMod) tickQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"tick\".\"chart_id\"=?", o.ID),
	)

	query := Ticks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"tick\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"tick\".*"})
	}

	return query
}

// LoadMarket allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (chartL) LoadMarket(e boil.Executor, singular bool, maybeChart interface{}) error {
	var slice []*Chart
	var object *Chart

	count := 1
	if singular {
		object = maybeChart.(*Chart)
	} else {
		slice = *maybeChart.(*[]*Chart)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &chartR{}
		}
		args[0] = object.MarketID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &chartR{}
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

	if len(chartAfterSelectHooks) != 0 {
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

// LoadTicks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (chartL) LoadTicks(e boil.Executor, singular bool, maybeChart interface{}) error {
	var slice []*Chart
	var object *Chart

	count := 1
	if singular {
		object = maybeChart.(*Chart)
	} else {
		slice = *maybeChart.(*[]*Chart)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &chartR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &chartR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"tick\" where \"chart_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load tick")
	}
	defer results.Close()

	var resultSlice []*Tick
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice tick")
	}

	if len(tickAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Ticks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChartID {
				local.R.Ticks = append(local.R.Ticks, foreign)
				break
			}
		}
	}

	return nil
}

// SetMarketG of the chart to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Charts.
// Uses the global database handle.
func (o *Chart) SetMarketG(insert bool, related *Market) error {
	return o.SetMarket(boil.GetDB(), insert, related)
}

// SetMarketP of the chart to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Charts.
// Panics on error.
func (o *Chart) SetMarketP(exec boil.Executor, insert bool, related *Market) {
	if err := o.SetMarket(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMarketGP of the chart to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Charts.
// Uses the global database handle and panics on error.
func (o *Chart) SetMarketGP(insert bool, related *Market) {
	if err := o.SetMarket(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMarket of the chart to the related item.
// Sets o.R.Market to related.
// Adds o to related.R.Charts.
func (o *Chart) SetMarket(exec boil.Executor, insert bool, related *Market) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chart\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"market_id"}),
		strmangle.WhereClause("\"", "\"", 2, chartPrimaryKeyColumns),
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
		o.R = &chartR{
			Market: related,
		}
	} else {
		o.R.Market = related
	}

	if related.R == nil {
		related.R = &marketR{
			Charts: ChartSlice{o},
		}
	} else {
		related.R.Charts = append(related.R.Charts, o)
	}

	return nil
}

// AddTicksG adds the given related objects to the existing relationships
// of the chart, optionally inserting them as new records.
// Appends related to o.R.Ticks.
// Sets related.R.Chart appropriately.
// Uses the global database handle.
func (o *Chart) AddTicksG(insert bool, related ...*Tick) error {
	return o.AddTicks(boil.GetDB(), insert, related...)
}

// AddTicksP adds the given related objects to the existing relationships
// of the chart, optionally inserting them as new records.
// Appends related to o.R.Ticks.
// Sets related.R.Chart appropriately.
// Panics on error.
func (o *Chart) AddTicksP(exec boil.Executor, insert bool, related ...*Tick) {
	if err := o.AddTicks(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddTicksGP adds the given related objects to the existing relationships
// of the chart, optionally inserting them as new records.
// Appends related to o.R.Ticks.
// Sets related.R.Chart appropriately.
// Uses the global database handle and panics on error.
func (o *Chart) AddTicksGP(insert bool, related ...*Tick) {
	if err := o.AddTicks(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddTicks adds the given related objects to the existing relationships
// of the chart, optionally inserting them as new records.
// Appends related to o.R.Ticks.
// Sets related.R.Chart appropriately.
func (o *Chart) AddTicks(exec boil.Executor, insert bool, related ...*Tick) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChartID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"tick\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"chart_id"}),
				strmangle.WhereClause("\"", "\"", 2, tickPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChartID = o.ID
		}
	}

	if o.R == nil {
		o.R = &chartR{
			Ticks: related,
		}
	} else {
		o.R.Ticks = append(o.R.Ticks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tickR{
				Chart: o,
			}
		} else {
			rel.R.Chart = o
		}
	}
	return nil
}

// ChartsG retrieves all records.
func ChartsG(mods ...qm.QueryMod) chartQuery {
	return Charts(boil.GetDB(), mods...)
}

// Charts retrieves all the records using an executor.
func Charts(exec boil.Executor, mods ...qm.QueryMod) chartQuery {
	mods = append(mods, qm.From("\"chart\""))
	return chartQuery{NewQuery(exec, mods...)}
}

// FindChartG retrieves a single record by ID.
func FindChartG(id int, selectCols ...string) (*Chart, error) {
	return FindChart(boil.GetDB(), id, selectCols...)
}

// FindChartGP retrieves a single record by ID, and panics on error.
func FindChartGP(id int, selectCols ...string) *Chart {
	retobj, err := FindChart(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindChart retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChart(exec boil.Executor, id int, selectCols ...string) (*Chart, error) {
	chartObj := &Chart{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chart\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(chartObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chart")
	}

	return chartObj, nil
}

// FindChartP retrieves a single record by ID with an executor, and panics on error.
func FindChartP(exec boil.Executor, id int, selectCols ...string) *Chart {
	retobj, err := FindChart(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Chart) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Chart) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Chart) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Chart) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chart provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chartColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	chartInsertCacheMut.RLock()
	cache, cached := chartInsertCache[key]
	chartInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			chartColumns,
			chartColumnsWithDefault,
			chartColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(chartType, chartMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chartType, chartMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"chart\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"chart\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into chart")
	}

	if !cached {
		chartInsertCacheMut.Lock()
		chartInsertCache[key] = cache
		chartInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Chart record. See Update for
// whitelist behavior description.
func (o *Chart) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Chart record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Chart) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Chart, and panics on error.
// See Update for whitelist behavior description.
func (o *Chart) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Chart.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Chart) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	chartUpdateCacheMut.RLock()
	cache, cached := chartUpdateCache[key]
	chartUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			chartColumns,
			chartPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update chart, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chart\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chartPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chartType, chartMapping, append(wl, chartPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update chart row")
	}

	if !cached {
		chartUpdateCacheMut.Lock()
		chartUpdateCache[key] = cache
		chartUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q chartQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q chartQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for chart")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ChartSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ChartSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ChartSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChartSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"chart\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, chartPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in chart slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Chart) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Chart) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Chart) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Chart) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chart provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chartColumnsWithDefault, o)

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

	chartUpsertCacheMut.RLock()
	cache, cached := chartUpsertCache[key]
	chartUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			chartColumns,
			chartColumnsWithDefault,
			chartColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			chartColumns,
			chartPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert chart, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chartPrimaryKeyColumns))
			copy(conflict, chartPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"chart\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(chartType, chartMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chartType, chartMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert chart")
	}

	if !cached {
		chartUpsertCacheMut.Lock()
		chartUpsertCache[key] = cache
		chartUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Chart record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Chart) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Chart record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Chart) DeleteG() error {
	if o == nil {
		return errors.New("models: no Chart provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Chart record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Chart) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Chart record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Chart) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Chart provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chartPrimaryKeyMapping)
	sql := "DELETE FROM \"chart\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from chart")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q chartQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q chartQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no chartQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chart")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ChartSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ChartSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Chart slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ChartSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChartSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Chart slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(chartBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"chart\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chartPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chart slice")
	}

	if len(chartAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Chart) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Chart) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Chart) ReloadG() error {
	if o == nil {
		return errors.New("models: no Chart provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Chart) Reload(exec boil.Executor) error {
	ret, err := FindChart(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChartSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChartSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChartSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ChartSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChartSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	charts := ChartSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"chart\".* FROM \"chart\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chartPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&charts)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChartSlice")
	}

	*o = charts

	return nil
}

// ChartExists checks if the Chart row exists.
func ChartExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"chart\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chart exists")
	}

	return exists, nil
}

// ChartExistsG checks if the Chart row exists.
func ChartExistsG(id int) (bool, error) {
	return ChartExists(boil.GetDB(), id)
}

// ChartExistsGP checks if the Chart row exists. Panics on error.
func ChartExistsGP(id int) bool {
	e, err := ChartExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ChartExistsP checks if the Chart row exists. Panics on error.
func ChartExistsP(exec boil.Executor, id int) bool {
	e, err := ChartExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
