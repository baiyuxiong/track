package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
)

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*BaseController).Before, revel.BEFORE)
}

type BaseController struct {
	*revel.Controller
	User *models.Users
}

// Before is called prior to the controller method
func (c *BaseController) Before() revel.Result {

	invalid := true
	if c.Params.Get(utils.URL_CLIENT_ID_KEY) == utils.URL_CLIENT_ID{
		noTokenPath := []string{"/auth/reg", "/auth/login", "/auth/get_password", "/comm/send_sms"}
		if !utils.StringInSlice(c.Request.URL.Path,noTokenPath){
			token := c.Params.Get(utils.URL_TOKEN_KEY)
			if token!= "" {
				var u = &models.Users{Token: token}
				has, _ := app.Engine.Get(u)
			
				if has {
					c.User = u
					invalid = false
				}
			}
		}else{
			invalid = false
		}
	}
	
	if invalid{
		return c.RenderJson(utils.WrapFailJsonResult("INVALID_REQUEST"))
	}else{
		return nil
	}
}

func (c *BaseController) Err(message string) revel.Result{
	return c.RenderJson(utils.WrapFailJsonResult(message))
}

func (c *BaseController) OK(o interface{}) revel.Result{
	return c.RenderJson(utils.WrapOKJsonResult(o))
}