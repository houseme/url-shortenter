// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

type (
	IShort interface {
		// CreateShort is the handler for CreateShort
		CreateShort(ctx context.Context, in *model.CreateShortInput) (out *model.CreateShortOutput, err error)
		// ModifyShort is the handler for ModifyShort
		ModifyShort(ctx context.Context, in *model.ModifyShortInput) (out *model.ModifyShortOutput, err error)
		// QueryShort is the handler for QueryShort
		QueryShort(ctx context.Context, in *model.QueryShortInput) (out *model.QueryShortOutput, err error)
		// QueryStat is the handler for QueryStat
		QueryStat(ctx context.Context, in *model.QueryStatInput) (out *model.QueryStatOutput, err error)
		// ShortDomain is the handler for ShortDomain
		ShortDomain(ctx context.Context, in *model.ShortDomainInput) (out *model.ShortDomainOutput, err error)
		// CreateTag is the handler for CreateTag.
		CreateTag(ctx context.Context, in *model.CreateTagInput) (out *model.CreateTagOutput, err error)
		// DelTag is the handler for DelTag.
		DelTag(ctx context.Context, in *model.DelTagInput) (out *model.DelTagOutput, err error)
		// AddTag is the handler for AddTag.
		AddTag(ctx context.Context, in *model.AddTagInput) (out *model.AddTagOutput, err error)
	}
)

var (
	localShort IShort
)

func Short() IShort {
	if localShort == nil {
		panic("implement not found for interface IShort, forgot register?")
	}
	return localShort
}

func RegisterShort(i IShort) {
	localShort = i
}
