package huawei

import (
	"strconv"
	"time"
	"xiangxiang/jackaroo/global"
	"xiangxiang/jackaroo/model"
)

func huaweiOrm(list []List) {
	jobs := []model.Job{}

	for i := 0; i < len(list); i++ {
		if list[i].JobRequire == "请您详见岗位意向中的岗位要求" {
			continue
		}
		job := model.Job{
			Id:          strconv.Itoa(list[i].JobId) + " " + strconv.Itoa(list[i].JobRequirementId),
			Company:     "华为",
			Title:       list[i].Jobname,
			JobDetail:   "Duty:\n" + list[i].MainBusiness + "\nRequire:\n" + list[i].JobRequire,
			JobCategory: list[i].JobFamilyName,
			JobTypeName: "校园招聘",
			JobLocation: list[i].WorkArea,
			PushTime:    "",
			FetchTime:   time.Now(),
		}

		jobs = append(jobs, job)
	}

	global.G_DB.Create(&jobs)
}

func huaweiOrmS(list List, list1 []*List1) {

	for i := 0; i < len(list1); i++ {

		job := model.Job{
			Id:          strconv.Itoa(list.JobId) + " " + strconv.Itoa(list.JobRequirementId),
			Company:     "华为",
			Title:       list.Jobname + list1[i].DISPLAYNAME,
			JobDetail:   "Duty:\n" + list1[i].MAINBUSINESS + "\nRequire:\n" + list1[i].DEMAND,
			JobCategory: list.JobFamilyName,
			JobTypeName: "校园招聘",
			JobLocation: list1[i].LOCDESCS,
			PushTime:    "",
			FetchTime:   time.Now(),
		}
		global.G_DB.Create(&job)
	}
}
