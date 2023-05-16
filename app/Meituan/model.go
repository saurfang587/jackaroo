package Meituan

import (
	"database/sql/driver"
	"encoding/json"
)

type Content struct {
	Data Content1 `json:"data"`
}
type Content1 struct {
	List []Meituan `json:"list"`
}
type Meituan struct {
	Id           string `json:"jobUnionId"` //工作id
	Title        string `json:"name"`       //工作名字
	Job_category string `json:"jobFamily"`  //技术类型
	Job_Obj      string `json:"highLight"`  //工作要求
	Job_Detail   string `json:"jobDuty"`    //工作内容 需要合并到上条中
	WorkPlace    []City `json:"cityList"`   //工作地点
}
type Meituan1 struct {
	Uuid          int    `gorm:"primaryKey;column:uuid"`
	ID            string `json:"id" column:"id"`
	Company       string `gorm:"column:company"`                   // 公司id
	Title         string ` gorm:"column:title"`                    //工作名字
	Job_category  string `gorm:"column:job_category"`              //工作类型
	Job_type_name string ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  string `gorm:"column:job_location"`
	Fetch_time    string `gorm:"column:fetch_time"`
}

type City struct {
	Name string `json:"name"`
}

func (t *City) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t City) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (Meituan1) TableName() string {
	return "meituan"
}
