package timer

import (
	"fmt"
	"time"
	"xiangxiang/jackaroo/router"
)

var Time1 *time.Ticker

func Timer() {
	router.Router()
	//定时器 使用NewTicker() 来每隔12小时执行一次，30秒后开启监听机制
	//当十二小时内没有执行路由组机制，就会返回超时
	Time1 = time.NewTicker(time.Hour * 12)
	for {
		select {
		case <-Time1.C:
			router.Router()
			break
		case <-time.After(time.Minute * 721):
			fmt.Println("超时了,在十二小时内没有监听到你的路由组的执行")
			return
		}
	}
}
func Close() {
	Time1.Stop()
}
