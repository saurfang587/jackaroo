package zijie

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Handler(cookie string) {
	token := GetCSRF()
	fmt.Println(token)
	if cookie == "" {
		cookie = "locale=zh-CN;channel=campus; platform=pc; atsx-csrf-token=" + token[:len(token)-1] + "%3D"
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Content-Type", "application/json")
		r.Headers.Set("Cookie", cookie)
		r.Headers.Set("Origin", "https://jobs.bytedance.com")
		r.Headers.Set("Portal-Channel", "campus")
		r.Headers.Set("Portal-Platform", "pc")
		r.Headers.Set("Referer", "https://jobs.bytedance.com/campus/position?keywords=&category=&location=&type=3&job_hot_flag=")
		r.Headers.Set("Sec-Fetch-Dest", "empty")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Set("accept-language", "zh-CN")
		r.Headers.Set("env", "undefined")
		r.Headers.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Set("website-path", "campus")
		r.Headers.Set("x-csrf-token", token)

	})
	list := []List{}
	rep := &Rep{}
	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		for i := 0; i < len(rep.Data.List); i++ {
			list = append(list, *rep.Data.List[i])
			fmt.Printf("data:%+V", rep.Data)
		}
		//fmt.Println(rep.Data)
	})
	i := 0
	for {
		req := &Req{
			Keyword:           "",
			Limit:             10,
			Offset:            i * 10,
			JobCategoryIdList: nil,
			LocationCodeList:  nil,
			SubjectIdList:     nil,
			RecruitmentIdList: []string{"202", "301"},
			PortalType:        3,
			JobFunctionIdList: nil,
			PortalEntrance:    1,
			_Signature:        " Yx.j8QAAAAA73VPrvjzRZWMf4-AAAdG",
		}
		b, _ := json.Marshal(req)

		err := c.PostRaw("https://jobs.bytedance.com/api/v1/search/job/posts", b)
		if err != nil {
			fmt.Println("-=-=", err)
			return
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	fmt.Println(list)
}

func GetCSRF() string {

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "application/json, text/plain, */*")
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/json")
		r.Headers.Add("Cookie", "device-id=; locale=zh-CN; channel=saas-career; platform=pc; s_v_web_id=verify_lhfvryo8_c2HbON8L_mPXK_40o2_BjR4_jz6OOta5rcaU")
		r.Headers.Add("Origin", "https://jobs.bytedance.com")
		r.Headers.Add("Portal-Channel", "saas-career")
		r.Headers.Add("Portal-Platform", "pc")
		r.Headers.Add("Referer", "https://jobs.bytedance.com/internship/")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-origin")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("accept-language", "zh-CN")
		r.Headers.Add("env", "undefined")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Add("website-path", "internship")
		r.Headers.Add("x-csrf-token", "undefined")

	})
	rep := &SCRFRep{}
	c.OnResponse(func(r *colly.Response) {

		err := json.Unmarshal(r.Body, rep)
		if err != nil {
			fmt.Println("=-=-", err)
		}
		//fmt.Printf("data:%+V", rep.Data)
	})

	_ = c.PostRaw("https://jobs.bytedance.com/api/v1/csrf/token", nil)
	return rep.Data.Token
}
