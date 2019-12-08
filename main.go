package main

import (
	"fmt"
	"goods/pkg/setting"
	"goods/routers"
	"net/http"
)

func main() {
	//注册一个默认的路由
	router:=routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		//http请求最大长度
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}