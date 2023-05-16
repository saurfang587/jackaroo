package _60

import (
	"time"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/model"
)

func _60Orm(list []List) {
	jobs := make([]model.Job, len(list))

	for i := 0; i < len(list); i++ {
		job := model.Job{
			Id:          list[i].Id,
			Company:     "360",
			Title:       list[i].JobAdName,
			JobDetail:   "Duty:\n" + list[i].Duty + "\nRequire:\n" + list[i].Require,
			JobCategory: list[i].Category,
			JobTypeName: list[i].ClassificationOne,
			JobLocation: list[i].LocNames[0],
			PushTime:    "",
			FetchTime:   time.Now(),
		}
		jobs[i] = job
	}
	global.G_DB.Create(jobs)
}
