package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//初始config
func init() {
	getConfig()
}

//设置config
func getConfig() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("../..")

	viper.AddConfigPath("../") // optionally look for config in the working directory
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

//取连接数据库所需字符串
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}

//端口
func GetPort() string {
	return viper.GetString("server.port")
}

//日志
func GetLog() string {
	return viper.GetString("server.log")
}

//进程
func GetPid() string {
	return viper.GetString("server.pid")
}

//监听文件路径
func GetWatchPath() []string {
	return viper.GetStringSlice("server.watch_path")
}
