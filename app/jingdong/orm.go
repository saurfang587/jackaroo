package jingdong

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/global"
)

func Jingdong_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(List1); i++ {
		utime := int64(List1[i].Pushtime)
		uloc := time.FixedZone("CST", 8*3600)
		t := time.Unix(utime/1000, 0).In(uloc)
		now := t.Format("2006-01-02 15:04:05")
		worklocation := removeDuplicates(List1[i].WorkCity)
		information := &global.Hello{
			ID:           List1[i].Id,
			Company:      "京东",
			Title:        List1[i].PositionName,
			JobCategory:  List1[i].JobCategory,
			JobTypeName:  "实习生",
			JobDetail:    List1[i].Job_detail + List1[i].Job_obsity,
			WorkLocation: worklocation,
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
