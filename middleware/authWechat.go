package middleware

import (
	"fmt"
	"net/http"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//WechatBaseModel response to http request
type WechatBaseModel struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
}

//AuthWechat if login with wechat, wechat session ID is stored in json named "sessionid"
func AuthWechat() gin.HandlerFunc {
	var status int
	var wechatsession WechatBaseModel
	return func(c *gin.Context) {
		if sessionid := extractSessionID(c); sessionid != "" {
			conn := c.MustGet(RedisConnKey).(redis.Conn)
			defer conn.Close()
			if b, err := redis.Bytes(conn.Do("GET", sessionid)); err == nil {
				openid := string(b)
				session := sessions.Default(c)
				session.Set("openid", openid)
				c.Next()
				conn.Do("EXPIRE", sessionid, 1800)
				return
			} else {
				status = http.StatusNotAcceptable
				wechatsession.Message = fmt.Sprintf("Can't get openid by sessionid %s, error is : %s", sessionid, err)
			}
		} else {
			status = http.StatusUnauthorized
			wechatsession.Message = fmt.Sprint("There is no sessionid in header.")
		}
		glog.Errorf("AuthWechat error: %s", wechatsession.Message)
		wechatsession.Status = "error"
		c.AbortWithStatusJSON(status, wechatsession)

	}
}

func extractSessionID(c *gin.Context) string {
	return c.Request.Header.Get("sessionid")
}
