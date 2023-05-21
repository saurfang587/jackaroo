package lilisi

import (
	"encoding/json"
	"fmt"
	"time"
	"xiangxiang/jackaroo/app/model"

	//"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

type Req struct {
	Keyword           string      `json:"keyword"`
	Limit             int         `json:"limit"`
	Offset            int         `json:"offset"`
	JobCategoryIdList interface{} `json:"job_category_id_list"`
	TagIdList         interface{} `json:"tag_id_list"`
	LocationCodeList  interface{} `json:"location_code_list"`
	SubjectIdList     interface{} `json:"subject_id_list"`
	RecruitmentIdList interface{} `json:"recruitment_id_list"`
	PortalType        int         `json:"portal_type"`
	JobFunctionIdList interface{} `json:"job_function_id_list"`
	PortalEntrance    int         `json:"portal_entrance"`
	_Signature        string      `json:"_signature"`
}

type Rep struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}
type SCRFRep struct {
	Code    int    `json:"code"`
	Data    Token  `json:"data"`
	Message string `json:"message"`
}

type Token struct {
	Token string `json:"token"`
}

type Data struct {
	Count int     `json:"count"`
	List  []*List `json:"job_post_list"`
}

type List struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	JobCategory interface{} `json:"job_category"`
	Description string      `json:"description"`
	CityList    []Work      `json:"city_list"`
	RecruitType interface{} `json:"recruit_Type"`
	Requirement string      `json:"requirement"`
}

type Work struct {
	Local string `json:"name"`
}

func Header(cookie string) (bool, error) {
	token := GetCSRF()
	if token == "" {
		return false, nil
	}
	if cookie == "" {
		cookie = " locale=zh-CN; channel=saas-career; platform=pc;atsx-csrf-token=" + token[:len(token)-1] + "%3D"
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("authority", "lilithgames.jobs.feishu.cn")
		r.Headers.Add("accept", "application/json, text/plain, */*")
		r.Headers.Add("accept-language", "zh-CN")
		r.Headers.Add("content-type", "application/json")
		r.Headers.Add("cookie", cookie)
		r.Headers.Add("env", "undefined")
		r.Headers.Add("origin", "https://lilithgames.jobs.feishu.cn")
		r.Headers.Add("portal-channel", "saas-career")
		r.Headers.Add("portal-platform", "pc")
		r.Headers.Add("referer", "https://lilithgames.jobs.feishu.cn/intern")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Add("sec-fetch-dest", "empty")
		r.Headers.Add("sec-fetch-mode", "cors")
		r.Headers.Add("sec-fetch-site", "same-origin")
		r.Headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("website-path", "intern")
		r.Headers.Add("x-csrf-token", token)

	})
	list := []List{}
	rep := &Rep{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, rep)
		if err != nil {

		}
		for i := 0; i < len(rep.Data.List); i++ {
			list = append(list, *rep.Data.List[i])
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

		err := c.PostRaw("https://lilithgames.jobs.feishu.cn/api/v1/search/job/posts", b)
		if err != nil {
			return false, err
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	time1 := time.Now().Format("2006-01-02 15:04:05")
	jobs := []*model.Job{}
	for i := 0; i < len(list); i++ {
		str := ""
		for j := 0; j < len(list[i].CityList); j++ {
			str += list[i].CityList[j].Local
		}
		job := &model.Job{
			ID:           list[i].Id,
			Company:      "lilisi",
			Title:        list[i].Title,
			JobDetail:    "Duty:\n" + list[i].Description + "\nRequire:\n" + list[i].Requirement,
			JobCategory:  fmt.Sprintf("%v", list[i].JobCategory),
			JobTypeName:  fmt.Sprintf("%v", list[i].RecruitType),
			WorkLocation: model.Work{str},
			PushTime:     "",
			FetchTime:    time1,
		}
		jobs = append(jobs, job)
	}

	pan, err1 := model.UpdateJobs(jobs)
	if pan == false {
		return false, err1
	}
	return true, nil
}

func GetCSRF() string {

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "application/json, text/plain, */*")
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/json")
		r.Headers.Add("Cookie", "channel=saas-career;platform=pc;")

		//r.Headers.Add("Cookie", "device-id=7230726326599222841;channel=saas-career;platform=pc;s_v_web_id=verify_lhfwm0nr_3iZm2ke4_MDU4_4pk2_B9iS_HL1BnUALfGsk;")
		r.Headers.Add("Origin", "https://lilithgames.jobs.feishu.cn")
		r.Headers.Add("Portal-Channel", "saas-career")
		r.Headers.Add("Portal-Platform", "pc")
		r.Headers.Add("Referer", "https://lilithgames.jobs.feishu.cn/intern")
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

	_ = c.PostRaw("https://lilithgames.jobs.feishu.cn/api/v1/csrf/token", nil)
	//fmt.Println(rep.Message)
	return rep.Data.Token
}
