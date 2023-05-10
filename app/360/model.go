package _60

type Req struct {
	Category      int      `json:"Category"`
	DisplayFields []string `json:"DisplayFields"`
	KeyWords      string   `json:"KeyWords"`
	PageIndex     int      `json:"PageIndex"`
	PageSize      int      `json:"PageSize"`
	PortalId      string   `json:"PortalId"`
	SpecialType   int      `json:"SpecialType"`
}

type Rep struct {
	Code    int     `json:"code"`
	Data    []*List `json:"data"`
	Message string  `json:"message"`
}

type List struct {
	Category          string   `json:"Category,omitempty"`
	CategoryId        string   `json:"CategoryId,omitempty"`
	ClassificationOne string   `json:"ClassificationOne,omitempty"`
	Duty              string   `json:"Duty,omitempty"`
	JobAdName         string   `json:"JobAdName,omitempty"`
	Kind              string   `json:"Kind,omitempty"`
	LocNames          []string `json:"LocNames,omitempty"`
	Org               string   `json:"Org,omitempty"`
	Require           string   `json:"Require,omitempty"`
	Salary            string   `json:"Salary,omitempty"`
}
