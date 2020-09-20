package controllers

import (
	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get(){
	c.TplName="product.html"
}