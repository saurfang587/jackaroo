package xiaohongshu

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"time"
	"xiangxiang/jackaroo/app/model"
)

type Req struct {
	RecruitType  string `json:"recruitType"`
	PositionName string `json:"positionName"`
	JobType      string `json:"jobType"`
	Workplace    string `json:"workplace"`
	TimeSlotType string `json:"timeSlotType"`
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
}

type Rep struct {
	AlertMsg   string      `json:"alertMsg"`
	Data       Data        `json:"data"`
	ErrorCode  interface{} `json:"errorCode"`
	ErrorMsg   interface{} `json:"errorMsg"`
	ExtMap     interface{} `json:"extMap"`
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	TraceLogID string      `json:"traceLogId"`
}

type List struct {
	AmountInNeed  string      `json:"amountInNeed"`  //岗位需求
	Duty          string      `json:"duty"`          //岗位职责
	JobType       string      `json:"jobType"`       //工作类型
	Labels        interface{} `json:"labels"`        //标签
	PositionID    int         `json:"positionId"`    //主ID
	PositionName  string      `json:"positionName"`  //岗位名称
	PublishTime   string      `json:"publishTime"`   //推送时间
	Qualification string      `json:"qualification"` //岗位要求
	RecruitStatus interface{} `json:"recruitStatus"`
	Workplace     string      `json:"workplace"` //工作地点
}
type Data struct {
	List      []*List `json:"list"`
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	Total     int     `json:"total"`
	TotalPage int     `json:"totalPage"`
}

func Header(cookie string) (bool, error) {

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Set("authority", "job.xiaohongshu.com")
		r.Headers.Set("accept", "application/json, text/plain, */*")
		r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,sm;q=0.8,en;q=0.7")
		r.Headers.Set("authorization", "")
		r.Headers.Set("content-type", "application/json;charset=UTF-8")
		r.Headers.Set("cookie", cookie)
		r.Headers.Set("origin", "https://job.xiaohongshu.com")
		r.Headers.Set("referer", "https://job.xiaohongshu.com/campus")
		r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "macOS")
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
		//r.Headers.Set("x-b3-traceid", "bbb360b113377a7f")
		//r.Headers.Set("x-s", "OjO6sB5LO6F+O6sKOBAKsjU6ZBTG1idvZjsGO6ak0js3")
		//r.Headers.Set("x-t", "1683441575791")

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

	i := 1
	for {
		req := &Req{
			RecruitType:  "campus",
			PositionName: "",
			JobType:      "all",
			Workplace:    "all",
			TimeSlotType: "all",
			PageNum:      i,
			PageSize:     10,
		}
		b, _ := json.Marshal(req)

		err := c.PostRaw("https://job.xiaohongshu.com/websiterecruit/position/pageQueryPosition", b)
		if err != nil {
			fmt.Println("=-=-=", err)
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
		job := &model.Job{
			ID:           strconv.Itoa(list[i].PositionID),
			Company:      "xiaohongshu",
			Title:        list[i].PositionName,
			JobDetail:    "Duty:\n" + list[i].Duty + "\nRequire:\n" + list[i].Qualification,
			JobCategory:  list[i].JobType,
			JobTypeName:  list[i].JobType,
			WorkLocation: model.Work{list[i].Workplace},
			PushTime:     list[i].PublishTime,
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
