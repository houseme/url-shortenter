// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/url-shortenter/internal/database/dao/internal"
)

// internalShortTagDao is internal type for wrapping internal DAO implements.
type internalShortTagDao = *internal.ShortTagDao

// shortTagDao is the data access object for table short_tag.
// You can define custom methods on it to extend its functionality as you wish.
type shortTagDao struct {
	internalShortTagDao
}

var (
	// ShortTag is globally public accessible object for table short_tag operations.
	ShortTag = shortTagDao{
		internal.NewShortTagDao(),
	}
)

// Fill with you ideas below.