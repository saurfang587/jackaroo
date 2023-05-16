package main

import (
	"database/sql"
	"fmt"
	"xiangxiang/jackaroo/app/Wangyi"
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
	//测试B站
	//Bilibili.Header("")
	//测试alibaba
	//Alibaba.Header("")
	//测试百度
	//Baidu.Header("")
	//测试美团
	//Meituan.Header("")
	//测试腾讯
	//Tencent.Header("")
	//测试网易
	Wangyi.Header("")
	//测试微软
	//Weiruan.Header("")
}
