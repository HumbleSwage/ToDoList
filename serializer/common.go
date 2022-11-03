package serializer

// Response 基础的一个序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TokenData 带有token的data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// BuildListResponse 带总数的返回
func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  item,
			Total: total,
		},
		Msg: "ok",
	}
}
