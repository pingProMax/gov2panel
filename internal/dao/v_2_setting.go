// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2SettingDao is internal type for wrapping internal DAO implements.
type internalV2SettingDao = *internal.V2SettingDao

// v2SettingDao is the data access object for table v2_setting.
// You can define custom methods on it to extend its functionality as you wish.
type v2SettingDao struct {
	internalV2SettingDao
}

var (
	// V2Setting is globally public accessible object for table v2_setting operations.
	V2Setting = v2SettingDao{
		internal.NewV2SettingDao(),
	}
)

// Fill with you ideas below.
