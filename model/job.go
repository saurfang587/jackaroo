package model

import "time"

type Job struct {
	Id          string    `json:"id"`
	Company     string    `json:"company"`
	Title       string    `json:"title"`
	JobDetail   string    `json:"job_detail"`
	JobCategory string    `json:"job_category"`
	JobTypeName string    `json:"job_type_name"`
	JobLocation string    `json:"job_location"`
	PushTime    string    `json:"push_time"`
	FetchTime   time.Time `json:"fetch_time"`
}

func (Job) TableName() string {
	return "Job"
}
