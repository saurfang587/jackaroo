package xiaohongshu

type Req struct {
	RecruitType  string `json:"recruitType"`
	PositionName string `json:"positionName"`
	JobType      string `json:"jobType"`
	Workplace    string `json:"workplace"`
	TimeSlotType string `json:"timeSlotType"`
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
}

type Rep struct {
	AlertMsg   string      `json:"alertMsg"`
	Data       Data        `json:"data"`
	ErrorCode  interface{} `json:"errorCode"`
	ErrorMsg   interface{} `json:"errorMsg"`
	ExtMap     interface{} `json:"extMap"`
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	TraceLogID string      `json:"traceLogId"`
}
type List struct {
	AmountInNeed  string      `json:"amountInNeed"`  //岗位需求
	Duty          string      `json:"duty"`          //岗位职责
	JobType       string      `json:"jobType"`       //工作类型
	Labels        interface{} `json:"labels"`        //标签
	PositionID    int         `json:"positionId"`    //主ID
	PositionName  string      `json:"positionName"`  //岗位名称
	PublishTime   string      `json:"publishTime"`   //推送时间
	Qualification string      `json:"qualification"` //岗位要求
	RecruitStatus interface{} `json:"recruitStatus"`
	Workplace     string      `json:"workplace"` //工作地点
}
type Data struct {
	List      []*List `json:"list"`
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	Total     int     `json:"total"`
	TotalPage int     `json:"totalPage"`
}
