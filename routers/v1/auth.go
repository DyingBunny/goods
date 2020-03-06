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
// @Summary token获取
// @Produce  json
// @Param name query string true "username,password"
// @Success 200 {string} json "{"code":code,"msg" : e.GetMsg(code),"data":{"token":token,"role":role,"login_time":time.Now()}}"
// @Failure 400 {string} json "{"code":code,"msg" : e.GetMsg(code)}"
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
				userId :=models.FindUserId(username,password)
				role:=models.FindRole(username,password)
				data["token"] = token
				data["role"]=role
				data["login_time"]=time.Now()
				code = e.SUCCESS
				models.UpdateSta(userId)
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
	})
}
// @Summary 登录
// @Produce  json
// @Param name query string true "username,password"
// @Success 200 {string} json "{"code":code,"msg" : e.GetMsg(code),	"data":{"login_time":time.Now()}}"
// @Failure 400 {string} json "{"code":code,"msg" : e.GetMsg(code)}"
// @Router /login [post]
func Login(context *gin.Context){
	var people table.Login
	if context.BindJSON(&people)==nil{
		if models.Check(people.Username,people.Password){
			userid:=models.FindUserId(people.Username,people.Password)
			models.UpdateSta(userid)
			data := make(map[string]interface{})
			code:=e.SUCCESS
			data["login_time"]=time.Now()
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
				"data":data,
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
// @Param name query string true "user_id"
// @Success 200 {string} json "{"code" : code,	"msg" : e.GetMsg(code),	"data":{"user_id":auth.User_id,	"username":auth.Username,	"password":auth.Password,	"phone_number":auth.PhoneNumber,	"gender":auth.Gender,	"email":auth.Email,	"login":auth.Login,	"role":auth.Role,"role_id":auth.RoleId}}"
// @Failure 400 {string} json "{"code" : code,	"msg" : e.GetMsg(code)}"
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
		data["phone_number"]=auth.PhoneNumber
		data["gender"]=auth.Gender
		data["email"]=auth.Email
		data["login"]=auth.Login
		data["role"]=auth.Role
		data["role_id"]=auth.RoleId
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
			//"data":auth.UserId,auth.Username,
		})
	}
}
// @Summary 注册
// @Produce  json
// @Param name query string true "username,password,phone_number,gender,email,role"
// @Success 200 {string} json "{"code":code,"msg" : e.GetMsg(code)}"
// @Failure 400 {string} json "{"code":code,"msg" : e.GetMsg(code)}"
// @Router /register [post]
func Register(context *gin.Context){
	var people table.Login
	if context.BindJSON(&people)==nil{
		if models.Check(people.Username,people.Password)==false{
			models.AddUser(people.Username,people.Password,people.PhoneNumber,people.Gender,people.Email,people.Role)
			code:=e.SUCCESS
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
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

