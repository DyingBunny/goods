package database

import (
	"fmt"
	"log"
	"os"
	"time"
)
var (
	LogSavePath = "runtime/mysqlLogs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)
//获取日志路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}
//全路径
func GetLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}
//打开日志文件
func OpenLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}
//建立目录
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}


