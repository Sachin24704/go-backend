// Code generated by ent, DO NOT EDIT.

package tweet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Sachin24704/go-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldID, id))
}

// Tweet applies equality check predicate on the "tweet" field. It's identical to TweetEQ.
func Tweet(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldTweet, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUserID, v))
}

// TweetEQ applies the EQ predicate on the "tweet" field.
func TweetEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldTweet, v))
}

// TweetNEQ applies the NEQ predicate on the "tweet" field.
func TweetNEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldTweet, v))
}

// TweetIn applies the In predicate on the "tweet" field.
func TweetIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldTweet, vs...))
}

// TweetNotIn applies the NotIn predicate on the "tweet" field.
func TweetNotIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldTweet, vs...))
}

// TweetGT applies the GT predicate on the "tweet" field.
func TweetGT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldTweet, v))
}

// TweetGTE applies the GTE predicate on the "tweet" field.
func TweetGTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldTweet, v))
}

// TweetLT applies the LT predicate on the "tweet" field.
func TweetLT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldTweet, v))
}

// TweetLTE applies the LTE predicate on the "tweet" field.
func TweetLTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldTweet, v))
}

// TweetContains applies the Contains predicate on the "tweet" field.
func TweetContains(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContains(FieldTweet, v))
}

// TweetHasPrefix applies the HasPrefix predicate on the "tweet" field.
func TweetHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasPrefix(FieldTweet, v))
}

// TweetHasSuffix applies the HasSuffix predicate on the "tweet" field.
func TweetHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasSuffix(FieldTweet, v))
}

// TweetEqualFold applies the EqualFold predicate on the "tweet" field.
func TweetEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldTweet, v))
}

// TweetContainsFold applies the ContainsFold predicate on the "tweet" field.
func TweetContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldTweet, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldUserID, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.NotPredicates(p))
}
