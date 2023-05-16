package xiaohongshu

import (
	"strconv"
	"time"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/model"
)

func xiaohongshuOrm(list []List) {
	jobs := make([]model.Job, len(list))

	for i := 0; i < len(list); i++ {
		job := model.Job{
			Id:          strconv.Itoa(list[i].PositionID),
			Company:     "xiaohongshu",
			Title:       list[i].PositionName,
			JobDetail:   "Duty:\n" + list[i].Duty + "\nRequire:\n" + list[i].Qualification,
			JobCategory: list[i].JobType,
			JobTypeName: list[i].JobType,
			JobLocation: list[i].Workplace,
			PushTime:    list[i].PublishTime,
			FetchTime:   time.Now(),
		}
		jobs[i] = job
	}
	global.G_DB.Create(jobs)
}
