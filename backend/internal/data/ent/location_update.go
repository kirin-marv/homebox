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
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LocationUpdate) SetUpdatedAt(t time.Time) *LocationUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetName sets the "name" field.
func (lu *LocationUpdate) SetName(s string) *LocationUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableName(s *string) *LocationUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// SetDescription sets the "description" field.
func (lu *LocationUpdate) SetDescription(s string) *LocationUpdate {
	lu.mutation.SetDescription(s)
	return lu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableDescription(s *string) *LocationUpdate {
	if s != nil {
		lu.SetDescription(*s)
	}
	return lu
}

// ClearDescription clears the value of the "description" field.
func (lu *LocationUpdate) ClearDescription() *LocationUpdate {
	lu.mutation.ClearDescription()
	return lu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (lu *LocationUpdate) SetGroupID(id uuid.UUID) *LocationUpdate {
	lu.mutation.SetGroupID(id)
	return lu
}

// SetGroup sets the "group" edge to the Group entity.
func (lu *LocationUpdate) SetGroup(g *Group) *LocationUpdate {
	return lu.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the Location entity by ID.
func (lu *LocationUpdate) SetParentID(id uuid.UUID) *LocationUpdate {
	lu.mutation.SetParentID(id)
	return lu
}

// SetNillableParentID sets the "parent" edge to the Location entity by ID if the given value is not nil.
func (lu *LocationUpdate) SetNillableParentID(id *uuid.UUID) *LocationUpdate {
	if id != nil {
		lu = lu.SetParentID(*id)
	}
	return lu
}

// SetParent sets the "parent" edge to the Location entity.
func (lu *LocationUpdate) SetParent(l *Location) *LocationUpdate {
	return lu.SetParentID(l.ID)
}

// AddChildIDs adds the "children" edge to the Location entity by IDs.
func (lu *LocationUpdate) AddChildIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.AddChildIDs(ids...)
	return lu
}

// AddChildren adds the "children" edges to the Location entity.
func (lu *LocationUpdate) AddChildren(l ...*Location) *LocationUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.AddChildIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (lu *LocationUpdate) AddItemIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.AddItemIDs(ids...)
	return lu
}

// AddItems adds the "items" edges to the Item entity.
func (lu *LocationUpdate) AddItems(i ...*Item) *LocationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return lu.AddItemIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (lu *LocationUpdate) ClearGroup() *LocationUpdate {
	lu.mutation.ClearGroup()
	return lu
}

// ClearParent clears the "parent" edge to the Location entity.
func (lu *LocationUpdate) ClearParent() *LocationUpdate {
	lu.mutation.ClearParent()
	return lu
}

// ClearChildren clears all "children" edges to the Location entity.
func (lu *LocationUpdate) ClearChildren() *LocationUpdate {
	lu.mutation.ClearChildren()
	return lu
}

// RemoveChildIDs removes the "children" edge to Location entities by IDs.
func (lu *LocationUpdate) RemoveChildIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.RemoveChildIDs(ids...)
	return lu
}

// RemoveChildren removes "children" edges to Location entities.
func (lu *LocationUpdate) RemoveChildren(l ...*Location) *LocationUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.RemoveChildIDs(ids...)
}

// ClearItems clears all "items" edges to the Item entity.
func (lu *LocationUpdate) ClearItems() *LocationUpdate {
	lu.mutation.ClearItems()
	return lu
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (lu *LocationUpdate) RemoveItemIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.RemoveItemIDs(ids...)
	return lu
}

// RemoveItems removes "items" edges to Item entities.
func (lu *LocationUpdate) RemoveItems(i ...*Item) *LocationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return lu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LocationUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := location.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LocationUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := location.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Location.name": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if _, ok := lu.mutation.GroupID(); lu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Location.group"`)
	}
	return nil
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
	}
	if lu.mutation.DescriptionCleared() {
		_spec.ClearField(location.FieldDescription, field.TypeString)
	}
	if lu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.GroupTable,
			Columns: []string{location.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.GroupTable,
			Columns: []string{location.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !lu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !lu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LocationUpdateOne) SetUpdatedAt(t time.Time) *LocationUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetName sets the "name" field.
func (luo *LocationUpdateOne) SetName(s string) *LocationUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableName(s *string) *LocationUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// SetDescription sets the "description" field.
func (luo *LocationUpdateOne) SetDescription(s string) *LocationUpdateOne {
	luo.mutation.SetDescription(s)
	return luo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableDescription(s *string) *LocationUpdateOne {
	if s != nil {
		luo.SetDescription(*s)
	}
	return luo
}

// ClearDescription clears the value of the "description" field.
func (luo *LocationUpdateOne) ClearDescription() *LocationUpdateOne {
	luo.mutation.ClearDescription()
	return luo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (luo *LocationUpdateOne) SetGroupID(id uuid.UUID) *LocationUpdateOne {
	luo.mutation.SetGroupID(id)
	return luo
}

// SetGroup sets the "group" edge to the Group entity.
func (luo *LocationUpdateOne) SetGroup(g *Group) *LocationUpdateOne {
	return luo.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the Location entity by ID.
func (luo *LocationUpdateOne) SetParentID(id uuid.UUID) *LocationUpdateOne {
	luo.mutation.SetParentID(id)
	return luo
}

// SetNillableParentID sets the "parent" edge to the Location entity by ID if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableParentID(id *uuid.UUID) *LocationUpdateOne {
	if id != nil {
		luo = luo.SetParentID(*id)
	}
	return luo
}

// SetParent sets the "parent" edge to the Location entity.
func (luo *LocationUpdateOne) SetParent(l *Location) *LocationUpdateOne {
	return luo.SetParentID(l.ID)
}

// AddChildIDs adds the "children" edge to the Location entity by IDs.
func (luo *LocationUpdateOne) AddChildIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.AddChildIDs(ids...)
	return luo
}

// AddChildren adds the "children" edges to the Location entity.
func (luo *LocationUpdateOne) AddChildren(l ...*Location) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.AddChildIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (luo *LocationUpdateOne) AddItemIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.AddItemIDs(ids...)
	return luo
}

// AddItems adds the "items" edges to the Item entity.
func (luo *LocationUpdateOne) AddItems(i ...*Item) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return luo.AddItemIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (luo *LocationUpdateOne) ClearGroup() *LocationUpdateOne {
	luo.mutation.ClearGroup()
	return luo
}

// ClearParent clears the "parent" edge to the Location entity.
func (luo *LocationUpdateOne) ClearParent() *LocationUpdateOne {
	luo.mutation.ClearParent()
	return luo
}

// ClearChildren clears all "children" edges to the Location entity.
func (luo *LocationUpdateOne) ClearChildren() *LocationUpdateOne {
	luo.mutation.ClearChildren()
	return luo
}

// RemoveChildIDs removes the "children" edge to Location entities by IDs.
func (luo *LocationUpdateOne) RemoveChildIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.RemoveChildIDs(ids...)
	return luo
}

// RemoveChildren removes "children" edges to Location entities.
func (luo *LocationUpdateOne) RemoveChildren(l ...*Location) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.RemoveChildIDs(ids...)
}

// ClearItems clears all "items" edges to the Item entity.
func (luo *LocationUpdateOne) ClearItems() *LocationUpdateOne {
	luo.mutation.ClearItems()
	return luo
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (luo *LocationUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.RemoveItemIDs(ids...)
	return luo
}

// RemoveItems removes "items" edges to Item entities.
func (luo *LocationUpdateOne) RemoveItems(i ...*Item) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return luo.RemoveItemIDs(ids...)
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LocationUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := location.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LocationUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := location.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Location.name": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if _, ok := luo.mutation.GroupID(); luo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Location.group"`)
	}
	return nil
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
	}
	if luo.mutation.DescriptionCleared() {
		_spec.ClearField(location.FieldDescription, field.TypeString)
	}
	if luo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.GroupTable,
			Columns: []string{location.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.GroupTable,
			Columns: []string{location.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !luo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !luo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
