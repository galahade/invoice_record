package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func TestSession(c *gin.Context) {
	session := sessions.Default(c)
	openid := session.Get("openid").(string)
	c.JSON(http.StatusOK, gin.H{
		"openid": openid,
	})
}