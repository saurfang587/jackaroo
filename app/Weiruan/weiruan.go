package Weiruan

import (
	"fmt"
	"github.com/gocolly/colly"
)

var Allhref []string

func Head(cookie string) {
	Gethref()
	if len(Allhref) == 0 {
		fmt.Println("此时解析失败了，请你重新尝试吧")
	}
	//fmt.Println(Allhref[0])
	for _, v := range Allhref {
		GetIndex(v)
	}

}
func Gethref() {
	c := colly.NewCollector()
	c.OnHTML("#jsApp > main > div:nth-child(2) > div > div > div.campus-company-main > div.jobs-wrap > div.left > div.tw-relative.dark\\:tw-bg-\\[\\#313540\\] > div.tw-h-full > div", func(div *colly.HTMLElement) {
		div.ForEach("a.job-message-boxs", func(_ int, element *colly.HTMLElement) {
			Allhref = append(Allhref, element.Attr("href"))
		})
	})
	err := c.Visit("https://www.nowcoder.com/enterprise/146")
	if err != nil {
		fmt.Println("url请求失败")
		return
	}
}

func GetIndex(url string) {
	c := colly.NewCollector()
	c.OnHTML("#app > main > div > div > div.tw-flex.tw-items-start.tw-pb-10 > div.tw-flex-auto", func(e *colly.HTMLElement) {
		//标题
		e.ForEach("h1.title", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
		//实习工资
		e.ForEach("div.salary", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
		//其他要求
		e.ForEach("div.info >div:nth-child(2) ", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
		//投递时间
		e.ForEach("div.deliver-range", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
		//职位要求
		e.ForEach("div:nth-child(2)", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
	})
	err := c.Visit("https://www.nowcoder.com/jobs/detail/122396?pageSource=5012&channel=newCompanyPage")
	if err != nil {
		fmt.Println("解析单个url出错了，")
		return
	}
}
