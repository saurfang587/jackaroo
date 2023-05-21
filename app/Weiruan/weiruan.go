package Weiruan

import (
	"fmt"
	"github.com/gocolly/colly"
)

var Allhref []string
var Title []string
var Job_category []string
var Job_detail []string
var WorkLocation []string

func Header(cookie string) (bool, error) {
	pan, err := Gethref()
	if pan == false {
		return false, err
	}
	for i, _ := range Allhref {
		GetIndex(Allhref[i])
	}
	pan1, err1 := Weiruan_orm()
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
