package jingdong

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"time"
	"xiangxiang/jackaroo/app/model"
)

var List1 []List

type Req struct {
	PageIndex int       `json:"pageIndex"`
	PageSize  int       `json:"pageSize"`
	Parameter parameter `json:"parameter"`
}

type parameter struct {
	jobDirectionCodeList []string
	planIdList           []string
	positionDeptList     []string
	positionName         string
	workCityCodeList     []string
}

type Rep struct {
	Data    Data   `json:"body"`
	Success string `json:"success"`
}

type Data struct {
	List []List `json:"items"`
}

type List struct {
	Id           string  `json:"jobCategoryCode"`
	PositionName string  `json:"positionName,omitempty"`
	JobCategory  string  `json:"jobCategory,omitempty"`
	JobObsity    string  `json:"qualification,omitempty"`
	JobDetail    string  `json:"workContent"`
	WorkCity     []Work1 `json:"requirementVoList"`
	Pushtime     int     `json:"publishTime"`
}

type Work1 struct {
	Location string `json:"workCity"`
}

type Location []string

func (l *Location) Scan(value interface{}) error {
	bytevalues := value.([]byte)
	return json.Unmarshal(bytevalues, l)
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func Header(cookie string) (bool, error) {
	pan, err := Get(cookie)
	if pan == false {
		return false, err
	}

	time1 := time.Now().Format("2006-01-02 15:04:05")
	jobs := []*model.Job{}
	for i := 0; i < len(List1); i++ {
		utime := int64(List1[i].Pushtime)
		uloc := time.FixedZone("CST", 8*3600)
		t := time.Unix(utime/1000, 0).In(uloc)
		now := t.Format("2006-01-02 15:04:05")
		worklocation := removeDuplicates(List1[i].WorkCity)
		job := &model.Job{
			ID:           List1[i].Id,
			Company:      "京东",
			Title:        List1[i].PositionName,
			JobCategory:  List1[i].JobCategory,
			JobTypeName:  "实习生",
			JobDetail:    List1[i].JobDetail + List1[i].JobObsity,
			WorkLocation: worklocation,
			PushTime:     now,
			FetchTime:    time1,
		}
		jobs = append(jobs, job)
	}
	pan1, err1 := model.UpdateJobs(jobs)
	if pan1 == false {
		return false, err1
	}
	return true, nil
}
func Get(cookie string) (bool, error) {
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
			List1 = append(List1, rep.Data.List[i])
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
			return false, err
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	return true, nil
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
