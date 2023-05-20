package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
)

var Alllist1 []List1

func Header(cookie string) (bool, error) {
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

	//temp := List{}
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
	list1 := []List1{}
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
		Alllist1 = append(Alllist1, list1...)

		//for a := 0; a < len(r); a++ {
		//list1 = append(list1, *rep1[].Data[a])
		//fmt.Printf("data:%+V\n", list1)
		//}
	})

	for a := 0; a < len(list); a++ {
		//if list[a].JobRequire == "请您详见岗位意向中的岗位要求" || list[a].JobRequire == "专业知识要求：\\n/" {
		//	continue
		//}
		//temp = list[a]
		err := d.Visit("https://career.huawei.com/reccampportal/services/portal/portaluser/findIntentListByJobRequirementId/newHr/zh_CN/" + strconv.Itoa(list[a].JobRequirementId) + "/2?dataSource=" + strconv.Itoa(list[a].DataSource) + "&jobId=" + strconv.Itoa(list[a].JobId))
		if err != nil {
			fmt.Println(err)
			return false, err
		}
	}
	pan, err1 := huaweiOrmS(list, Alllist1)
	if pan == false {
		return false, err1
	}
	return true, nil
}
