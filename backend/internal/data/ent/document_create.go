// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
)

// DocumentCreate is the builder for creating a Document entity.
type DocumentCreate struct {
	config
	mutation *DocumentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DocumentCreate) SetCreatedAt(t time.Time) *DocumentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DocumentCreate) SetNillableCreatedAt(t *time.Time) *DocumentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DocumentCreate) SetUpdatedAt(t time.Time) *DocumentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DocumentCreate) SetNillableUpdatedAt(t *time.Time) *DocumentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetTitle sets the "title" field.
func (dc *DocumentCreate) SetTitle(s string) *DocumentCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetPath sets the "path" field.
func (dc *DocumentCreate) SetPath(s string) *DocumentCreate {
	dc.mutation.SetPath(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DocumentCreate) SetID(u uuid.UUID) *DocumentCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DocumentCreate) SetNillableID(u *uuid.UUID) *DocumentCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (dc *DocumentCreate) SetGroupID(id uuid.UUID) *DocumentCreate {
	dc.mutation.SetGroupID(id)
	return dc
}

// SetGroup sets the "group" edge to the Group entity.
func (dc *DocumentCreate) SetGroup(g *Group) *DocumentCreate {
	return dc.SetGroupID(g.ID)
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (dc *DocumentCreate) AddAttachmentIDs(ids ...uuid.UUID) *DocumentCreate {
	dc.mutation.AddAttachmentIDs(ids...)
	return dc
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (dc *DocumentCreate) AddAttachments(a ...*Attachment) *DocumentCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddAttachmentIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (dc *DocumentCreate) Mutation() *DocumentMutation {
	return dc.mutation
}

// Save creates the Document in the database.
func (dc *DocumentCreate) Save(ctx context.Context) (*Document, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DocumentCreate) SaveX(ctx context.Context) *Document {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DocumentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DocumentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DocumentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := document.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := document.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := document.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DocumentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Document.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Document.updated_at"`)}
	}
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Document.title"`)}
	}
	if v, ok := dc.mutation.Title(); ok {
		if err := document.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Document.title": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Document.path"`)}
	}
	if v, ok := dc.mutation.Path(); ok {
		if err := document.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Document.path": %w`, err)}
		}
	}
	if _, ok := dc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Document.group"`)}
	}
	return nil
}

func (dc *DocumentCreate) sqlSave(ctx context.Context) (*Document, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DocumentCreate) createSpec() (*Document, *sqlgraph.CreateSpec) {
	var (
		_node = &Document{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(document.Table, sqlgraph.NewFieldSpec(document.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(document.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(document.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Title(); ok {
		_spec.SetField(document.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := dc.mutation.Path(); ok {
		_spec.SetField(document.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := dc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   document.GroupTable,
			Columns: []string{document.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_documents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AttachmentsTable,
			Columns: []string{document.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DocumentCreateBulk is the builder for creating many Document entities in bulk.
type DocumentCreateBulk struct {
	config
	err      error
	builders []*DocumentCreate
}

// Save creates the Document entities in the database.
func (dcb *DocumentCreateBulk) Save(ctx context.Context) ([]*Document, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Document, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocumentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DocumentCreateBulk) SaveX(ctx context.Context) []*Document {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DocumentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DocumentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
