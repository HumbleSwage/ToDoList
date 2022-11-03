package model

import (
	"ToDoList/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"` //0-finished，1-unfinished
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}

func (Task *Task) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
}
