package service

import (
	"ToDoList/consts"
	"ToDoList/pkg/ctl"
	"ToDoList/pkg/utils"
	"ToDoList/repository/db/dao"
	"ToDoList/repository/db/model"
	"ToDoList/types"
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (ts *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	taskDao := dao.NewTaskDaoByDB(userDao.DB)
	task := &model.Task{
		UserName:  user.UserName,
		Uid:       user.Id,
		Title:     req.Title,
		Status:    req.Status,
		Content:   req.Content,
		View:      0,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	if err := taskDao.CreateTask(task); err != nil {
		return nil, err
	}
	return ctl.RespSuccess(), nil
}

func (ts *TaskSrv) ListTask(ctx context.Context, req *types.ListTaskReq) (resp interface{}, err error) {
	// 解析用户信息
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	taskDao := dao.NewTaskDao(ctx)
	tasks, total, err := taskDao.GetTaskByUserId(req.PageNum, req.PageSize, userInfo.Id)
	if err != nil {
		return nil, err
	}
	data := make([]*types.ListTaskResp, 0)
	for _, task := range tasks {
		data = append(data, &types.ListTaskResp{
			Id:        task.Id,
			Uid:       task.Uid,
			UserName:  task.UserName,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			View:      task.GetView(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		})
	}

	return ctl.RespSuccessWithData(ctl.DataList{Item: data, Total: int(total)}), nil
}

func (ts *TaskSrv) GetTask(ctx context.Context, req *types.GetTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	taskDao := dao.NewTaskDao(ctx)
	task, err := taskDao.GetTask(userInfo.Id, req.Id)
	switch err {
	case gorm.ErrRecordNotFound:
		return ctl.RespSuccessWithData("该用户没有该条任务记录"), nil
	case nil:
		task.AddView()
		data := types.GetTaskResp{
			Id:        task.Id,
			UserName:  task.UserName,
			Title:     task.Title,
			Status:    task.Status,
			Content:   task.Content,
			View:      task.GetView(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		}
		return ctl.RespSuccessWithData(data), nil
	default:
		utils.LogrusObj.Info(err)
		return nil, errors.New("获取任务发生未知错误")
	}
}

func (ts *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	taskDao := dao.NewTaskDao(ctx)
	task, err := taskDao.GetTask(userInfo.Id, req.Id)
	switch err {
	case gorm.ErrRecordNotFound:
		return ctl.RespSuccessWithData("该用户没有该任务记录"), nil
	case nil:
		if req.Title != "" && task.Title != req.Title {
			task.Title = req.Title
		} else if req.Content != "" && task.Content != req.Content {
			task.Content = req.Content
		} else if req.Status != 0 && req.Status != task.Status {
			task.Status = req.Status
		} else if req.StartTime != 0 && req.StartTime != task.StartTime {
			task.StartTime = req.StartTime
		} else if req.EndTime != 0 && req.StartTime != task.EndTime {
			task.EndTime = req.EndTime
		}
		err := taskDao.UpdateTaskById(task)
		if err != nil {
			return nil, err
		}
		return ctl.RespSuccess(), nil
	default:
		return nil, errors.New("更新任务时发生错误")
	}
}

func (ts *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return nil, err
	}
	taskDao := dao.NewTaskDao(ctx)
	err = taskDao.UserIsHadTask(userInfo.Id, req.Id)
	switch err {
	case gorm.ErrRecordNotFound:
		return ctl.RespSuccessWithData("该用户没有该任务记录"), nil
	case nil:
		err = taskDao.DeleteTaskById(req.Id)
		if err != nil {
			utils.LogrusObj.Info(err)
			return nil, errors.New("删除用户时发生未知错误")
		}
		return ctl.RespSuccess(), nil
	default:
		utils.LogrusObj.Info(err)
		return nil, errors.New("检查用户与任务关系时发生未知错误")
	}
}

func (ts *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	taskDao := dao.NewTaskDao(ctx)
	tasks, total, err := taskDao.SearchTask(userInfo.Id, req.Key)
	switch err {
	case nil:
		data := make([]*types.SearchTaskResp, 0)
		if total > 0 {
			for _, task := range tasks {
				data = append(data, &types.SearchTaskResp{
					Id:        task.Id,
					Uid:       task.Uid,
					UserName:  task.UserName,
					Title:     task.Title,
					Content:   task.Content,
					Status:    task.Status,
					View:      task.GetView(),
					StartTime: task.StartTime,
					EndTime:   task.EndTime,
				})
			}
		}
		return ctl.RespSuccessWithData(data), nil
	default:
		return nil, errors.New("用户搜索任务时发生未知错误")
	}
}
