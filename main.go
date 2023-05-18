package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
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
	err := global.G_DB.AutoMigrate(&Alibaba.Hello{})
	if err != nil {
		fmt.Println("表创建失败")
		return
	}
	Hello()
}

// 并发爬取信息
func Hello() {
	var wg sync.WaitGroup
	wg.Add(9)
	go func() {
		defer wg.Done()
		for {
			if pan, err := Alibaba.Header(""); pan != false {
				fmt.Println("alibaba爬完了")
				break
			} else {
				fmt.Println("alibaba重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Baidu.Header(""); pan != false {
				fmt.Println("百度爬完了")
				break
			} else {
				fmt.Println("百度重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Bilibili.Header(""); pan != false {
				fmt.Println("Bilibili爬完了")
				break
			} else {
				fmt.Println("bilibili重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := jingdong.Header(""); pan != false {
				fmt.Println("jingdong爬完了")
				break
			} else {
				fmt.Println("京东重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Meituan.Header(""); pan != false {
				fmt.Println("美团爬完了")
				break
			} else {
				fmt.Println("美团重试", err)
				time.Sleep(3 * time.Second)
				continue
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := mihoyo.Header(""); pan != false {
				fmt.Println("米哈游爬完了")
				break
			} else {
				fmt.Println("米哈游重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Tencent.Header(""); pan != false {
				fmt.Println("腾讯爬完了")
				break
			} else {
				fmt.Println("腾讯重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Wangyi.Header(""); pan != false {
				fmt.Println("网易爬完了")
				break
			} else {
				fmt.Println("网易重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Weiruan.Header(""); pan != false {
				fmt.Println("微软爬完了")
				break
			} else {
				fmt.Println("微软重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	wg.Wait()
	fmt.Println("所有网站都爬取完毕了")
}
