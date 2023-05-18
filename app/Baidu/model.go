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

// -------------------------------
type Response struct {
	RecruitType string `json:"recruitType"`
	PageSize    int    `json:"pageSize"`
	KeyWord     string `json:"keyWord"`
	CurPage     int    `json:"curPage"`
	ProjectType string `json:"projectType"`
}
