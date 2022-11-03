package api

import (
	"ToDoList/pkg/utils"
	"ToDoList/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//	CreateTask @Tags TASK
//
// @Summary create a task
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body service.CreateTaskService true  "title"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [post]
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.CheckToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

//	ShowTask @Tags TASK
//
// @Summary show the detail of one task
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body service.ShowTaskService true  "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [get]
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//	ListTask @Tags TASK
//
// @Summary get the task list
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body  service.ListTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /tasks [get]
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.CheckToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		utils.LogrusObj.Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//	UpdateTask @Tags TASK
//
// @Summary update a task
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body  service.UpdateTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [put]
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	//claim, _ := utils.CheckToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(updateTask.Title, updateTask.Content, updateTask.Status)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//	SearchTask @Tags TASK
//
// @Summary uncertainly search in all the task by key information
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body  service.SearchTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /search [post]
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.CheckToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		utils.LogrusObj.Info(err)
		c.JSON(400, ErrorResponse(err))
	}

}

//	DeleteTask @Tags TASK
//
// @Summary uncertainly search in all the task by key information
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "must"
// @Param data body  service.SearchTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [delete]
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	//claim, _ := utils.CheckToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		utils.LogrusObj.Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}
