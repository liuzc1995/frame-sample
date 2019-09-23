package main

import (
	"gggo/controller"
	"gggo/model"
	"net/http"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//连接数据库
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	//开启路由
	controller.Startup()
	//监听
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
