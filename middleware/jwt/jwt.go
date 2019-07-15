package jwt

import (
	"ginapi/pkg/e"
	"ginapi/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//import (
//"time"
//"net/http"
//
//"github.com/gin-gonic/gin"
//
//"gin-blog/pkg/util"
//"gin-blog/pkg/e"
//)

func JWT() (gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		code := 1
		token := ctx.Query("token")
		claims, err := util.ParseToken(token)
		if err {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
		if code != 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMssg(code),
				"data": make(map[string]interface{}),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
