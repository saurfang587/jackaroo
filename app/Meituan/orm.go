package Meituan

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
	"xiangxiang/jackaroo/global"
)

func Meituan_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(AllMeituan); i++ {
		utime := int64(AllMeituan[i].PushTime)
		uloc := time.FixedZone("CST", 8*3600)
		t := time.Unix(utime/1000, 0).In(uloc)
		now := t.Format("2006-01-02 15:04:05")
		str, _ := strconv.Atoi(AllMeituan[i].Id)
		if len(AllMeituan[i].WorkPlace) < 1 {
			continue
		}
		information := &global.Hello{
			ID:           str,
			Company:      "美团",
			Title:        AllMeituan[i].Title,
			JobCategory:  AllMeituan[i].Job_category,
			JobTypeName:  "实习生",
			JobDetail:    AllMeituan[i].Job_Detail + AllMeituan[i].Job_Obj,
			WorkLocation: global.Work{AllMeituan[i].WorkPlace[0].City},
			PushTime:     now,
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
