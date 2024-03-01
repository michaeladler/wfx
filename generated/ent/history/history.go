// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package history

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDefinition holds the string denoting the definition field in the database.
	FieldDefinition = "definition"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"
	// Table holds the table name of the history in the database.
	Table = "history"
	// JobTable is the table that holds the job relation/edge.
	JobTable = "history"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "job"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_history"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldMtime,
	FieldStatus,
	FieldDefinition,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "history"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"job_history",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the History queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMtime orders the results by the mtime field.
func ByMtime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMtime, opts...).ToFunc()
}

// ByJobField orders the results by job field.
func ByJobField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobStep(), sql.OrderByField(field, opts...))
	}
}
func newJobStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
	)
}