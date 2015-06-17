// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBaseController struct {}
var BaseController tBaseController


func (_ tBaseController) Before(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.Before", args).Url
}

func (_ tBaseController) Err(
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("BaseController.Err", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tProject struct {}
var Project tProject


func (_ tProject) ListByOwner(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Project.ListByOwner", args).Url
}

func (_ tProject) ListByCompany(
		companyId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	return revel.MainRouter.Reverse("Project.ListByCompany", args).Url
}

func (_ tProject) Id(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Project.Id", args).Url
}

func (_ tProject) Add(
		companyId int,
		name string,
		info string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	return revel.MainRouter.Reverse("Project.Add", args).Url
}

func (_ tProject) Update(
		id int,
		name string,
		info string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	return revel.MainRouter.Reverse("Project.Update", args).Url
}


type tTaskTransfer struct {}
var TaskTransfer tTaskTransfer


func (_ tTaskTransfer) Add(
		companyId int,
		taskId int,
		assignTo int,
		info string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "taskId", taskId)
	revel.Unbind(args, "assignTo", assignTo)
	revel.Unbind(args, "info", info)
	return revel.MainRouter.Reverse("TaskTransfer.Add", args).Url
}

func (_ tTaskTransfer) UpdateProgress(
		companyId int,
		taskId int,
		progress int,
		info string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "taskId", taskId)
	revel.Unbind(args, "progress", progress)
	revel.Unbind(args, "info", info)
	return revel.MainRouter.Reverse("TaskTransfer.UpdateProgress", args).Url
}

func (_ tTaskTransfer) Read(
		companyId int,
		taskId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "taskId", taskId)
	return revel.MainRouter.Reverse("TaskTransfer.Read", args).Url
}

func (_ tTaskTransfer) ListByTaskId(
		taskId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "taskId", taskId)
	return revel.MainRouter.Reverse("TaskTransfer.ListByTaskId", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) Me(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Me", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Reg(
		username string,
		password string,
		smsCode string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "smsCode", smsCode)
	return revel.MainRouter.Reverse("Auth.Reg", args).Url
}

func (_ tAuth) Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Auth.Login", args).Url
}

func (_ tAuth) ChangePassword(
		oldPassword string,
		newPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "oldPassword", oldPassword)
	revel.Unbind(args, "newPassword", newPassword)
	return revel.MainRouter.Reverse("Auth.ChangePassword", args).Url
}

func (_ tAuth) GetPassword(
		username string,
		newPassword string,
		smsCode string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "newPassword", newPassword)
	revel.Unbind(args, "smsCode", smsCode)
	return revel.MainRouter.Reverse("Auth.GetPassword", args).Url
}

func (_ tAuth) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Logout", args).Url
}


type tComm struct {}
var Comm tComm


func (_ tComm) SendSms(
		username string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	return revel.MainRouter.Reverse("Comm.SendSms", args).Url
}


type tCompany struct {}
var Company tCompany


func (_ tCompany) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Company.List", args).Url
}

func (_ tCompany) Id(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Company.Id", args).Url
}

func (_ tCompany) Add(
		name string,
		info string,
		phone string,
		address string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	revel.Unbind(args, "phone", phone)
	revel.Unbind(args, "address", address)
	return revel.MainRouter.Reverse("Company.Add", args).Url
}

func (_ tCompany) Update(
		id int,
		name string,
		info string,
		phone string,
		address string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	revel.Unbind(args, "phone", phone)
	revel.Unbind(args, "address", address)
	return revel.MainRouter.Reverse("Company.Update", args).Url
}

func (_ tCompany) UpdateLogo(
		id int,
		logo string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "logo", logo)
	return revel.MainRouter.Reverse("Company.UpdateLogo", args).Url
}


type tCompanyUsers struct {}
var CompanyUsers tCompanyUsers


func (_ tCompanyUsers) List(
		companyId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	return revel.MainRouter.Reverse("CompanyUsers.List", args).Url
}

func (_ tCompanyUsers) Add(
		companyId int,
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("CompanyUsers.Add", args).Url
}

func (_ tCompanyUsers) Check(
		companyId int,
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("CompanyUsers.Check", args).Url
}

func (_ tCompanyUsers) Delete(
		companyId int,
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("CompanyUsers.Delete", args).Url
}


type tTask struct {}
var Task tTask


func (_ tTask) ListTodo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Task.ListTodo", args).Url
}

func (_ tTask) Add(
		companyId int,
		projectId int,
		priority int,
		inChargeUserId int,
		name string,
		info string,
		deadline interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "projectId", projectId)
	revel.Unbind(args, "priority", priority)
	revel.Unbind(args, "inChargeUserId", inChargeUserId)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	revel.Unbind(args, "deadline", deadline)
	return revel.MainRouter.Reverse("Task.Add", args).Url
}

func (_ tTask) Done(
		companyId int,
		taskId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "companyId", companyId)
	revel.Unbind(args, "taskId", taskId)
	return revel.MainRouter.Reverse("Task.Done", args).Url
}


