package service

import (
	"ToDoList/pkg/ctl"
	"ToDoList/pkg/utils"
	"ToDoList/repository/db/dao"
	"ToDoList/repository/db/model"
	"ToDoList/types"
	"context"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (us *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, err = userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		user := &model.User{
			UserName: req.UserName,
		}
		fmt.Println(req.PassWord)
		if err := user.SetPassword(req.PassWord); err != nil {
			utils.LogrusObj.Info(err)
			return nil, err
		}
		if err = userDao.CreateUser(user); err != nil {
			utils.LogrusObj.Info(err)
			return nil, err
		}
		return ctl.RespSuccess(), nil
	case nil:
		return ctl.RespSuccessWithData("用户已经存在"), nil
	default:
		return
	}
}

func (us *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		return ctl.RespSuccessWithData("请先注册再登录"), nil
	case nil:
		if user.CheckPassword(req.PassWord) {
			// 分发token
			token, err := utils.GenerateToken(user.Id, user.UserName)
			if err != nil {
				return nil, err
			}
			data := ctl.TokenData{
				User: types.UserLoginResp{
					Id:       user.Id,
					UserName: user.UserName,
					CreateAt: user.CreatedAt.Unix(),
				},
				AccessToken: token,
			}
			return ctl.RespSuccessWithData(data), nil
		} else {
			return ctl.RespSuccessWithData("账号/密码错误"), nil
		}
	default:
		return ctl.RespSuccessWithData("登录时发生未知错误"), nil
	}
}
