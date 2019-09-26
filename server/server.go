package main

import (
	"flag"
	"net/http"
	"os"
	"runtime"

	"golanger.com/log"

	"github.com/542213314/frame-sample/config"
	"github.com/542213314/frame-sample/server/controller"
	"github.com/542213314/frame-sample/server/model"
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
	logFile, logErr := os.OpenFile(config.GetLog(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		log.Debug("Fail to find", *logFile, "Server start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	if *autoWatch != "false" {
		watchApp()
	} else {
		//连接数据库
		db := model.ConnectToDB()
		defer db.Close()
		model.SetDB(db)
		//开启路由
		controller.Startup()
		//监听
		http.ListenAndServe(":"+config.GetPort(), context.ClearHandler(http.DefaultServeMux))
	}
}
