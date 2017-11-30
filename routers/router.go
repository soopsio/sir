package routers

import (
	"github.com/soopsio/sir/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.TaskController{})
	beego.SetStaticPath("/public", "public")
}
