package xiaomi

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Handler(cookie string) {
	token := GetCSRF()
	if cookie == "" {
		cookie = "device-id=; locale=zh-CN; channel=saas-career; platform=pc; s_v_web_id=verify_lhe9e1pw_KV7VxS41_R6aN_4Riw_BglK_Q5ouWlR3Id6U; atsx-csrf-token=" + token[:len(token)-1] + "%3D"
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Content-Type", "application/json")
		r.Headers.Set("Cookie", cookie)
		r.Headers.Set("Origin", "https://xiaomi.jobs.f.mioffice.cn")
		r.Headers.Set("Portal-Channel", "saas-career")
		r.Headers.Set("Portal-Platform", "pc")
		r.Headers.Set("Referer", "https://xiaomi.jobs.f.mioffice.cn/internship/?spread=6AA3R7B")
		r.Headers.Set("Sec-Fetch-Dest", "empty")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Set("accept-language", "zh-CN")
		r.Headers.Set("env", "undefined")
		r.Headers.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Set("website-path", "internship")
		r.Headers.Set("x-csrf-token", token)

	})
	list := []List{}
	rep := &Rep{}
	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		for i := 0; i < len(rep.Data.List); i++ {
			list = append(list, *rep.Data.List[i])
			//fmt.Printf("data:%+V", rep.Data)
		}
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
			RecruitmentIdList: nil,
			PortalType:        0,
			JobFunctionIdList: nil,
			PortalEntrance:    0,
			_Signature:        "",
		}
		b, _ := json.Marshal(req)

		_ = c.PostRaw("https://xiaomi.jobs.f.mioffice.cn/api/v1/search/job/posts?keyword=&limit=10&offset=0&job_category_id_list=&location_code_list=&subject_id_list=&recruitment_id_list=&portal_type=6&job_function_id_list=&portal_entrance=1&_signature=LGyWRwAAAAB0riZdHOafHSxsllAAEg5", b)
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
		r.Headers.Add("Origin", "https://xiaomi.jobs.f.mioffice.cn")
		r.Headers.Add("Portal-Channel", "saas-career")
		r.Headers.Add("Portal-Platform", "pc")
		r.Headers.Add("Referer", "https://xiaomi.jobs.f.mioffice.cn/internship/")
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
		fmt.Printf("data:%+V", rep.Data)
	})

	_ = c.PostRaw("https://xiaomi.jobs.f.mioffice.cn/api/v1/csrf/token", nil)
	return rep.Data.Token
}
