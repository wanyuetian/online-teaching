// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"online-teaching/internal/ent/courseinfo"
	"online-teaching/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseInfoUpdate is the builder for updating CourseInfo entities.
type CourseInfoUpdate struct {
	config
	hooks    []Hook
	mutation *CourseInfoMutation
}

// Where appends a list predicates to the CourseInfoUpdate builder.
func (ciu *CourseInfoUpdate) Where(ps ...predicate.CourseInfo) *CourseInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetClickCount sets the "click_count" field.
func (ciu *CourseInfoUpdate) SetClickCount(i int) *CourseInfoUpdate {
	ciu.mutation.ResetClickCount()
	ciu.mutation.SetClickCount(i)
	return ciu
}

// AddClickCount adds i to the "click_count" field.
func (ciu *CourseInfoUpdate) AddClickCount(i int) *CourseInfoUpdate {
	ciu.mutation.AddClickCount(i)
	return ciu
}

// SetLearnCount sets the "learn_count" field.
func (ciu *CourseInfoUpdate) SetLearnCount(i int) *CourseInfoUpdate {
	ciu.mutation.ResetLearnCount()
	ciu.mutation.SetLearnCount(i)
	return ciu
}

// AddLearnCount adds i to the "learn_count" field.
func (ciu *CourseInfoUpdate) AddLearnCount(i int) *CourseInfoUpdate {
	ciu.mutation.AddLearnCount(i)
	return ciu
}

// SetTotalDuration sets the "total_duration" field.
func (ciu *CourseInfoUpdate) SetTotalDuration(i int) *CourseInfoUpdate {
	ciu.mutation.ResetTotalDuration()
	ciu.mutation.SetTotalDuration(i)
	return ciu
}

// AddTotalDuration adds i to the "total_duration" field.
func (ciu *CourseInfoUpdate) AddTotalDuration(i int) *CourseInfoUpdate {
	ciu.mutation.AddTotalDuration(i)
	return ciu
}

// SetSectionCount sets the "section_count" field.
func (ciu *CourseInfoUpdate) SetSectionCount(i int) *CourseInfoUpdate {
	ciu.mutation.ResetSectionCount()
	ciu.mutation.SetSectionCount(i)
	return ciu
}

// AddSectionCount adds i to the "section_count" field.
func (ciu *CourseInfoUpdate) AddSectionCount(i int) *CourseInfoUpdate {
	ciu.mutation.AddSectionCount(i)
	return ciu
}

// SetPrice sets the "price" field.
func (ciu *CourseInfoUpdate) SetPrice(f float64) *CourseInfoUpdate {
	ciu.mutation.ResetPrice()
	ciu.mutation.SetPrice(f)
	return ciu
}

// AddPrice adds f to the "price" field.
func (ciu *CourseInfoUpdate) AddPrice(f float64) *CourseInfoUpdate {
	ciu.mutation.AddPrice(f)
	return ciu
}

// SetDetail sets the "detail" field.
func (ciu *CourseInfoUpdate) SetDetail(s string) *CourseInfoUpdate {
	ciu.mutation.SetDetail(s)
	return ciu
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (ciu *CourseInfoUpdate) SetNillableDetail(s *string) *CourseInfoUpdate {
	if s != nil {
		ciu.SetDetail(*s)
	}
	return ciu
}

// SetState sets the "state" field.
func (ciu *CourseInfoUpdate) SetState(s string) *CourseInfoUpdate {
	ciu.mutation.SetState(s)
	return ciu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ciu *CourseInfoUpdate) SetNillableState(s *string) *CourseInfoUpdate {
	if s != nil {
		ciu.SetState(*s)
	}
	return ciu
}

// SetIsDeleted sets the "is_deleted" field.
func (ciu *CourseInfoUpdate) SetIsDeleted(b bool) *CourseInfoUpdate {
	ciu.mutation.SetIsDeleted(b)
	return ciu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ciu *CourseInfoUpdate) SetNillableIsDeleted(b *bool) *CourseInfoUpdate {
	if b != nil {
		ciu.SetIsDeleted(*b)
	}
	return ciu
}

// SetCreatedAt sets the "created_at" field.
func (ciu *CourseInfoUpdate) SetCreatedAt(t time.Time) *CourseInfoUpdate {
	ciu.mutation.SetCreatedAt(t)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CourseInfoUpdate) SetNillableCreatedAt(t *time.Time) *CourseInfoUpdate {
	if t != nil {
		ciu.SetCreatedAt(*t)
	}
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CourseInfoUpdate) SetUpdatedAt(t time.Time) *CourseInfoUpdate {
	ciu.mutation.SetUpdatedAt(t)
	return ciu
}

// SetDeletedAt sets the "deleted_at" field.
func (ciu *CourseInfoUpdate) SetDeletedAt(t time.Time) *CourseInfoUpdate {
	ciu.mutation.SetDeletedAt(t)
	return ciu
}

// Mutation returns the CourseInfoMutation object of the builder.
func (ciu *CourseInfoUpdate) Mutation() *CourseInfoMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CourseInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ciu.defaults()
	if len(ciu.hooks) == 0 {
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			if ciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CourseInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CourseInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CourseInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciu *CourseInfoUpdate) defaults() {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		v := courseinfo.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
}

func (ciu *CourseInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseinfo.Table,
			Columns: courseinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseinfo.FieldID,
			},
		},
	}
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.ClickCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldClickCount,
		})
	}
	if value, ok := ciu.mutation.AddedClickCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldClickCount,
		})
	}
	if value, ok := ciu.mutation.LearnCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldLearnCount,
		})
	}
	if value, ok := ciu.mutation.AddedLearnCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldLearnCount,
		})
	}
	if value, ok := ciu.mutation.TotalDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldTotalDuration,
		})
	}
	if value, ok := ciu.mutation.AddedTotalDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldTotalDuration,
		})
	}
	if value, ok := ciu.mutation.SectionCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldSectionCount,
		})
	}
	if value, ok := ciu.mutation.AddedSectionCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldSectionCount,
		})
	}
	if value, ok := ciu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: courseinfo.FieldPrice,
		})
	}
	if value, ok := ciu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: courseinfo.FieldPrice,
		})
	}
	if value, ok := ciu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseinfo.FieldDetail,
		})
	}
	if value, ok := ciu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseinfo.FieldState,
		})
	}
	if value, ok := ciu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: courseinfo.FieldIsDeleted,
		})
	}
	if value, ok := ciu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CourseInfoUpdateOne is the builder for updating a single CourseInfo entity.
type CourseInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseInfoMutation
}

// SetClickCount sets the "click_count" field.
func (ciuo *CourseInfoUpdateOne) SetClickCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.ResetClickCount()
	ciuo.mutation.SetClickCount(i)
	return ciuo
}

// AddClickCount adds i to the "click_count" field.
func (ciuo *CourseInfoUpdateOne) AddClickCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.AddClickCount(i)
	return ciuo
}

// SetLearnCount sets the "learn_count" field.
func (ciuo *CourseInfoUpdateOne) SetLearnCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.ResetLearnCount()
	ciuo.mutation.SetLearnCount(i)
	return ciuo
}

// AddLearnCount adds i to the "learn_count" field.
func (ciuo *CourseInfoUpdateOne) AddLearnCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.AddLearnCount(i)
	return ciuo
}

// SetTotalDuration sets the "total_duration" field.
func (ciuo *CourseInfoUpdateOne) SetTotalDuration(i int) *CourseInfoUpdateOne {
	ciuo.mutation.ResetTotalDuration()
	ciuo.mutation.SetTotalDuration(i)
	return ciuo
}

// AddTotalDuration adds i to the "total_duration" field.
func (ciuo *CourseInfoUpdateOne) AddTotalDuration(i int) *CourseInfoUpdateOne {
	ciuo.mutation.AddTotalDuration(i)
	return ciuo
}

// SetSectionCount sets the "section_count" field.
func (ciuo *CourseInfoUpdateOne) SetSectionCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.ResetSectionCount()
	ciuo.mutation.SetSectionCount(i)
	return ciuo
}

// AddSectionCount adds i to the "section_count" field.
func (ciuo *CourseInfoUpdateOne) AddSectionCount(i int) *CourseInfoUpdateOne {
	ciuo.mutation.AddSectionCount(i)
	return ciuo
}

// SetPrice sets the "price" field.
func (ciuo *CourseInfoUpdateOne) SetPrice(f float64) *CourseInfoUpdateOne {
	ciuo.mutation.ResetPrice()
	ciuo.mutation.SetPrice(f)
	return ciuo
}

// AddPrice adds f to the "price" field.
func (ciuo *CourseInfoUpdateOne) AddPrice(f float64) *CourseInfoUpdateOne {
	ciuo.mutation.AddPrice(f)
	return ciuo
}

// SetDetail sets the "detail" field.
func (ciuo *CourseInfoUpdateOne) SetDetail(s string) *CourseInfoUpdateOne {
	ciuo.mutation.SetDetail(s)
	return ciuo
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (ciuo *CourseInfoUpdateOne) SetNillableDetail(s *string) *CourseInfoUpdateOne {
	if s != nil {
		ciuo.SetDetail(*s)
	}
	return ciuo
}

// SetState sets the "state" field.
func (ciuo *CourseInfoUpdateOne) SetState(s string) *CourseInfoUpdateOne {
	ciuo.mutation.SetState(s)
	return ciuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ciuo *CourseInfoUpdateOne) SetNillableState(s *string) *CourseInfoUpdateOne {
	if s != nil {
		ciuo.SetState(*s)
	}
	return ciuo
}

// SetIsDeleted sets the "is_deleted" field.
func (ciuo *CourseInfoUpdateOne) SetIsDeleted(b bool) *CourseInfoUpdateOne {
	ciuo.mutation.SetIsDeleted(b)
	return ciuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ciuo *CourseInfoUpdateOne) SetNillableIsDeleted(b *bool) *CourseInfoUpdateOne {
	if b != nil {
		ciuo.SetIsDeleted(*b)
	}
	return ciuo
}

// SetCreatedAt sets the "created_at" field.
func (ciuo *CourseInfoUpdateOne) SetCreatedAt(t time.Time) *CourseInfoUpdateOne {
	ciuo.mutation.SetCreatedAt(t)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CourseInfoUpdateOne) SetNillableCreatedAt(t *time.Time) *CourseInfoUpdateOne {
	if t != nil {
		ciuo.SetCreatedAt(*t)
	}
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CourseInfoUpdateOne) SetUpdatedAt(t time.Time) *CourseInfoUpdateOne {
	ciuo.mutation.SetUpdatedAt(t)
	return ciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ciuo *CourseInfoUpdateOne) SetDeletedAt(t time.Time) *CourseInfoUpdateOne {
	ciuo.mutation.SetDeletedAt(t)
	return ciuo
}

// Mutation returns the CourseInfoMutation object of the builder.
func (ciuo *CourseInfoUpdateOne) Mutation() *CourseInfoMutation {
	return ciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CourseInfoUpdateOne) Select(field string, fields ...string) *CourseInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CourseInfo entity.
func (ciuo *CourseInfoUpdateOne) Save(ctx context.Context) (*CourseInfo, error) {
	var (
		err  error
		node *CourseInfo
	)
	ciuo.defaults()
	if len(ciuo.hooks) == 0 {
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			if ciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ciuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CourseInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CourseInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CourseInfoUpdateOne) SaveX(ctx context.Context) *CourseInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CourseInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CourseInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuo *CourseInfoUpdateOne) defaults() {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		v := courseinfo.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
}

func (ciuo *CourseInfoUpdateOne) sqlSave(ctx context.Context) (_node *CourseInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   courseinfo.Table,
			Columns: courseinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: courseinfo.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, courseinfo.FieldID)
		for _, f := range fields {
			if !courseinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != courseinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.ClickCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldClickCount,
		})
	}
	if value, ok := ciuo.mutation.AddedClickCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldClickCount,
		})
	}
	if value, ok := ciuo.mutation.LearnCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldLearnCount,
		})
	}
	if value, ok := ciuo.mutation.AddedLearnCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldLearnCount,
		})
	}
	if value, ok := ciuo.mutation.TotalDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldTotalDuration,
		})
	}
	if value, ok := ciuo.mutation.AddedTotalDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldTotalDuration,
		})
	}
	if value, ok := ciuo.mutation.SectionCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldSectionCount,
		})
	}
	if value, ok := ciuo.mutation.AddedSectionCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: courseinfo.FieldSectionCount,
		})
	}
	if value, ok := ciuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: courseinfo.FieldPrice,
		})
	}
	if value, ok := ciuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: courseinfo.FieldPrice,
		})
	}
	if value, ok := ciuo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseinfo.FieldDetail,
		})
	}
	if value, ok := ciuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: courseinfo.FieldState,
		})
	}
	if value, ok := ciuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: courseinfo.FieldIsDeleted,
		})
	}
	if value, ok := ciuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: courseinfo.FieldDeletedAt,
		})
	}
	_node = &CourseInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courseinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
