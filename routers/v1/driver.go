package v1

import (
	"github.com/gin-gonic/gin"
	"goods/models"
	"goods/pkg/e"
	"net/http"
)

// @Summary 评价
// @Produce  json
// @Param name query string true "user_id,mark"
// @Success 200 {string} json "{"code":code,	"msg" : e.GetMsg(code)}"
// @Failure 400 {string} json "{"code":code,	"msg" : e.GetMsg(code)}"
// @Router /evaluate [get]
func Evaluate(context *gin.Context){
	UserId := context.Query("user_id")
	Mark:=context.Query("mark")
	if models.EvaDriver(UserId,Mark) {
		code:=e.SUCCESS
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			//"data":auth.UserId,auth.Username,
		})
	} else{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg" : e.GetMsg(code),
		})
	}
}

// @Summary 查看司机评价
// @Produce  json
// @Param name query string true "user_id"
// @Success 200 {string} json "{"code":code,"msg" : e.GetMsg(code),"data":{"mark":mark}}"
// @Failure 400 {string} json "{"code":code,"msg" : e.GetMsg(code)}"
// @Router /comprehensive [get]
func Comprehensive(context *gin.Context){
	UserId := context.Query("user_id")
	driver, _ :=models.FindDriver(UserId)
	mark:=driver.Evaluation/driver.Count
	code:=e.ERROR
	data:=make(map[string]interface{})
	data["mark"]=mark
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg" : e.GetMsg(code),
		"mark":mark,
	})
}