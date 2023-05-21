package Meituan

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

var AllMeituan []Meituan

func Header(cookie string) (bool, error) {
	pan, err := Get()
	if pan == false {
		return false, err
	}
	pan1, err1 := Meituan_orm()
	if pan1 == false {
		return false, err1
	}
	return true, nil
}
func Get() (bool, error) {
	c := colly.NewCollector()
	// 在请求前设置Header字段
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "application/json")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Content-Type", "application/json")
		r.Headers.Set("Origin", "https://zhaopin.meituan.com")
		r.Headers.Set("Referer", "https://zhaopin.meituan.com/web/campus")
		r.Headers.Set("Sec-Fetch-Dest", "empty")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("sec-ch-ua", "\"Microsoft Edge\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
	})
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		Delay:       time.Second * 1,
		RandomDelay: time.Millisecond * 500,
	})
	if err != nil {
		return false, err
	}
	test := Content{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("json数据解析失败", err)
		}
		if len(test.Data.List) < 1 {
			return
		}
		AllMeituan = append(AllMeituan, test.Data.List...)
	})
	i := 1
	for {
		data := map[string]interface{}{
			"page": map[string]int{
				"pageNo":   i,
				"pageSize": 10,
			},
			"jobShareType": "1",
			"keywords":     "",
			"cityList":     []interface{}{},
			"department":   []interface{}{},
			"jfJgList":     []interface{}{},
			"jobType": []map[string]interface{}{
				{
					"code":    "1",
					"subCode": []interface{}{},
				},
				{
					"code":    "2",
					"subCode": []interface{}{},
				},
			},
			"u_query_id": "",
			"r_query_id": "",
		}
		Data, err2 := json.Marshal(data)
		if err2 != nil {
			fmt.Println("将json解压为byte数组出错")
			return false, err2
		}
		err1 := c.PostRaw("https://zhaopin.meituan.com/api/official/job/getJobList", Data)
		if err1 != nil {
			fmt.Println("解析url请求地址出错，请你继续尝试", err1)
			return false, err1
		}
		i++
		if len(test.Data.List) < 1 {
			return true, nil
		}
	}
}
