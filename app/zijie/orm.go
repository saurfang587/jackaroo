package zijie

import (
	"fmt"
	"strings"
	"time"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/model"
)

func zijieOrm(list []List) {
	jobs := make([]model.Job, len(list))

	for i := 0; i < len(list); i++ {
		job := model.Job{
			Id:          list[i].Id,
			Company:     "zijie",
			Title:       list[i].Title,
			JobDetail:   "Duty:\n" + list[i].Description + "\nRequire:\n" + list[i].Requirement,
			JobCategory: list[i].JobCategory.Name,
			JobTypeName: list[i].JobSubject.Name.Name,
			JobLocation: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list[i].City)), ","), "[]"),
			PushTime:    "",
			FetchTime:   time.Now(),
		}
		jobs[i] = job
	}
	global.G_DB.Create(jobs)
}
