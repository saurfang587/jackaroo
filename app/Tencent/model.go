package Tencent

import (
	"database/sql/driver"
	"encoding/json"
)

type Content struct {
	Data Content1 `json:"data"`
}
type Content1 struct {
	PositionList []Each `json:"positionList"`
}

// 用来存储id信息
type Each struct {
	Pid int `json:"projectId"`
	ID  int `json:"position"`
	Tid int `json:"positionFamily"`
}

// 用来存储工作
type Content2 struct {
	Data Content3 `json:"data"`
}
type Content3 struct {
	Id            int    `json:"id"`           //工作id
	Title         string `json:"title"`        //工作名字
	Job_type_name string `json:"projectName"`  //类型
	Job_Obj       string `json:"desc"`         //工作要求
	Job_Detail    string `json:"request"`      //工作内容 需要合并到上条中
	WorkPlace     Work   `json:"workCityList"` //工作地点
}
type Tencent1 struct {
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
func (Tencent1) TableName() string {
	return "tencent"
}
