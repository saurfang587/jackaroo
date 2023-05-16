package Wangyi

import (
	"database/sql/driver"
	"encoding/json"
)

type Content struct {
	Data Content1 `json:"data"`
}
type Content1 struct {
	List []Wangyi `json:"list"`
	Page int      `json:"pages"`
}
type Wangyi struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Job_detail        string `json:"description"`
	Job_Obisity       string `json:"requirement"`
	Job_category      string `json:"reqEducationName"` //
	FirstPostTypeName string `json:"firstPostTypeName"`
	WorkPlaceNameList Work   `json:"workPlaceNameList"`
}
type Wangyi1 struct {
	Uuid          int    `gorm:"primaryKey;column:uuid"`
	ID            int    `json:"id" column:"id"`
	Company       string `gorm:"column:company"`                   // 公司id
	Title         string ` gorm:"column:title"`                    //工作名字
	Job_category  string `gorm:"column:job_category"`              //工作类型
	Job_type_name string ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  string `gorm:"column:job_location"`
	Fetch_time    string `gorm:"column:fetch_time"`
}
type Work []string

func (w *Work) Scan(value interface{}) error {
	bytevalue := value.([]byte)
	return json.Unmarshal(bytevalue, w)
}
func (w Work) Value() (driver.Value, error) {
	return json.Marshal(w)
}
func (Wangyi1) TableName() string {
	return "wangyi"
}
