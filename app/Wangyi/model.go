package Wangyi

type Content struct {
	Data Content1 `json:"data"`
}
type Content1 struct {
	List []Wangyi `json:"list"`
	Page int      `json:"pages"`
}
type Wangyi struct {
	Id                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Requirement       string   `json:"requirement"`
	ReqEducationName  string   `json:"reqEducationName"` // 学历
	FirstPostTypeName string   `json:"firstPostTypeName"`
	UpdateTime        int      `json:"updateTime"`
	WorkPlaceNameList []string `json:"workPlaceNameList"`
}
