// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/curiosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/view"
	"github.com/google/uuid"
)

// View is the model entity for the View schema.
type View struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ViewQuery when eager-loading is set.
	Edges        ViewEdges `json:"edges"`
	user_id      *uuid.UUID
	curiosity_id *uuid.UUID
	selectValues sql.SelectValues
}

// ViewEdges holds the relations/edges for other nodes in the graph.
type ViewEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Curiosity holds the value of the curiosity edge.
	Curiosity *Curiosity `json:"curiosity,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ViewEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CuriosityOrErr returns the Curiosity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ViewEdges) CuriosityOrErr() (*Curiosity, error) {
	if e.Curiosity != nil {
		return e.Curiosity, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: curiosity.Label}
	}
	return nil, &NotLoadedError{edge: "curiosity"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*View) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case view.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case view.FieldID:
			values[i] = new(uuid.UUID)
		case view.ForeignKeys[0]: // user_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case view.ForeignKeys[1]: // curiosity_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the View fields.
func (v *View) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case view.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				v.ID = *value
			}
		case view.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case view.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				v.user_id = new(uuid.UUID)
				*v.user_id = *value.S.(*uuid.UUID)
			}
		case view.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field curiosity_id", values[i])
			} else if value.Valid {
				v.curiosity_id = new(uuid.UUID)
				*v.curiosity_id = *value.S.(*uuid.UUID)
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the View.
// This includes values selected through modifiers, order, etc.
func (v *View) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the View entity.
func (v *View) QueryUser() *UserQuery {
	return NewViewClient(v.config).QueryUser(v)
}

// QueryCuriosity queries the "curiosity" edge of the View entity.
func (v *View) QueryCuriosity() *CuriosityQuery {
	return NewViewClient(v.config).QueryCuriosity(v)
}

// Update returns a builder for updating this View.
// Note that you need to call View.Unwrap() before calling this method if this View
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *View) Update() *ViewUpdateOne {
	return NewViewClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the View entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *View) Unwrap() *View {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: View is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *View) String() string {
	var builder strings.Builder
	builder.WriteString("View(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Views is a parsable slice of View.
type Views []*View
