package main

import (
	"database/sql"
	"fmt"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/app/Baidu"
	"xiangxiang/jackaroo/app/Meituan"
	"xiangxiang/jackaroo/app/Tencent"
	"xiangxiang/jackaroo/app/Wangyi"
	"xiangxiang/jackaroo/app/Weiruan"
	Bilibili "xiangxiang/jackaroo/app/bilibili"
	"xiangxiang/jackaroo/app/jingdong"
	"xiangxiang/jackaroo/app/mihoyo"
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
	fmt.Print("数据库连接成功")
	Alibaba.Header("")
	Baidu.Header("")
	Bilibili.Header("")
	jingdong.Handler("")
	Meituan.Header("")
	mihoyo.Handler("")
	Tencent.Header("")
	Wangyi.Header("")
	Weiruan.Header("")
}
