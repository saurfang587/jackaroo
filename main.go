package main

import (
	"database/sql"
	"xiangxiang/jackaroo/app/huawei"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/inital"
)

func main() {
	global.G_VP = inital.Viper()
	global.G_DB = inital.Gorm()
	if global.G_DB != nil {
		db, _ := global.G_DB.DB()
		defer func(db *sql.DB) {
			db.Close()
		}(db)
	}

	huawei.Handler("")
}
