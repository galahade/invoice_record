package controller

import (
	"github.com/galahade/invoice_record/middleware"
	"github.com/galahade/invoice_record/domain"
	"time"
)

type InvoiceResponseModle struct {
	middleware.WechatBaseModel
	//Sessionid string `json:"sessionid"`
//	Status    string `json:"status"`
	//	Message   string `json:"message"`
	invoiceModle
}

type invoiceModle struct {
	InvoiceCode string `json:"code"`
	Number string `json:"number"`
	Amount float64 `json:"amount"`
	Date domain.JsonTime `json:"date"`
	CreateDate time.Time `json:"createDate"`
}

type InvoiceListResponseModel struct {
	middleware.WechatBaseModel
	InvoiceList []invoiceModle `json:"invoiceList"`
}
