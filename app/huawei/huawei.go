package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
)

func Handler(cookie string) {

	list := []List{}
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
			//fmt.Printf("data:%+V\n", rep.Data)
		}
	})

	for {

		//req := &Req{
		//	CurPage:       i,
		//	PageSize:      10,
		//	JobTypes:      j,
		//	JobType:       0,
		//	JobFamClsCode: "",
		//	SearchText:    "",
		//	CityCode:      "",
		//	CountryCode:   "",
		//	DeptCode:      "",
		//	GraduateItem:  "",
		//	ReqTime:       "",
		//	Language:      "zh_CN",
		//	OrderBy:       "",
		//}

		//b, _ := json.Marshal(req)
		a := strconv.Itoa(i)
		b := strconv.Itoa(j)
		url := "https://career.huawei.com/reccampportal/services/portal/portalpub/getJob/newHr/page/10/" + a + "?curPage=" + a + "&pageSize=10&jobTypes=" + b + "&jobType=0&jobFamClsCode=&searchText=&cityCode=&countryCode=&deptCode=&graduateItem=&language=zh_CN"
		err := c.Visit(url)
		if err != nil {
			fmt.Println("-=-=", err)
			return
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

	fmt.Println(len(list))

}
