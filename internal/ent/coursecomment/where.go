// Code generated by ent, DO NOT EDIT.

package coursecomment

import (
	"online-teaching/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsShow), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComment), v))
	})
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldComment), v...))
	})
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldComment), v...))
	})
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComment), v))
	})
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComment), v))
	})
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComment), v))
	})
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComment), v))
	})
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComment), v))
	})
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComment), v))
	})
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComment), v))
	})
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComment), v))
	})
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComment), v))
	})
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrder), v))
	})
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrder), v...))
	})
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrder), v...))
	})
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrder), v))
	})
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrder), v))
	})
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrder), v))
	})
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrder), v))
	})
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsShow), v))
	})
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsShow), v))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CourseComment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
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
func Not(p predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		p(s.Not())
	})
}
