package Alibaba

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/json-iterator/go"
	"gopkg.in/headzoo/surf.v1"
	"strconv"
)

var Allali []Alibaba

// 阿里实习生
func Header(cookie string) (bool, error) {
	// 创建一个新的浏览器对象
	browser := surf.NewBrowser()
	// 打开目标页面
	err := browser.Open("https://talent.alibaba.com/campus/position-list?campusType=internship&lang=zh")
	if err != nil {
		return false, err
	}
	// 获取 Cookie
	cookies := browser.SiteCookies()
	cookie1 := cookies[0].String()[11:]
	pan, err2 := Get(cookie1)
	pan1 := Get1(cookie1)
	if pan == false || pan1 == false {
		return false, err2
	}
	pan2, err3 := Alibaba_orm()
	if pan2 == false {
		return false, err3
	}
	return true, nil
}

// 获取实习生
func Get(cookie1 string) (bool, error) {
	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("authority", "talent.alibaba.com")
		req.Headers.Set("accept", "application/json, text/plain, */*")
		req.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Headers.Set("bx-v", "2.2.3")
		req.Headers.Set("content-type", "application/json")
		req.Headers.Set("cookie", "XSRF-TOKEN="+cookie1)
		req.Headers.Set("origin", "https://talent.alibaba.com")
		req.Headers.Set("referer", "https://talent.alibaba.com/campus/position-list?campusType=internship&lang=zh")
		req.Headers.Set("sec-ch-ua", "\"Microsoft Edge\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		req.Headers.Set("sec-ch-ua-mobile", "?0")
		req.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		req.Headers.Set("sec-fetch-dest", "empty")
		req.Headers.Set("sec-fetch-mode", "cors")
		req.Headers.Set("sec-fetch-site", "same-origin")
		req.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})
	test := Context1{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("json解析错误", err)
			return
		}
		if len(test.Data.Data1) < 1 {
			return
		}
		Allali = append(Allali, test.Data.Data1...)
	})
	i := 1
	for {
		str := strconv.Itoa(i)
		res := map[string]string{
			"channel":       "campus_group_official_site",
			"language":      "zh",
			"pageSize":      "10",
			"batchId":       "",
			"subCategories": "",
			"regions":       "",
			"bgCode":        "",
			"corpCode":      "",
			"pageIndex":     str,
			"key":           "",
			"categoryType":  "internship",
		}
		Data, _ := json.Marshal(res)
		err := c.PostRaw("https://talent.alibaba.com/position/search?_csrf="+cookie1, Data)
		if err != nil {
			fmt.Printf("抓取第: %d出错", i)
			return false, err
		}
		i++
		if len(test.Data.Data1) < 1 {
			return true, nil
		}
	}
}

// 项目制实习生
func Get1(cookie string) bool {
	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("authority", "talent.alibaba.com")
		req.Headers.Set("accept", "application/json, text/plain, */*")
		req.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Headers.Set("bx-v", "2.2.3")
		req.Headers.Set("content-type", "application/json")
		req.Headers.Set("cookie", "XSRF-TOKEN="+cookie)
		req.Headers.Set("origin", "https://talent.alibaba.com")
		req.Headers.Set("referer", "https://talent.alibaba.com/campus/position-list?campusType=internship&lang=zh")
		req.Headers.Set("sec-ch-ua", "\"Microsoft Edge\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		req.Headers.Set("sec-ch-ua-mobile", "?0")
		req.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		req.Headers.Set("sec-fetch-dest", "empty")
		req.Headers.Set("sec-fetch-mode", "cors")
		req.Headers.Set("sec-fetch-site", "same-origin")
		req.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})
	test := Context1{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("json解析错误", err)
			return
		}
		if len(test.Data.Data1) < 1 {
			return
		}
		Allali = append(Allali, test.Data.Data1...)
	})
	i := 1
	for {
		str := strconv.Itoa(i)
		res := map[string]string{
			"channel":       "campus_group_official_site",
			"language":      "zh",
			"pageSize":      "10",
			"batchId":       "",
			"subCategories": "",
			"regions":       "",
			"bgCode":        "",
			"corpCode":      "",
			"pageIndex":     str,
			"key":           "",
			"categoryType":  "talentPlan",
		}
		Data, _ := json.Marshal(res)
		err := c.PostRaw("https://talent.alibaba.com/position/search?_csrf="+cookie, Data)
		if err != nil {
			fmt.Printf("抓取第: %d出错", i)
			return false
		}
		i++
		if len(test.Data.Data1) < 1 {
			return true
		}
	}
}
