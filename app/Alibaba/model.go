package Alibaba

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
	Id           int      `json:"batchId"`       //工作id
	Title        string   `json:"name"`          //工作名字
	Posttype     string   `json:"categoryType"`  //类型
	Job_category string   `json:"categoryName"`  //技术类型
	Push_time    string   `json:"publishTime"`   //发布时间
	Job_Obj      string   `json:"description"`   //工作要求
	Job_Detail   string   `json:"requirement"`   //工作内容 需要合并到上条中
	WorkPlace    []string `json:"workLocations"` //工作地点
}
