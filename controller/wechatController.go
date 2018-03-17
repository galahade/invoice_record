package controller

import (
	"fmt"
	"net/http"
	"github.com/galahade/invoice_record/domain"
	"github.com/galahade/invoice_record/util"
	"github.com/gin-gonic/gin"
)

type wechatModel struct {
	Code string `json:"code" binding:"required"`
}

// login() ...
func Login(c *gin.Context) {
	var wechat wechatModel
	var message string
	if err := c.BindJSON(&wechat); err == nil {
		if sessionID, err := util.GenerateNewSessionID(); err == nil {
			request := new(domain.WechatSessionRequest)
			request.JsCode = wechat.Code
			if session, err := request.GetWechatSession(sessionID); err == nil {
				c.JSON(http.StatusOK, session)
				return
			} else {
				message = fmt.Sprintf("Send code to wechat api err : %s", err)
			}
		} else {
			message = fmt.Sprintf("Generate session id err : %s", err)
		}
	} else {
		message = fmt.Sprintf("JSON binging error with : %s ", err)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  message,
	})
}
