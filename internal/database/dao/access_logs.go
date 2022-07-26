// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/url-shortenter/internal/database/dao/internal"
)

// internalAccessLogsDao is internal type for wrapping internal DAO implements.
type internalAccessLogsDao = *internal.AccessLogsDao

// accessLogsDao is the data access object for table access_logs.
// You can define custom methods on it to extend its functionality as you wish.
type accessLogsDao struct {
	internalAccessLogsDao
}

var (
	// AccessLogs is globally public accessible object for table access_logs operations.
	AccessLogs = accessLogsDao{
		internal.NewAccessLogsDao(),
	}
)

// Fill with you ideas below.
