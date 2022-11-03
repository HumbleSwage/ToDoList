package serializer

import (
	"ToDoList/model"
)

type Task struct {
	ID        uint   `json:"id" form:"id" example:"1"`  //任务ID
	Title     string `json:"title" example:"吃饭"`        // 题目
	Content   string `json:"content" example:"红烧肉，炒白菜"` //内容
	View      uint64 `json:"view" example:"32"`         //浏览量
	Status    int    `json:"status" example:"0"`        //状态（0未完成，1已完成）
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_at"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func BuildTasks(tasks []model.Task) []Task {
	var res []Task
	for _, v := range tasks {
		res = append(res, Task{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Status:    v.Status,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return res
}
