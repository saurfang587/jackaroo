package _60

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Header(cookie string) (bool, error) {
	list := []List{}
	rep := &Rep{}
	i := 0
	j := 2
	if cookie == "" {
		cookie = "acw_tc=2760820416836328501224677eec1e85f4fd0a7ea63a3eacfa8050bf904c24"
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "application/json")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/json")
		//r.Headers.Add("Cookie", cookie)
		//r.Headers.Add("EagleEye-TraceID", "fb469dd0-082d-cbe6-1693-a5d8d7571f5e")
		r.Headers.Add("Origin", "https://360campus.zhiye.com")
		r.Headers.Add("Referer", "https://360campus.zhiye.com/jobs")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-origin")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("X-Requested-With", "xmlhttprequest")
		r.Headers.Add("langType", "zh_CN")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")

	})

	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		for i := 0; i < len(rep.Data); i++ {
			list = append(list, *rep.Data[i])
		}
	})

	for {
		req := &Req{
			Category: j,
			DisplayFields: []string{"Category",
				"Kind",
				"LocId",
				"Org",
				"HeadCount",
				"Station",
				"EndTime",
				"PostDate",
				"Salary",
				"ClassificationOne",
				"ClassificationTwo",
				"WorkWeChatQrCode"},
			KeyWords:    "",
			PageIndex:   i,
			PageSize:    20,
			PortalId:    "",
			SpecialType: 0,
		}
		b, _ := json.Marshal(req)

		err := c.PostRaw("https://360campus.zhiye.com/api/Jobad/GetJobAdPageList", b)
		if err != nil {
			fmt.Println("-=-=", err)
			return false, err
		}
		i++
		if len(rep.Data) < 1 {

			if j == 3 {
				break
			}
			j++
			i = 0
		}
	}
	pan, err := _60Orm(list)
	if pan == false {
		return false, err
	}
	return true, nil
}
