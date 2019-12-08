package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"goods/pkg/logging/serve"
	"goods/pkg/util"
	"net/http"
	"time"
)
// @Summary token获取b
// @Produce  json
// @Param name query string true "username,password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"fail"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := table.Login{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.Check(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				userid:=models.FindUserId(username,password)
				role:=models.FindRole(username,password)
				data["token"] = token
				data["role"]=role
				code = e.SUCCESS
				models.UpdateSta(userid)
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
		"Logintime":time.Now(),
	})
}
// @Summary 登录
// @Produce  json
// @Param name query string true "username,password"
// @Success 200 {string} json "{"code":code,	"msg" : e.GetMsg(code),	"Logintime":time.Now()}"
// @Failure 400 {string} json "{"code":code,	"msg" : e.GetMsg(code),}"
// @Router /login [post]
func Login(context *gin.Context){
	var people table.Login
	if context.BindJSON(&people)==nil{
		if models.Check(people.Username,people.Password){
			userid:=models.FindUserId(people.Username,people.Password)
			models.UpdateSta(userid)
			code:=e.SUCCESS
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
				"Logintime":time.Now(),
			})
		}else{
			code:=e.ERROR
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
			})
		}
	}
}
// @Summary 信息查看
// @Produce  json
// @Param name query string true "username"
// @Success 200 {string} json "{"code" : code,	"msg" : e.GetMsg(code),	"user_id":auth.User_id,	"username":auth.Username,	"password":auth.Password,	"gender":auth.Gender,	"email":auth.Email,	"login":auth.Login,	"role":auth.Role,}"
// @Failure 400 {string} json "{"code" : code,	"msg" : e.GetMsg(code),}"
// @Router /profile [get]
func Profile(context *gin.Context){
	UserId := context.Query("user_id")
	auth,isExist := models.Find(UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		data := make(map[string]interface{})
		data["user_id"]=auth.UserId
		data["username"]=auth.Username
		data["password"]=auth.Password
		data["gender"]=auth.Gender
		data["email"]=auth.Email
		data["login"]=auth.Login
		data["role"]=auth.Role
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
			//"data":auth.UserId,auth.Username,
		})
	}
}

