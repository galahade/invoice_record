package controller

import (
	"github.com/galahade/invoice_record/middleware"
	"github.com/galahade/invoice_record/domain"
)

type InvoiceResponseModel struct {
	middleware.WechatBaseModel
    InvoiceModel
	CreateDate domain.JsonTime `json:"createDate"`
}

type InvoiceModel struct {
	InvoiceCode string `json:"code"`
	Number string `json:"number"`
	Amount string `json:"amount"`
	Date *domain.JsonTime `json:"date, omitempty"`
	SubmitPerson string `json:"submit_person, omitempty"`
	Note string `json:"note, omitempty"`
}

type InvoiceListResponseModel struct {
	middleware.WechatBaseModel
	InvoiceList []InvoiceResponseModel `json:"invoiceList"`
}
