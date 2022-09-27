// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"online-teaching/internal/ent/coursechapter"
	"online-teaching/internal/ent/coursesection"
	"online-teaching/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseChapterUpdate is the builder for updating CourseChapter entities.
type CourseChapterUpdate struct {
	config
	hooks    []Hook
	mutation *CourseChapterMutation
}

// Where appends a list predicates to the CourseChapterUpdate builder.
func (ccu *CourseChapterUpdate) Where(ps ...predicate.CourseChapter) *CourseChapterUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetTitle sets the "title" field.
func (ccu *CourseChapterUpdate) SetTitle(s string) *CourseChapterUpdate {
	ccu.mutation.SetTitle(s)
	return ccu
}

// SetIsDeleted sets the "is_deleted" field.
func (ccu *CourseChapterUpdate) SetIsDeleted(b bool) *CourseChapterUpdate {
	ccu.mutation.SetIsDeleted(b)
	return ccu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ccu *CourseChapterUpdate) SetNillableIsDeleted(b *bool) *CourseChapterUpdate {
	if b != nil {
		ccu.SetIsDeleted(*b)
	}
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *CourseChapterUpdate) SetCreatedAt(t time.Time) *CourseChapterUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *CourseChapterUpdate) SetNillableCreatedAt(t *time.Time) *CourseChapterUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *CourseChapterUpdate) SetUpdatedAt(t time.Time) *CourseChapterUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// SetDeletedAt sets the "deleted_at" field.
func (ccu *CourseChapterUpdate) SetDeletedAt(t time.Time) *CourseChapterUpdate {
	ccu.mutation.SetDeletedAt(t)
	return ccu
}

// AddCourseSectionIDs adds the "course_section" edge to the CourseSection entity by IDs.
func (ccu *CourseChapterUpdate) AddCourseSectionIDs(ids ...int) *CourseChapterUpdate {
	ccu.mutation.AddCourseSectionIDs(ids...)
	return ccu
}

// AddCourseSection adds the "course_section" edges to the CourseSection entity.
func (ccu *CourseChapterUpdate) AddCourseSection(c ...*CourseSection) *CourseChapterUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccu.AddCourseSectionIDs(ids...)
}

// Mutation returns the CourseChapterMutation object of the builder.
func (ccu *CourseChapterUpdate) Mutation() *CourseChapterMutation {
	return ccu.mutation
}

// ClearCourseSection clears all "course_section" edges to the CourseSection entity.
func (ccu *CourseChapterUpdate) ClearCourseSection() *CourseChapterUpdate {
	ccu.mutation.ClearCourseSection()
	return ccu
}

// RemoveCourseSectionIDs removes the "course_section" edge to CourseSection entities by IDs.
func (ccu *CourseChapterUpdate) RemoveCourseSectionIDs(ids ...int) *CourseChapterUpdate {
	ccu.mutation.RemoveCourseSectionIDs(ids...)
	return ccu
}

// RemoveCourseSection removes "course_section" edges to CourseSection entities.
func (ccu *CourseChapterUpdate) RemoveCourseSection(c ...*CourseSection) *CourseChapterUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccu.RemoveCourseSectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CourseChapterUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ccu.defaults()
	if len(ccu.hooks) == 0 {
		if err = ccu.check(); err != nil {
			return 0, err
		}
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseChapterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccu.check(); err != nil {
				return 0, err
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			if ccu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CourseChapterUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CourseChapterUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CourseChapterUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *CourseChapterUpdate) defaults() {
	if _, ok := ccu.mutation.UpdatedAt(); !ok {
		v := coursechapter.UpdateDefaultUpdatedAt()
		ccu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *CourseChapterUpdate) check() error {
	if v, ok := ccu.mutation.Title(); ok {
		if err := coursechapter.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CourseChapter.title": %w`, err)}
		}
	}
	return nil
}

func (ccu *CourseChapterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coursechapter.Table,
			Columns: coursechapter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coursechapter.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursechapter.FieldTitle,
		})
	}
	if value, ok := ccu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coursechapter.FieldIsDeleted,
		})
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldUpdatedAt,
		})
	}
	if value, ok := ccu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldDeletedAt,
		})
	}
	if ccu.mutation.CourseSectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.RemovedCourseSectionIDs(); len(nodes) > 0 && !ccu.mutation.CourseSectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.CourseSectionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursechapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CourseChapterUpdateOne is the builder for updating a single CourseChapter entity.
type CourseChapterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseChapterMutation
}

// SetTitle sets the "title" field.
func (ccuo *CourseChapterUpdateOne) SetTitle(s string) *CourseChapterUpdateOne {
	ccuo.mutation.SetTitle(s)
	return ccuo
}

// SetIsDeleted sets the "is_deleted" field.
func (ccuo *CourseChapterUpdateOne) SetIsDeleted(b bool) *CourseChapterUpdateOne {
	ccuo.mutation.SetIsDeleted(b)
	return ccuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ccuo *CourseChapterUpdateOne) SetNillableIsDeleted(b *bool) *CourseChapterUpdateOne {
	if b != nil {
		ccuo.SetIsDeleted(*b)
	}
	return ccuo
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *CourseChapterUpdateOne) SetCreatedAt(t time.Time) *CourseChapterUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *CourseChapterUpdateOne) SetNillableCreatedAt(t *time.Time) *CourseChapterUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *CourseChapterUpdateOne) SetUpdatedAt(t time.Time) *CourseChapterUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ccuo *CourseChapterUpdateOne) SetDeletedAt(t time.Time) *CourseChapterUpdateOne {
	ccuo.mutation.SetDeletedAt(t)
	return ccuo
}

// AddCourseSectionIDs adds the "course_section" edge to the CourseSection entity by IDs.
func (ccuo *CourseChapterUpdateOne) AddCourseSectionIDs(ids ...int) *CourseChapterUpdateOne {
	ccuo.mutation.AddCourseSectionIDs(ids...)
	return ccuo
}

// AddCourseSection adds the "course_section" edges to the CourseSection entity.
func (ccuo *CourseChapterUpdateOne) AddCourseSection(c ...*CourseSection) *CourseChapterUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccuo.AddCourseSectionIDs(ids...)
}

// Mutation returns the CourseChapterMutation object of the builder.
func (ccuo *CourseChapterUpdateOne) Mutation() *CourseChapterMutation {
	return ccuo.mutation
}

// ClearCourseSection clears all "course_section" edges to the CourseSection entity.
func (ccuo *CourseChapterUpdateOne) ClearCourseSection() *CourseChapterUpdateOne {
	ccuo.mutation.ClearCourseSection()
	return ccuo
}

// RemoveCourseSectionIDs removes the "course_section" edge to CourseSection entities by IDs.
func (ccuo *CourseChapterUpdateOne) RemoveCourseSectionIDs(ids ...int) *CourseChapterUpdateOne {
	ccuo.mutation.RemoveCourseSectionIDs(ids...)
	return ccuo
}

// RemoveCourseSection removes "course_section" edges to CourseSection entities.
func (ccuo *CourseChapterUpdateOne) RemoveCourseSection(c ...*CourseSection) *CourseChapterUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccuo.RemoveCourseSectionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CourseChapterUpdateOne) Select(field string, fields ...string) *CourseChapterUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CourseChapter entity.
func (ccuo *CourseChapterUpdateOne) Save(ctx context.Context) (*CourseChapter, error) {
	var (
		err  error
		node *CourseChapter
	)
	ccuo.defaults()
	if len(ccuo.hooks) == 0 {
		if err = ccuo.check(); err != nil {
			return nil, err
		}
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseChapterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccuo.check(); err != nil {
				return nil, err
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			if ccuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CourseChapterUpdateOne) SaveX(ctx context.Context) *CourseChapter {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CourseChapterUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CourseChapterUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *CourseChapterUpdateOne) defaults() {
	if _, ok := ccuo.mutation.UpdatedAt(); !ok {
		v := coursechapter.UpdateDefaultUpdatedAt()
		ccuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *CourseChapterUpdateOne) check() error {
	if v, ok := ccuo.mutation.Title(); ok {
		if err := coursechapter.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CourseChapter.title": %w`, err)}
		}
	}
	return nil
}

func (ccuo *CourseChapterUpdateOne) sqlSave(ctx context.Context) (_node *CourseChapter, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coursechapter.Table,
			Columns: coursechapter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coursechapter.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseChapter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coursechapter.FieldID)
		for _, f := range fields {
			if !coursechapter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coursechapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursechapter.FieldTitle,
		})
	}
	if value, ok := ccuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coursechapter.FieldIsDeleted,
		})
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldUpdatedAt,
		})
	}
	if value, ok := ccuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursechapter.FieldDeletedAt,
		})
	}
	if ccuo.mutation.CourseSectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.RemovedCourseSectionIDs(); len(nodes) > 0 && !ccuo.mutation.CourseSectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.CourseSectionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CourseChapter{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursechapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}