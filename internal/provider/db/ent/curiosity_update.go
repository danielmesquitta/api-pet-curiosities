// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/curiosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/predicate"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/usercuriosity"
	"github.com/google/uuid"
)

// CuriosityUpdate is the builder for updating Curiosity entities.
type CuriosityUpdate struct {
	config
	hooks    []Hook
	mutation *CuriosityMutation
}

// Where appends a list predicates to the CuriosityUpdate builder.
func (cu *CuriosityUpdate) Where(ps ...predicate.Curiosity) *CuriosityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *CuriosityUpdate) SetTitle(s string) *CuriosityUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *CuriosityUpdate) SetNillableTitle(s *string) *CuriosityUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetContent sets the "content" field.
func (cu *CuriosityUpdate) SetContent(s string) *CuriosityUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CuriosityUpdate) SetNillableContent(s *string) *CuriosityUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CuriosityUpdate) SetUpdatedAt(t time.Time) *CuriosityUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (cu *CuriosityUpdate) SetPetID(id uuid.UUID) *CuriosityUpdate {
	cu.mutation.SetPetID(id)
	return cu
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (cu *CuriosityUpdate) SetNillablePetID(id *uuid.UUID) *CuriosityUpdate {
	if id != nil {
		cu = cu.SetPetID(*id)
	}
	return cu
}

// SetPet sets the "pet" edge to the Pet entity.
func (cu *CuriosityUpdate) SetPet(p *Pet) *CuriosityUpdate {
	return cu.SetPetID(p.ID)
}

// AddUserCuriosityIDs adds the "user_curiosities" edge to the UserCuriosity entity by IDs.
func (cu *CuriosityUpdate) AddUserCuriosityIDs(ids ...uuid.UUID) *CuriosityUpdate {
	cu.mutation.AddUserCuriosityIDs(ids...)
	return cu
}

// AddUserCuriosities adds the "user_curiosities" edges to the UserCuriosity entity.
func (cu *CuriosityUpdate) AddUserCuriosities(u ...*UserCuriosity) *CuriosityUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserCuriosityIDs(ids...)
}

// Mutation returns the CuriosityMutation object of the builder.
func (cu *CuriosityUpdate) Mutation() *CuriosityMutation {
	return cu.mutation
}

// ClearPet clears the "pet" edge to the Pet entity.
func (cu *CuriosityUpdate) ClearPet() *CuriosityUpdate {
	cu.mutation.ClearPet()
	return cu
}

// ClearUserCuriosities clears all "user_curiosities" edges to the UserCuriosity entity.
func (cu *CuriosityUpdate) ClearUserCuriosities() *CuriosityUpdate {
	cu.mutation.ClearUserCuriosities()
	return cu
}

// RemoveUserCuriosityIDs removes the "user_curiosities" edge to UserCuriosity entities by IDs.
func (cu *CuriosityUpdate) RemoveUserCuriosityIDs(ids ...uuid.UUID) *CuriosityUpdate {
	cu.mutation.RemoveUserCuriosityIDs(ids...)
	return cu
}

// RemoveUserCuriosities removes "user_curiosities" edges to UserCuriosity entities.
func (cu *CuriosityUpdate) RemoveUserCuriosities(u ...*UserCuriosity) *CuriosityUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserCuriosityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CuriosityUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CuriosityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CuriosityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CuriosityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CuriosityUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := curiosity.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CuriosityUpdate) check() error {
	if v, ok := cu.mutation.Title(); ok {
		if err := curiosity.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Curiosity.title": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Content(); ok {
		if err := curiosity.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Curiosity.content": %w`, err)}
		}
	}
	return nil
}

func (cu *CuriosityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(curiosity.Table, curiosity.Columns, sqlgraph.NewFieldSpec(curiosity.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(curiosity.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(curiosity.FieldContent, field.TypeString, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(curiosity.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   curiosity.PetTable,
			Columns: []string{curiosity.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   curiosity.PetTable,
			Columns: []string{curiosity.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCuriositiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUserCuriositiesIDs(); len(nodes) > 0 && !cu.mutation.UserCuriositiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserCuriositiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{curiosity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CuriosityUpdateOne is the builder for updating a single Curiosity entity.
type CuriosityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CuriosityMutation
}

// SetTitle sets the "title" field.
func (cuo *CuriosityUpdateOne) SetTitle(s string) *CuriosityUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *CuriosityUpdateOne) SetNillableTitle(s *string) *CuriosityUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CuriosityUpdateOne) SetContent(s string) *CuriosityUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CuriosityUpdateOne) SetNillableContent(s *string) *CuriosityUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CuriosityUpdateOne) SetUpdatedAt(t time.Time) *CuriosityUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (cuo *CuriosityUpdateOne) SetPetID(id uuid.UUID) *CuriosityUpdateOne {
	cuo.mutation.SetPetID(id)
	return cuo
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (cuo *CuriosityUpdateOne) SetNillablePetID(id *uuid.UUID) *CuriosityUpdateOne {
	if id != nil {
		cuo = cuo.SetPetID(*id)
	}
	return cuo
}

// SetPet sets the "pet" edge to the Pet entity.
func (cuo *CuriosityUpdateOne) SetPet(p *Pet) *CuriosityUpdateOne {
	return cuo.SetPetID(p.ID)
}

// AddUserCuriosityIDs adds the "user_curiosities" edge to the UserCuriosity entity by IDs.
func (cuo *CuriosityUpdateOne) AddUserCuriosityIDs(ids ...uuid.UUID) *CuriosityUpdateOne {
	cuo.mutation.AddUserCuriosityIDs(ids...)
	return cuo
}

// AddUserCuriosities adds the "user_curiosities" edges to the UserCuriosity entity.
func (cuo *CuriosityUpdateOne) AddUserCuriosities(u ...*UserCuriosity) *CuriosityUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserCuriosityIDs(ids...)
}

// Mutation returns the CuriosityMutation object of the builder.
func (cuo *CuriosityUpdateOne) Mutation() *CuriosityMutation {
	return cuo.mutation
}

// ClearPet clears the "pet" edge to the Pet entity.
func (cuo *CuriosityUpdateOne) ClearPet() *CuriosityUpdateOne {
	cuo.mutation.ClearPet()
	return cuo
}

// ClearUserCuriosities clears all "user_curiosities" edges to the UserCuriosity entity.
func (cuo *CuriosityUpdateOne) ClearUserCuriosities() *CuriosityUpdateOne {
	cuo.mutation.ClearUserCuriosities()
	return cuo
}

// RemoveUserCuriosityIDs removes the "user_curiosities" edge to UserCuriosity entities by IDs.
func (cuo *CuriosityUpdateOne) RemoveUserCuriosityIDs(ids ...uuid.UUID) *CuriosityUpdateOne {
	cuo.mutation.RemoveUserCuriosityIDs(ids...)
	return cuo
}

// RemoveUserCuriosities removes "user_curiosities" edges to UserCuriosity entities.
func (cuo *CuriosityUpdateOne) RemoveUserCuriosities(u ...*UserCuriosity) *CuriosityUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserCuriosityIDs(ids...)
}

// Where appends a list predicates to the CuriosityUpdate builder.
func (cuo *CuriosityUpdateOne) Where(ps ...predicate.Curiosity) *CuriosityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CuriosityUpdateOne) Select(field string, fields ...string) *CuriosityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Curiosity entity.
func (cuo *CuriosityUpdateOne) Save(ctx context.Context) (*Curiosity, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CuriosityUpdateOne) SaveX(ctx context.Context) *Curiosity {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CuriosityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CuriosityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CuriosityUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := curiosity.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CuriosityUpdateOne) check() error {
	if v, ok := cuo.mutation.Title(); ok {
		if err := curiosity.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Curiosity.title": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Content(); ok {
		if err := curiosity.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Curiosity.content": %w`, err)}
		}
	}
	return nil
}

func (cuo *CuriosityUpdateOne) sqlSave(ctx context.Context) (_node *Curiosity, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(curiosity.Table, curiosity.Columns, sqlgraph.NewFieldSpec(curiosity.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Curiosity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, curiosity.FieldID)
		for _, f := range fields {
			if !curiosity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != curiosity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(curiosity.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(curiosity.FieldContent, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(curiosity.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   curiosity.PetTable,
			Columns: []string{curiosity.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   curiosity.PetTable,
			Columns: []string{curiosity.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCuriositiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUserCuriositiesIDs(); len(nodes) > 0 && !cuo.mutation.UserCuriositiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserCuriositiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   curiosity.UserCuriositiesTable,
			Columns: []string{curiosity.UserCuriositiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Curiosity{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{curiosity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
