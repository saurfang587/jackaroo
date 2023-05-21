package global

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
	"xiangxiang/jackaroo/config"
)

var (
	G_CONFIG config.Server
	G_VP     *viper.Viper
	G_DB     *gorm.DB
)

// 定义公共的结构体
type Hello struct {
	Uuid         int    `gorm:"primaryKey;column:uuid"`
	ID           int    `json:"id" column:"id"`
	Company      string `gorm:"column:company"`                   // 公司id
	Title        string ` gorm:"column:title"`                    //工作名字
	JobCategory  string `gorm:"column:job_category"`              //工作类型
	JobTypeName  string ` gorm:"column:job_type_name"`            //工作种类
	JobDetail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation Work   `gorm:"column:job_location"`
	PushTime     string `gorm:"push_time"`
	FetchTime    string `gorm:"column:fetch_time"`
}

func (Hello) TableName() string {
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

// 删除过期数据
func DeleteStaleRecords(db *gorm.DB) {
	cutoff := time.Now().Add(-12 * time.Hour)
	var information []Hello
	db.Where("fetch_time < ?", cutoff).Find(&information)
	for _, user := range information {
		db.Delete(&user)
	}
}
