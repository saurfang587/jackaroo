package _60

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/global"
)

func _60Orm(list []List) (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(list); i++ {
		job := &global.Hello{
			ID:           list[i].Id,
			Company:      "360",
			Title:        list[i].JobAdName,
			JobCategory:  list[i].Category,
			JobDetail:    "Duty:\n" + list[i].Duty + "\nRequire:\n" + list[i].Require,
			JobTypeName:  list[i].ClassificationOne,
			WorkLocation: list[i].LocNames,
			PushTime:     "",
			FetchTime:    time1,
		}
		//首先查询是否存在 不存在就创建，存在的话就更新时间  对于时间超过1小时未做任何更改的数据，进行删除
		err3 := global.G_DB.Where("title= ? AND job_category=? AND job_type_name = ? AND job_location=?", job.Title, job.JobCategory, job.JobTypeName, job.WorkLocation).First(&global.Hello{}).Error
		if err3 == gorm.ErrRecordNotFound {
			//创建成功的话 就不用更新抓取时间了
			err1 := global.G_DB.Create(job).Error
			if err1 != nil {
				fmt.Println("插入数据失败了，请查看并修改错误")
				return false, err1
			}
			continue
		}
		err1 := global.G_DB.Where("title=?", job.Title).First(&global.Hello{}).Set("fetch_time", time1).Error
		if err1 != nil {
			fmt.Println("更新数据库中表的时间出错")
			return false, err1
		}
	}
	return true, nil
}
