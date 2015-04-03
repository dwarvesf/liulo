package routers

import (
	"github.com/dwarvesf/liulo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
