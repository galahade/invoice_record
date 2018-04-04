package domain

import (
	"fmt"
	"strings"
	"time"
	"github.com/galahade/invoice_record/util"
	"github.com/golang/glog"
	"github.com/json-iterator/go"
)

type JsonTime time.Time
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (t *JsonTime) UnmarshalJSON(b []byte) error {
	//use slic to remove "" in begin and end of string
	glog.Infof("JsonTime unmarshalJSON byte array value is : %s", string(b))
	fmt.Printf("JsonTime unmarshalJSON byte array value is : %s", string(b))
	var jsonValue string
	if strings.HasPrefix(string(b), "\"") {
		jsonValue = string(b[1 : len(b)-1])
	} else {
		jsonValue = string(b)
	}

	glog.Info(jsonValue)
	value, err := time.Parse(util.DateString, jsonValue)
	glog.Infof("UnmarshalJSON value is : %s", value)
	if err == nil {
		*t = JsonTime(value)
		return nil
	}
	glog.Info(err)
	return err
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	value := time.Time(t)
	glog.Infof("json time is : %s", value)
	result := value.Format(util.DateString)
	glog.Info(result)
	return []byte(fmt.Sprintf(`"%s"`, result)), nil
}

type JsonDashTime time.Time

func (t *JsonDashTime) UnmarshalJSON(b []byte) error {
	//use slic to remove "" in begin and end of string
	glog.Infof("JsonTime unmarshalJSON byte array value is : %s", string(b))
	fmt.Printf("JsonTime unmarshalJSON byte array value is : %s", string(b))
	var jsonValue string
	if strings.HasPrefix(string(b), "\"") {
		jsonValue = string(b[1 : len(b)-1])
	} else {
		jsonValue = string(b)
	}
	glog.V(3).Info(jsonValue)
	value, err := time.Parse(util.DateDashString, jsonValue)
	glog.V(3).Infof("UnmarshalJSON value is : %s", value)
	if err == nil {
		*t = JsonDashTime(value)
		return nil
	}
	glog.Info(err)
	return err
}

func (t JsonDashTime) MarshalJSON() ([]byte, error) {
	value := time.Time(t)
	glog.V(3).Infof("json time is : %s", value)
	result := value.Format(util.DateDashString)
	glog.V(3).Info(result)
	return []byte(fmt.Sprintf(`"%s"`, result)), nil
}