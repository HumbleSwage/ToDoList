package model

import (
	"ToDoList/repository/cache"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model
	Id        uint   `gorm:"column:id;primary_key"`
	UserName  string `gorm:"column:user_name"`
	Uid       uint   `gorm:"column:uid"`
	Title     string `gorm:"title"`
	Status    int    `gorm:"column:status"`
	Content   string `gorm:"column:content"`
	View      int    `gorm:"column:view"`
	StartTime int64  `gorm:"start_time"`
	EndTime   int64  `gorm:"end_time"`
}

func (t *Task) GetView() int {
	count, _ := cache.RedisClient.Get(cache.TaskViewKey(t.Id)).Result()
	return cast.ToInt(count)
}

func (t *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(t.Id))
}
