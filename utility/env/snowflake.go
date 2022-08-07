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
	ctx        context.Context
}

// Datacenter .
func (s *SnowflakeEnv) Datacenter(ctx context.Context) int64 {
	s.ctx = ctx
	return s.datacenter
}

// Worker .
func (s *SnowflakeEnv) Worker(ctx context.Context) int64 {
	s.ctx = ctx
	return s.worker
}

// String .
func (s *SnowflakeEnv) String(ctx context.Context) string {
	s.ctx = ctx
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

	if err != nil {
		err = gerror.Wrap(err, "config snowflake get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config snowflake is empty")
		return nil, err
	}

	var (
		config = v.MapStrVar()
		env    = &SnowflakeEnv{
			worker:     config["worker"].Int64(),
			datacenter: config["datacenter"].Int64(),
			ctx:        ctx,
			config:     config,
		}
	)
	g.Log(logger).Debug(ctx, " config snowflake:", env.String(ctx))
	return env, nil
}
