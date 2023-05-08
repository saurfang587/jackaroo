package Bilibili

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

var Allid []string

func Header() {
	c := colly.NewCollector()
	// 设置请求头
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("authority", "jobs.bilibili.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("origin", "https://jobs.bilibili.com")
		r.Headers.Set("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
		r.Headers.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
		r.Headers.Set("x-appkey", "ops.ehr-api.auth")
		r.Headers.Set("x-channel", "campus")
		r.Headers.Set("x-csrf", "728e3c6c-e981-412e-be94-ab1f76557972")
		r.Headers.Set("x-usertype", "2")
	})
	c.OnResponse(func(r *colly.Response) {
		var result Context
		err := json.Unmarshal(r.Body, &result)
		if err != nil {
			fmt.Println("解析json数据失败：", err)
			return
		}
		fmt.Println(result.Data.List)
		for i := 0; i < 10; i++ {
			if len(result.Data.List) == i {
				return
			}
			Allid = append(Allid, result.Data.List[i].Id)
		}
	})
	//爬每个id
	for i := 1; i <= 19; i++ {
		str := i
		// 发送POST请求
		data := map[string]interface{}{
			"pageSize":         10,
			"pageNum":          str,
			"positionName":     "",
			"postCode":         []interface{}{},
			"postCodeList":     []interface{}{},
			"workLocationList": []interface{}{},
			"workTypeList":     []string{"0"},
			"positionTypeList": []string{"0"},
			"deptCodeList":     []interface{}{},
			"recruitType":      nil,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = c.PostRaw("https://jobs.bilibili.com/api/campus/position/positionList", jsonData)
		if err != nil {
			fmt.Println("请求第", i, "页失败：", err)
		}
	}
	FetchInformation()
}

// 每个页面信息抓取
func FetchInformation() {
	c := colly.NewCollector()
	// 设置请求头
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("authority", "jobs.bilibili.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("origin", "https://jobs.bilibili.com")
		r.Headers.Set("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
		r.Headers.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
		r.Headers.Set("x-appkey", "ops.ehr-api.auth")
		r.Headers.Set("x-channel", "campus")
		r.Headers.Set("x-csrf", "728e3c6c-e981-412e-be94-ab1f76557972")
		r.Headers.Set("x-usertype", "2")
	})
	c.OnResponse(func(r *colly.Response) {
		var res Context1
		err := json.Unmarshal(r.Body, &res)
		if err != nil {
			fmt.Println("解析json数据失败：", err)
			return
		}
		fmt.Println(res.Data)
	})
	for i := 0; i < len(Allid); i++ {
		url := "https://jobs.bilibili.com/api/campus/position/detail/" + Allid[i]
		err := c.Visit(url)
		if err != nil {
			fmt.Println("请求url地址过程出错，请修改重试")
			return
		}
	}
}
