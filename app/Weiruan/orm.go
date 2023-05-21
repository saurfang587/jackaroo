package Weiruan

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/global"
)

func Weiruan_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(Allhref); i++ {
		information := &global.Hello{
			ID:           0,
			Company:      "微软",
			Title:        Title[i],
			JobCategory:  Job_category[i],
			JobTypeName:  "实习生",
			JobDetail:    Job_detail[i],
			WorkLocation: WorkLocation,
			FetchTime:    time1,
		}
		err3 := global.G_DB.Where("title= ? AND job_category=? AND job_type_name = ? AND job_location=?", information.Title, information.JobCategory, information.JobTypeName, information.WorkLocation).First(&global.Hello{}).Error
		if err3 == gorm.ErrRecordNotFound {
			//创建成功的话 就不用更新抓取时间了
			err1 := global.G_DB.Create(information).Error
			if err1 != nil {
				fmt.Println("插入数据失败了，请查看并修改错误")
				return false, err1
			}
			continue
		}
		err1 := global.G_DB.Where("title=?", information.Title).First(&global.Hello{}).Set("fetch_time", time1).Error
		if err1 != nil {
			fmt.Println("更新数据库中表的时间出错")
			return false, err1
		}
	}
	return true, nil
}
