// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"sync"

	"github.com/10gen/mcli/internal/config"
)

type Config interface {
	Service() string
	SetService(string)
	PublicAPIKey() string
	SetPublicAPIKey(string)
	PrivateAPIKey() string
	SetPrivateAPIKey(string)
	OpsManagerURL() string
	SetOpsManagerURL(string)
	ProjectID() string
	SetProjectID(string)
	Save() error
}

type globalOpts struct {
	Config
	profile   string
	projectID string
	once      sync.Once
}

// newGlobalOpts returns an globalOpts
func newGlobalOpts() *globalOpts {
	return new(globalOpts)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) ProjectID() string {
	_ = opts.loadConfig()
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = opts.Config.ProjectID()
	return opts.projectID
}

func (opts *globalOpts) loadConfig() error {
	var err error
	opts.once.Do(func() {
		opts.Config, err = config.New(opts.profile)
	})
	return err
}