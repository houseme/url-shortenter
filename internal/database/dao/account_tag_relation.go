// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/url-shortenter/internal/database/dao/internal"
)

// internalAccountTagRelationDao is internal type for wrapping internal DAO implements.
type internalAccountTagRelationDao = *internal.AccountTagRelationDao

// accountTagRelationDao is the data access object for table account_tag_relation.
// You can define custom methods on it to extend its functionality as you wish.
type accountTagRelationDao struct {
	internalAccountTagRelationDao
}

var (
	// AccountTagRelation is globally public accessible object for table account_tag_relation operations.
	AccountTagRelation = accountTagRelationDao{
		internal.NewAccountTagRelationDao(),
	}
)

// Fill with you ideas below.
