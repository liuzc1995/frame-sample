package main

import (
	"flag"
	"fmt"
	"gggo/config"
	"gggo/controller"
	"gggo/model"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	logFileName = flag.String("log", "run/server.log", "Log file name") //日志
)

func init() {
	//生成,更新日志
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	//连接数据库
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	//开启路由
	controller.Startup()
	//监听
	http.ListenAndServe(":"+config.GetPort(), context.ClearHandler(http.DefaultServeMux))

}
