package types

type CreateTaskReq struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

type ListTaskReq struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type GetTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type UpdateTaskReq struct {
	Id        uint   `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

type DeleteTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type SearchTaskReq struct {
	Key string `json:"key" form:"key"`
}

type ListTaskResp struct {
	Id        uint   `json:"id" form:"id"`
	Uid       uint   `json:"uid" form:"uid"`
	UserName  string `json:"user_name" form:"user_name"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"`
	View      int    `json:"view" form:"view"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

type GetTaskResp struct {
	Id        uint   `json:"id" form:"id"`
	UserName  string `json:"user_name" form:"user_name"`
	Title     string `json:"title" form:"json"`
	Status    int    `json:"status" form:"json"`
	Content   string `json:"content" form:"json"`
	View      int    `json:"view" form:"view"`
	StartTime int64  `json:"start_time" form:"end_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

type SearchTaskResp struct {
	Id        uint   `json:"id" form:"id"`
	Uid       uint   `json:"uid" form:"uid"`
	UserName  string `json:"user_name" form:"user_name"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"`
	View      int    `json:"view" form:"view"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}
