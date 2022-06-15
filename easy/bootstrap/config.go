package bootstrap

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	MysqlHost     string
	MysqlPort     string
	MysqlDb       string
	MysqlUser     string
	MysqlPassword string
	MysqlCharset  string
	MysqlPrefix   string
)

func init() {
	file, err := ini.Load("../config/config.ini")
	if err != nil {
		log.Println("缺少config.ini配置文件", err)
	}

	MysqlHost = file.Section("database").Key("host").MustString("")
	MysqlPort = file.Section("database").Key("port").MustString("3306")
	MysqlDb = file.Section("database").Key("db").MustString("")
	MysqlUser = file.Section("database").Key("username").MustString("")
	MysqlPassword = file.Section("database").Key("password").MustString("")
	MysqlCharset = file.Section("database").Key("charset").MustString("utf8mb4")
	MysqlPrefix = file.Section("database").Key("prefix").MustString("")

}
