// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/curiosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/predicate"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/usercuriosity"
	"github.com/google/uuid"
)

// UserCuriosityQuery is the builder for querying UserCuriosity entities.
type UserCuriosityQuery struct {
	config
	ctx           *QueryContext
	order         []usercuriosity.OrderOption
	inters        []Interceptor
	predicates    []predicate.UserCuriosity
	withUser      *UserQuery
	withCuriosity *CuriosityQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCuriosityQuery builder.
func (ucq *UserCuriosityQuery) Where(ps ...predicate.UserCuriosity) *UserCuriosityQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit the number of records to be returned by this query.
func (ucq *UserCuriosityQuery) Limit(limit int) *UserCuriosityQuery {
	ucq.ctx.Limit = &limit
	return ucq
}

// Offset to start from.
func (ucq *UserCuriosityQuery) Offset(offset int) *UserCuriosityQuery {
	ucq.ctx.Offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserCuriosityQuery) Unique(unique bool) *UserCuriosityQuery {
	ucq.ctx.Unique = &unique
	return ucq
}

// Order specifies how the records should be ordered.
func (ucq *UserCuriosityQuery) Order(o ...usercuriosity.OrderOption) *UserCuriosityQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryUser chains the current query on the "user" edge.
func (ucq *UserCuriosityQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercuriosity.Table, usercuriosity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usercuriosity.UserTable, usercuriosity.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCuriosity chains the current query on the "curiosity" edge.
func (ucq *UserCuriosityQuery) QueryCuriosity() *CuriosityQuery {
	query := (&CuriosityClient{config: ucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercuriosity.Table, usercuriosity.FieldID, selector),
			sqlgraph.To(curiosity.Table, curiosity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usercuriosity.CuriosityTable, usercuriosity.CuriosityColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCuriosity entity from the query.
// Returns a *NotFoundError when no UserCuriosity was found.
func (ucq *UserCuriosityQuery) First(ctx context.Context) (*UserCuriosity, error) {
	nodes, err := ucq.Limit(1).All(setContextOp(ctx, ucq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercuriosity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserCuriosityQuery) FirstX(ctx context.Context) *UserCuriosity {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCuriosity ID from the query.
// Returns a *NotFoundError when no UserCuriosity ID was found.
func (ucq *UserCuriosityQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ucq.Limit(1).IDs(setContextOp(ctx, ucq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercuriosity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserCuriosityQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCuriosity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCuriosity entity is found.
// Returns a *NotFoundError when no UserCuriosity entities are found.
func (ucq *UserCuriosityQuery) Only(ctx context.Context) (*UserCuriosity, error) {
	nodes, err := ucq.Limit(2).All(setContextOp(ctx, ucq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercuriosity.Label}
	default:
		return nil, &NotSingularError{usercuriosity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserCuriosityQuery) OnlyX(ctx context.Context) *UserCuriosity {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCuriosity ID in the query.
// Returns a *NotSingularError when more than one UserCuriosity ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserCuriosityQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ucq.Limit(2).IDs(setContextOp(ctx, ucq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercuriosity.Label}
	default:
		err = &NotSingularError{usercuriosity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserCuriosityQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCuriosities.
func (ucq *UserCuriosityQuery) All(ctx context.Context) ([]*UserCuriosity, error) {
	ctx = setContextOp(ctx, ucq.ctx, ent.OpQueryAll)
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserCuriosity, *UserCuriosityQuery]()
	return withInterceptors[[]*UserCuriosity](ctx, ucq, qr, ucq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserCuriosityQuery) AllX(ctx context.Context) []*UserCuriosity {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCuriosity IDs.
func (ucq *UserCuriosityQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ucq.ctx.Unique == nil && ucq.path != nil {
		ucq.Unique(true)
	}
	ctx = setContextOp(ctx, ucq.ctx, ent.OpQueryIDs)
	if err = ucq.Select(usercuriosity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserCuriosityQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserCuriosityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ucq.ctx, ent.OpQueryCount)
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ucq, querierCount[*UserCuriosityQuery](), ucq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserCuriosityQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserCuriosityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ucq.ctx, ent.OpQueryExist)
	switch _, err := ucq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserCuriosityQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCuriosityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserCuriosityQuery) Clone() *UserCuriosityQuery {
	if ucq == nil {
		return nil
	}
	return &UserCuriosityQuery{
		config:        ucq.config,
		ctx:           ucq.ctx.Clone(),
		order:         append([]usercuriosity.OrderOption{}, ucq.order...),
		inters:        append([]Interceptor{}, ucq.inters...),
		predicates:    append([]predicate.UserCuriosity{}, ucq.predicates...),
		withUser:      ucq.withUser.Clone(),
		withCuriosity: ucq.withCuriosity.Clone(),
		// clone intermediate query.
		sql:  ucq.sql.Clone(),
		path: ucq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCuriosityQuery) WithUser(opts ...func(*UserQuery)) *UserCuriosityQuery {
	query := (&UserClient{config: ucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucq.withUser = query
	return ucq
}

// WithCuriosity tells the query-builder to eager-load the nodes that are connected to
// the "curiosity" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCuriosityQuery) WithCuriosity(opts ...func(*CuriosityQuery)) *UserCuriosityQuery {
	query := (&CuriosityClient{config: ucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucq.withCuriosity = query
	return ucq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Viewed bool `json:"viewed,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserCuriosity.Query().
//		GroupBy(usercuriosity.FieldViewed).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ucq *UserCuriosityQuery) GroupBy(field string, fields ...string) *UserCuriosityGroupBy {
	ucq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserCuriosityGroupBy{build: ucq}
	grbuild.flds = &ucq.ctx.Fields
	grbuild.label = usercuriosity.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Viewed bool `json:"viewed,omitempty"`
//	}
//
//	client.UserCuriosity.Query().
//		Select(usercuriosity.FieldViewed).
//		Scan(ctx, &v)
func (ucq *UserCuriosityQuery) Select(fields ...string) *UserCuriositySelect {
	ucq.ctx.Fields = append(ucq.ctx.Fields, fields...)
	sbuild := &UserCuriositySelect{UserCuriosityQuery: ucq}
	sbuild.label = usercuriosity.Label
	sbuild.flds, sbuild.scan = &ucq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserCuriositySelect configured with the given aggregations.
func (ucq *UserCuriosityQuery) Aggregate(fns ...AggregateFunc) *UserCuriositySelect {
	return ucq.Select().Aggregate(fns...)
}

func (ucq *UserCuriosityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ucq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ucq); err != nil {
				return err
			}
		}
	}
	for _, f := range ucq.ctx.Fields {
		if !usercuriosity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserCuriosityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCuriosity, error) {
	var (
		nodes       = []*UserCuriosity{}
		withFKs     = ucq.withFKs
		_spec       = ucq.querySpec()
		loadedTypes = [2]bool{
			ucq.withUser != nil,
			ucq.withCuriosity != nil,
		}
	)
	if ucq.withUser != nil || ucq.withCuriosity != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usercuriosity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserCuriosity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserCuriosity{config: ucq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucq.withUser; query != nil {
		if err := ucq.loadUser(ctx, query, nodes, nil,
			func(n *UserCuriosity, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ucq.withCuriosity; query != nil {
		if err := ucq.loadCuriosity(ctx, query, nodes, nil,
			func(n *UserCuriosity, e *Curiosity) { n.Edges.Curiosity = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucq *UserCuriosityQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserCuriosity, init func(*UserCuriosity), assign func(*UserCuriosity, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserCuriosity)
	for i := range nodes {
		if nodes[i].user_id == nil {
			continue
		}
		fk := *nodes[i].user_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ucq *UserCuriosityQuery) loadCuriosity(ctx context.Context, query *CuriosityQuery, nodes []*UserCuriosity, init func(*UserCuriosity), assign func(*UserCuriosity, *Curiosity)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserCuriosity)
	for i := range nodes {
		if nodes[i].curiosity_id == nil {
			continue
		}
		fk := *nodes[i].curiosity_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(curiosity.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "curiosity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ucq *UserCuriosityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	_spec.Node.Columns = ucq.ctx.Fields
	if len(ucq.ctx.Fields) > 0 {
		_spec.Unique = ucq.ctx.Unique != nil && *ucq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserCuriosityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usercuriosity.Table, usercuriosity.Columns, sqlgraph.NewFieldSpec(usercuriosity.FieldID, field.TypeUUID))
	_spec.From = ucq.sql
	if unique := ucq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ucq.path != nil {
		_spec.Unique = true
	}
	if fields := ucq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercuriosity.FieldID)
		for i := range fields {
			if fields[i] != usercuriosity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserCuriosityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(usercuriosity.Table)
	columns := ucq.ctx.Fields
	if len(columns) == 0 {
		columns = usercuriosity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucq.ctx.Unique != nil && *ucq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCuriosityGroupBy is the group-by builder for UserCuriosity entities.
type UserCuriosityGroupBy struct {
	selector
	build *UserCuriosityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserCuriosityGroupBy) Aggregate(fns ...AggregateFunc) *UserCuriosityGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the selector query and scans the result into the given value.
func (ucgb *UserCuriosityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucgb.build.ctx, ent.OpQueryGroupBy)
	if err := ucgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCuriosityQuery, *UserCuriosityGroupBy](ctx, ucgb.build, ucgb, ucgb.build.inters, v)
}

func (ucgb *UserCuriosityGroupBy) sqlScan(ctx context.Context, root *UserCuriosityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ucgb.flds)+len(ucgb.fns))
		for _, f := range *ucgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ucgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserCuriositySelect is the builder for selecting fields of UserCuriosity entities.
type UserCuriositySelect struct {
	*UserCuriosityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ucs *UserCuriositySelect) Aggregate(fns ...AggregateFunc) *UserCuriositySelect {
	ucs.fns = append(ucs.fns, fns...)
	return ucs
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserCuriositySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucs.ctx, ent.OpQuerySelect)
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCuriosityQuery, *UserCuriositySelect](ctx, ucs.UserCuriosityQuery, ucs, ucs.inters, v)
}

func (ucs *UserCuriositySelect) sqlScan(ctx context.Context, root *UserCuriosityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ucs.fns))
	for _, fn := range ucs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ucs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
