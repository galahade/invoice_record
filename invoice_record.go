package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/galahade/invoice_record/controller"
	_ "flag"
	"github.com/gin-contrib/sessions"
	"github.com/galahade/invoice_record/middleware"
	"github.com/galahade/invoice_record/util"
	"os"
	"fmt"
	"flag"
)

func main() {
	var env string
	var port int
	flag.StringVar(&env, "env", "", "application enviroment")
	flag.IntVar(&port, "p", 8080, "application port number")
	flag.Parse()
	setConfigFile(env)
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
	router.Run(fmt.Sprintf(":%d",port))
}

func setConfigFile(env string) {
	path, _ := os.Getwd()
	var configFilePath string
	switch env {
	case "":
		configFilePath = "config.yml"
	case "test":
		configFilePath = "config_test.yml"
	case "prod":
		configFilePath = "config_prod.yml"
	default:
		configFilePath = "config.yml"
	}
	util.Config = util.LoadYamflConfigFile(fmt.Sprintf("%s/%s",path, configFilePath))
}