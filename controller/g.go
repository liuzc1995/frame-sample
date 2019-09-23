package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	flashName      string
	store          *sessions.CookieStore
)

//初始化模板
func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("something very secret"))
}

//开启注册路由
func Startup() {
	homeController.registerRoutes()
}
