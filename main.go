package main

import (
	"fmt"
	"xiangxiang/jackaroo/app/model"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/inital"
	"xiangxiang/jackaroo/timer"
)

func main() {
	//关闭数据库
	defer func() {
		db, _ := model.GlobalDb.DB()
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
func init() {
	global.G_VP = inital.Viper()
	model.GlobalDb = model.New()
	fmt.Print("数据库连接成功")
	err := model.GlobalDb.AutoMigrate(&model.Job{})
	if err != nil {
		fmt.Println("表创建失败")
		return
	}
}
