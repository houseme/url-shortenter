package test

import (
	"context"
	"testing"

	"github.com/houseme/url-shortenter/utility/env"
)

func TestNewSnowflakeEnv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *env.SnowflakeEnv
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    &env.SnowflakeEnv{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := env.NewSnowflakeEnv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSnowflakeEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewSnowflakeEnv() got = %v, want %v", got, tt.want)
			// }
			t.Log("got params:", got.String(tt.args.ctx))
		})
	}
}
