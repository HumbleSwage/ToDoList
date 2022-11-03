package service

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/pkg/utils"
	"ToDoList/serializer"
	"github.com/jinzhu/gorm"
)

// UserService 用户注册服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	// 表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	// 加密密码
	if err := user.SetPassWord(service.PassWord); err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
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

// Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			utils.LogrusObj.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		// 如果不是用户不存在，是其他不可抗拒的因素导致的错误
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if user.CheckPassWord(service.PassWord) == false {
		//密码验证成功
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误，登陆失败",
		}
	}
	// 为了其他功能，这里给前端分发一个token
	// 比如说创建一个备忘录，就需要知道是谁创建的这个备忘录
	token, err := utils.GenerateToken(user.ID, service.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登陆成功",
	}
}
