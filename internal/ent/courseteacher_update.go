// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"online-teaching/internal/ent/course"
	"online-teaching/internal/ent/courseteacher"
	"online-teaching/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseTeacherUpdate is the builder for updating CourseTeacher entities.
type CourseTeacherUpdate struct {
	config
	hooks    []Hook
	mutation *CourseTeacherMutation
}

// Where appends a list predicates to the CourseTeacherUpdate builder.
func (ctu *CourseTeacherUpdate) Where(ps ...predicate.CourseTeacher) *CourseTeacherUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetName sets the "name" field.
func (ctu *CourseTeacherUpdate) SetName(s string) *CourseTeacherUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// SetDesc sets the "desc" field.
func (ctu *CourseTeacherUpdate) SetDesc(s string) *CourseTeacherUpdate {
	ctu.mutation.SetDesc(s)
	return ctu
}

// SetAvatar sets the "avatar" field.
func (ctu *CourseTeacherUpdate) SetAvatar(s string) *CourseTeacherUpdate {
	ctu.mutation.SetAvatar(s)
	return ctu
}

// SetTitle sets the "title" field.
func (ctu *CourseTeacherUpdate) SetTitle(s string) *CourseTeacherUpdate {
	ctu.mutation.SetTitle(s)
	return ctu
}

// SetIsDeleted sets the "is_deleted" field.
func (ctu *CourseTeacherUpdate) SetIsDeleted(b bool) *CourseTeacherUpdate {
	ctu.mutation.SetIsDeleted(b)
	return ctu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ctu *CourseTeacherUpdate) SetNillableIsDeleted(b *bool) *CourseTeacherUpdate {
	if b != nil {
		ctu.SetIsDeleted(*b)
	}
	return ctu
}

// SetCreatedAt sets the "created_at" field.
func (ctu *CourseTeacherUpdate) SetCreatedAt(t time.Time) *CourseTeacherUpdate {
	ctu.mutation.SetCreatedAt(t)
	return ctu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctu *CourseTeacherUpdate) SetNillableCreatedAt(t *time.Time) *CourseTeacherUpdate {
	if t != nil {
		ctu.SetCreatedAt(*t)
	}
	return ctu
}

// SetUpdatedAt sets the "updated_at" field.
func (ctu *CourseTeacherUpdate) SetUpdatedAt(t time.Time) *CourseTeacherUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// SetDeletedAt sets the "deleted_at" field.
func (ctu *CourseTeacherUpdate) SetDeletedAt(t time.Time) *CourseTeacherUpdate {
	ctu.mutation.SetDeletedAt(t)
	return ctu
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (ctu *CourseTeacherUpdate) AddCourseIDs(ids ...int) *CourseTeacherUpdate {
	ctu.mutation.AddCourseIDs(ids...)
	return ctu
}

// AddCourse adds the "course" edges to the Course entity.
func (ctu *CourseTeacherUpdate) AddCourse(c ...*Course) *CourseTeacherUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddCourseIDs(ids...)
}

// Mutation returns the CourseTeacherMutation object of the builder.
func (ctu *CourseTeacherUpdate) Mutation() *CourseTeacherMutation {
	return ctu.mutation
}

// ClearCourse clears all "course" edges to the Course entity.
func (ctu *CourseTeacherUpdate) ClearCourse() *CourseTeacherUpdate {
	ctu.mutation.ClearCourse()
	return ctu
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (ctu *CourseTeacherUpdate) RemoveCourseIDs(ids ...int) *CourseTeacherUpdate {
	ctu.mutation.RemoveCourseIDs(ids...)
	return ctu
}

// RemoveCourse removes "course" edges to Course entities.
func (ctu *CourseTeacherUpdate) RemoveCourse(c ...*Course) *CourseTeacherUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveCourseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CourseTeacherUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ctu.defaults()
	if len(ctu.hooks) == 0 {
		if err = ctu.check(); err != nil {
			return 0, err
		}
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseTeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctu.check(); err != nil {
				return 0, err
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CourseTeacherUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CourseTeacherUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CourseTeacherUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CourseTeacherUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := courseteacher.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *CourseTeacherUpdate) check() error {
	if v, ok := ctu.mutation.Name(); ok {
		if err := courseteacher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CourseTeacher.name": %w`, err)}
		}
	}
	return nil
}

func (ctu *CourseTeacherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseteacher.Table,
			Columns: courseteacher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseteacher.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldName,
		})
	}
	if value, ok := ctu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldDesc,
		})
	}
	if value, ok := ctu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldAvatar,
		})
	}
	if value, ok := ctu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldTitle,
		})
	}
	if value, ok := ctu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: courseteacher.FieldIsDeleted,
		})
	}
	if value, ok := ctu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldCreatedAt,
		})
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldUpdatedAt,
		})
	}
	if value, ok := ctu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldDeletedAt,
		})
	}
	if ctu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedCourseIDs(); len(nodes) > 0 && !ctu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseteacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CourseTeacherUpdateOne is the builder for updating a single CourseTeacher entity.
type CourseTeacherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseTeacherMutation
}

// SetName sets the "name" field.
func (ctuo *CourseTeacherUpdateOne) SetName(s string) *CourseTeacherUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// SetDesc sets the "desc" field.
func (ctuo *CourseTeacherUpdateOne) SetDesc(s string) *CourseTeacherUpdateOne {
	ctuo.mutation.SetDesc(s)
	return ctuo
}

// SetAvatar sets the "avatar" field.
func (ctuo *CourseTeacherUpdateOne) SetAvatar(s string) *CourseTeacherUpdateOne {
	ctuo.mutation.SetAvatar(s)
	return ctuo
}

// SetTitle sets the "title" field.
func (ctuo *CourseTeacherUpdateOne) SetTitle(s string) *CourseTeacherUpdateOne {
	ctuo.mutation.SetTitle(s)
	return ctuo
}

// SetIsDeleted sets the "is_deleted" field.
func (ctuo *CourseTeacherUpdateOne) SetIsDeleted(b bool) *CourseTeacherUpdateOne {
	ctuo.mutation.SetIsDeleted(b)
	return ctuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ctuo *CourseTeacherUpdateOne) SetNillableIsDeleted(b *bool) *CourseTeacherUpdateOne {
	if b != nil {
		ctuo.SetIsDeleted(*b)
	}
	return ctuo
}

// SetCreatedAt sets the "created_at" field.
func (ctuo *CourseTeacherUpdateOne) SetCreatedAt(t time.Time) *CourseTeacherUpdateOne {
	ctuo.mutation.SetCreatedAt(t)
	return ctuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctuo *CourseTeacherUpdateOne) SetNillableCreatedAt(t *time.Time) *CourseTeacherUpdateOne {
	if t != nil {
		ctuo.SetCreatedAt(*t)
	}
	return ctuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ctuo *CourseTeacherUpdateOne) SetUpdatedAt(t time.Time) *CourseTeacherUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ctuo *CourseTeacherUpdateOne) SetDeletedAt(t time.Time) *CourseTeacherUpdateOne {
	ctuo.mutation.SetDeletedAt(t)
	return ctuo
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (ctuo *CourseTeacherUpdateOne) AddCourseIDs(ids ...int) *CourseTeacherUpdateOne {
	ctuo.mutation.AddCourseIDs(ids...)
	return ctuo
}

// AddCourse adds the "course" edges to the Course entity.
func (ctuo *CourseTeacherUpdateOne) AddCourse(c ...*Course) *CourseTeacherUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddCourseIDs(ids...)
}

// Mutation returns the CourseTeacherMutation object of the builder.
func (ctuo *CourseTeacherUpdateOne) Mutation() *CourseTeacherMutation {
	return ctuo.mutation
}

// ClearCourse clears all "course" edges to the Course entity.
func (ctuo *CourseTeacherUpdateOne) ClearCourse() *CourseTeacherUpdateOne {
	ctuo.mutation.ClearCourse()
	return ctuo
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (ctuo *CourseTeacherUpdateOne) RemoveCourseIDs(ids ...int) *CourseTeacherUpdateOne {
	ctuo.mutation.RemoveCourseIDs(ids...)
	return ctuo
}

// RemoveCourse removes "course" edges to Course entities.
func (ctuo *CourseTeacherUpdateOne) RemoveCourse(c ...*Course) *CourseTeacherUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveCourseIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CourseTeacherUpdateOne) Select(field string, fields ...string) *CourseTeacherUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CourseTeacher entity.
func (ctuo *CourseTeacherUpdateOne) Save(ctx context.Context) (*CourseTeacher, error) {
	var (
		err  error
		node *CourseTeacher
	)
	ctuo.defaults()
	if len(ctuo.hooks) == 0 {
		if err = ctuo.check(); err != nil {
			return nil, err
		}
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseTeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctuo.check(); err != nil {
				return nil, err
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CourseTeacher)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CourseTeacherMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CourseTeacherUpdateOne) SaveX(ctx context.Context) *CourseTeacher {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CourseTeacherUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CourseTeacherUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CourseTeacherUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := courseteacher.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *CourseTeacherUpdateOne) check() error {
	if v, ok := ctuo.mutation.Name(); ok {
		if err := courseteacher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CourseTeacher.name": %w`, err)}
		}
	}
	return nil
}

func (ctuo *CourseTeacherUpdateOne) sqlSave(ctx context.Context) (_node *CourseTeacher, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseteacher.Table,
			Columns: courseteacher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseteacher.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseTeacher.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, courseteacher.FieldID)
		for _, f := range fields {
			if !courseteacher.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != courseteacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldName,
		})
	}
	if value, ok := ctuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldDesc,
		})
	}
	if value, ok := ctuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldAvatar,
		})
	}
	if value, ok := ctuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseteacher.FieldTitle,
		})
	}
	if value, ok := ctuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: courseteacher.FieldIsDeleted,
		})
	}
	if value, ok := ctuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldCreatedAt,
		})
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldUpdatedAt,
		})
	}
	if value, ok := ctuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseteacher.FieldDeletedAt,
		})
	}
	if ctuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedCourseIDs(); len(nodes) > 0 && !ctuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   courseteacher.CourseTable,
			Columns: []string{courseteacher.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CourseTeacher{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseteacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}