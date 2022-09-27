// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"online-teaching/internal/ent/courseinfo"
	"online-teaching/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseInfoDelete is the builder for deleting a CourseInfo entity.
type CourseInfoDelete struct {
	config
	hooks    []Hook
	mutation *CourseInfoMutation
}

// Where appends a list predicates to the CourseInfoDelete builder.
func (cid *CourseInfoDelete) Where(ps ...predicate.CourseInfo) *CourseInfoDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CourseInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cid.hooks) == 0 {
		affected, err = cid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cid.mutation = mutation
			affected, err = cid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cid.hooks) - 1; i >= 0; i-- {
			if cid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CourseInfoDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CourseInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: courseinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseinfo.FieldID,
			},
		},
	}
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CourseInfoDeleteOne is the builder for deleting a single CourseInfo entity.
type CourseInfoDeleteOne struct {
	cid *CourseInfoDelete
}

// Exec executes the deletion query.
func (cido *CourseInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{courseinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CourseInfoDeleteOne) ExecX(ctx context.Context) {
	cido.cid.ExecX(ctx)
}
