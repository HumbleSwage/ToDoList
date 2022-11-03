package service

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/pkg/utils"
	"ToDoList/serializer"
	"time"
)

// CreateTaskService 创建任务的服务
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

// ShowTaskService 展示一条任务服务
type ShowTaskService struct {
}

// ListTaskService 展示所有任务服务
type ListTaskService struct {
	Limit int `json:"limit" form:"limit"`
	Start int `json:"start" form:"start"`
}

// UpdateTaskService 更新任务服务
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

// SearchTaskService 模糊搜索服务
type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageSize int    `json:"page_size" from:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

// DeleteTaskService 删除任务服务
type DeleteTaskService struct {
}

// Create 新增一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := e.SUCCESS
	model.DB.Model(&model.User{}).First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTask(task),
	}
}

// Show 展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	if err := model.DB.First(&task, tid).Error; err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	task.AddView() //增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

// List 展示该用户所有的备忘录
func (service *ListTaskService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total)),
	}
}

// Update 更新指定的用户
func (service *UpdateTaskService) Update(title string, content string, status int) serializer.Response {
	var task model.Task
	model.DB.First(&task)
	task.Title = title
	task.Content = content
	task.Status = status
	code := e.SUCCESS
	if err := model.DB.Save(&task).Error; err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTask(task),
	}
}

// Search 关键字搜索
func (service *SearchTaskService) Search(id uint) serializer.Response {
	var tasks []model.Task
	code := e.SUCCESS
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", id).Count(&count).
		Where("title Like ? Or content Like ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count)),
	}
}

// Delete 删除任务
func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.Delete(&task, id).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
