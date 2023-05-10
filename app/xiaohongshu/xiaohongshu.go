package xiaohongshu

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Handler(cookie string) {
	//if cookie == "" {
	//	cookie = "_pk_id.10.4f72=5fb50925a34d5c9e.1682607118.; Hm_lvt_900d393eff703909946efe28447affd3=1682607118; _pk_ref.10.4f72=%5B%22%22%2C%22%22%2C1683441558%2C%22https%3A%2F%2Flezhiyuan.feishu.cn%2F%22%5D; _pk_ses.10.4f72=1; Hm_lpvt_900d393eff703909946efe28447affd3=1683441558"
	//}
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
	//defer func() { // 必须要先声明defer，否则不能捕获到panic异常
	//	fmt.Println("捕获小红书页数越界")
	//	if err := recover(); err != nil {
	//		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	//	}
	//	fmt.Println("捕获成功")
	//}()
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
		}
		i++
		if len(rep.Data.List) < 1 {
			break
		}
	}
	fmt.Println(list)

}
