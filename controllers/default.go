package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
<<<<<<< HEAD
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "beego.geamil.com"
=======
	c.Data["Website"] = "2438121546@qq.com"/*修改后的自己的邮箱*/
	c.Data["Email"] = "www.baidu.com"/*修改后的网址为百度*/
>>>>>>> 289c93d... revied
	c.TplName = "index.tpl"
}
