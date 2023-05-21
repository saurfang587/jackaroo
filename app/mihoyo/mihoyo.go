package mihoyo

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

var list []string
var list2 []Data

func Header(cookie string) (bool, error) {
	pan, err := Get(cookie)
	if pan == false {
		return false, err
	}
	pan1, err1 := Mihoyo_orm()
	if pan1 == false {
		return false, err1
	}
	return true, nil
}
func Get(cookie string) (bool, error) {
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
			return false, err
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	pan, err := Fetch(list)
	if pan == false {
		return false, err
	}
	return true, nil
}
func Fetch(list []string) (bool, error) {
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
			return false, err
		}
	}
	return true, nil
}
