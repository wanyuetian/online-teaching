// Code generated by ent, DO NOT EDIT.

package courseswiper

import (
	"online-teaching/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsShow), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrder), v))
	})
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrder), v...))
	})
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrder), v...))
	})
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrder), v))
	})
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrder), v))
	})
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrder), v))
	})
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrder), v))
	})
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImage), v))
	})
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldImage), v...))
	})
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldImage), v...))
	})
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImage), v))
	})
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImage), v))
	})
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImage), v))
	})
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImage), v))
	})
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImage), v))
	})
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImage), v))
	})
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImage), v))
	})
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImage), v))
	})
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImage), v))
	})
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsShow), v))
	})
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsShow), v))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CourseSwiper {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CourseSwiper) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CourseSwiper) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CourseSwiper) predicate.CourseSwiper {
	return predicate.CourseSwiper(func(s *sql.Selector) {
		p(s.Not())
	})
}
