package Weiruan

import (
	"fmt"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/app/Alibaba"
	"xiangxiang/jackaroo/global"
)

var Allhref []string
var Title []string
var Job_category []string
var Job_detail []string
var WorkLocation []string

func Header(cookie string) (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	pan, err := Gethref()
	if pan == false {
		return false, err
	}
	for i, _ := range Allhref {
		GetIndex(Allhref[i])
	}
	for i := 0; i < len(Allhref); i++ {
		information := &Alibaba.Hello{
			ID:            0,
			Company:       "微软",
			Title:         Title[i],
			Job_category:  Job_category[i],
			Job_type_name: "实习生",
			Job_detail:    Job_detail[i],
			WorkLocation:  WorkLocation,
			Fetch_time:    time1,
		}
		time1 := time.Now().Format("2006-01-02 15:04:05")
		//首先查询是否存在 不存在就创建，存在的话就更新时间  对于时间超过1小时未做任何更改的数据，进行删除
		err3 := global.G_DB.Where("title=?", information.Title).First(&Alibaba.Hello{}).Error
		if err3 == gorm.ErrRecordNotFound {
			err1 := global.G_DB.Create(information).Error
			if err1 != nil {
				fmt.Println("插入数据失败了，请查看并修改错误")
				return false, err1
			}
		}
		err1 := global.G_DB.Where("title=?", information.Title).First(&Alibaba.Hello{}).Set("fetch_time", time1).Error
		if err1 != nil {
			fmt.Println("更新数据库中表的时间出错")
			return false, err1
		}
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
			Job_category = append(Job_category, element.Text)
			//fmt.Println(element.Text)
		})
		//实习地点
		e.ForEach(" div.info > div.extra.flex-row > span.el-tooltip", func(_ int, element *colly.HTMLElement) {
			WorkLocation = append(WorkLocation, element.Text)
			//fmt.Println(element.Text)
		})

		//职位要求
		e.ForEach("div.card.job-detail-word.mt-2.tw-flex > div > div:nth-child(5)", func(_ int, element *colly.HTMLElement) {
			Job_detail = append(Job_detail, element.Text)

		})
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("解析单个url出错了，")
		return
	}
}
