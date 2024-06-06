// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/siemens/wfx/generated/ent/job"
	"github.com/siemens/wfx/generated/ent/workflow"
	"github.com/siemens/wfx/generated/model"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// start time
	Stime time.Time `json:"stime,omitempty"`
	// modification time
	Mtime time.Time `json:"mtime,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// Definition holds the value of the "definition" field.
	Definition map[string]interface{} `json:"definition,omitempty"`
	// Status holds the value of the "status" field.
	Status model.JobStatus `json:"status,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobQuery when eager-loading is set.
	Edges         JobEdges `json:"edges"`
	workflow_jobs *int
	selectValues  sql.SelectValues
}

// JobEdges holds the relations/edges for other nodes in the graph.
type JobEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// History holds the value of the history edge.
	History []*History `json:"history,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) WorkflowOrErr() (*Workflow, error) {
	if e.Workflow != nil {
		return e.Workflow, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: workflow.Label}
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// HistoryOrErr returns the History value or an error if the edge
// was not loaded in eager-loading.
func (e JobEdges) HistoryOrErr() ([]*History, error) {
	if e.loadedTypes[1] {
		return e.History, nil
	}
	return nil, &NotLoadedError{edge: "history"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e JobEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldDefinition, job.FieldStatus:
			values[i] = new([]byte)
		case job.FieldID, job.FieldClientID, job.FieldGroup:
			values[i] = new(sql.NullString)
		case job.FieldStime, job.FieldMtime:
			values[i] = new(sql.NullTime)
		case job.ForeignKeys[0]: // workflow_jobs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				j.ID = value.String
			}
		case job.FieldStime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field stime", values[i])
			} else if value.Valid {
				j.Stime = value.Time
			}
		case job.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				j.Mtime = value.Time
			}
		case job.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				j.ClientID = value.String
			}
		case job.FieldDefinition:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field definition", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &j.Definition); err != nil {
					return fmt.Errorf("unmarshal field definition: %w", err)
				}
			}
		case job.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &j.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case job.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				j.Group = value.String
			}
		case job.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workflow_jobs", value)
			} else if value.Valid {
				j.workflow_jobs = new(int)
				*j.workflow_jobs = int(value.Int64)
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Job.
// This includes values selected through modifiers, order, etc.
func (j *Job) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// QueryWorkflow queries the "workflow" edge of the Job entity.
func (j *Job) QueryWorkflow() *WorkflowQuery {
	return NewJobClient(j.config).QueryWorkflow(j)
}

// QueryHistory queries the "history" edge of the Job entity.
func (j *Job) QueryHistory() *HistoryQuery {
	return NewJobClient(j.config).QueryHistory(j)
}

// QueryTags queries the "tags" edge of the Job entity.
func (j *Job) QueryTags() *TagQuery {
	return NewJobClient(j.config).QueryTags(j)
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return NewJobClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("stime=")
	builder.WriteString(j.Stime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(j.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(j.ClientID)
	builder.WriteString(", ")
	builder.WriteString("definition=")
	builder.WriteString(fmt.Sprintf("%v", j.Definition))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", j.Status))
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(j.Group)
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job