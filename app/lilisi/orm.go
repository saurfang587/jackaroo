package lilisi

import (
	"fmt"
	"time"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/model"
)

func lilisiOrm(list []List) {
	jobs := make([]model.Job, len(list))

	for i := 0; i < len(list); i++ {
		job := model.Job{
			Id:          list[i].Id,
			Company:     "lilisi",
			Title:       list[i].Title,
			JobDetail:   "Duty:\n" + list[i].Description + "\nRequire:\n" + list[i].Requirement,
			JobCategory: fmt.Sprintf("%v", list[i].JobCategory),
			JobTypeName: fmt.Sprintf("%v", list[i].RecruitType),
			JobLocation: fmt.Sprintf("%v", list[i].CityList),
			PushTime:    "",
			FetchTime:   time.Now(),
		}
		jobs[i] = job
	}
	global.G_DB.Create(jobs)
}
