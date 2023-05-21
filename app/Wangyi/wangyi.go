package Wangyi

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

var AllWangyi []Wangyi

func Header(cookie string) (bool, error) {
	pan, err := Get()
	if pan == false {
		return false, err
	}
	pan1, err1 := Wangyi_orm()
	if pan1 == false {
		return false, err1
	}
	return true, nil
}
func Get() (bool, error) {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("authority", "hr.163.com")
		request.Headers.Set("accept", "application/json, text/plain, */*")
		request.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		request.Headers.Set("authtype", "ursAuth")
		request.Headers.Set("content-type", "application/json;charset=UTF-8")
		request.Headers.Set("lang", "zh")
		request.Headers.Set("origin", "https://hr.163.com")
		request.Headers.Set("referer", "https://hr.163.com/job-list.html?workType=1")
		request.Headers.Set("sec-ch-ua", "\"Microsoft Edge\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		request.Headers.Set("sec-ch-ua-mobile", "?0")
		request.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		request.Headers.Set("sec-fetch-dest", "empty")
		request.Headers.Set("sec-fetch-mode", "cors")
		request.Headers.Set("sec-fetch-site", "same-origin")
		request.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42")
	})
	test := Content{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("解析json数据失败，请重新尝试")
			return
		}
		AllWangyi = append(AllWangyi, test.Data.List...)
	})
	i := 1
	for {
		str := strconv.Itoa(i)
		res := map[string]string{
			"currentPage": str,
			"pageSize":    "10",
			"workType":    "1",
		}
		Data, err1 := json.Marshal(res)
		if err1 != nil {
			fmt.Println("将数组数据转化为json格式")
			return false, err1
		}
		err := c.PostRaw("https://hr.163.com/api/hr163/position/queryPage", Data)
		if err != nil {
			fmt.Println("请求url时出现错误啦！！")
			return false, err
		}
		if i > test.Data.Page {
			return true, nil
		}
		i++
	}
}
