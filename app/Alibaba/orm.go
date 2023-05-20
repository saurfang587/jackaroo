package Alibaba

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/global"
)

// 阿里巴巴 向数据库表中插入数据
func Alibaba_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(Allali); i++ {
		information := &global.Hello{
			ID:            Allali[i].Id,
			Company:       "阿里巴巴",
			Title:         Allali[i].Title,
			Job_category:  Allali[i].Job_category,
			Job_type_name: Allali[i].Job_type_name,
			Job_detail:    Allali[i].Job_Detail + Allali[i].Job_Obj,
			WorkLocation:  Allali[i].WorkLocation,
			PushTime:      Allali[i].PushTime,
			Fetch_time:    time1,
		}
		//首先查询是否存在 不存在就创建，存在的话就更新时间  对于时间超过1小时未做任何更改的数据，进行删除
		err3 := global.G_DB.Where("title= ? AND job_category=? AND job_type_name = ? AND job_location=?", information.Title, information.Job_category, information.Job_type_name, information.WorkLocation).First(&global.Hello{}).Error
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
