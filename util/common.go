package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/olebedev/config"
)

const (
	DateString string = "2006-01-02"
	TimeString string = "2006-01-02T15:04:05.000Z"
)

var Config config.Config

//This method used to generate 16 byte random session id
func GenerateNewSessionID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err == nil {
		return base64.URLEncoding.EncodeToString(b), nil
	} else {
		return string(""), err
	}
}

func LoadYamflConfigFile(path string) config.Config {
	if cfg, err := config.ParseYamlFile(path); err != nil {
		panic(fmt.Sprintf("Fail to load config file in path : %s \n", path))
	} else {
		return *cfg
	}
}
