// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"online-teaching/internal/ent/coursecomment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseCommentCreate is the builder for creating a CourseComment entity.
type CourseCommentCreate struct {
	config
	mutation *CourseCommentMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (ccc *CourseCommentCreate) SetUsername(s string) *CourseCommentCreate {
	ccc.mutation.SetUsername(s)
	return ccc
}

// SetComment sets the "comment" field.
func (ccc *CourseCommentCreate) SetComment(s string) *CourseCommentCreate {
	ccc.mutation.SetComment(s)
	return ccc
}

// SetOrder sets the "order" field.
func (ccc *CourseCommentCreate) SetOrder(i int) *CourseCommentCreate {
	ccc.mutation.SetOrder(i)
	return ccc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (ccc *CourseCommentCreate) SetNillableOrder(i *int) *CourseCommentCreate {
	if i != nil {
		ccc.SetOrder(*i)
	}
	return ccc
}

// SetIsShow sets the "is_show" field.
func (ccc *CourseCommentCreate) SetIsShow(b bool) *CourseCommentCreate {
	ccc.mutation.SetIsShow(b)
	return ccc
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (ccc *CourseCommentCreate) SetNillableIsShow(b *bool) *CourseCommentCreate {
	if b != nil {
		ccc.SetIsShow(*b)
	}
	return ccc
}

// SetIsDeleted sets the "is_deleted" field.
func (ccc *CourseCommentCreate) SetIsDeleted(b bool) *CourseCommentCreate {
	ccc.mutation.SetIsDeleted(b)
	return ccc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ccc *CourseCommentCreate) SetNillableIsDeleted(b *bool) *CourseCommentCreate {
	if b != nil {
		ccc.SetIsDeleted(*b)
	}
	return ccc
}

// SetCreatedAt sets the "created_at" field.
func (ccc *CourseCommentCreate) SetCreatedAt(t time.Time) *CourseCommentCreate {
	ccc.mutation.SetCreatedAt(t)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *CourseCommentCreate) SetNillableCreatedAt(t *time.Time) *CourseCommentCreate {
	if t != nil {
		ccc.SetCreatedAt(*t)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *CourseCommentCreate) SetUpdatedAt(t time.Time) *CourseCommentCreate {
	ccc.mutation.SetUpdatedAt(t)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *CourseCommentCreate) SetNillableUpdatedAt(t *time.Time) *CourseCommentCreate {
	if t != nil {
		ccc.SetUpdatedAt(*t)
	}
	return ccc
}

// Mutation returns the CourseCommentMutation object of the builder.
func (ccc *CourseCommentCreate) Mutation() *CourseCommentMutation {
	return ccc.mutation
}

// Save creates the CourseComment in the database.
func (ccc *CourseCommentCreate) Save(ctx context.Context) (*CourseComment, error) {
	var (
		err  error
		node *CourseComment
	)
	ccc.defaults()
	if len(ccc.hooks) == 0 {
		if err = ccc.check(); err != nil {
			return nil, err
		}
		node, err = ccc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccc.check(); err != nil {
				return nil, err
			}
			ccc.mutation = mutation
			if node, err = ccc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ccc.hooks) - 1; i >= 0; i-- {
			if ccc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CourseComment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CourseCommentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *CourseCommentCreate) SaveX(ctx context.Context) *CourseComment {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CourseCommentCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CourseCommentCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *CourseCommentCreate) defaults() {
	if _, ok := ccc.mutation.Order(); !ok {
		v := coursecomment.DefaultOrder
		ccc.mutation.SetOrder(v)
	}
	if _, ok := ccc.mutation.IsShow(); !ok {
		v := coursecomment.DefaultIsShow
		ccc.mutation.SetIsShow(v)
	}
	if _, ok := ccc.mutation.IsDeleted(); !ok {
		v := coursecomment.DefaultIsDeleted
		ccc.mutation.SetIsDeleted(v)
	}
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		v := coursecomment.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		v := coursecomment.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CourseCommentCreate) check() error {
	if _, ok := ccc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "CourseComment.username"`)}
	}
	if _, ok := ccc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "CourseComment.comment"`)}
	}
	if _, ok := ccc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "CourseComment.order"`)}
	}
	if _, ok := ccc.mutation.IsShow(); !ok {
		return &ValidationError{Name: "is_show", err: errors.New(`ent: missing required field "CourseComment.is_show"`)}
	}
	if _, ok := ccc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "CourseComment.is_deleted"`)}
	}
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CourseComment.created_at"`)}
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CourseComment.updated_at"`)}
	}
	return nil
}

func (ccc *CourseCommentCreate) sqlSave(ctx context.Context) (*CourseComment, error) {
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ccc *CourseCommentCreate) createSpec() (*CourseComment, *sqlgraph.CreateSpec) {
	var (
		_node = &CourseComment{config: ccc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coursecomment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coursecomment.FieldID,
			},
		}
	)
	if value, ok := ccc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursecomment.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ccc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursecomment.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := ccc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursecomment.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := ccc.mutation.IsShow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coursecomment.FieldIsShow,
		})
		_node.IsShow = value
	}
	if value, ok := ccc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coursecomment.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursecomment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursecomment.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// CourseCommentCreateBulk is the builder for creating many CourseComment entities in bulk.
type CourseCommentCreateBulk struct {
	config
	builders []*CourseCommentCreate
}

// Save creates the CourseComment entities in the database.
func (cccb *CourseCommentCreateBulk) Save(ctx context.Context) ([]*CourseComment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CourseComment, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseCommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *CourseCommentCreateBulk) SaveX(ctx context.Context) []*CourseComment {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CourseCommentCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CourseCommentCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}
