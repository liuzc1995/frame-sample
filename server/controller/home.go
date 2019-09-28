package controller

import (
	"net/http"

	"github.com/542213314/frame-sample/server/vm"

	"github.com/gorilla/mux"
)

type home struct{}

//设置路由
func (h home) registerRoutes() {
	r := mux.NewRouter()
	iPath := r.PathPrefix("/").Subrouter()
	initUserPath(iPath)

	//首页路由
	r.HandleFunc("/", indexHandler)
	//静态资源配置
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)
}

//配置可对 /user下发送请求的路由
func initUserPath(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/list", UserListHandler)
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

// /user/list
func UserListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
