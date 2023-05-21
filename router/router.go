package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
	"xiangxiang/jackaroo/app/logic/360"
	"xiangxiang/jackaroo/app/logic/Alibaba"
	"xiangxiang/jackaroo/app/logic/Baidu"
	"xiangxiang/jackaroo/app/logic/Meituan"
	"xiangxiang/jackaroo/app/logic/Tencent"
	"xiangxiang/jackaroo/app/logic/Wangyi"
	"xiangxiang/jackaroo/app/logic/Weiruan"
	"xiangxiang/jackaroo/app/logic/bilibili"
	"xiangxiang/jackaroo/app/logic/huawei"
	"xiangxiang/jackaroo/app/logic/jingdong"
	"xiangxiang/jackaroo/app/logic/lilisi"
	"xiangxiang/jackaroo/app/logic/mihoyo"
	"xiangxiang/jackaroo/app/logic/xiaohongshu"
	"xiangxiang/jackaroo/app/logic/xiaomi"
	"xiangxiang/jackaroo/app/logic/zijie"
	"xiangxiang/jackaroo/app/model"
)

func Router() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		Hello()
	})
	err1 := r.Run()
	if err1 != nil {
		fmt.Println("引擎启动失败", err1)
	}
	//定时删除未更新的数据：即招聘公司已经撤销的招聘信息
	model.DeleteStaleRecords(model.GlobalDb)
}

// 并发爬取信息
func Hello() {
	var wg sync.WaitGroup
	wg.Add(15)
	go func() {
		defer wg.Done()
		for {
			if pan, err := zijie.Header(""); pan != false {
				fmt.Println("zijie爬完了")
				break
			} else {
				fmt.Println("zijie重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := xiaomi.Header(""); pan != false {
				fmt.Println("xiaomi爬完了")
				break
			} else {
				fmt.Println("xiaomi重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := xiaohongshu.Header(""); pan != false {
				fmt.Println("xiaohongshu爬完了")
				break
			} else {
				fmt.Println("xiaohongshu重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := lilisi.Header(""); pan != false {
				fmt.Println("lilisi爬完了")
				break
			} else {
				fmt.Println("lilisi重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := huawei.Header(""); pan != false {
				fmt.Println("huawei爬完了")
				break
			} else {
				fmt.Println("huawei重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
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
	go func() {
		defer wg.Done()
		for {
			if pan, err := _60.Header(""); pan != false {
				fmt.Println("360爬完了")
				break
			} else {
				fmt.Println("360重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	wg.Wait()
	fmt.Println("所有网站都爬取完毕了")
}
