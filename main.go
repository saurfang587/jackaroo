package main

import (
	"fmt"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/inital"
	"xiangxiang/jackaroo/timer"
)

func main() {
	//初始化相关配置 数据库和日志
	Init()
	//关闭数据库
	defer func() {
		db, _ := global.G_DB.DB()
		err := db.Close()
		if err != nil {
			fmt.Println("数据库关闭失败")
			return
		}
	}()
	//关闭定时器
	defer timer.Close()
	//开启定时器
	timer.Timer()
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
