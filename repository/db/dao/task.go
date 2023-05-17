package dao

import (
	"ToDoList/repository/db/model"
	"context"
	"gorm.io/gorm"
	"strings"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewMySQLClient(ctx)}
}

func NewTaskDaoByDB(db *gorm.DB) *TaskDao {
	return &TaskDao{db}
}

func (td *TaskDao) CreateTask(task *model.Task) (err error) {
	err = td.DB.Model(&model.Task{}).Create(task).Error
	return
}

func (td *TaskDao) DeleteTaskById(id uint) (err error) {
	err = td.DB.Delete(&model.Task{}, id).Error
	return err
}

func (td *TaskDao) UpdateTaskById(task *model.Task) (err error) {
	err = td.DB.Where("id = ?", task.Id).Save(task).Error
	return
}

func (td *TaskDao) GetTaskById(id uint) (task *model.Task, err error) {
	err = td.DB.Model(&model.Task{}).Where("id = ?", id).First(id).Error
	return
}

func (td *TaskDao) GetTask(uid, id uint) (task *model.Task, err error) {
	err = td.DB.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uid).First(&task).Error
	return
}

func (td *TaskDao) GetTaskByUserId(start, size int, id uint) (task []*model.Task, total int64, err error) {
	err = td.DB.Model(&model.Task{}).Where("uid = ?", id).Count(&total).Error
	if err != nil {
		return
	}
	err = td.DB.Model(&model.Task{}).Where("uid = ?", id).Limit(size).Offset((start - 1) * size).Find(&task).Error
	return
}

func (td *TaskDao) UserIsHadTask(uid, id uint) (err error) {
	err = td.DB.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uid).Error
	return
}

func (td *TaskDao) SearchTask(uid uint, key string) (task []*model.Task, total int64, err error) {
	key = strings.Join([]string{"%", key, "%"}, "")
	err = td.DB.Model(&model.Task{}).Where("uid = ?", uid).Where("title Like ? OR content Like ?", key, key).Count(&total).Error
	if err != nil {
		return
	}
	// TODO:这种搜索方式不太好
	err = td.DB.Model(&model.Task{}).Where("uid = ?", uid).Where("title Like ? OR content Like ?", key, key).Find(&task).Error
	return
}
