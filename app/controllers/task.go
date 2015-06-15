package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/lib"
	"time"
)

type Task struct {
	BaseController
}

type TodoList  struct {
	TasksInChange  []models.Task
	TasksOwnedByMe []models.Task
	Users          map[int]models.UserProfiles
	Projects       map[int]models.Project
}

//我负责的、未完成的
func (c Task) ListTodo() revel.Result {
	tasksInChange := make([]models.Task, 0)
	//我负责的、未完成的
	err := app.Engine.Where("status=?", utils.TASK_STATUS_DOING).Where("in_charge_id = ?", c.User.Id).OrderBy("priority,updated_at desc").Find(&tasksInChange)
	if err != nil {
		return c.Err(err.Error())
	}

	tasksOwnedByMe := make([]models.Task, 0)
	err = app.Engine.Where("status=?", utils.TASK_STATUS_DOING).Where("owner_id <> in_charge_id").Where("owner_id = ?", c.User.Id).OrderBy("priority,updated_at desc").Find(&tasksOwnedByMe)
	if err != nil {
		return c.Err(err.Error())
	}

	todoList := TodoList{
		TasksInChange:make([]models.Task),
		TasksOwnedByMe:make([]models.Task),
		Users:make(map[int]models.UserProfiles),
		Projects:make(map[int]models.Project),
	}
	for _, tic := range tasksInChange {
		//TODO 判断存在，避免重复取库
		_, exists := todoList.Users[tic.OwnerId]
		if !exists{
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(tic.OwnerId).Get(UserProfiles)
			if has{
				todoList.Users[tic.OwnerId] =userProfile
			}
		}

		_, exists = todoList.Projects[tic.ProjectId]
		if !exists {
			project := new(models.Project)
			has, _ := app.Engine.Id(tic.ProjectId).Get(project)
			if has {
				todoList.Projects[tic.ProjectId] =project
			}
		}
	}

	for _, tobm := range tasksOwnedByMe {
		_, exists := todoList.Users[tobm.OwnerId]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(tobm.InChargeId).Get(UserProfiles)
			if has {
				todoList.Users[tobm.InChargeId] =userProfile
			}
		}

		_, exists = todoList.Projects[tobm.ProjectId]
		if !exists {
			project := new(models.Project)
			has, _ := app.Engine.Id(tobm.ProjectId).Get(project)
			if has {
				todoList.Projects[tobm.ProjectId] =project
			}
		}
	}
	return c.OK(todoList)
}


func (c Task) Add(companyId,projectId,priority,InChargeId int,name,info string,deadline time.Time) revel.Result {
	if !lib.IsCompanyCheckedUser(companyId,c.User.Id){
		return c.Err("没有权限")
	}
	if !lib.IsCompanyCheckedUser(companyId,InChargeId){
		return c.Err("不能指派给不同单位的人员")
	}
	if !lib.IsProjectBelongToCompany(nil,projectId,companyId){
		return c.Err("没有权限")
	}

	task := &models.Task{
		CompanyId:companyId,
		ProjectId:projectId,
		OwnerId:c.User.Id,
		InChargeId:InChargeId,
		InChargeProgress:0,
		Priority:priority,
		Status:utils.TASK_STATUS_DOING,
		Name:name,
		Info:info,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}

	_, err := app.Engine.Insert(task)
	if err != nil{
		return c.Err("添加失败，请联系管理员")
	}

	taskTransfer := &models.TaskTransfer{
		TaskId:task.Id,
		AssignTo:InChargeId,
		AssignFrom:c.User.Id,
		IsRead:utils.TASK_UN_READ,
		Progress:0,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}
	_, err = app.Engine.Insert(taskTransfer)
	if err != nil{
		app.Engine.Id(task.Id).Delete(task)
		return c.Err("添加失败，请联系管理员")
	}
	return c.OK(nil)
}

