package controller

import (
	"fmt"
	"net/http"
	"github.com/galahade/invoice_record/domain"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"github.com/galahade/invoice_record/middleware"
)

const (
	ErrorStatus = "error"
)

func GetInvoicesList(c *gin.Context) {
	status := http.StatusBadRequest
	response := new(InvoiceListResponseModel)
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	response.Status = "OK"
	if invoices, err := domain.QueryAllInvoices(getOpenID(c), conn); err == nil {
		invoiceModleList := make([]InvoiceResponseModel, len(invoices))
		response.InvoiceList = invoiceModleList
		for i, invoice := range invoices {
			invoiceM := new(InvoiceResponseModel)
			invoiceM.InvoiceCode = invoice.Code
			invoiceM.Number = invoice.No
			invoiceM.Amount = invoice.Amount
			invoiceM.Date =  invoice.Date
			invoiceM.CreateDate = domain.JsonDashTime(invoice.CreateDate)
			invoiceModleList[i] = *invoiceM
		}
		status = http.StatusOK
	} else {
		response.Status = ErrorStatus
		response.Message = fmt.Sprintf("Query Invoices from DB err : %s", err)
	}
	c.JSON(status, response)
}

// Get invoice info by invoice number
func GetInvoiceInfoByNo(c *gin.Context) {
	invoiceCode := c.Param("invoice_code")
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	invoice, ok := domain.QueryByNo(invoiceCode, getOpenID(c), conn)
	if ok {
		c.JSON(http.StatusOK, invoice)
	} else {
		c.JSON(http.StatusNotFound, "{}")
	}
}

func AddInvoice(c *gin.Context) {
	status := http.StatusBadRequest
	invoiceModel := new(invoiceModel)
	invoiceResponseModel := new(InvoiceResponseModel)
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	if err := c.ShouldBindJSON(invoiceModel); err == nil {
		invoice := new(domain.Invoice)
		invoice.Code = invoiceModel.InvoiceCode
		invoice.No = invoiceModel.Number
		invoice.Amount = invoiceModel.Amount
		invoice.Date = domain.JsonDashTime(invoiceModel.Date)
		if ok, err := invoice.CreateNewInvoice(getOpenID(c), conn); err == nil {
			if !ok {
				status = http.StatusFound
				invoiceResponseModel.Status = "Found"
			} else {
				status = http.StatusCreated
				invoiceResponseModel.Status = "OK"
				invoiceResponseModel.InvoiceCode = invoice.Code
				invoiceResponseModel.Number = invoice.No
				invoiceResponseModel.Amount = invoice.Amount
				invoiceResponseModel.Date = invoice.Date
				invoiceResponseModel.CreateDate = domain.JsonDashTime(invoice.CreateDate)
			}
		} else {
			status = http.StatusBadRequest
			invoiceResponseModel.Status = ErrorStatus
			invoiceResponseModel.Message = fmt.Sprintf("Store invoice err : %s", err)
		}
	} else {
		invoiceResponseModel.Status = ErrorStatus
		invoiceResponseModel.Message = fmt.Sprintf("invoice json struct error : %s", err)
	}
	c.JSON(status, invoiceResponseModel)

}

func getOpenID(c *gin.Context) string {
	return sessions.Default(c).Get("openid").(string)
}
