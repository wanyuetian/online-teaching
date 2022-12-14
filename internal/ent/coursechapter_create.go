// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"online-teaching/internal/ent/coursechapter"
	"online-teaching/internal/ent/coursesection"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseChapterCreate is the builder for creating a CourseChapter entity.
type CourseChapterCreate struct {
	config
	mutation *CourseChapterMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (ccc *CourseChapterCreate) SetTitle(s string) *CourseChapterCreate {
	ccc.mutation.SetTitle(s)
	return ccc
}

// SetIsDeleted sets the "is_deleted" field.
func (ccc *CourseChapterCreate) SetIsDeleted(b bool) *CourseChapterCreate {
	ccc.mutation.SetIsDeleted(b)
	return ccc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ccc *CourseChapterCreate) SetNillableIsDeleted(b *bool) *CourseChapterCreate {
	if b != nil {
		ccc.SetIsDeleted(*b)
	}
	return ccc
}

// SetCreatedAt sets the "created_at" field.
func (ccc *CourseChapterCreate) SetCreatedAt(t time.Time) *CourseChapterCreate {
	ccc.mutation.SetCreatedAt(t)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *CourseChapterCreate) SetNillableCreatedAt(t *time.Time) *CourseChapterCreate {
	if t != nil {
		ccc.SetCreatedAt(*t)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *CourseChapterCreate) SetUpdatedAt(t time.Time) *CourseChapterCreate {
	ccc.mutation.SetUpdatedAt(t)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *CourseChapterCreate) SetNillableUpdatedAt(t *time.Time) *CourseChapterCreate {
	if t != nil {
		ccc.SetUpdatedAt(*t)
	}
	return ccc
}

// AddCourseSectionIDs adds the "course_section" edge to the CourseSection entity by IDs.
func (ccc *CourseChapterCreate) AddCourseSectionIDs(ids ...int) *CourseChapterCreate {
	ccc.mutation.AddCourseSectionIDs(ids...)
	return ccc
}

// AddCourseSection adds the "course_section" edges to the CourseSection entity.
func (ccc *CourseChapterCreate) AddCourseSection(c ...*CourseSection) *CourseChapterCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccc.AddCourseSectionIDs(ids...)
}

// Mutation returns the CourseChapterMutation object of the builder.
func (ccc *CourseChapterCreate) Mutation() *CourseChapterMutation {
	return ccc.mutation
}

// Save creates the CourseChapter in the database.
func (ccc *CourseChapterCreate) Save(ctx context.Context) (*CourseChapter, error) {
	var (
		err  error
		node *CourseChapter
	)
	ccc.defaults()
	if len(ccc.hooks) == 0 {
		if err = ccc.check(); err != nil {
			return nil, err
		}
		node, err = ccc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseChapterMutation)
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
		nv, ok := v.(*CourseChapter)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CourseChapterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *CourseChapterCreate) SaveX(ctx context.Context) *CourseChapter {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CourseChapterCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CourseChapterCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *CourseChapterCreate) defaults() {
	if _, ok := ccc.mutation.IsDeleted(); !ok {
		v := coursechapter.DefaultIsDeleted
		ccc.mutation.SetIsDeleted(v)
	}
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		v := coursechapter.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		v := coursechapter.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CourseChapterCreate) check() error {
	if _, ok := ccc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "CourseChapter.title"`)}
	}
	if v, ok := ccc.mutation.Title(); ok {
		if err := coursechapter.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CourseChapter.title": %w`, err)}
		}
	}
	if _, ok := ccc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "CourseChapter.is_deleted"`)}
	}
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CourseChapter.created_at"`)}
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CourseChapter.updated_at"`)}
	}
	return nil
}

func (ccc *CourseChapterCreate) sqlSave(ctx context.Context) (*CourseChapter, error) {
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

func (ccc *CourseChapterCreate) createSpec() (*CourseChapter, *sqlgraph.CreateSpec) {
	var (
		_node = &CourseChapter{config: ccc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coursechapter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coursechapter.FieldID,
			},
		}
	)
	if value, ok := ccc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursechapter.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ccc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coursechapter.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ccc.mutation.CourseSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coursechapter.CourseSectionTable,
			Columns: []string{coursechapter.CourseSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coursesection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CourseChapterCreateBulk is the builder for creating many CourseChapter entities in bulk.
type CourseChapterCreateBulk struct {
	config
	builders []*CourseChapterCreate
}

// Save creates the CourseChapter entities in the database.
func (cccb *CourseChapterCreateBulk) Save(ctx context.Context) ([]*CourseChapter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CourseChapter, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseChapterMutation)
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
func (cccb *CourseChapterCreateBulk) SaveX(ctx context.Context) []*CourseChapter {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CourseChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CourseChapterCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}
