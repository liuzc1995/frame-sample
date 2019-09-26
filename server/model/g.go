package model

import (
	"github.com/542213314/frame-sample/config"
	"golanger.com/log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//设置数据库
func SetDB(database *gorm.DB) {
	db = database
}

//连接数据库
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connet to db...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}
