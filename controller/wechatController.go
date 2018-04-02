package controller

import (
	"fmt"
	"net/http"
	"github.com/galahade/invoice_record/domain"
	"github.com/galahade/invoice_record/middleware"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/olebedev/config"
)

type wechatModel struct {
	Code string `json:"code" binding:"required"`
}

//Login method
func Login(c *gin.Context) {
	var wechat wechatModel
	var message string
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	cfg := c.MustGet(middleware.ProjectConfigKey).(config.Config)
	if err := c.BindJSON(&wechat); err == nil {
		request := new(domain.WechatSessionRequest)
		request.JsCode = wechat.Code
		if session, err := request.GetWechatSession(conn, cfg); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"sessionid": session.SessionID,
			})
			return
		} else {
			message = fmt.Sprintf("Send code to wechat api err : %s", err)
		}

	} else {
		message = fmt.Sprintf("JSON binging error with : %s ", err)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  message,
	})
}
