package middleware

import (
	"fmt"
	"net/http"

	"github.com/galahade/invoice_record/util"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type WechatBaseModel struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
}

//Check if login with wechat, wechat session ID is stored in json named "sessionid"
func AuthWechat() gin.HandlerFunc {
	var wechatsession WechatBaseModel
	return func(c *gin.Context) {
		if sessionid := extractSessionID(c); sessionid != "" {
			conn := util.GetRedisClient(util.Config).Get()
			defer conn.Close()
			if b, err := redis.Bytes(conn.Do("GET", sessionid)); err == nil {
				openid := string(b)
				session := sessions.Default(c)
				session.Set("openid", openid)
				c.Next()
				return
			} else {
				wechatsession.Message = fmt.Sprintf("Can't get openid by sessionid %s, error is : %s", sessionid, err)
			}
		} else {
			wechatsession.Message = fmt.Sprint("There is no sessionid in header.")
		}
		wechatsession.Status = "error"
		c.AbortWithStatusJSON(http.StatusUnauthorized, wechatsession)

	}
}

func extractSessionID(c *gin.Context) string {
	return c.Request.Header.Get("sessionid")
}
