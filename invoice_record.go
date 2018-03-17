package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/galahade/invoice_record/controller"
	_ "flag"
	"github.com/gin-contrib/sessions"
	"github.com/galahade/invoice_record/middleware"
	"github.com/galahade/invoice_record/util"
)

func main() {
	util.Config = util.LoadYamflConfigFile("config.yml")
	//	flag.Parse()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("wechat", store))

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
	router.Run(":8080")

}