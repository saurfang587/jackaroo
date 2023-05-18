package mihoyo

type IdRequest struct {
	ChannelDetailIds []int `json:"channelDetailIds"`
	PageNo           int   `json:"pageNo"`
	PageSize         int   `json:"pageSize"`
}

type Request struct {
	ChannelDetailIds []int  `json:"channelDetailIds"`
	Id               string `json:"id"`
}

type IdResponse struct {
	Code    int    `json:"code"`
	Data    IdData `json:"data"`
	Message string `json:"message"`
}

type DataResponse struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type IdData struct {
	List []*List `json:"list"`
}
type List struct {
	Id string `json:"id"`
}

type Data struct {
	ObjectId        int    `json:"objectId"`
	WorkLocation    []Work `json:"addressDetailList"`
	Job_Description string `json:"description"`
	Job_type_name   string `json:"jobNature"`
	Job_Require     string `json:"jobRequire"`
	Job_ObjectName  string `json:"objectName"`
	Title           string `json:"title"`
	Job_category    string `json:"competencyType"`
}
type Work struct {
	Location string `json:"addressDetail"`
}
type Mihoyo struct {
	Uuid          int    `gorm:"primaryKey;column:uuid"`
	ID            string `json:"id" column:"id"`
	Company       string `gorm:"column:company"`                   // 公司id
	Title         string ` gorm:"column:title"`                    //工作名字
	Job_category  string `gorm:"column:job_category"`              //工作类型
	Job_type_name string ` gorm:"column:job_type_name"`            //工作种类
	Job_detail    string ` gorm:"column:job_detail;type:longtext"` //工作职责
	WorkLocation  string `gorm:"column:job_location"`
	Fetch_time    string `gorm:"column:fetch_time"`
}

type AddressDetailList struct {
	AddressDetail string `json:"addressDetail"`
}
