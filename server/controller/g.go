package controller

import (
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
)

//初始化模板
func init() {
	templates = PopulateTemplates()
}

//开启注册路由
func Startup() {
	homeController.registerRoutes()
}
