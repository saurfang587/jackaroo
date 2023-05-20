package Bilibili

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
	"xiangxiang/jackaroo/global"
)

func Bilibili_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(AllBilibili); i++ {
		str, _ := strconv.Atoi(AllBilibili[i].ID)
		information := &global.Hello{
			ID:            str,
			Company:       "哔哩哔哩",
			Title:         AllBilibili[i].Title,
			Job_category:  AllBilibili[i].Job_category,
			Job_type_name: AllBilibili[i].Job_type_name,
			Job_detail:    AllBilibili[i].Job_detail,
			WorkLocation:  global.Work{AllBilibili[i].WorkLocation},
			PushTime:      AllBilibili[i].PushTime,
			Fetch_time:    time1,
		}
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
