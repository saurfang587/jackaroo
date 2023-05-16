package Alibaba

import (
	"database/sql/driver"
	"encoding/json"
)

// ----------------------------
// 存储每页信息用到
type Context1 struct {
	Data Context2 `json:"content"`
	//ErrorCode string   `json:"errorCode"`
	//ErrorMsg  string   `json:"errorMsg"`
	//Success   string   `json:"success"`
}
type Context2 struct {
	//CurrentPage int       `json:"currentPage"`
	Data1 []Alibaba `json:"datas"`
	//PageSize    int       `json:"pageSize"`
	//TotalCount  int       `json:"totalCount"`
}
type Alibaba struct {
	Id            int      `json:"batchId"`       //工作id
	Title         string   `json:"name"`          //工作名字
	Job_type_name string   `json:"batchName"`     //类型
	Job_category  string   `json:"categoryName"`  //技术类型
	Job_Obj       string   `json:"description"`   //工作要求
	Job_Detail    string   `json:"requirement"`   //工作内容 需要合并到上条中
	WorkLocation  []string `json:"workLocations"` //工作地点
}
type Alibaba1 struct {
	Uuid          int    `gorm:"primaryKey;column:uuid"`
	ID            int    `json:"id" column:"id"`
	Company       string `gorm:"column:company"`                   // 公司id
	Title         string ` gorm:"column:title"`                    //工作名字
	Job_category  string `gorm:"column:job_category"`              //工作类型
	Job_type_name string ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  Work   `gorm:"column:job_location"`
	//Information BiliBili `gorm:"foreignKey:BiliBiliID_id"`
	//BiliBiliID  string   `gorm:"column:id"` // 外键
	Fetch_time string `gorm:"column:fetch_time"`
}
type Work []string

func (t *Work) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Work) Value() (driver.Value, error) {
	return json.Marshal(t)
}
func (Alibaba1) TableName() string {
	return "alibaba"
}
