package Bilibili

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

var AllBilibili []BiliBili

func Header(cookie string) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	Get()
	err := global.G_DB.AutoMigrate(&Alibaba.Hello{})
	if err != nil {
		fmt.Println("数据库迁移失败")
	}
	for i := 0; i < len(AllBilibili); i++ {
		str, _ := strconv.Atoi(AllBilibili[i].ID)
		information := &Alibaba.Hello{
			ID:            str,
			Company:       "哔哩哔哩",
			Title:         AllBilibili[i].Title,
			Job_category:  AllBilibili[i].Job_category,
			Job_type_name: AllBilibili[i].Job_type_name,
			Job_detail:    AllBilibili[i].Job_detail,
			WorkLocation:  Alibaba.Work{AllBilibili[i].WorkLocation},
			Fetch_time:    time1,
		}
		err1 := global.G_DB.Create(information).Error
		if err1 != nil {
			fmt.Println("插入数据失败了，请查看并修改错误")
			return
		}
	}
}
func Get() {
	token := GetAdd()
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
		//fmt.Println(test.Data.List)
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
			return
		}
		err = c.PostRaw("https://jobs.bilibili.com/api/campus/position/positionList", jsonData)
		if err != nil {
			fmt.Println("请求失败：", err)
		}
		i++
		if len(test.Data.List) < 1 {
			return
		}
	}
}

//// 爬取每个id
//func FetchEach(num int) {
//	token := GetAdd()
//	c := colly.NewCollector()
//	// 设置请求头
//	c.OnRequest(func(r *colly.Request) {
//		r.Headers.Set("authority", "jobs.bilibili.com")
//		r.Headers.Set("accept", "application/json, text/plain, */*")
//		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
//		r.Headers.Set("content-type", "application/json;charset=UTF-8")
//		r.Headers.Set("origin", "https://jobs.bilibili.com")
//		r.Headers.Set("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
//		r.Headers.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
//		r.Headers.Set("sec-ch-ua-mobile", "?0")
//		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
//		r.Headers.Set("sec-fetch-dest", "empty")
//		r.Headers.Set("sec-fetch-mode", "cors")
//		r.Headers.Set("sec-fetch-site", "same-origin")
//		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
//		r.Headers.Set("x-appkey", "ops.ehr-api.auth")
//		r.Headers.Set("x-channel", "campus")
//		r.Headers.Set("x-csrf", token)
//		r.Headers.Set("x-usertype", "2")
//	})
//	c.OnResponse(func(r *colly.Response) {
//		var result Context
//		err := json.Unmarshal(r.Body, &result)
//		if err != nil {
//			fmt.Println("解析json数据失败：", err)
//			return
//		}
//		for i := 0; i < 10; i++ {
//			if len(result.Data.List) == i {
//				return
//			}
//			Allid = append(Allid, result.Data.List[i].Id)
//		}
//		//爬取每页信息
//		FetchInformation()
//	})
//	for i := 1; i <= num; i++ {
//		str := i
//		// 发送POST请求
//		data := map[string]interface{}{
//			"pageSize":         10,
//			"pageNum":          str,
//			"positionName":     "",
//			"postCode":         []interface{}{},
//			"postCodeList":     []interface{}{},
//			"workLocationList": []interface{}{},
//			"workTypeList":     []string{"0"},
//			"positionTypeList": []string{"0"},
//			"deptCodeList":     []interface{}{},
//			"recruitType":      nil,
//		}
//		jsonData, err := json.Marshal(data)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		err = c.PostRaw("https://jobs.bilibili.com/api/campus/position/positionList", jsonData)
//		if err != nil {
//			fmt.Println("请求第", i, "页失败：", err)
//		}
//	}
//}
//
//// 每个页面信息抓取
//func FetchInformation() {
//	c := colly.NewCollector()
//	// 设置请求头
//	c.OnRequest(func(r *colly.Request) {
//		r.Headers.Set("authority", "jobs.bilibili.com")
//		r.Headers.Set("accept", "application/json, text/plain, */*")
//		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
//		r.Headers.Set("content-type", "application/json;charset=UTF-8")
//		r.Headers.Set("origin", "https://jobs.bilibili.com")
//		r.Headers.Set("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
//		r.Headers.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
//		r.Headers.Set("sec-ch-ua-mobile", "?0")
//		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
//		r.Headers.Set("sec-fetch-dest", "empty")
//		r.Headers.Set("sec-fetch-mode", "cors")
//		r.Headers.Set("sec-fetch-site", "same-origin")
//		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
//		r.Headers.Set("x-appkey", "ops.ehr-api.auth")
//		r.Headers.Set("x-channel", "campus")
//		r.Headers.Set("x-csrf", "728e3c6c-e981-412e-be94-ab1f76557972")
//		r.Headers.Set("x-usertype", "2")
//	})
//	c.OnResponse(func(r *colly.Response) {
//		var res Context1
//		err := json.Unmarshal(r.Body, &res)
//		if err != nil {
//			fmt.Println("解析json数据失败：", err)
//			return
//		}
//		AllBilibili.Information = append(AllBilibili.Information, res.Data)
//	})
//	//len(Allid)
//	for i := 0; i < len(Allid); i++ {
//		url := "https://jobs.bilibili.com/api/campus/position/detail/" + Allid[i]
//		err := c.Visit(url)
//		if err != nil {
//			fmt.Println("请求url地址过程出错，请修改重试")
//			return
//		}
//	}
//}

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
