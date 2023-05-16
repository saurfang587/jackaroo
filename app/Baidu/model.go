package Baidu

type Contont struct {
	Code string `json:"status"`
	Data Kind   `json:"data"`
}
type Kind struct {
	List  []Baidu `json:"list"`
	Pages int     `json:"pages"`
}
type Baidu struct {
	Id           string `json:"jobId"`            //工作id
	Title        string `json:"name"`             //工作名字
	Job_category string `json:"postType"`         //技术类型
	Job_Obj      string `json:"serviceCondition"` //工作要求
	Job_Detail   string `json:"workContent"`      //工作内容
	WorkPlace    string `json:"workPlace"`        //工作地点
}
type Baidu1 struct {
	Uuid          int    `gorm:"primaryKey;column:uuid"`
	ID            string `json:"id" column:"id"`
	Company       string `gorm:"column:company"`                   // 公司id
	Title         string ` gorm:"column:title"`                    //工作名字
	Job_category  string `gorm:"column:job_category"`              //工作类型
	Job_type_name string ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  string `gorm:"column:job_location"`
	//Information BiliBili `gorm:"foreignKey:BiliBiliID_id"`
	//BiliBiliID  string   `gorm:"column:id"` // 外键
	Fetch_time string `gorm:"column:fetch_time"`
}

// -------------------------------
type Response struct {
	RecruitType string `json:"recruitType"`
	PageSize    int    `json:"pageSize"`
	KeyWord     string `json:"keyWord"`
	CurPage     int    `json:"curPage"`
	ProjectType string `json:"projectType"`
}

func (Baidu1) TableName() string {
	return "baidu"
}
