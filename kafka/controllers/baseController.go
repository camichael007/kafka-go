package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) IsPost() bool {
	return c.Ctx.Request.Method == beego.HTTPMETHOD["POST"]
}

func (c *BaseController) IsGet() bool {
	return c.Ctx.Request.Method == beego.HTTPMETHOD["GET"]
}

func (c *BaseController) SetJson(code int, msg string, data interface{}) {
	c.Data["json"] = map[string]interface{}{"code":code, "msg":msg, "data":data}
	c.ServeJSON()
}

func (c *BaseController) IsOptions() bool {
	return c.Ctx.Request.Method == beego.HTTPMETHOD["OPTIONS"];
}

func (c *BaseController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
	return
}




