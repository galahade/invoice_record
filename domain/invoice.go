package domain

import (
	"fmt"
	"time"
	"github.com/galahade/invoice_record/util"
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
	Amount     float64
	Date       JsonTime
	CreateDate time.Time
}


func QueryAllInvoices(openid string) (invoiceList []Invoice, err error){
	conn := util.GetRedisClient().Get()
	defer conn.Close()
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

func QueryByNo(code, openid string) (Invoice, bool) {
	ok := true
	invoiceKey := fmt.Sprintf(invoiceKeyPattern, openid, code)
	invoice := new(Invoice)
	invoice.Code = code
	redisClient := util.GetRedisClient()
	conn := redisClient.Get()
	defer conn.Close()
	if b, err := redis.Bytes(conn.Do("GET", invoiceKey)); err == nil {
		json.Unmarshal(b, &invoice)
	} else {
		ok = false
	}
	return *invoice, ok
}

func CreateNewInvoice(invoice *Invoice, openid string) (ok bool, err error) {
	invoice.CreateDate = time.Now()
	redisClient := util.GetRedisClient()
	conn := redisClient.Get()
	defer conn.Close()
	if b, err := json.Marshal(invoice); err == nil {
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
			return false, err
	}
	return false, err
}
