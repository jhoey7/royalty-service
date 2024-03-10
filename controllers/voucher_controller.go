package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"royalty-service/repositories"
	"royalty-service/services"
	"time"
)

// VoucherController operation
type VoucherController struct {
	beego.Controller
}

// Create controller for create voucher.
// @Title Create controller for create voucher.
// @Description Create controller for create voucher.
// @Success 200 {object} models.Response
// @router [post]
func (v VoucherController) Create() {
	identifier := time.Now().UnixNano()

	vRepo := repositories.NewVoucherRepository(orm.NewOrm())
	vdRepo := repositories.NewVoucherDetailRepository(orm.NewOrm())
	svc := services.NewVoucherService(vRepo, vdRepo, orm.NewOrm(), identifier)
	resp := svc.Create(v.Ctx.Input.RequestBody)

	v.Data["json"] = resp
	v.ServeJSON()
	return
}
