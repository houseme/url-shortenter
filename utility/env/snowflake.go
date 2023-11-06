// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package env

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// SnowflakeEnv .
type SnowflakeEnv struct {
	datacenter int64
	worker     int64
	config     map[string]*gvar.Var
}

// Datacenter .
func (e *SnowflakeEnv) Datacenter(_ context.Context) int64 {
	return e.datacenter
}

// Worker .
func (e *SnowflakeEnv) Worker(_ context.Context) int64 {
	return e.worker
}

// Config .
func (e *SnowflakeEnv) Config(_ context.Context) map[string]*gvar.Var {
	return e.config
}

// String .
func (e *SnowflakeEnv) String(_ context.Context) string {
	return `{"datacenter":"` + gconv.String(e.datacenter) + `","worker":"` + gconv.String(e.worker) + `"}`
}

// NewSnowflakeEnv .
func NewSnowflakeEnv(ctx context.Context) (*SnowflakeEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewSnowflakeEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "snowflake")
	if err != nil {
		return nil, gerror.Wrap(err, "config snowflake get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config snowflake is empty")
	}

	var config = v.MapStrVar()
	return &SnowflakeEnv{
		worker:     config["worker"].Int64(),
		datacenter: config["datacenter"].Int64(),
		config:     config,
	}, nil
}
