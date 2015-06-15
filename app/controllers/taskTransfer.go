package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app"
	"time"
)

type TaskTransfer struct {
	BaseController
}

func (c TaskTransfer) Add(companyId,taskId ,assignTo int) revel.Result {
	if !c.hasTaskAuth(companyId,taskId){
		return c.Err("没有权限")
	}

	//获取最新的assign
	transfer := new(models.TaskTransfer)
	has, _ := app.Engine.Where("task_id=?", taskId).OrderBy("created_at desc").Get(transfer)

	if !has{
		return c.Err("记录不存在")
	}

	if transfer.AssignTo != c.User.Id{
		return c.Err("您当前不在负责此任务，不可指派他人")
	}

	taskTransfer := &models.TaskTransfer{
		TaskId:taskId,
		AssignTo:assignTo,
		AssignFrom:c.User.Id,
		IsRead:utils.TASK_UN_READ,
		Progress:0,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}
	_, err := app.Engine.Insert(taskTransfer)
	if err != nil{
		return c.Err("添加失败，请联系管理员")
	}

	//task表
	//TODO
	return c.OK("")
}

func (c TaskTransfer) UpdateProgress(companyId,taskId,progress int,info string) revel.Result {
	if !c.hasTaskAuth(companyId,taskId){
		return c.Err("没有权限")
	}

	if progress < 0{
		progress = 0;
	}
	if progress > 100{
		progress = 100;
	}

	//获取最新的assign
	transfer := new(models.TaskTransfer)
	has, _ := app.Engine.Where("task_id=?", taskId).OrderBy("created_at desc").Get(transfer)

	if !has{
		return c.Err("记录不存在")
	}

	if transfer.AssignTo != c.User.Id{
		return c.Err("您当前不在负责此任务，不可更新进度")
	}

	transfer.Progress = progress
	transfer.Info = info

	app.Engine.Id(transfer.Id).Cols("progress").Cols("info").Update(transfer)

	//task表
	//TODO

	return c.OK("")
}

//修改为已读
func (c TaskTransfer) Read(companyId, taskId ,id int) revel.Result {
	if !c.hasTaskAuth(companyId,taskId){
		return c.Err("没有权限")
	}
	//获取最新的assign
	transfer := new(models.TaskTransfer)
	has, _ := app.Engine.Where("task_id=?", taskId).OrderBy("created_at desc").Get(transfer)

	if !has{
		return c.Err("记录不存在")
	}

	if transfer.AssignTo != c.User.Id{
		return c.Err("您当前不在负责此任务，不可修改阅读状态")
	}

	if transfer.Id != id {
		return c.Err("不可修改阅读状态")
	}

	transfer.IsRead = utils.TASK_IS_READ

	app.Engine.Id(transfer.Id).Cols("is_read").Update(transfer)
	return c.OK("")

}

func (c TaskTransfer) ListByTaskId(taskId int) revel.Result {
	//TODO 权限
	transfer := make([]models.TaskTransfer, 0)
	err := app.Engine.Where("task_id = ?", taskId).OrderBy("created_at desc").Find(&transfer)
	if err != nil {
		return c.OK(transfer)
	}
	return c.Err(err.Error())
}

func (c TaskTransfer) hasTaskAuth(companyId, taskId int)  (bool) {
	if !lib.IsCompanyCheckedUser(companyId,c.User.Id){
		return false
	}

	if !lib.IsTaskBelongToCompany(nil,taskId,companyId){
		return false
	}
	return true;
}