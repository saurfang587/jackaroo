package Weiruan

type Weiruan1 struct {
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

func (Weiruan1) TableName() string {
	return "weiruan"
}
