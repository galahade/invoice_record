package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/olebedev/config"
	"github.com/golang/glog"
	"os"
)

const (
	//DateString is date format
	DateString string = "20060102"

	DateDashString string = "2006-01-02"
	//TimeString is datetime format
	TimeString string = "2006-01-02T15:04:05.000Z"
)

//GenerateNewSessionID method used to generate 16 byte random session id
func GenerateNewSessionID() (string) {
	b := make([]byte, 16)
	var sessionID string
	if _, err := rand.Read(b); err == nil {
		sessionID =  base64.URLEncoding.EncodeToString(b)
	}
	return sessionID
}

// LoadYamlConfigFile load yaml config by given file path
func LoadYamlConfigFile(path string) config.Config {
	if cfg, err := config.ParseYamlFile(path); err != nil {
		panic(fmt.Sprintf("Fail to load config file in path : %s \n", path))
	} else {
		glog.V(2).Infof("Load Yaml config file %s success\n", path)
		return *cfg
	}
}

// GetRootPath get golang $GOPATH PATH
func GetRootPath() string {
	goPath, ok := os.LookupEnv("GOPATH")
	if !ok {
		glog.Fatal("You may don't have GOPATH env or you may need run command by sudo -E")
	} else {
		goPath += "/src/github.com/galahade/invoice_record"
	}
	glog.V(2).Infof("Runtime root path is %s", goPath)
	return goPath
}



