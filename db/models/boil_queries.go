// This file is generated by SQLBoiler (https://github.com/volatiletech/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dialect = queries.Dialect{
	LQ:                0x22,
	RQ:                0x22,
	IndexPlaceholders: true,
	UseTopClause:      false,
}

// NewQueryG initializes a new Query using the passed in QueryMods
func NewQueryG(mods ...qm.QueryMod) *queries.Query {
	return NewQuery(boil.GetDB(), mods...)
}

// NewQuery initializes a new Query using the passed in QueryMods
func NewQuery(exec boil.Executor, mods ...qm.QueryMod) *queries.Query {
	q := &queries.Query{}
	queries.SetExecutor(q, exec)
	queries.SetDialect(q, &dialect)
	qm.Apply(q, mods...)

	return q
}
