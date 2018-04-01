package main

import (
	"fmt"
	"log"
	"net/http"
	c "github.com/galahade/invoice_record/controller"
	"github.com/galahade/invoice_record/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-glog"
	"time"
	"github.com/galahade/invoice_record/util"
)

var env string
var port int

func main() {
	getParams()
	initProjectConfit()
	initRedis()
	router :=gin.New()
	router.Use(ginglog.Logger(3 * time.Second), gin.Recovery())
	router.Use(middleware.SetupConfig(cfg))
	router.Use(middleware.SetupRedisConn(pool))
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("wechat", store))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "welcome to my site.",
		})
	})
	router.POST("/wechat/login", c.Login)

	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthWechat())
	{
		authorized.GET("/test", c.TestSession)
		authorized.GET("/invoices/:invoice_code", c.GetInvoiceInfoByNo)
		authorized.POST("/invoice", c.AddInvoice)
		authorized.GET("/invoices", c.GetInvoicesList)
		//		authorized.POST("/read", readEndpoint)

		// nested group
		//	testing := authorized.Group("testing")
		//		testing.GET("/analytics", analyticsEndpoint)
	}
	//router.Run(fmt.Sprintf(":%d",port))
	certfffificaton := fmt.Sprintf("%s/%s",util.GetRootPath(),"2_wechat.yuboxuan.club.crt")
	key := fmt.Sprintf("%s/%s", util.GetRootPath(), "3_wechat.yuboxuan.club.key")
	err := router.RunTLS(fmt.Sprintf(":%d", port), certfffificaton, key)
	if err != nil {
		log.Fatal(err)
	}
	pool.Close()
}

