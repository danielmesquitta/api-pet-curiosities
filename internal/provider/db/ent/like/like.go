// Code generated by ent, DO NOT EDIT.

package like

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the like type in the database.
	Label = "like"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCuriosity holds the string denoting the curiosity edge name in mutations.
	EdgeCuriosity = "curiosity"
	// Table holds the table name of the like in the database.
	Table = "likes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CuriosityTable is the table that holds the curiosity relation/edge.
	CuriosityTable = "likes"
	// CuriosityInverseTable is the table name for the Curiosity entity.
	// It exists in this package in order to avoid circular dependency with the "curiosity" package.
	CuriosityInverseTable = "curiosities"
	// CuriosityColumn is the table column denoting the curiosity relation/edge.
	CuriosityColumn = "curiosity_id"
)

// Columns holds all SQL columns for like fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "likes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
	"curiosity_id",
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

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Like queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCuriosityField orders the results by curiosity field.
func ByCuriosityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCuriosityStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newCuriosityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CuriosityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CuriosityTable, CuriosityColumn),
	)
}
