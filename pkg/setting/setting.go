//配置文件读取
package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret string
)
//初始化
func init() {
	var err error
	Cfg, err = ini.Load("conf/goods.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadGoods()
}
//基本
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
//服务器信息
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
//日志
func LoadGoods() {
	sec, err := Cfg.GetSection("goods")
	if err != nil {
		log.Fatalf("Fail to get section 'goods': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
