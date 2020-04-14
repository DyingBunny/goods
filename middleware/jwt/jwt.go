package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/pkg/e"
	"goods/pkg/util"
	"net/http"
	"time"
)
type token struct{
	Token  string `json:"token"`
}

//验证token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token token
		var code int
		var data interface{}
		code = e.SUCCESS
		_ = c.ShouldBindBodyWith(&token,binding.JSON)
		//token := c.Query("token")
		if token.Token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token.Token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}