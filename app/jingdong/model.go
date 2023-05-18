package jingdong

import (
	"database/sql/driver"
	"encoding/json"
)

type Req struct {
	PageIndex int       `json:"pageIndex"`
	PageSize  int       `json:"pageSize"`
	Parameter parameter `json:"parameter"`
}

type parameter struct {
	jobDirectionCodeList []string
	planIdList           []string
	positionDeptList     []string
	positionName         string
	workCityCodeList     []string
}

type Rep struct {
	Data    Data   `json:"body"`
	Success string `json:"success"`
}

type Data struct {
	List []List `json:"items"`
}

type List struct {
	Id           int     `json:"jobCategoryCode"`
	PositionName string  `json:"positionName,omitempty"`
	JobCategory  string  `json:"jobCategory,omitempty"`
	Job_obsity   string  `json:"qualification,omitempty"`
	Job_detail   string  `json:"workContent"`
	WorkCity     []Work1 `json:"requirementVoList"`
}

type Work1 struct {
	Location string `json:"workCity"`
}

type Jingdong struct {
	Uuid          int      `gorm:"primaryKey;column:uuid"`
	ID            string   `json:"id" column:"id"`
	Company       string   `gorm:"column:company"`                   // 公司id
	Title         string   ` gorm:"column:title"`                    //工作名字
	Job_category  string   `gorm:"column:job_category"`              //工作类型
	Job_type_name string   ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string   ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  Location `gorm:"column:job_location"`
	Fetch_time    string   `gorm:"column:fetch_time"`
}
type Location []string

func (l *Location) Scan(value interface{}) error {
	bytevalues := value.([]byte)
	return json.Unmarshal(bytevalues, l)
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}
func (Jingdong) TableName() string {
	return "jingdong"
}
