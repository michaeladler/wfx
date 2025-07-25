package validate

/*
 * SPDX-FileCopyrightText: 2023 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Southclaws/fault"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/siemens/wfx/generated/api"
	"github.com/siemens/wfx/workflow"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a workflow",
		Long:  "Offline validation of a workflow",
		Example: `
wfxctl workflow validate wfx.workflow.dau.direct.yml
`,
		TraverseChildren: true,
		Args:             cobra.OnlyValidArgs,
		ValidArgsFunction: func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
			return []string{"yaml", "yml"}, cobra.ShellCompDirectiveFilterFileExt
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			allWorkflows := make([]*api.Workflow, 0, len(args))
			n := len(args)
			if n == 1 && args[0] == "-" {
				log.Debug().Msg("Reading workflow from stdin...")
				b, err := io.ReadAll(cmd.InOrStdin())
				if err != nil {
					return fault.Wrap(err)
				}
				wf, err := unmarshal(b)
				if err != nil {
					return fault.Wrap(err)
				}
				allWorkflows = append(allWorkflows, wf)
			} else {
				if n == 0 {
					return errors.New("workflow must be provided either via file or stdin")
				}
				for _, fname := range args {
					b, err := os.ReadFile(fname)
					if err != nil {
						return fault.Wrap(err)
					}
					wf, err := unmarshal(b)
					if err != nil {
						return fault.Wrap(err)
					}
					allWorkflows = append(allWorkflows, wf)
				}
			}

			out := cmd.OutOrStdout()
			for _, wf := range allWorkflows {
				err := workflow.ValidateWorkflow(wf)
				if err != nil {
					fmt.Fprintln(out, "ERROR:", err)
				} else {
					fmt.Fprintln(out, wf.Name, ": OK")
				}
			}
			return nil
		},
	}
	return cmd
}

func unmarshal(raw []byte) (*api.Workflow, error) {
	// try YAML
	var wf api.Workflow
	err := yaml.Unmarshal(raw, &wf)
	if err != nil {
		return nil, fault.Wrap(err)
	}
	return &wf, nil
}
