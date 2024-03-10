package routers

import (
	"github.com/astaxie/beego"
	"royalty-service/controllers"
)

func init() {
	ns :=
		beego.NewNamespace("/1.0",
			beego.NSNamespace("/users",
				beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			),

			beego.NSNamespace("/vouchers",
				beego.NSRouter("", &controllers.VoucherController{}, "post:Create"),
			),

			beego.NSNamespace("/transactions",
				beego.NSRouter("", &controllers.Transaction{}, "post:Create"),
			),
		)

	beego.AddNamespace(ns)
}
