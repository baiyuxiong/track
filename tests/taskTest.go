package tests

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"net/url"
	"strconv"
	"time"
)

func (t *AppTest) StartTestTask() {
//	token = ""
//	t.GetToken()
//	t.ClearCompanyUsersTable()
	t.ClearTaskTable()
	t.ClearTaskTransferTable()
	t.AddTaskTest()
	t.AddTaskTransferTest()
	t.UpdateProgressTest()
	t.ReadTaskTest()
	t.DoneTaskTest()
	t.InitTaskData()
}


func (t *AppTest) AddTaskTest(){
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["priority"] = []string{strconv.Itoa(utils.TASK_PRIORITY_LOW)}
	data["inChargeUserId"] = []string{strconv.Itoa(userId1)}
	data["name"] = []string{taskName}
	data["info"] = []string{taskInfo}
	data["deadline"] = []string{time.Now().Add(time.Hour*24).String()}

	t.PostForm(t.GenUrl("/task/add",token1), data)
	t.AssertContains("没有权限")

	t.PostForm(t.GenUrl("/task/add",token), data)
	t.AssertContains("被指派用户不是团队成员，不能指派")

	//让用户1通过审核，变成团队成员
	data["userId"] = []string{strconv.Itoa(userId1)}
	t.PostForm(t.GenUrl("/companyUsers/check",token), data)
	t.AssertContains("200")

	data["projectId"] = []string{strconv.Itoa(projectId+1)}
	t.PostForm(t.GenUrl("/task/add",token), data)
	t.AssertContains("项目不属于本团队")

	data["projectId"] = []string{strconv.Itoa(projectId)}
	t.PostForm(t.GenUrl("/task/add",token), data)
	t.AssertContains("200")

	var task = &models.Task{Name: taskName}
	has, _ := app.Engine.Get(task)
	t.AssertEqual(has,true)

	taskId = task.Id

	taskTransfer := new(models.TaskTransfer)
	hasT, err := app.Engine.Where("task_id = ?",taskId).Get(taskTransfer)
	t.AssertEqual(nil,err)
	t.AssertEqual(hasT,true)

	t.Get(t.GenUrl("/task/listTodo",token))
	t.AssertContains(taskName)
	t.AssertContains(username1)

	t.Get(t.GenUrl("/task/listTodo",token1))
	t.AssertContains(taskName)
	t.AssertContains(username)
}

func (t *AppTest)AddTaskTransferTest() {
	//前面已指派给user1，
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["taskId"] = []string{strconv.Itoa(taskId)}
	data["info"] = []string{taskTransferInfo}
	data["assignTo"] = []string{strconv.Itoa(userId1)}

	t.PostForm(t.GenUrl("/taskTransfer/add",token), data)
	t.AssertContains("您当前不在负责此任务，不可指派他人")

	data["assignTo"] = []string{strconv.Itoa(userId+userId1)}
	t.PostForm(t.GenUrl("/taskTransfer/add",token1), data)
	t.AssertContains("被指派用户不是团队成员")

	//指派给user
	data["assignTo"] = []string{strconv.Itoa(userId)}
	t.PostForm(t.GenUrl("/taskTransfer/add",token1), data)
	t.AssertContains("200")

	t.Get(t.GenUrl("/task/listTodo",token))
	t.AssertContains(taskName)
	t.AssertContains(username)

	t.Get(t.GenUrl("/task/listTodo",token1))
	t.AssertContains(taskName)
	t.AssertContains(username)
}

func (t *AppTest)UpdateProgressTest() {
	//指派给user了
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["taskId"] = []string{strconv.Itoa(taskId)}
	data["info"] = []string{taskTransferInfo}
	data["progress"] = []string{strconv.Itoa(50)}

	t.PostForm(t.GenUrl("/taskTransfer/updateProgress",token1), data)
	t.AssertContains("您当前不在负责此任务，不可修改进度")

	t.PostForm(t.GenUrl("/taskTransfer/updateProgress",token), data)
	t.AssertContains("200")

	taskTransfer := &models.TaskTransfer{}
	has, _ := app.Engine.Where("progress=?",50).Get(taskTransfer)

	t.AssertEqual(has,true)
}

func (t *AppTest) ReadTaskTest() {
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["taskId"] = []string{strconv.Itoa(taskId)}

	t.PostForm(t.GenUrl("/taskTransfer/read",token1), data)
	t.AssertContains("您当前不在负责此任务")

	t.PostForm(t.GenUrl("/taskTransfer/read",token), data)
	t.AssertContains("200")

	taskTransfer := &models.TaskTransfer{}
	has, _ := app.Engine.Where("is_read=?",1).Get(taskTransfer)

	t.AssertEqual(has,true)
}

func (t *AppTest) DoneTaskTest(){
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["taskId"] = []string{strconv.Itoa(taskId)}

	t.PostForm(t.GenUrl("/task/done",token1), data)
	t.AssertContains("没有权限")

	t.PostForm(t.GenUrl("/task/done",token), data)
	t.AssertContains("200")

	task := &models.Task{}
	has, _ := app.Engine.Where("status=?",utils.TASK_STATUS_DONE).Get(task)

	t.AssertEqual(has,true)
}

func (t *AppTest) InitTaskData() {
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["projectId"] = []string{strconv.Itoa(projectId)}
	data["priority"] = []string{strconv.Itoa(utils.TASK_PRIORITY_HIGH)}
	data["inChargeUserId"] = []string{strconv.Itoa(userId1)}
	data["name"] = []string{"user1帮我写代码吧"}
	data["info"] = []string{"快点写..."}
	data["deadline"] = []string{time.Now().Add(time.Hour*24).String()}

	t.PostForm(t.GenUrl("/task/add",token), data)
	t.AssertContains("200")

	data["priority"] = []string{strconv.Itoa(utils.TASK_PRIORITY_NORMAL)}
	data["inChargeUserId"] = []string{strconv.Itoa(userId)}
	data["name"] = []string{"user帮我做设计吧"}
	data["info"] = []string{"弄漂亮点"}
	data["deadline"] = []string{time.Now().Add(time.Hour*24).String()}

	t.PostForm(t.GenUrl("/task/add",token1), data)
	t.AssertContains("200")
}