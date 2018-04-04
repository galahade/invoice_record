package controller

import (
	"github.com/galahade/invoice_record/middleware"
	"github.com/galahade/invoice_record/domain"
	"time"
)

type InvoiceResponseModel struct {
	middleware.WechatBaseModel
	InvoiceCode string `json:"code"`
	Number string `json:"number"`
	Amount string `json:"amount"`
	Date domain.JsonDashTime `json:"date"`
	CreateDate domain.JsonDashTime `json:"createDate"`
}

type invoiceModel struct {
	InvoiceCode string `json:"code"`
	Number string `json:"number"`
	Amount string `json:"amount"`
	Date domain.JsonTime `json:"date"`
	CreateDate time.Time `json:"createDate"`
}

type InvoiceListResponseModel struct {
	middleware.WechatBaseModel
	InvoiceList []InvoiceResponseModel `json:"invoiceList"`
}
