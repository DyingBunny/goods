package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goods/pkg/logging/database"
	"goods/pkg/setting"
	"log"
)
var Db *gorm.DB
//数据库连接
func Connect() {
	var (
		err error
		dbType, dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	Db=db
	//defer db.Close()
	if err != nil {
		log.Fatal("Connect error:", err)
		return
	}
}
//初始化
func Init(){
	Connect()
	//显示详细日志
	filePath := database.GetLogFileFullPath()
	F:=database.OpenLogFile(filePath)
	DefaultPrefix := ""
	logger:=log.New(F, DefaultPrefix, log.LstdFlags)
	Db.SetLogger(logger)
	Db.LogMode(true)
	Db.SingularTable(true)
}









