package Tencent

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
)

// var AllTencent []Each
var timestamp string
var each *[]Each
var AllInformation []Content3

func Header(cookie string) (bool, error) {
	each, pan, err := Get()
	if pan == false {
		return false, err
	}
	Get1(each)
	pan1, err1 := Tencent_orm()
	if pan1 == false {
		return false, err1
	}
	return true, nil
}

// 获取所有页面的链接
func Get() (AllTencent []Each, b1 bool, err error) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = "POST"
		r.Headers.Set("authority", "join.qq.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("origin", "https://join.qq.com")
		r.Headers.Set("referer", "https://join.qq.com/post.html")
		r.Headers.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "Windows")
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})
	test := &Content{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, test)
		if err != nil {
			return
		}
		if len(test.Data.PositionList) < 1 {
			return
		}
		AllTencent = append(AllTencent, test.Data.PositionList...)
	})
	i := 1
	for {
		requestBody := map[string]interface{}{
			"projectIdList":   []int{1, 2, 8, 11},
			"keyword":         "",
			"bgList":          []interface{}{},
			"workCountryType": 0,
			"workCityList":    []interface{}{},
			"recruitCityList": []interface{}{},
			"positionFidList": []interface{}{},
			"pageIndex":       i,
			"pageSize":        10,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, false, err
		}
		timestamp = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		err1 := c.PostRaw("https://join.qq.com/api/v1/position/searchPosition?"+timestamp, jsonBody)
		if err1 != nil {
			fmt.Println("访问链接地址出错了")
			return nil, false, err1
		}
		i++
		if len(test.Data.PositionList) < 1 {
			return AllTencent, true, nil
		}
	}
}

func Get1(each []Each) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = "GET"
		r.Headers.Set("authority", "join.qq.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("origin", "https://join.qq.com")
		r.Headers.Set("referer", "https://join.qq.com/post.html")
		r.Headers.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "Windows")
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})
	test := Content2{}
	c.OnResponse(func(r *colly.Response) {
		//var test interface{}
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("json解析失败1")
		}
		//if len(test.Data.PositionList) < 1 {
		//	return
		//}
		//fmt.Println(test)
		AllInformation = append(AllInformation, test.Data)
	})
	//fmt.Println(each)
	i := 0
	for {
		id := strconv.Itoa(each[i].ID)
		pid := strconv.Itoa(each[i].Pid)
		timestamp = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		err := c.Visit("https://join.qq.com/api/v1/jobDetails/getJobDetailsByPidAndId?" + timestamp + "&id=" + id + "&pid=" + pid)
		if err != nil {
			return
		}
		if err != nil {
			fmt.Println("访问链接地址出错了1")
		}
		if i == len(each)-1 {
			break
		}
		i++

	}
}
