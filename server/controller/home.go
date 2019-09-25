package controller

import (
	"github.com/542213314/frame-sample/server/vm"
	"net/http"

	"github.com/gorilla/mux"
)

type home struct{}

//设置路由
func (h home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	staticHandler()
}

//静态资源路径配置
func staticHandler() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

//tpPath模板路径,相对template目录文件路径
//GetVM() 取模板页面内容
//templates[tpName].Execute 执行渲染模板

//首页
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpPath := "index/index.html"
	vop := vm.IndexViewModelOp{}
	if r.Method == http.MethodGet {
		v := vop.GetVM()
		templates[tpPath].Execute(w, &v)
	}
}
