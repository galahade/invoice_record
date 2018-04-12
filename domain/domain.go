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
	glog.V(3).Infof("JsonTime unmarshalJSON byte array value is : %s\n", string(b))
	var jsonValue string
	if strings.HasPrefix(string(b), "\"") {
		jsonValue = string(b[1 : len(b)-1])
	} else {
		jsonValue = string(b)
	}
	glog.V(3).Info(jsonValue)
	var tValue time.Time
	var err error
	if len(jsonValue) == 8 {

	} else if len(jsonValue) == 10 {

	}
	switch len(jsonValue) {
	case 8:
		tValue, err = time.Parse(util.DateString, jsonValue)
	case 10:
		tValue, err = time.Parse(util.DateDashString, jsonValue)
	default:
		err = fmt.Errorf("Date format error, should use 20060102 or 2006-01-02 but %s\n", jsonValue)
	}
	if err == nil {
		*t = JsonTime(tValue)
		return nil
	}
	glog.Errorf("JsonTime unmarshalJson err with message : %s\n", err)
	return err
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	value := time.Time(t)
	glog.V(3).Infof("JsonTime marshal jsone value is : %s\n", value)
	result := value.Format(util.DateDashString)
	glog.V(2).Info(result)
	return []byte(fmt.Sprintf(`"%s"`, result)), nil
}







