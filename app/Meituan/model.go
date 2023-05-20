package Meituan

type Content struct {
	Data Content1 `json:"data"`
}
type Content1 struct {
	List []Meituan `json:"list"`
}
type Meituan struct {
	Id           string `json:"jobUnionId"` //工作id
	Title        string `json:"name"`       //工作名字
	Job_category string `json:"jobFamily"`  //技术类型
	Job_Obj      string `json:"highLight"`  //工作要求
	Job_Detail   string `json:"jobDuty"`    //工作内容 需要合并到上条中
	WorkPlace    []Work `json:"cityList"`   //工作地点
	PushTime     int    `json:"refreshTime"`
}

type Work struct {
	City string `json:"name"`
}
