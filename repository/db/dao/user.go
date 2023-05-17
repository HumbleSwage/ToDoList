package dao

import (
	"ToDoList/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewMySQLClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (ud *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = ud.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}

func (ud *UserDao) FindUserByUserName(name string) (user *model.User, err error) {
	err = ud.DB.Model(&model.User{}).Where("user_name = ?", name).First(&user).Error
	return
}

func (ud *UserDao) DeleteUserById(id uint) (err error) {
	err = ud.DB.Delete(&model.User{}, id).Error
	return
}

func (ud *UserDao) CreateUser(user *model.User) (err error) {
	err = ud.DB.Model(&model.User{}).Create(&user).Error
	return
}

func (ud *UserDao) UpdateUser(user *model.User) (err error) {
	err = ud.DB.Where("id = ?", user.Id).Save(user).Error
	return
}
