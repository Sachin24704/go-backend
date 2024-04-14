// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sachin24704/go-backend/ent/predicate"
	"github.com/Sachin24704/go-backend/ent/tweet"
	"github.com/Sachin24704/go-backend/ent/user"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTweet sets the "tweet" field.
func (tu *TweetUpdate) SetTweet(s string) *TweetUpdate {
	tu.mutation.SetTweet(s)
	return tu
}

// SetNillableTweet sets the "tweet" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableTweet(s *string) *TweetUpdate {
	if s != nil {
		tu.SetTweet(*s)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TweetUpdate) SetCreatedAt(t time.Time) *TweetUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableCreatedAt(t *time.Time) *TweetUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TweetUpdate) SetUserID(s string) *TweetUpdate {
	tu.mutation.SetUserID(s)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableUserID(s *string) *TweetUpdate {
	if s != nil {
		tu.SetUserID(*s)
	}
	return tu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (tu *TweetUpdate) SetAuthorID(id string) *TweetUpdate {
	tu.mutation.SetAuthorID(id)
	return tu
}

// SetAuthor sets the "author" edge to the User entity.
func (tu *TweetUpdate) SetAuthor(u *User) *TweetUpdate {
	return tu.SetAuthorID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (tu *TweetUpdate) ClearAuthor() *TweetUpdate {
	tu.mutation.ClearAuthor()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TweetUpdate) check() error {
	if _, ok := tu.mutation.AuthorID(); tu.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tweet.author"`)
	}
	return nil
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Tweet(); ok {
		_spec.SetField(tweet.FieldTweet, field.TypeString, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(tweet.FieldCreatedAt, field.TypeTime, value)
	}
	if tu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetTweet sets the "tweet" field.
func (tuo *TweetUpdateOne) SetTweet(s string) *TweetUpdateOne {
	tuo.mutation.SetTweet(s)
	return tuo
}

// SetNillableTweet sets the "tweet" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableTweet(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetTweet(*s)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TweetUpdateOne) SetCreatedAt(t time.Time) *TweetUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableCreatedAt(t *time.Time) *TweetUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TweetUpdateOne) SetUserID(s string) *TweetUpdateOne {
	tuo.mutation.SetUserID(s)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableUserID(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetUserID(*s)
	}
	return tuo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (tuo *TweetUpdateOne) SetAuthorID(id string) *TweetUpdateOne {
	tuo.mutation.SetAuthorID(id)
	return tuo
}

// SetAuthor sets the "author" edge to the User entity.
func (tuo *TweetUpdateOne) SetAuthor(u *User) *TweetUpdateOne {
	return tuo.SetAuthorID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (tuo *TweetUpdateOne) ClearAuthor() *TweetUpdateOne {
	tuo.mutation.ClearAuthor()
	return tuo
}

// Where appends a list predicates to the TweetUpdate builder.
func (tuo *TweetUpdateOne) Where(ps ...predicate.Tweet) *TweetUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TweetUpdateOne) check() error {
	if _, ok := tuo.mutation.AuthorID(); tuo.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tweet.author"`)
	}
	return nil
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Tweet(); ok {
		_spec.SetField(tweet.FieldTweet, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(tweet.FieldCreatedAt, field.TypeTime, value)
	}
	if tuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
