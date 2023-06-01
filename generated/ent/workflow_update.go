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
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/siemens/wfx/generated/ent/job"
	"github.com/siemens/wfx/generated/ent/predicate"
	"github.com/siemens/wfx/generated/ent/workflow"
	"github.com/siemens/wfx/generated/model"
)

// WorkflowUpdate is the builder for updating Workflow entities.
type WorkflowUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowMutation
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wu *WorkflowUpdate) Where(ps ...predicate.Workflow) *WorkflowUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetName sets the "name" field.
func (wu *WorkflowUpdate) SetName(s string) *WorkflowUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetStates sets the "states" field.
func (wu *WorkflowUpdate) SetStates(m []*model.State) *WorkflowUpdate {
	wu.mutation.SetStates(m)
	return wu
}

// AppendStates appends m to the "states" field.
func (wu *WorkflowUpdate) AppendStates(m []*model.State) *WorkflowUpdate {
	wu.mutation.AppendStates(m)
	return wu
}

// SetTransitions sets the "transitions" field.
func (wu *WorkflowUpdate) SetTransitions(m []*model.Transition) *WorkflowUpdate {
	wu.mutation.SetTransitions(m)
	return wu
}

// AppendTransitions appends m to the "transitions" field.
func (wu *WorkflowUpdate) AppendTransitions(m []*model.Transition) *WorkflowUpdate {
	wu.mutation.AppendTransitions(m)
	return wu
}

// SetGroups sets the "groups" field.
func (wu *WorkflowUpdate) SetGroups(m []*model.Group) *WorkflowUpdate {
	wu.mutation.SetGroups(m)
	return wu
}

// AppendGroups appends m to the "groups" field.
func (wu *WorkflowUpdate) AppendGroups(m []*model.Group) *WorkflowUpdate {
	wu.mutation.AppendGroups(m)
	return wu
}

// AddJobIDs adds the "jobs" edge to the Job entity by IDs.
func (wu *WorkflowUpdate) AddJobIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.AddJobIDs(ids...)
	return wu
}

// AddJobs adds the "jobs" edges to the Job entity.
func (wu *WorkflowUpdate) AddJobs(j ...*Job) *WorkflowUpdate {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return wu.AddJobIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wu *WorkflowUpdate) Mutation() *WorkflowMutation {
	return wu.mutation
}

// ClearJobs clears all "jobs" edges to the Job entity.
func (wu *WorkflowUpdate) ClearJobs() *WorkflowUpdate {
	wu.mutation.ClearJobs()
	return wu
}

// RemoveJobIDs removes the "jobs" edge to Job entities by IDs.
func (wu *WorkflowUpdate) RemoveJobIDs(ids ...string) *WorkflowUpdate {
	wu.mutation.RemoveJobIDs(ids...)
	return wu
}

// RemoveJobs removes "jobs" edges to Job entities.
func (wu *WorkflowUpdate) RemoveJobs(j ...*Job) *WorkflowUpdate {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return wu.RemoveJobIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowUpdate) check() error {
	if v, ok := wu.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Workflow.name": %w`, err)}
		}
	}
	return nil
}

func (wu *WorkflowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
	}
	if value, ok := wu.mutation.States(); ok {
		_spec.SetField(workflow.FieldStates, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedStates(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldStates, value)
		})
	}
	if value, ok := wu.mutation.Transitions(); ok {
		_spec.SetField(workflow.FieldTransitions, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedTransitions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldTransitions, value)
		})
	}
	if value, ok := wu.mutation.Groups(); ok {
		_spec.SetField(workflow.FieldGroups, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldGroups, value)
		})
	}
	if wu.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedJobsIDs(); len(nodes) > 0 && !wu.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
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
	if nodes := wu.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorkflowUpdateOne is the builder for updating a single Workflow entity.
type WorkflowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowMutation
}

// SetName sets the "name" field.
func (wuo *WorkflowUpdateOne) SetName(s string) *WorkflowUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetStates sets the "states" field.
func (wuo *WorkflowUpdateOne) SetStates(m []*model.State) *WorkflowUpdateOne {
	wuo.mutation.SetStates(m)
	return wuo
}

// AppendStates appends m to the "states" field.
func (wuo *WorkflowUpdateOne) AppendStates(m []*model.State) *WorkflowUpdateOne {
	wuo.mutation.AppendStates(m)
	return wuo
}

// SetTransitions sets the "transitions" field.
func (wuo *WorkflowUpdateOne) SetTransitions(m []*model.Transition) *WorkflowUpdateOne {
	wuo.mutation.SetTransitions(m)
	return wuo
}

// AppendTransitions appends m to the "transitions" field.
func (wuo *WorkflowUpdateOne) AppendTransitions(m []*model.Transition) *WorkflowUpdateOne {
	wuo.mutation.AppendTransitions(m)
	return wuo
}

// SetGroups sets the "groups" field.
func (wuo *WorkflowUpdateOne) SetGroups(m []*model.Group) *WorkflowUpdateOne {
	wuo.mutation.SetGroups(m)
	return wuo
}

// AppendGroups appends m to the "groups" field.
func (wuo *WorkflowUpdateOne) AppendGroups(m []*model.Group) *WorkflowUpdateOne {
	wuo.mutation.AppendGroups(m)
	return wuo
}

// AddJobIDs adds the "jobs" edge to the Job entity by IDs.
func (wuo *WorkflowUpdateOne) AddJobIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.AddJobIDs(ids...)
	return wuo
}

// AddJobs adds the "jobs" edges to the Job entity.
func (wuo *WorkflowUpdateOne) AddJobs(j ...*Job) *WorkflowUpdateOne {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return wuo.AddJobIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wuo *WorkflowUpdateOne) Mutation() *WorkflowMutation {
	return wuo.mutation
}

// ClearJobs clears all "jobs" edges to the Job entity.
func (wuo *WorkflowUpdateOne) ClearJobs() *WorkflowUpdateOne {
	wuo.mutation.ClearJobs()
	return wuo
}

// RemoveJobIDs removes the "jobs" edge to Job entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveJobIDs(ids ...string) *WorkflowUpdateOne {
	wuo.mutation.RemoveJobIDs(ids...)
	return wuo
}

// RemoveJobs removes "jobs" edges to Job entities.
func (wuo *WorkflowUpdateOne) RemoveJobs(j ...*Job) *WorkflowUpdateOne {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return wuo.RemoveJobIDs(ids...)
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wuo *WorkflowUpdateOne) Where(ps ...predicate.Workflow) *WorkflowUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowUpdateOne) Select(field string, fields ...string) *WorkflowUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflow entity.
func (wuo *WorkflowUpdateOne) Save(ctx context.Context) (*Workflow, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) SaveX(ctx context.Context) *Workflow {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowUpdateOne) check() error {
	if v, ok := wuo.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Workflow.name": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorkflowUpdateOne) sqlSave(ctx context.Context) (_node *Workflow, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Workflow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for _, f := range fields {
			if !workflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.States(); ok {
		_spec.SetField(workflow.FieldStates, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedStates(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldStates, value)
		})
	}
	if value, ok := wuo.mutation.Transitions(); ok {
		_spec.SetField(workflow.FieldTransitions, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedTransitions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldTransitions, value)
		})
	}
	if value, ok := wuo.mutation.Groups(); ok {
		_spec.SetField(workflow.FieldGroups, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldGroups, value)
		})
	}
	if wuo.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedJobsIDs(); len(nodes) > 0 && !wuo.mutation.JobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
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
	if nodes := wuo.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
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
	_node = &Workflow{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
