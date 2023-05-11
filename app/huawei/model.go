package huawei

type Req struct {
	CurPage       int    `json:"curPage,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	JobTypes      int    `json:"jobTypes,omitempty"`
	JobType       int    `json:"jobType,omitempty"`
	JobFamClsCode string `json:"jobFamClsCode,omitempty"`
	SearchText    string `json:"searchText,omitempty"`
	CityCode      string `json:"cityCode,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	DeptCode      string `json:"deptCode,omitempty"`
	GraduateItem  string `json:"graduateItem,omitempty"`
	ReqTime       string `json:"reqTime,omitempty"`
	Language      string `json:"language,omitempty"`
	OrderBy       string `json:"orderBy,omitempty"`
}

type Rep struct {
	Data []*List `json:"result"`
}

type List struct {
	JobId         string `json:"jobId"`
	NameCn        string `json:"nameCn"`
	Jobname       string `json:"jobname"`
	MainBusiness  string `json:"mainBusiness"`
	JobRequire    string `json:"jobRequire"`
	JobFamilyName string `json:"jobFamilyName"`
	JobAddressId  string `json:"jobAddressId"`
	CityIds       string `json:"cityIds"`
	JobArea       string `json:"jobArea"`
}
