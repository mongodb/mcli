// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blockstore

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/admin/backupstore"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var updateTemplate = "Blockstore configuration '{{.ID}}' updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	backupstore.AdminOpts
	store store.BlockstoresUpdater
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateBlockstore(opts.NewBackupStore())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// mongocli ops-manager admin backup blockstore(s) update <ID> [--assignment][--encryptedCredentials]
// [--label label][--loadFactor loadFactor][--maxCapacityGB maxCapacityGB][--uri uri][--ssl][--writeConcern writeConcern]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: update,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Assignment, flag.Assignment, false, usage.Assignment)
	cmd.Flags().BoolVar(&opts.EncryptedCredentials, flag.EncryptedCredentials, false, usage.EncryptedCredentials)
	cmd.Flags().StringSliceVar(&opts.Label, flag.Label, []string{}, usage.Label)
	cmd.Flags().Int64Var(&opts.LoadFactor, flag.LoadFactor, 0, usage.LoadFactor)
	cmd.Flags().Int64Var(&opts.MaxCapacityGB, flag.MaxCapacityGB, 0, usage.MaxCapacityGB)
	cmd.Flags().StringVar(&opts.URI, flag.URI, "", usage.BlockstoreURI)
	cmd.Flags().BoolVar(&opts.SSL, flag.SSL, false, usage.BlockstoreSSL)
	cmd.Flags().StringVar(&opts.WriteConcern, flag.WriteConcern, "", usage.WriteConcern)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}