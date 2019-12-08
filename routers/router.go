package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goods/docs"
	"goods/middleware/jwt"
	"goods/models"
	"goods/pkg/logging/serve"
	"goods/pkg/setting"
	"goods/routers/v1"
)
//初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(setting.RunMode)
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	logging.Init()
	models.Init()
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/auth", v1.GetAuth)
	//r.POST("/upload", api.UploadImage)
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/v1")
	api.Use(jwt.JWT())
	{
		//登录
		api.POST("/login",v1.Login)
		//查看信息
		api.GET("/profile",v1.Profile)
	}
	return r
}
