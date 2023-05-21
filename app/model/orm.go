package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"xiangxiang/jackaroo/global"
)

var GlobalDb *gorm.DB

func New() *gorm.DB {
	if GlobalDb != nil {
		return GlobalDb
	}
	return GormMysql()
}

func GormMysql() *gorm.DB {
	msq := global.G_CONFIG.Mysql
	if msq.Dbname == "" {
		return nil
	}
	mysqlInstance := mysql.New(mysql.Config{
		DSN:                       global.G_CONFIG.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         256,                         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                       // 根据版本自动配置
	})
	if db, err := gorm.Open(mysqlInstance, gormConfig()); err != nil {
		return nil
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(msq.MaxIdleConns)
		sqlDb.SetMaxOpenConns(msq.MaxOpenConns)
		sqlDb.SetConnMaxLifetime(msq.ConnMaxLifetime)
		return db
	}

}

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Silent),
	}
	return config
}
