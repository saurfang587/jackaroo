package mihoyo

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

var list []string
var list2 []Data

func Handler(cookie string) {
	Get(cookie)
	time1 := time.Now().Format("2006-01-02 15:04:05")
	err1 := global.G_DB.AutoMigrate(&Alibaba.Hello{})
	if err1 != nil {
		fmt.Println("数据库迁移失败")
	}
	for i := 0; i < len(list2); i++ {
		information := &Alibaba.Hello{
			ID:            list2[i].ObjectId,
			Company:       "米哈游",
			Title:         list2[i].Title,
			Job_category:  list2[i].Job_category,
			Job_type_name: list2[i].Job_type_name,
			Job_detail:    list2[i].Job_Description + list2[i].Job_Require + list2[i].Job_ObjectName,
			WorkLocation:  Alibaba.Work{list2[i].WorkLocation[0].Location},
			Fetch_time:    time1,
		}
		err1 := global.G_DB.Create(information).Error
		if err1 != nil {
			fmt.Println("插入数据失败了，请查看并修改错误")
			return
		}
	}

}
func Get(cookie string) {
	rep := &IdResponse{}
	i := 1

	if cookie == "" {
		cookie = ""
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "application/json")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Add("Authorization", "null")
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Add("Origin", "https://campus.mihoyo.com")
		r.Headers.Add("Referer", "https://campus.mihoyo.com/")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-site")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("bucket-name", "undefined")
		r.Headers.Add("current-request", "request")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		//r.Headers.Add("x-request-id", "front-1683791534755-4452-1370-407263067")
	})

	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		for i := 0; i < len(rep.Data.List); i++ {
			list = append(list, rep.Data.List[i].Id)
		}
	})

	for {

		req := &IdRequest{
			ChannelDetailIds: []int{1},
			PageNo:           i,
			PageSize:         10,
		}
		b, _ := json.Marshal(req)
		url := "https://ats.openout.mihoyo.com/ats-portal/v1/job/list"
		err := c.PostRaw(url, b)
		if err != nil {
			fmt.Println("-=-=", err)
			return
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	Fetch(list)
}
func Fetch(list []string) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "application/json")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Add("Authorization", "null")
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Add("Origin", "https://campus.mihoyo.com")
		r.Headers.Add("Referer", "https://campus.mihoyo.com/")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-site")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("bucket-name", "undefined")
		r.Headers.Add("current-request", "request")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		//r.Headers.Add("x-request-id", "front-1683791534755-4452-1370-407263067")
	})
	test := DataResponse{}
	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, &test)
		list2 = append(list2, test.Data)
	})

	for j := 0; j < len(list); j++ {
		req := &Request{
			ChannelDetailIds: []int{1},
			Id:               list[j],
		}

		b, _ := json.Marshal(req)
		url := "https://ats.openout.mihoyo.com/ats-portal/v1/job/info"
		err := c.PostRaw(url, b)
		if err != nil {
			return
		}
		//url := "https://ats.openout.mihoyo.com/ats-portal/v1/job/info"

	}
}
