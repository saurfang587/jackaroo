package xiaomi

type Req struct {
	Keyword           string      `json:"keyword"`
	Limit             int         `json:"limit"`
	Offset            int         `json:"offset"`
	JobCategoryIdList interface{} `json:"job_category_id_list"`
	LocationCodeList  interface{} `json:"location_code_list"`
	SubjectIdList     interface{} `json:"subject_id_list"`
	RecruitmentIdList interface{} `json:"recruitment_id_list"`
	PortalType        int         `json:"portal_type"`
	JobFunctionIdList interface{} `json:"job_function_id_list"`
	PortalEntrance    int         `json:"portal_entrance"`
	_Signature        string      `json:"_signature"`
}

type Rep struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type SCRFRep struct {
	Code    int    `json:"code"`
	Data    Token  `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Count int     `json:"count"`
	List  []*List `json:"job_post_list"`
}
type Token struct {
	Token string `json:"token"`
}

type List struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	JobCategory JobCategory `json:"job_category"`
	Description string      `json:"description"`
	JobFunction JobFunction `json:"job_Function"`
	CityList    []City      `json:"city_list"`
	JobSubject  string      `json:"job_Subject"`
	Name        string      `json:"name"`
	RecruitType string      `json:"recruit_Type"`
	Requirement string      `json:"requirement"`
}

type City struct {
	Name string `json:"name,omitempty"`
}

type JobCategory struct {
	Name string `json:"name"`
}

type JobFunction struct {
	Name string `json:"name"`
}
