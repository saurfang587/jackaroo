package Bilibili

// ----------------------------
// 存储每页信息用到
type Context1 struct {
	Code    int      `json:"code"`
	Data    BiliBili `json:"data"`
	Message string   `json:"message"`
}
type BiliBili struct {
	id                  int    `gorm:"primaryKey";column:"ID"`
	PositionName        string `json:"positionName" gorm:"column:positionName"`               //工作名字
	PositionTypeName    string `json:"positionTypeName" gorm:"column:positionTypeName"`       //工作类型
	PostCodeName        string `json:"postCodeName" gorm:"column:postCodeName"`               //工作种类
	PositionDescription string `json:"positionDescription" gorm:"column:positionDescription"` //工作职责
	WorkLocation        string `json:"workLocation" gorm:"column:workLocation"`               //工作地点
	PushTime            string `json:"pushTime" gorm:"column:pushTime"`                       //发布时间
	WebApplyEndTime     string `json:"webApplyEndTime" gorm:"column:webApplyEndTime"`         //网申结束时间
	WebApplyStartTime   string `json:"webApplyStartTime" gorm:"column:webApplyStartTime"`     //网申开始时间
}

// ---------------------------
// 存储id时候用到
// 把读取到的数据存储到结构体中
// 请求到数据
type Context struct {
	Code    int    `json:"code"`
	Data    Kind   `json:"data"`
	Message string `json:"message"`
}

// Kind 和 ListObj  两个结构体的作用就是将Data中的数据一层一层刨析开，得到我们想要的数据
type Kind struct {
	PageNum  int       `json:"pageNum"`
	PageSize int       `json:"pageSize"`
	List     []ListObj `json:"list"`
}
type ListObj struct {
	Id string `json:"id"`
}
