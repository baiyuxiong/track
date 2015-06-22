package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/lib"
	"time"
	"fmt"
)

type Task struct {
	BaseController
}

type TodoList  struct {
	TasksInCharge  []models.Task `json:"tasksInCharge"`
	TasksOwnedByMe []models.Task `json:"tasksOwnedByMe"`
	Users          map[string]models.UserProfiles `json:"users"`
	Projects       map[string]models.Project `json:"projects"`
	Transfers      map[string]models.TaskTransfer `json:"transfers"`
}

//我负责的、未完成的
func (c Task) ListTodo() revel.Result {
	tasksInChange := make([]models.Task, 0)
	//我负责的、未完成的
	err := app.Engine.Where("status=?", utils.TASK_STATUS_DOING).Where("in_charge_user_id = ?", c.User.Id).OrderBy("priority,updated_at desc").Find(&tasksInChange)
	if err != nil {
		return c.Err(err.Error())
	}

	tasksOwnedByMe := make([]models.Task, 0)
	err = app.Engine.Where("status=?", utils.TASK_STATUS_DOING).Where("owner_id <> in_charge_user_id").Where("owner_id = ?", c.User.Id).OrderBy("priority,updated_at desc").Find(&tasksOwnedByMe)
	if err != nil {
		return c.Err(err.Error())
	}

	todoList := TodoList{
		TasksInCharge:tasksInChange,
		TasksOwnedByMe:tasksOwnedByMe,
		Users:make(map[string]models.UserProfiles, 0),
		Projects:make(map[string]models.Project, 0),
		Transfers:make(map[string]models.TaskTransfer, 0),
	}
	for _, tic := range tasksInChange {
		transfer := new(models.TaskTransfer)
		has, err := app.Engine.Id(tic.LatestTransferId).Get(transfer)
		if !has {
			return c.Err("数据有误，请联系管理员"+err.Error())
		}
		todoList.Transfers[fmt.Sprint(tic.LatestTransferId)] =*transfer

		_, exists := todoList.Users[fmt.Sprint(tic.OwnerId)]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(tic.OwnerId).Get(userProfile)
			if has {
				todoList.Users[fmt.Sprint(tic.OwnerId)] =*userProfile
			}
		}

		_, exists = todoList.Projects[fmt.Sprint(tic.ProjectId)]
		if !exists {
			project := new(models.Project)
			has, _ := app.Engine.Id(tic.ProjectId).Get(project)
			if has {
				todoList.Projects[fmt.Sprint(tic.ProjectId)] =*project
			}
		}
	}

	for _, tobm := range tasksOwnedByMe {
		transfer := new(models.TaskTransfer)
		has, err := app.Engine.Id(tobm.LatestTransferId).Get(transfer)
		if !has {
			return c.Err("数据有误，请联系管理员"+err.Error())
		}
		todoList.Transfers[fmt.Sprint(tobm.LatestTransferId)] =*transfer


		_, exists := todoList.Users[fmt.Sprint(transfer.AssignTo)]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(transfer.AssignTo).Get(userProfile)
			if has {
				todoList.Users[fmt.Sprint(transfer.AssignTo)] = *userProfile
			}
		}

		_, exists = todoList.Projects[fmt.Sprint(tobm.ProjectId)]
		if !exists {
			project := new(models.Project)
			has, _ := app.Engine.Id(tobm.ProjectId).Get(project)
			if has {
				todoList.Projects[fmt.Sprint(tobm.ProjectId)] = *project
			}
		}
	}
	return c.OK(todoList)
}


func (c Task) Add(companyId, projectId, priority, inChargeUserId int, name, info string, deadline time.Time) revel.Result {
	if !lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		return c.Err("没有权限")
	}
	if !lib.IsCompanyCheckedUser(companyId, inChargeUserId) {
		return c.Err("被指派用户不是团队成员，不能指派")
	}
	if res, _ := lib.IsProjectBelongToCompany(nil, projectId, companyId); !res {
		return c.Err("项目不属于本团队")
	}

	task := &models.Task{
		CompanyId:companyId,
		ProjectId:projectId,
		OwnerId:c.User.Id,
		LatestTransferId:0,
		InChargeUserId:inChargeUserId,
		Priority:priority,
		Status:utils.TASK_STATUS_DOING,
		Name:name,
		Info:info,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}

	_, err := app.Engine.Insert(task)
	if err != nil {
		return c.Err("添加失败，请联系管理员")
	}

	taskTransfer := &models.TaskTransfer{
		TaskId:task.Id,
		AssignTo:inChargeUserId,
		AssignFr:c.User.Id,
		IsRead:utils.TASK_UN_READ,
		Progress:0,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}
	_, err = app.Engine.Insert(taskTransfer)
	if err != nil {
		app.Engine.Id(task.Id).Delete(task)
		return c.Err("添加失败，请联系管理员")
	}

	task.LatestTransferId = taskTransfer.Id
	app.Engine.Id(task.Id).Cols("latest_transfer_id").Update(task)

	return c.OK(nil)
}

type TaskDetail struct {
	Task models.Task `json:"task"`
	TaskTransfers []models.TaskTransfer `json:"taskTransfers"`
	TaskTransferUsers map[string]models.UserProfiles `json:"taskTransferUsers"`
	Owner models.UserProfiles `json:"owner"`
	InChargeUser models.UserProfiles`json:"inChargeUser"`
	Company models.Company `json:"company"`
	Project models.Project `json:"project"`
}
func (c Task) Detail(companyId, taskId int) revel.Result{
	if !lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		return c.Err("没有权限")
	}
	taskTransferUsers := make(map[string]models.UserProfiles,0)

	taskDetail := TaskDetail{}

	task := &models.Task{}
	has, _ := app.Engine.Id(taskId).Get(task)
	if !has {
		return c.Err("任务记录不存在")
	}
	taskDetail.Task = *task

	user := &models.UserProfiles{}
	has, _ = app.Engine.Id(task.OwnerId).Get(user)
	if !has {
		return c.Err("用户记录不存在")
	}
	taskDetail.Owner = *user
	taskTransferUsers[fmt.Sprint(task.OwnerId)] = *user

	project := &models.Project{}
	has, _ = app.Engine.Id(task.ProjectId).Get(project)
	if !has {
		return c.Err("项目记录不存在")
	}
	taskDetail.Project = *project

	company := &models.Company{}
	has, _ = app.Engine.Id(task.CompanyId).Get(company)
	if !has {
		return c.Err("组织记录不存在")
	}
	taskDetail.Company = *company

	if task.OwnerId == task.InChargeUserId{
		taskDetail.InChargeUser = *user
	}else{
		userInCharge := &models.UserProfiles{}
		has, _ = app.Engine.Id(task.InChargeUserId).Get(userInCharge)
		if !has {
			return c.Err("用户记录不存在")
		}
		taskDetail.InChargeUser = *userInCharge
		taskTransferUsers[fmt.Sprint(task.InChargeUserId)] = *userInCharge
	}

	//指派过程
	transfers := make([]models.TaskTransfer, 0)
	err := app.Engine.Where("task_id = ?", task.Id).OrderBy("updated_at desc").Find(&transfers)
	if err != nil {
		return c.Err(err.Error())
	}
	taskDetail.TaskTransfers = transfers

	//指派涉及用户信息
	for _,transfer := range transfers{
		_, exists := taskTransferUsers[fmt.Sprint(transfer.AssignTo)]
		if !exists {
			transferUser := &models.UserProfiles{}
			has, _ = app.Engine.Id(transfer.AssignTo).Get(transferUser)
			if has {
				taskTransferUsers[fmt.Sprint(transfer.AssignTo)] = *transferUser
			}
		}

	}
	taskDetail.TaskTransferUsers = taskTransferUsers

	return c.OK(taskDetail)
}

func (c Task) Done(companyId, taskId int) revel.Result {
	if !lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		return c.Err("没有权限")
	}

	task := &models.Task{}
	has, _ := app.Engine.Id(taskId).Get(task)
	if !has {
		return c.Err("任务记录不存在")
	}

	if task.OwnerId != c.User.Id {
		return c.Err("不是任务发起人，没有权限")
	}

	task.Status = utils.TASK_STATUS_DONE
	app.Engine.Id(taskId).Cols("status").Update(task)
	return c.OK(nil)
}


