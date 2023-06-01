// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/siemens/wfx/generated/ent/job"
	"github.com/siemens/wfx/generated/ent/predicate"
	"github.com/siemens/wfx/generated/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// AddJobIDs adds the "jobs" edge to the Job entity by IDs.
func (tu *TagUpdate) AddJobIDs(ids ...string) *TagUpdate {
	tu.mutation.AddJobIDs(ids...)
	return tu
}

// AddJobs adds the "jobs" edges to the Job entity.
func (tu *TagUpdate) AddJobs(j ...*Job) *TagUpdate {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tu.AddJobIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearJobs clears all "jobs" edges to the Job entity.
func (tu *TagUpdate) ClearJobs() *TagUpdate {
	tu.mutation.ClearJobs()
	return tu
}

// RemoveJobIDs removes the "jobs" edge to Job entities by IDs.
func (tu *TagUpdate) RemoveJobIDs(ids ...string) *TagUpdate {
	tu.mutation.RemoveJobIDs(ids...)
	return tu
}

// RemoveJobs removes "jobs" edges to Job entities.
func (tu *TagUpdate) RemoveJobs(j ...*Job) *TagUpdate {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tu.RemoveJobIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if tu.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedJobsIDs(); len(nodes) > 0 && !tu.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// AddJobIDs adds the "jobs" edge to the Job entity by IDs.
func (tuo *TagUpdateOne) AddJobIDs(ids ...string) *TagUpdateOne {
	tuo.mutation.AddJobIDs(ids...)
	return tuo
}

// AddJobs adds the "jobs" edges to the Job entity.
func (tuo *TagUpdateOne) AddJobs(j ...*Job) *TagUpdateOne {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tuo.AddJobIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearJobs clears all "jobs" edges to the Job entity.
func (tuo *TagUpdateOne) ClearJobs() *TagUpdateOne {
	tuo.mutation.ClearJobs()
	return tuo
}

// RemoveJobIDs removes the "jobs" edge to Job entities by IDs.
func (tuo *TagUpdateOne) RemoveJobIDs(ids ...string) *TagUpdateOne {
	tuo.mutation.RemoveJobIDs(ids...)
	return tuo
}

// RemoveJobs removes "jobs" edges to Job entities.
func (tuo *TagUpdateOne) RemoveJobs(j ...*Job) *TagUpdateOne {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return tuo.RemoveJobIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (tuo *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
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
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if tuo.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedJobsIDs(); len(nodes) > 0 && !tuo.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.JobsTable,
			Columns: tag.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
