package entgo

/*
 * SPDX-FileCopyrightText: 2023 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/ftag"
	"github.com/siemens/wfx/generated/api"
	"github.com/siemens/wfx/generated/ent"
	"github.com/siemens/wfx/middleware/logging"
)

// CreateWorkflow creates a new workflow.
func (db Database) CreateWorkflow(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error) {
	log := logging.LoggerFromCtx(ctx)

	builder := db.client.Workflow.
		Create().
		SetName(workflow.Name).
		SetStates(workflow.States).
		SetTransitions(workflow.Transitions).
		SetGroups(workflow.Groups).
		SetDescription(workflow.Description)
	entity, err := builder.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			log.Error().Err(err).Msg("Failed to persist workflow due to constraints")
			return nil, fault.Wrap(err, ftag.With(ftag.AlreadyExists))
		}
		log.Error().Err(err).Msg("Failed to persist workflow due to internal problem")
		return nil, fault.Wrap(err, ftag.With(ftag.Internal))
	}
	wf := convertWorkflow(entity)
	return &wf, nil
}
