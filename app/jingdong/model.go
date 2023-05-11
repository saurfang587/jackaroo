package jingdong

type Req struct {
	PageIndex int       `json:"pageIndex"`
	PageSize  int       `json:"pageSize"`
	Parameter parameter `json:"parameter"`
}

type parameter struct {
	jobDirectionCodeList []string
	planIdList           []string
	positionDeptList     []string
	positionName         string
	workCityCodeList     []string
}

type Rep struct {
	Data    Data   `json:"body"`
	Success string `json:"success"`
}

type Data struct {
	List []*List `json:"items"`
}

type List struct {
	PositionName  string `json:"positionName,omitempty"`
	JobCategory   string `json:"jobCategory,omitempty"`
	JobDirection  string `json:"jobDirection,omitempty"`
	Qualification string `json:"qualification,omitempty"`
	WorkContent   string `json:"workContent,omitempty"`
	//WorkCity          string            `json:"workCity,omitempty"`
	RequirementVoList requirementVoList `json:"requirementVoList"`
}

type requirementVoList struct {
	InterviewCity string `json:"interviewCity,omitempty"`
	PositionBg    string `json:"positionBg,omitempty"`
	PositionDept  string `json:"positionDept,omitempty"`
	WorkCity      string `json:"workCity,omitempty"`
}
