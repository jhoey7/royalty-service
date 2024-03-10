package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"royalty-service/repositories"
	"royalty-service/services"
	"time"
)

// UserController operation
type UserController struct {
	beego.Controller
}

// Register controller for register new user.
// @Title Register controller for register new user.
// @Description Register controller for register new user.
// @Success 200 {object} models.Response
// @router /register [put]
func (c UserController) Register() {
	identifier := time.Now().UnixNano()

	userRepo := repositories.NewUserRepository(orm.NewOrm())
	svc := services.NewUserService(userRepo, identifier)
	resp := svc.Register(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}
