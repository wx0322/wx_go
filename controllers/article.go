package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get(){
	username := c.GetSession("username")
	c.Data["username"] = username
	c.TplName="article.html"
}