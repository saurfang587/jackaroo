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
	Job_detail        string   `json:"description"`
	Job_Obisity       string   `json:"requirement"`
	Job_category      string   `json:"reqEducationName"` //
	FirstPostTypeName string   `json:"firstPostTypeName"`
	WorkPlaceNameList []string `json:"workPlaceNameList"`
	PushTime          int      `json:"updateTime"`
}
