package routers

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "goods/docs"
	"goods/middleware/core"
	"goods/models"
	"goods/pkg/logging/serve"
	_ "goods/pkg/sdk"
	"goods/pkg/setting"
)
//初始化路由
//func InitRouter() *gin.Engine{
//	gin.SetMode(setting.RunMode)
//	r:=gin.New()
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//	r.Use(core.Cors())
//	//跨域名
//	//r.Use(core.Cors())
//	logging.Init()
//	models.Init()
//	r.POST("/auth",v1.Login)
//	r.GET("/auth", v1.GetAuth)
//	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
//	//127.0.0.1:8080/swagger/index.html
//	api := r.Group("/api/v1")
//	api.Use(jwt.JWT())
//	{
//		//登录
//		api.POST("/login",v1.Login)
//		//注册
//		api.POST("/register",v1.Register)
//		//查看信息
//		api.GET("/profile",v1.Profile)
//		//获取用户发货位置
//		api.GET("/deliver_address",v1.DeliverAddress)
//		//获取用户收货位置
//		api.GET("/receive_address",v1.ReceiveAddress)
//		//评价
//		api.GET("/evaluate",v1.Evaluate)
//		//查看评价
//		api.GET("/comprehensive",v1.Comprehensive)
//		//ip获取位置
//		api.GET("/get_address",v1.GetAddress)
//	}
//	return r
//}

func InitRouter() *gin.Engine{
	gin.SetMode(setting.RunMode)
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(core.Cors())
	logging.Init()
	models.Init()
	return r
}
