// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/url-shortenter/internal/database/dao/internal"
)

// internalUsersMerchantDao is internal type for wrapping internal DAO implements.
type internalUsersMerchantDao = *internal.UsersMerchantDao

// usersMerchantDao is the data access object for table users_merchant.
// You can define custom methods on it to extend its functionality as you wish.
type usersMerchantDao struct {
	internalUsersMerchantDao
}

var (
	// UsersMerchant is globally public accessible object for table users_merchant operations.
	UsersMerchant = usersMerchantDao{
		internal.NewUsersMerchantDao(),
	}
)

// Fill with you ideas below.
