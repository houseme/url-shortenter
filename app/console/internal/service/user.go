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
	IUser interface {
		// CreateMerchant creates a new merchant.
		CreateMerchant(ctx context.Context, in *model.CreateMerchantInput) (out *model.CreateMerchantOutput, err error)
		// QueryMerchant queries merchant by id.
		QueryMerchant(ctx context.Context, in *model.QueryMerchantInput) (out *model.QueryMerchantOutput, err error)
		// Detail is the handler for Detail
		Detail(ctx context.Context, in *model.UserDetailInput) (out *model.UserDetailOutput, err error)
		// Update is the handler for Update
		Update(ctx context.Context, in *model.UpdateUserInput) (out *model.UpdateUserOutput, err error)
		// UpdatePassword is the handler for UpdatePassword
		UpdatePassword(ctx context.Context, in *model.UpdatePasswordInput) (out *model.UpdatePasswordOutput, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
