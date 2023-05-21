package huawei

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
	CurPage       int    `json:"curPage,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	JobTypes      int    `json:"jobTypes,omitempty"`
	JobType       int    `json:"jobType,omitempty"`
	JobFamClsCode string `json:"jobFamClsCode,omitempty"`
	SearchText    string `json:"searchText,omitempty"`
	CityCode      string `json:"cityCode,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	DeptCode      string `json:"deptCode,omitempty"`
	GraduateItem  string `json:"graduateItem,omitempty"`
	ReqTime       string `json:"reqTime,omitempty"`
	Language      string `json:"language,omitempty"`
	OrderBy       string `json:"orderBy,omitempty"`
}

type Rep struct {
	Data []*List `json:"result"`
}

type Rep1 struct {
	Data []*List1
}

type List struct {
	GraduateItem     int    `json:"graduateItem"`
	DataSource       int    `json:"dataSource"`
	JobId            int    `json:"jobId"`
	JobRequirementId int    `json:"jobRequirementId"`
	NameCn           string `json:"nameCn"`
	Jobname          string `json:"jobname"`
	MainBusiness     string `json:"mainBusiness"`
	JobRequire       string `json:"jobRequire"`
	JobFamilyName    string `json:"jobFamilyName"`
	JobAddressId     string `json:"jobAddressId"`
	CityIds          string `json:"cityIds"`
	JobArea          string `json:"jobArea"`
	WorkArea         string `json:"workArea"`
	CreationDate     string `json:"creationDate"`
}

type List1 struct {
	DEMAND       string `json:"DEMAND"`
	DEPTCODES    string `json:"DEPTCODES"`
	DISPLAYNAME  string `json:"DISPLAYNAME"`
	LOCDESCS     string `json:"LOCDESCS"`
	MAINBUSINESS string `json:"MAINBUSINESS"`
}

func Header(cookie string) (bool, error) {
	list := []List{}
	list1 := []List1{}
	rep := &Rep{}
	i := 1
	j := 0

	if cookie == "" {
		cookie = "channel_name=cn.bing.com; hwsso_login="
	}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "*/*")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Add("Content-Type", "application/json")
		r.Headers.Add("Referer", "https://career.huawei.com/reccampportal/portal5/campus-recruitment.html?jobTypes=2")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-origin")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Add("x-csrf-token", "undefined")
	})

	c.OnResponse(func(r *colly.Response) {
		_ = json.Unmarshal(r.Body, rep)
		for i := 0; i < len(rep.Data); i++ {
			list = append(list, *rep.Data[i])
			//fmt.Printf("data:%+V\n", rep.Data[i].JobId)
			//fmt.Printf("ID:%+V\n", rep.Data[i].JobRequirementId)
		}
	})

	for {
		a := strconv.Itoa(i)
		b := strconv.Itoa(j)
		url := "https://career.huawei.com/reccampportal/services/portal/portalpub/getJob/newHr/page/10/" + a + "?curPage=" + a + "&pageSize=10&jobTypes=" + b + "&jobType=0&jobFamClsCode=&searchText=&cityCode=&countryCode=&deptCode=&graduateItem=&language=zh_CN"
		err := c.Visit(url)
		if err != nil {
			return false, err
		}

		i++
		if len(rep.Data) < 1 {

			if j == 2 {
				break
			}
			j = 2
			i = 1
		}
	}
	//huaweiOrm(list)

	temp := List{}
	d := colly.NewCollector()
	d.OnRequest(func(r *colly.Request) {
		r.Method = http.MethodPost
		r.Headers.Add("Accept", "*/*")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Add("Content-Type", "application/json")
		r.Headers.Add("Referer", "https://career.huawei.com/reccampportal/portal5/campus-recruitment.html?jobTypes=2")
		r.Headers.Add("Sec-Fetch-Dest", "empty")
		r.Headers.Add("Sec-Fetch-Mode", "cors")
		r.Headers.Add("Sec-Fetch-Site", "same-origin")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
		r.Headers.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		r.Headers.Add("sec-ch-ua-mobile", "?0")
		r.Headers.Add("sec-ch-ua-platform", "\"Windows\"")
		r.Headers.Add("x-csrf-token", "undefined")
	})
	//rep1 := &Rep1{}

	d.OnResponse(func(r *colly.Response) {

		//fmt.Println(r.Body)
		err := json.Unmarshal(r.Body, &list1)
		if err != nil {
			fmt.Println(err)
			return
		}
		if list1 == nil {
			return
		}

	})

	for a := 0; a < len(list); a++ {

		err := d.Visit("https://career.huawei.com/reccampportal/services/portal/portaluser/findIntentListByJobRequirementId/newHr/zh_CN/" + strconv.Itoa(list[a].JobRequirementId) + "/2?dataSource=" + strconv.Itoa(list[a].DataSource) + "&jobId=" + strconv.Itoa(list[a].JobId))
		if err != nil {
			fmt.Println(err)
			return false, err
		}
	}

	time1 := time.Now().Format("2006-01-02 15:04:05")
	jobs := []*model.Job{}
	for i := 0; i < len(list1); i++ {
		job := &model.Job{
			ID:           strconv.Itoa(temp.JobId) + " " + strconv.Itoa(temp.JobRequirementId),
			Company:      "华为",
			Title:        temp.Jobname + list1[i].DISPLAYNAME,
			JobDetail:    "Duty:\n" + list1[i].MAINBUSINESS + "\nRequire:\n" + list1[i].DEMAND,
			JobCategory:  temp.JobFamilyName,
			JobTypeName:  "校园招聘",
			WorkLocation: model.Work{list1[i].LOCDESCS},
			PushTime:     temp.CreationDate,
			FetchTime:    time1,
		}
		fmt.Println()
		jobs = append(jobs, job)
	}
	pan, err1 := model.UpdateJobs(jobs)
	if pan == false {
		return false, err1
	}
	return true, nil
}
