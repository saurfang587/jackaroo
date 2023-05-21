package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func UpdateJobs(jobs []*Job) (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		//首先查询是否存在 不存在就创建，存在的话就更新时间  对于时间超过1小时未做任何更改的数据，进行删除
		err3 := GlobalDb.Where("id = ? AND title= ? AND job_category=? AND job_type_name = ? AND job_location=?", job.ID, job.Title, job.JobCategory, job.JobTypeName, job.WorkLocation).First(&Job{}).Error
		if err3 == gorm.ErrRecordNotFound {
			//创建成功的话 就不用更新抓取时间了
			err1 := GlobalDb.Create(job).Error
			if err1 != nil {
				fmt.Println("插入数据失败了，请查看并修改错误")
				return false, err1
			}
			continue
		}
		err1 := GlobalDb.Where("title=?", job.Title).First(&Job{}).Set("fetch_time", time1).Error
		if err1 != nil {
			fmt.Println("更新数据库中表的时间出错")
			return false, err1
		}
	}
	return true, nil
}

type Job struct {
	Uuid         int    `gorm:"primaryKey;column:uuid"`
	ID           string `json:"id" column:"id"`
	Company      string `gorm:"column:company"`                   // 公司id
	Title        string ` gorm:"column:title"`                    //工作名字
	JobCategory  string `gorm:"column:job_category"`              //工作类型
	JobTypeName  string ` gorm:"column:job_type_name"`            //工作种类
	JobDetail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation Work   `gorm:"column:job_location"`
	PushTime     string `gorm:"push_time"`
	FetchTime    string `gorm:"column:fetch_time"`
}

func (Job) TableName() string {
	return "job"
}

type Work []string

func (t *Work) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Work) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// DeleteStaleRecords 删除过期数据
func DeleteStaleRecords(db *gorm.DB) {
	cutoff := time.Now().Add(-12 * time.Hour)
	var information []Job
	db.Where("fetch_time < ?", cutoff).Find(&information)
	for _, user := range information {
		db.Delete(&user)
	}
}
