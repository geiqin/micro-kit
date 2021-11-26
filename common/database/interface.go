package database

import (
	"gorm.io/gorm"
)

// Database 数据库配置
type Databases interface {
	Setup(cfg *Configure) (*gorm.DB, error)
	Open(conn string, cfg *gorm.Config) (*gorm.DB, error)
	//GetConnect() string
	//GetDriver() string
}
