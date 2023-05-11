package jingdong

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Handler(cookie string) {

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
			list = append(list, *rep.Data.List[i])
			fmt.Printf("data:%+V\n", rep.Data)
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
			return
		}

		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}

	fmt.Println(len(list))

}
