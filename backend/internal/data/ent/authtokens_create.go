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
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

// AuthTokensCreate is the builder for creating a AuthTokens entity.
type AuthTokensCreate struct {
	config
	mutation *AuthTokensMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (atc *AuthTokensCreate) SetCreatedAt(t time.Time) *AuthTokensCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableCreatedAt(t *time.Time) *AuthTokensCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *AuthTokensCreate) SetUpdatedAt(t time.Time) *AuthTokensCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableUpdatedAt(t *time.Time) *AuthTokensCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetToken sets the "token" field.
func (atc *AuthTokensCreate) SetToken(b []byte) *AuthTokensCreate {
	atc.mutation.SetToken(b)
	return atc
}

// SetExpiresAt sets the "expires_at" field.
func (atc *AuthTokensCreate) SetExpiresAt(t time.Time) *AuthTokensCreate {
	atc.mutation.SetExpiresAt(t)
	return atc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableExpiresAt(t *time.Time) *AuthTokensCreate {
	if t != nil {
		atc.SetExpiresAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *AuthTokensCreate) SetID(u uuid.UUID) *AuthTokensCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableID(u *uuid.UUID) *AuthTokensCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atc *AuthTokensCreate) SetUserID(id uuid.UUID) *AuthTokensCreate {
	atc.mutation.SetUserID(id)
	return atc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableUserID(id *uuid.UUID) *AuthTokensCreate {
	if id != nil {
		atc = atc.SetUserID(*id)
	}
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *AuthTokensCreate) SetUser(u *User) *AuthTokensCreate {
	return atc.SetUserID(u.ID)
}

// SetRolesID sets the "roles" edge to the AuthRoles entity by ID.
func (atc *AuthTokensCreate) SetRolesID(id int) *AuthTokensCreate {
	atc.mutation.SetRolesID(id)
	return atc
}

// SetNillableRolesID sets the "roles" edge to the AuthRoles entity by ID if the given value is not nil.
func (atc *AuthTokensCreate) SetNillableRolesID(id *int) *AuthTokensCreate {
	if id != nil {
		atc = atc.SetRolesID(*id)
	}
	return atc
}

// SetRoles sets the "roles" edge to the AuthRoles entity.
func (atc *AuthTokensCreate) SetRoles(a *AuthRoles) *AuthTokensCreate {
	return atc.SetRolesID(a.ID)
}

// Mutation returns the AuthTokensMutation object of the builder.
func (atc *AuthTokensCreate) Mutation() *AuthTokensMutation {
	return atc.mutation
}

// Save creates the AuthTokens in the database.
func (atc *AuthTokensCreate) Save(ctx context.Context) (*AuthTokens, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AuthTokensCreate) SaveX(ctx context.Context) *AuthTokens {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AuthTokensCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AuthTokensCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AuthTokensCreate) defaults() {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := authtokens.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := authtokens.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ExpiresAt(); !ok {
		v := authtokens.DefaultExpiresAt()
		atc.mutation.SetExpiresAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := authtokens.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AuthTokensCreate) check() error {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthTokens.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AuthTokens.updated_at"`)}
	}
	if _, ok := atc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "AuthTokens.token"`)}
	}
	if _, ok := atc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "AuthTokens.expires_at"`)}
	}
	return nil
}

func (atc *AuthTokensCreate) sqlSave(ctx context.Context) (*AuthTokens, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
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
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AuthTokensCreate) createSpec() (*AuthTokens, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthTokens{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(authtokens.Table, sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID))
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(authtokens.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.SetField(authtokens.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := atc.mutation.Token(); ok {
		_spec.SetField(authtokens.FieldToken, field.TypeBytes, value)
		_node.Token = value
	}
	if value, ok := atc.mutation.ExpiresAt(); ok {
		_spec.SetField(authtokens.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authtokens.UserTable,
			Columns: []string{authtokens.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_auth_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   authtokens.RolesTable,
			Columns: []string{authtokens.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authroles.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthTokensCreateBulk is the builder for creating many AuthTokens entities in bulk.
type AuthTokensCreateBulk struct {
	config
	err      error
	builders []*AuthTokensCreate
}

// Save creates the AuthTokens entities in the database.
func (atcb *AuthTokensCreateBulk) Save(ctx context.Context) ([]*AuthTokens, error) {
	if atcb.err != nil {
		return nil, atcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AuthTokens, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthTokensMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AuthTokensCreateBulk) SaveX(ctx context.Context) []*AuthTokens {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AuthTokensCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AuthTokensCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
