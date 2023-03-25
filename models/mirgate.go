package models

import (
	"lite_blog/pkg/interfaces"
)

// MigrateData TODO
// @author
// @date
//
func MigrateData(sdk interfaces.MysqlClient) {

	tmp := MysqlAutoData{}

	tmp.AutoData(sdk.GetClient())
}
