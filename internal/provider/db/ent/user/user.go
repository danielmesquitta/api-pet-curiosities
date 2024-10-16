// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldTier holds the string denoting the tier field in the database.
	FieldTier = "tier"
	// FieldSubscriptionExpiresAt holds the string denoting the subscription_expires_at field in the database.
	FieldSubscriptionExpiresAt = "subscription_expires_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeUserCuriosities holds the string denoting the user_curiosities edge name in mutations.
	EdgeUserCuriosities = "user_curiosities"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge. The primary key declared below.
	PetsTable = "user_pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// UserCuriositiesTable is the table that holds the user_curiosities relation/edge.
	UserCuriositiesTable = "user_curiosities"
	// UserCuriositiesInverseTable is the table name for the UserCuriosity entity.
	// It exists in this package in order to avoid circular dependency with the "usercuriosity" package.
	UserCuriositiesInverseTable = "user_curiosities"
	// UserCuriositiesColumn is the table column denoting the user_curiosities relation/edge.
	UserCuriositiesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldTier,
	FieldSubscriptionExpiresAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// PetsPrimaryKey and PetsColumn2 are the table columns denoting the
	// primary key for the pets relation (M2M).
	PetsPrimaryKey = []string{"user_id", "pet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Tier defines the type for the "tier" enum field.
type Tier string

// TierFREE is the default value of the Tier enum.
const DefaultTier = TierFREE

// Tier values.
const (
	TierPRO  Tier = "PRO"
	TierFREE Tier = "FREE"
)

func (t Tier) String() string {
	return string(t)
}

// TierValidator is a validator for the "tier" field enum values. It is called by the builders before save.
func TierValidator(t Tier) error {
	switch t {
	case TierPRO, TierFREE:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for tier field: %q", t)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByTier orders the results by the tier field.
func ByTier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTier, opts...).ToFunc()
}

// BySubscriptionExpiresAt orders the results by the subscription_expires_at field.
func BySubscriptionExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionExpiresAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPetsCount orders the results by pets count.
func ByPetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetsStep(), opts...)
	}
}

// ByPets orders the results by pets terms.
func ByPets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCuriositiesCount orders the results by user_curiosities count.
func ByUserCuriositiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserCuriositiesStep(), opts...)
	}
}

// ByUserCuriosities orders the results by user_curiosities terms.
func ByUserCuriosities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserCuriositiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PetsTable, PetsPrimaryKey...),
	)
}
func newUserCuriositiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserCuriositiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserCuriositiesTable, UserCuriositiesColumn),
	)
}
