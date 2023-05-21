package Weiruan

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
	"xiangxiang/jackaroo/app/model"
)

type Weiruan1 struct {
	Uuid         int      `gorm:"primaryKey;column:uuid"`
	ID           int      `json:"id" column:"id"`
	Company      string   `gorm:"column:company"`                   // 公司id
	Title        string   ` gorm:"column:title"`                    //工作名字
	JobCategory  string   `gorm:"column:job_category"`              //工作类型
	JobTypeName  string   ` gorm:"column:job_type_name"`            //工作种类
	JobDetail    string   ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation []string `gorm:"column:job_location"`
	FetchTime    string   `gorm:"column:fetch_time"`
}

var Allhref []string
var Title []string
var JobCategory []string
var JobDetail []string
var WorkLocation []string

func Header(cookie string) (bool, error) {
	pan, err := Gethref()
	if pan == false {
		return false, err
	}
	for i, _ := range Allhref {
		GetIndex(Allhref[i])
	}
	time1 := time.Now().Format("2006-01-02 15:04:05")
	jobs := []*model.Job{}
	for i := 0; i < len(Allhref); i++ {
		job := &model.Job{
			ID:           "",
			Company:      "微软",
			Title:        Title[i],
			JobCategory:  JobCategory[i],
			JobTypeName:  "实习生",
			JobDetail:    JobDetail[i],
			WorkLocation: WorkLocation,
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
func Gethref() (bool, error) {
	c := colly.NewCollector()
	c.OnHTML("#jsApp > main > div:nth-child(2) > div > div > div.campus-company-main > div.jobs-wrap > div.left > div.tw-relative.dark\\:tw-bg-\\[\\#313540\\] > div.tw-h-full > div", func(div *colly.HTMLElement) {
		div.ForEach("a.job-message-boxs", func(_ int, element *colly.HTMLElement) {
			Allhref = append(Allhref, element.Attr("href"))
		})
	})
	err := c.Visit("https://www.nowcoder.com/enterprise/146")
	if err != nil {
		fmt.Println("url请求失败")
		return false, err
	}
	return true, nil
}

func GetIndex(url string) {
	c := colly.NewCollector()
	c.OnHTML("#app > main > div > div > div.tw-flex.tw-items-start.tw-pb-10 > div.tw-flex-auto", func(e *colly.HTMLElement) {
		//标题
		e.ForEach("h1.title", func(_ int, element *colly.HTMLElement) {
			Title = append(Title, element.Text)
			//fmt.Println(element.Text)
		})
		//实习岗位
		e.ForEach(" div.info > div.extra.flex-row >span:nth-child(1)", func(_ int, element *colly.HTMLElement) {
			JobCategory = append(JobCategory, element.Text)
			//fmt.Println(element.Text)
		})
		//实习地点
		e.ForEach(" div.info > div.extra.flex-row > span.el-tooltip", func(_ int, element *colly.HTMLElement) {
			WorkLocation = append(WorkLocation, element.Text)
			//fmt.Println(element.Text)
		})

		//职位要求
		e.ForEach("div.card.job-detail-word.mt-2.tw-flex > div > div:nth-child(5)", func(_ int, element *colly.HTMLElement) {
			JobDetail = append(JobDetail, element.Text)

		})
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("解析单个url出错了，")
		return
	}
}
