package xiaohongshu

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Handler(cookie string) {
	if cookie == "" {
		cookie = "_pk_id.10.4f72=5fb50925a34d5c9e.1682607118.; Hm_lvt_900d393eff703909946efe28447affd3=1682607118; _pk_ref.10.4f72=%5B%22%22%2C%22%22%2C1683441558%2C%22https%3A%2F%2Flezhiyuan.feishu.cn%2F%22%5D; _pk_ses.10.4f72=1; Hm_lpvt_900d393eff703909946efe28447affd3=1683441558"
	}
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

	c.OnResponse(func(r *colly.Response) {
		rep := &Rep{}
		_ = json.Unmarshal(r.Body, rep)
		fmt.Printf("data:%+V", rep.Data)
	})
	req := &Req{
		RecruitType:  "campus",
		PositionName: "",
		JobType:      "all",
		Workplace:    "all",
		TimeSlotType: "all",
		PageNum:      1,
		PageSize:     10,
	}
	b, _ := json.Marshal(req)

	_ = c.PostRaw("https://job.xiaohongshu.com/websiterecruit/position/pageQueryPosition", b)
}
