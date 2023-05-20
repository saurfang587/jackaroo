package Bilibili

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

var AllBilibili []BiliBili

func Header(cookie string) (bool, error) {
	pan, err := Get()
	if pan == false {
		return false, err
	}
	pan1, err1 := Bilibili_orm()
	if pan1 == false {
		return false, err1
	}
	return true, nil
}
func Get() (bool, error) {
	token := GetAdd()
	if token == "" {
		return false, nil
	}
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
		r.Headers.Set("x-csrf", token)
		r.Headers.Set("x-usertype", "2")
	})
	test := Context{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("解析json数据失败：", err)
			return
		}
		if len(test.Data.List) < 1 {
			return
		}
		AllBilibili = append(AllBilibili, test.Data.List...)
	})
	i := 1
	for {
		// 发送POST请求
		data := map[string]interface{}{
			"pageSize":         10,
			"pageNum":          i,
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
			return false, err
		}
		err = c.PostRaw("https://jobs.bilibili.com/api/campus/position/positionList", jsonData)
		if err != nil {
			fmt.Println("请求失败：", err)
			return false, err
		}
		if i > 2 {
			if i > test.Data.Pages {
				return true, nil
			}
		}
		i++
	}
}

// 爬取页面的token 因为这个校招信息没有cookie
func GetAdd() string {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = "GET"
		r.Headers.Set("authority", "jobs.bilibili.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("origin", "https://jobs.bilibili.com")
		r.Headers.Set("referer", "https://jobs.bilibili.com/campus/positions?type=0")
		r.Headers.Set("sec-ch-ua", `""Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24""`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
		r.Headers.Set("x-appkey", "ops.ehr-api.auth")
		r.Headers.Set("x-usertype", "2")
	})
	rep := &SCRFRep{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, rep)
		if err != nil {
			fmt.Println(err)
		}
	})
	_ = c.Visit("https://jobs.bilibili.com/api/auth/v1/csrf/token")
	return rep.Data
}
