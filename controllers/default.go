package controllers

import (
	"beeprojects2/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "beego.geamil.com"
    Username:=c.Ctx.Input .Query("user")
	Password:=c.Ctx.Input .Query("psd")

    if Username !="admin"||Password !="123456"{
    	c.Ctx.ResponseWriter.Write([]byte("对不起数据校验错误"))//向前端请求数据
		return
	}else {
		c.Ctx.ResponseWriter.Write([]byte("数据校验成功！！"))
	};
	c.Data["Website"] = "2438121546@qq.com"/*修改后的自己的邮箱*/
	c.Data["Email"] = "www.baidu.com"/*修改后的网址为百度*/
	c.TplName = "index.tpl"
}





//func(c *MainController)Post(){
//	name :=c.Ctx.Request.FormValue("name")
//	age :=c.Ctx.Request.FormValue("age")
//	sex :=c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	if name != "admin"&& age !="18"&& sex !="male"{
//		c.Ctx.WriteString("数据校验失败!")
//		return
//	}else {
//		c.Ctx.WriteString("数据校验成功！")
//	}
//}

//func (c *MainController)Post(){
//	//解析前端json的 格式
//	var Person models.Person
//	dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
//	if err !=nil{
//		c.Ctx.WriteString("数据解析1失败")
//		return
//	}
//	err = json.Unmarshal(dataBytes,&Person)
//	if err !=nil{
//		fmt.Println(err.Error())
//		c.Ctx.WriteString("数据解2析失败")
//		return
//	}
//	fmt.Println("姓名",Person.Name)
//	fmt.Println("姓名",Person.Age)
//	fmt.Println("姓名",Person.Sex)
//	c.Ctx.WriteString("数据解析成功")
//}

func (c*MainController) Post() {
   //解析前端json格式
	var information models.Information//面向对象编程
	dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("解析数据1失败")//判断解析是否正确
		return
	}
	err =json.Unmarshal(dataBytes,&information)//
	if err !=nil {
		fmt.Println(err.Error())
		c.Ctx.WriteString("解析数据2失败！")
		return
	}
	fmt.Println("姓名：",information.Name)
	fmt.Println("生日：",information.Birthday)
	fmt.Println("地址：",information.Address)
	fmt.Println("昵称：",information.Nick)
	c.Ctx.WriteString("数据解析成功")
}