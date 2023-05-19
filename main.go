package main

import (
	"fmt"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/inital"
	"xiangxiang/jackaroo/router"
)

func main() {
	Init()
	defer func() {
		db, _ := global.G_DB.DB()
		err := db.Close()
		if err != nil {
			fmt.Println("数据库关闭失败")
			return
		}
	}()
	router.Router()
}

// 配置初始化
func Init() {
	global.G_VP = inital.Viper()
	global.G_DB = inital.Gorm()
	fmt.Print("数据库连接成功")
	err := global.G_DB.AutoMigrate(&Alibaba.Hello{})
	if err != nil {
		fmt.Println("表创建失败")
		return
	}
}
