package Bilibili

// ----------------------------
// 存储每页信息用到
type Context1 struct {
	Code    int      `json:"code"`
	Data    BiliBili `json:"data"`
	Message string   `json:"message"`
}
type BiliBili struct {
	ID            string `json:"id" column:"id"`                               // 公司id
	Title         string `json:"positionName" gorm:"column:title"`             //工作名字
	Job_category  string `json:"postCodeName" gorm:"column:job_category"`      //工作类型
	Job_type_name string `json:"positionTypeName" gorm:"column:job_type_name"` //工作种类
	Job_detail    string `json:"positionDescription" gorm:"column:job_detail"` //工作职责
	WorkLocation  string `json:"workLocation" gorm:"column:job_location"`      //工作地点
	PushTime      string `json:"pushTime"`
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
	PageNum  int        `json:"pageNum"`
	PageSize int        `json:"pageSize"`
	List     []BiliBili `json:"list"`
	Pages    int        `json:"pages"`
}

// 获取token
type SCRFRep struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
