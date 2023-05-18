package Meituan

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

var AllMeituan []Meituan

func Header(cookie string) {
	Get()
	time1 := time.Now().Format("2006-01-02 15:04:05")
	err1 := global.G_DB.AutoMigrate(&Alibaba.Hello{})
	if err1 != nil {
		fmt.Println("数据库迁移失败")
	}
	for i := 0; i < len(AllMeituan); i++ {
		str, _ := strconv.Atoi(AllMeituan[i].Id)
		if len(AllMeituan[i].WorkPlace) < 1 {
			continue
		}
		information := &Alibaba.Hello{
			ID:            str,
			Company:       "美团",
			Title:         AllMeituan[i].Title,
			Job_category:  AllMeituan[i].Job_category,
			Job_type_name: "实习生",
			Job_detail:    AllMeituan[i].Job_Detail + AllMeituan[i].Job_Obj,
			WorkLocation:  Alibaba.Work{AllMeituan[i].WorkPlace[0].City},
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
	test := Content{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &test)
		if err != nil {
			fmt.Println("json数据解析失败", err)
		}
		if len(test.Data.List) < 1 {
			return
		}
		//if len(test.Data.List) < 10 {
		//	for i := 0; i < len(test.Data.List); i++ {
		//		if test.Data.List[i].WorkPlace[0].City == "" {
		//			return
		//		}
		//	}
		//}
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
		Data, err := json.Marshal(data)
		if err != nil {
			fmt.Println("将json解压为byte数组出错")
			return
		}
		err1 := c.PostRaw("https://zhaopin.meituan.com/api/official/job/getJobList", Data)
		if err1 != nil {
			fmt.Println("解析url请求地址出错，请你继续尝试", err1)
			return
		}
		i++
		if len(test.Data.List) < 1 {
			return
		}
	}

}
