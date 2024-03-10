package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"royalty-service/repositories"
	"royalty-service/services"
	"time"
)

// Transaction operation
type Transaction struct {
	beego.Controller
}

// Create controller for create transaction.
// @Title Create controller for create transaction.
// @Description Create controller for create transaction.
// @Success 200 {object} models.Response
// @router [post]
func (t Transaction) Create() {
	identifier := time.Now().UnixNano()

	tRepo := repositories.NewTransactionRepository(orm.NewOrm())
	uRepo := repositories.NewUserRepository(orm.NewOrm())
	vdRepo := repositories.NewVoucherDetailRepository(orm.NewOrm())
	vRepo := repositories.NewVoucherRepository(orm.NewOrm())
	svc := services.NewTransactionService(tRepo, uRepo, vdRepo, vRepo, orm.NewOrm(), identifier)
	resp := svc.Create(t.Ctx.Input.RequestBody)

	t.Data["json"] = resp
	t.ServeJSON()
	return
}
