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
	Data    *Data  `json:"data"`
	Message string `json:"message"`
}

type IdData struct {
	List []*List `json:"list"`
}
type List struct {
	Id string `json:"id"`
}

type Data struct {
	ObjectId          string            `json:"objectId"`
	AddressDetailList AddressDetailList `json:"addressDetailList"`
	Description       string            `json:"description"`
	JobNature         string            `json:"jobNature"`
	JobRequire        string            `json:"jobRequire"`
	ObjectName        string            `json:"objectName"`
	Title             string            `json:"title"`
}

type AddressDetailList struct {
	AddressDetail string `json:"addressDetail"`
}
