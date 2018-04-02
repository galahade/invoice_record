package domain

import (
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
)

const (
	invoiceKeyPattern = "invoice::%s::%s"
	invoiceKeysPattern = "invoice::%s::*"
)

/*
invoice domain
*/
type Invoice struct {
	Code       string
	No         string
	Amount     string
	Date       JsonTime
	CreateDate time.Time
}


func QueryAllInvoices(openid string, conn redis.Conn) (invoiceList []Invoice, err error){
	if result, err1 := conn.Do("KEYS", fmt.Sprintf(invoiceKeysPattern, openid)); err1 == nil {
		keyList := result.([]interface{})
		if len(keyList) > 0 {
			invoiceList = make([]Invoice,len(keyList))
			for i, invoiceKey := range keyList {
				invoiceB, _ := redis.Bytes(conn.Do("GET", invoiceKey))
				invoice := new(Invoice)
				json.Unmarshal(invoiceB, invoice)
				invoiceList[i] = *invoice
			}
		}
	} else {
		err = err1
	}
	return
//	redis.Bytes(conn.Do("", args ...interface{}), err error)
}

func QueryByNo(code, openid string, conn redis.Conn) (Invoice, bool) {
	ok := true
	invoiceKey := fmt.Sprintf(invoiceKeyPattern, openid, code)
	invoice := new(Invoice)
	invoice.Code = code
	if b, err := redis.Bytes(conn.Do("GET", invoiceKey)); err == nil {
		json.Unmarshal(b, invoice)
	} else {
		ok = false
	}
	return *invoice, ok
}

func CreateNewInvoice(invoice *Invoice, openid string, conn redis.Conn) (bool, error) {
	invoice.CreateDate = time.Now()
	var err error
	if b, err1 := json.Marshal(invoice); err1 == nil {
		invoiceKey := fmt.Sprintf(invoiceKeyPattern, openid, invoice.Code)
		if status, err := conn.Do("SETNX", invoiceKey, b); err == err {
			glog.Infof("SETNX status is %d", status)
			switch status {
			case int64(0):
				glog.Infof("status 0 is %d", status)
				return false, nil
			case int64(1):
				glog.Infof("status 1 is %d", status)
				return true, nil
			}
		}
	} else {
			return false, err1
	}
	return false, err
}
