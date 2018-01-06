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

// Exchange is an object representing the database table.
type Exchange struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *exchangeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L exchangeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExchangeColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// exchangeR is where relationships are stored.
type exchangeR struct {
}

// exchangeL is where Load methods for each relationship are stored.
type exchangeL struct{}

var (
	exchangeColumns               = []string{"id", "name"}
	exchangeColumnsWithoutDefault = []string{"name"}
	exchangeColumnsWithDefault    = []string{"id"}
	exchangePrimaryKeyColumns     = []string{"id"}
)

type (
	// ExchangeSlice is an alias for a slice of pointers to Exchange.
	// This should generally be used opposed to []Exchange.
	ExchangeSlice []*Exchange
	// ExchangeHook is the signature for custom Exchange hook methods
	ExchangeHook func(boil.Executor, *Exchange) error

	exchangeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	exchangeType                 = reflect.TypeOf(&Exchange{})
	exchangeMapping              = queries.MakeStructMapping(exchangeType)
	exchangePrimaryKeyMapping, _ = queries.BindMapping(exchangeType, exchangeMapping, exchangePrimaryKeyColumns)
	exchangeInsertCacheMut       sync.RWMutex
	exchangeInsertCache          = make(map[string]insertCache)
	exchangeUpdateCacheMut       sync.RWMutex
	exchangeUpdateCache          = make(map[string]updateCache)
	exchangeUpsertCacheMut       sync.RWMutex
	exchangeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var exchangeBeforeInsertHooks []ExchangeHook
var exchangeBeforeUpdateHooks []ExchangeHook
var exchangeBeforeDeleteHooks []ExchangeHook
var exchangeBeforeUpsertHooks []ExchangeHook

var exchangeAfterInsertHooks []ExchangeHook
var exchangeAfterSelectHooks []ExchangeHook
var exchangeAfterUpdateHooks []ExchangeHook
var exchangeAfterDeleteHooks []ExchangeHook
var exchangeAfterUpsertHooks []ExchangeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Exchange) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Exchange) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Exchange) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Exchange) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Exchange) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Exchange) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Exchange) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Exchange) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Exchange) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range exchangeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExchangeHook registers your hook function for all future operations.
func AddExchangeHook(hookPoint boil.HookPoint, exchangeHook ExchangeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		exchangeBeforeInsertHooks = append(exchangeBeforeInsertHooks, exchangeHook)
	case boil.BeforeUpdateHook:
		exchangeBeforeUpdateHooks = append(exchangeBeforeUpdateHooks, exchangeHook)
	case boil.BeforeDeleteHook:
		exchangeBeforeDeleteHooks = append(exchangeBeforeDeleteHooks, exchangeHook)
	case boil.BeforeUpsertHook:
		exchangeBeforeUpsertHooks = append(exchangeBeforeUpsertHooks, exchangeHook)
	case boil.AfterInsertHook:
		exchangeAfterInsertHooks = append(exchangeAfterInsertHooks, exchangeHook)
	case boil.AfterSelectHook:
		exchangeAfterSelectHooks = append(exchangeAfterSelectHooks, exchangeHook)
	case boil.AfterUpdateHook:
		exchangeAfterUpdateHooks = append(exchangeAfterUpdateHooks, exchangeHook)
	case boil.AfterDeleteHook:
		exchangeAfterDeleteHooks = append(exchangeAfterDeleteHooks, exchangeHook)
	case boil.AfterUpsertHook:
		exchangeAfterUpsertHooks = append(exchangeAfterUpsertHooks, exchangeHook)
	}
}

// OneP returns a single exchange record from the query, and panics on error.
func (q exchangeQuery) OneP() *Exchange {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single exchange record from the query.
func (q exchangeQuery) One() (*Exchange, error) {
	o := &Exchange{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for exchange")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Exchange records from the query, and panics on error.
func (q exchangeQuery) AllP() ExchangeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Exchange records from the query.
func (q exchangeQuery) All() (ExchangeSlice, error) {
	var o []*Exchange

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Exchange slice")
	}

	if len(exchangeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Exchange records in the query, and panics on error.
func (q exchangeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Exchange records in the query.
func (q exchangeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count exchange rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q exchangeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q exchangeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if exchange exists")
	}

	return count > 0, nil
}

// ExchangesG retrieves all records.
func ExchangesG(mods ...qm.QueryMod) exchangeQuery {
	return Exchanges(boil.GetDB(), mods...)
}

// Exchanges retrieves all the records using an executor.
func Exchanges(exec boil.Executor, mods ...qm.QueryMod) exchangeQuery {
	mods = append(mods, qm.From("\"exchange\""))
	return exchangeQuery{NewQuery(exec, mods...)}
}

// FindExchangeG retrieves a single record by ID.
func FindExchangeG(id int, selectCols ...string) (*Exchange, error) {
	return FindExchange(boil.GetDB(), id, selectCols...)
}

// FindExchangeGP retrieves a single record by ID, and panics on error.
func FindExchangeGP(id int, selectCols ...string) *Exchange {
	retobj, err := FindExchange(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindExchange retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExchange(exec boil.Executor, id int, selectCols ...string) (*Exchange, error) {
	exchangeObj := &Exchange{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"exchange\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(exchangeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from exchange")
	}

	return exchangeObj, nil
}

// FindExchangeP retrieves a single record by ID with an executor, and panics on error.
func FindExchangeP(exec boil.Executor, id int, selectCols ...string) *Exchange {
	retobj, err := FindExchange(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Exchange) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Exchange) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Exchange) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Exchange) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no exchange provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exchangeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	exchangeInsertCacheMut.RLock()
	cache, cached := exchangeInsertCache[key]
	exchangeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			exchangeColumns,
			exchangeColumnsWithDefault,
			exchangeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(exchangeType, exchangeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"exchange\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"exchange\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into exchange")
	}

	if !cached {
		exchangeInsertCacheMut.Lock()
		exchangeInsertCache[key] = cache
		exchangeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Exchange record. See Update for
// whitelist behavior description.
func (o *Exchange) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Exchange record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Exchange) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Exchange, and panics on error.
// See Update for whitelist behavior description.
func (o *Exchange) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Exchange.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Exchange) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	exchangeUpdateCacheMut.RLock()
	cache, cached := exchangeUpdateCache[key]
	exchangeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			exchangeColumns,
			exchangePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update exchange, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, exchangePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, append(wl, exchangePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update exchange row")
	}

	if !cached {
		exchangeUpdateCacheMut.Lock()
		exchangeUpdateCache[key] = cache
		exchangeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q exchangeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q exchangeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for exchange")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ExchangeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ExchangeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ExchangeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExchangeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, exchangePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in exchange slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Exchange) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Exchange) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Exchange) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Exchange) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no exchange provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exchangeColumnsWithDefault, o)

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

	exchangeUpsertCacheMut.RLock()
	cache, cached := exchangeUpsertCache[key]
	exchangeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			exchangeColumns,
			exchangeColumnsWithDefault,
			exchangeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			exchangeColumns,
			exchangePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert exchange, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(exchangePrimaryKeyColumns))
			copy(conflict, exchangePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"exchange\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(exchangeType, exchangeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert exchange")
	}

	if !cached {
		exchangeUpsertCacheMut.Lock()
		exchangeUpsertCache[key] = cache
		exchangeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Exchange record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Exchange) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Exchange record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Exchange) DeleteG() error {
	if o == nil {
		return errors.New("models: no Exchange provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Exchange record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Exchange) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Exchange record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Exchange) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Exchange provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), exchangePrimaryKeyMapping)
	sql := "DELETE FROM \"exchange\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from exchange")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q exchangeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q exchangeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no exchangeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from exchange")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ExchangeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ExchangeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Exchange slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ExchangeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExchangeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Exchange slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(exchangeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, exchangePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from exchange slice")
	}

	if len(exchangeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Exchange) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Exchange) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Exchange) ReloadG() error {
	if o == nil {
		return errors.New("models: no Exchange provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Exchange) Reload(exec boil.Executor) error {
	ret, err := FindExchange(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ExchangeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ExchangeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExchangeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ExchangeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExchangeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	exchanges := ExchangeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"exchange\".* FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, exchangePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&exchanges)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExchangeSlice")
	}

	*o = exchanges

	return nil
}

// ExchangeExists checks if the Exchange row exists.
func ExchangeExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"exchange\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if exchange exists")
	}

	return exists, nil
}

// ExchangeExistsG checks if the Exchange row exists.
func ExchangeExistsG(id int) (bool, error) {
	return ExchangeExists(boil.GetDB(), id)
}

// ExchangeExistsGP checks if the Exchange row exists. Panics on error.
func ExchangeExistsGP(id int) bool {
	e, err := ExchangeExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExchangeExistsP checks if the Exchange row exists. Panics on error.
func ExchangeExistsP(exec boil.Executor, id int) bool {
	e, err := ExchangeExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
