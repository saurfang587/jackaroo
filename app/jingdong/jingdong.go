package jingdong

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

func Header(cookie string) (bool, error) {
	list1, pan, err := Get(cookie)
	if pan == false {
		return false, err
	}
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(list1); i++ {
		worklocation := removeDuplicates(list1[i].WorkCity)
		information := &Alibaba.Hello{
			ID:            list1[i].Id,
			Company:       "京东",
			Title:         list1[i].PositionName,
			Job_category:  list1[i].JobCategory,
			Job_type_name: "实习生",
			Job_detail:    list1[i].Job_detail + list1[i].Job_obsity,
			WorkLocation:  worklocation,
			Fetch_time:    time1,
		}
		err1 := global.G_DB.Create(information).Error
		if err1 != nil {
			fmt.Println("插入数据失败了，请查看并修改错误")
			return false, err
		}
	}
	return true, nil
}
func Get(cookie string) ([]List, bool, error) {
	list := []List{}
	rep := &Rep{}
	i := 0
	if cookie == "" {
		cookie = ""
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("authority", "campus.jd.com")
		r.Headers.Add("accept", "*/*")
		r.Headers.Add("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Add("content-type", "application/json; charset=UTF-8")
		r.Headers.Add("origin", "https://campus.jd.com")
		r.Headers.Add("referer", "https://campus.jd.com/home")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Add("sec-fetch-dest", "empty")
		r.Headers.Add("sec-fetch-mode", "cors")
		r.Headers.Add("sec-fetch-site", "same-origin")
		r.Headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("x-requested-with", "XMLHttpRequest")
	})

	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		//fmt.Println("=-=-=-=", rep.Success)
		for i := 0; i < len(rep.Data.List); i++ {
			list = append(list, rep.Data.List[i])
		}
	})

	for {
		req := &Req{
			PageIndex: i,
			PageSize:  10,
			Parameter: parameter{
				jobDirectionCodeList: nil,
				planIdList:           nil,
				positionDeptList:     nil,
				positionName:         "",
				workCityCodeList:     nil,
			},
		}
		b, _ := json.Marshal(req)

		url := "https://campus.jd.com/api/wx/position/page?type=present"
		err := c.PostRaw(url, b)
		if err != nil {
			fmt.Println("-=-=", err)
			return nil, false, err
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	return list, true, nil
}

// 去除重复元素
func removeDuplicates(nums []Work1) []string {
	m := make(map[string]bool)
	result := []string{}
	for _, num := range nums {
		if !m[num.Location] {
			m[num.Location] = true
			result = append(result, num.Location)
		}
	}
	return result
}
