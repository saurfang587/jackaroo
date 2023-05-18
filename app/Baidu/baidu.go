package Baidu

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

var AllInformation []Baidu

func Header(cookie string) (bool, error) {
	Get_time, pan, err := Get()
	if pan == false {
		return false, err
	}
	for i := 0; i < len(AllInformation); i++ {
		str, _ := strconv.Atoi(AllInformation[i].Id)
		information := &Alibaba.Hello{
			ID:            str,
			Company:       "百度",
			Title:         AllInformation[i].Title,
			Job_category:  AllInformation[i].Job_category,
			Job_type_name: "校招",
			Job_detail:    AllInformation[i].Job_Detail + AllInformation[i].Job_Obj,
			WorkLocation:  Alibaba.Work{AllInformation[i].WorkPlace},
			Fetch_time:    Get_time,
		}
		err1 := global.G_DB.Create(information).Error
		if err1 != nil {
			fmt.Println("插入数据失败了，请查看并修改错误")
			return false, err
		}
	}
	return true, nil
}
func Get() (string, bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	//创建请求
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	// 设置请求头
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("connection", "keep-alive")
		r.Headers.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
		r.Headers.Set("origin", "https://talent.baidu.com")
		r.Headers.Set("referer", "https://talent.baidu.com/jobs/list")
		r.Headers.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})

	test1 := Contont{}
	//返回数据
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test1)
		if err != nil {
			fmt.Println("解析失败，", err)
			return
		}
		if len(test1.Data.List) < 1 {
			return
		}
		AllInformation = append(AllInformation, test1.Data.List...)
		//fmt.Println(AllInformation
		//fmt.Println("这是输出", len(AllInformation))
		//在编程中，c.Wait()通常是指等待某个事件的发生并阻塞当前线程的执行，直到该事件完成或超时
	})
	//校招生爬取
	i := 1
	for {
		str := strconv.Itoa(i)
		c.Post("https://talent.baidu.com/httservice/getPostListNew", map[string]string{
			"recruitType": "INTERN",
			"pageSize":    "10",
			"keyWord":     "",
			"curPage":     str,
			"projectType": "",
		})
		i++
		if len(test1.Data.List) < 1 {
			break
		}
	}
	pan, err := Fetch1()
	if pan == false {
		return "", false, err
	}
	return time1, true, nil
}

// 实习生招聘
func Fetch1() (bool, error) {
	c := colly.NewCollector()
	// 设置请求头
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("connection", "keep-alive")
		r.Headers.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
		r.Headers.Set("origin", "https://talent.baidu.com")
		r.Headers.Set("referer", "https://talent.baidu.com/jobs/list")
		r.Headers.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	})
	//返回数据
	test := Contont{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("解析失败，", err)
			return
		}
		if len(test.Data.List) < 1 {
			return
		}
		AllInformation = append(AllInformation, test.Data.List...)
	})
	i := 1
	for {
		str := strconv.Itoa(i)
		c.Post("https://talent.baidu.com/httservice/getPostListNew", map[string]string{
			"recruitType": "GRADUATE",
			"pageSize":    "10",
			"keyWord":     "",
			"curPage":     str,
			"projectType": "",
		})
		i++
		if len(test.Data.List) < 1 {
			return true, nil
		}
	}
}
