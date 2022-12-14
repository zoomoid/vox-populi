// Code generated by ent, DO NOT EDIT.

package reactiontemplate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Reaction applies equality check predicate on the "reaction" field. It's identical to ReactionEQ.
func Reaction(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReaction), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ReactionEQ applies the EQ predicate on the "reaction" field.
func ReactionEQ(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReaction), v))
	})
}

// ReactionNEQ applies the NEQ predicate on the "reaction" field.
func ReactionNEQ(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReaction), v))
	})
}

// ReactionIn applies the In predicate on the "reaction" field.
func ReactionIn(vs ...string) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReaction), v...))
	})
}

// ReactionNotIn applies the NotIn predicate on the "reaction" field.
func ReactionNotIn(vs ...string) predicate.ReactionTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReaction), v...))
	})
}

// ReactionGT applies the GT predicate on the "reaction" field.
func ReactionGT(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReaction), v))
	})
}

// ReactionGTE applies the GTE predicate on the "reaction" field.
func ReactionGTE(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReaction), v))
	})
}

// ReactionLT applies the LT predicate on the "reaction" field.
func ReactionLT(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReaction), v))
	})
}

// ReactionLTE applies the LTE predicate on the "reaction" field.
func ReactionLTE(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReaction), v))
	})
}

// ReactionContains applies the Contains predicate on the "reaction" field.
func ReactionContains(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReaction), v))
	})
}

// ReactionHasPrefix applies the HasPrefix predicate on the "reaction" field.
func ReactionHasPrefix(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReaction), v))
	})
}

// ReactionHasSuffix applies the HasSuffix predicate on the "reaction" field.
func ReactionHasSuffix(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReaction), v))
	})
}

// ReactionEqualFold applies the EqualFold predicate on the "reaction" field.
func ReactionEqualFold(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReaction), v))
	})
}

// ReactionContainsFold applies the ContainsFold predicate on the "reaction" field.
func ReactionContainsFold(v string) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReaction), v))
	})
}

// HasPoll applies the HasEdge predicate on the "poll" edge.
func HasPoll() predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PollTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PollTable, PollColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPollWith applies the HasEdge predicate on the "poll" edge with a given conditions (other predicates).
func HasPollWith(preds ...predicate.Poll) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PollInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PollTable, PollColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReactionTemplate) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReactionTemplate) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
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
func Not(p predicate.ReactionTemplate) predicate.ReactionTemplate {
	return predicate.ReactionTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
