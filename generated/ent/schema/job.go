// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/siemens/wfx/generated/model"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "job"},
	}
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		// override built-in id field since we want to use UUIDs (generated by wfx)
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			DefaultFunc(uuid.NewString).
			Annotations(entsql.Annotation{Size: 36}).
			Immutable(),
		// note: we do not use field.Time because this results in column type 'timestamp'; for some databases (e.g. MySQL)
		// this implies truncating to whole seconds; instead, we store milliseconds since the UNIX epoch.
		field.Time("stime").
			Comment("start time").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "TIMESTAMP(6)", // microsecond precision
			}),
		field.Time("mtime").
			Comment("modification time").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "TIMESTAMP(6)", // microsecond precision
			}),
		field.String("client_id"),
		field.JSON("definition", map[string]any{}).Optional(),
		// JobStatus
		field.JSON("status", model.JobStatus{}),
		field.String("group").Optional(),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "workflow" of type `Workflow`
		// and reference it to the "job" edge (in Workflow schema)
		// explicitly using the `Ref` method.
		edge.From("workflow", Workflow.Type).
			Ref("jobs").
			Unique(),
		// a job can have zero or more history entries
		edge.To("history", History.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("tags", Tag.Type).
			Ref("jobs"),
	}
}

func (Job) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("stime"),
		index.Fields("client_id"),
		index.Fields("group"),
	}
}

type JSONIndex struct {
	Column string
	Path   string
}
