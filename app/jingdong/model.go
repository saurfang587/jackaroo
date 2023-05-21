package jingdong

import (
	"database/sql/driver"
	"encoding/json"
)

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
	List []List `json:"items"`
}

type List struct {
	Id           int     `json:"jobCategoryCode"`
	PositionName string  `json:"positionName,omitempty"`
	JobCategory  string  `json:"jobCategory,omitempty"`
	Job_obsity   string  `json:"qualification,omitempty"`
	Job_detail   string  `json:"workContent"`
	WorkCity     []Work1 `json:"requirementVoList"`
	Pushtime     int     `json:"publishTime"`
}

type Work1 struct {
	Location string `json:"workCity"`
}

type Location []string

func (l *Location) Scan(value interface{}) error {
	bytevalues := value.([]byte)
	return json.Unmarshal(bytevalues, l)
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}
