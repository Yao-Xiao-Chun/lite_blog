package interfaces

import "github.com/jinzhu/gorm"

type MysqlClient interface {
	GetClient() *gorm.DB
}

type MysqlAutoData interface {
	AutoData() bool
}
