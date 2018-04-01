package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/galahade/invoice_record/util"
	"flag"
	"fmt"
	"github.com/olebedev/config"
	_"github.com/golang/glog"
)


var pool *redis.Pool
var cfg config.Config

func initRedis() {
	pool = util.GetRedisPool(cfg)
}

func initProjectConfit() {
	cfg = util.LoadYamlConfigFile(getConfigFileByEnv())
}

func getParams() {
	flag.StringVar(&env, "env", "", "application enviroment")
	flag.IntVar(&port, "p", 8080, "application port number")
	flag.Parse()
}

func getConfigFileByEnv() string {
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

	return fmt.Sprintf("%s/%s", util.GetRootPath(), configFilePath)
}




