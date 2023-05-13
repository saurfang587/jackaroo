package Tencent

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
	Id         int      `json:"id"`           //工作id
	Title      string   `json:"title"`        //工作名字
	Posttype   string   `json:"projectName"`  //类型
	Job_Obj    string   `json:"desc"`         //工作要求
	Job_Detail string   `json:"request"`      //工作内容 需要合并到上条中
	WorkPlace  []string `json:"workCityList"` //工作地点
}
