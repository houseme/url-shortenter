package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// SnowflakeEnv .
type SnowflakeEnv struct {
	datacenter int64 `json:"datacenter"`
	worker     int64 `json:"worker"`
}

// Datacenter .
func (s *SnowflakeEnv) Datacenter(ctx context.Context) int64 {
	return s.datacenter
}

// Worker .
func (s *SnowflakeEnv) Worker(ctx context.Context) int64 {
	return s.worker
}

// String .
func (s *SnowflakeEnv) String() string {
	return `{"datacenter":"` + gconv.String(s.datacenter) + `","worker":"` + gconv.String(s.worker) + `"}`
}

// NewSnowflakeEnv .
func NewSnowflakeEnv(ctx context.Context) (*SnowflakeEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewSnowflakeEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, "snowflake")
		logger = gconv.String(ctx.Value("logger"))
	)

	defer func() {
		span.RecordError(err)
	}()

	if err != nil {
		err = gerror.Wrap(err, "config snowflake get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config snowflake is empty")
		return nil, err
	}

	var env *SnowflakeEnv
	if err = v.Scan(&env); err != nil {
		err = gerror.Wrap(err, "config snowflake scan failed")
		return nil, err
	}
	g.Log(logger).Debug(ctx, " config snowflake:", env)
	return env, nil
}
