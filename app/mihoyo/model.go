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

type AddressDetailList struct {
	AddressDetail string `json:"addressDetail"`
}
