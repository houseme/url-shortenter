// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package test

import (
	"context"
	"testing"

	"github.com/houseme/url-shortenter/utility/env"
)

func TestNewAlibabaEnv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *env.AlibabaEnv
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
			want:    &env.AlibabaEnv{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := env.NewAlibabaEnv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAlibabaEnv() error = %v1, wantErr %v1", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewAlibabaEnv() got = %v1, want %v1", got, tt.want)
			// }
			t.Log("got params:", got.String(tt.args.ctx))
		})
	}
}
