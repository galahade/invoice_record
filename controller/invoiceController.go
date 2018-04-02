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

func GetInvoicesList(c *gin.Context) {
	status := http.StatusBadRequest
	response := new(InvoiceListResponseModel)
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	response.Status = "OK"
	if invoices, err := domain.QueryAllInvoices(getOpenID(c), conn); err == nil {
		invoiceModleList := make([]invoiceModle, len(invoices))
		response.InvoiceList = invoiceModleList
		for i, invoice := range invoices {
			invoiceM := new(invoiceModle)
			invoiceM.InvoiceCode = invoice.Code
			invoiceM.Number = invoice.No
			invoiceM.Amount = invoice.Amount
			invoiceM.Date =  invoice.Date
			invoiceM.CreateDate = invoice.CreateDate
			invoiceModleList[i] = *invoiceM
		}
		status = http.StatusOK
	} else {
		response.Status = "error"
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
	invoiceModle := new(InvoiceResponseModle)
	conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
	defer conn.Close()
	if err := c.ShouldBindJSON(invoiceModle); err == nil {
		invoice := new(domain.Invoice)
		invoice.Code = invoiceModle.InvoiceCode
		invoice.No = invoiceModle.Number
		invoice.Amount = invoiceModle.Amount
		invoice.Date = invoiceModle.Date
		if ok, err := domain.CreateNewInvoice(invoice, getOpenID(c), conn); err == nil {
			if !ok {
				status = http.StatusFound
				invoiceModle.Status = "This invoice info has been stored."

			} else {
				status = http.StatusCreated
				invoiceModle.Status = "OK"
				invoiceModle.InvoiceCode = invoice.Code
				invoiceModle.Number = invoice.No
				invoiceModle.Amount = invoice.Amount
				invoiceModle.Date = invoice.Date
				invoiceModle.CreateDate = invoice.CreateDate
			}
		} else {
			status = http.StatusBadRequest
			invoiceModle.Status = "error"
			invoiceModle.Message = fmt.Sprintf("Store invoice err : %s", err)
		}
	} else {
		invoiceModle.Status = "error"
		invoiceModle.Message = fmt.Sprintf("invoice json struct error : %s", err)
	}
	c.JSON(status, invoiceModle)

}

func getOpenID(c *gin.Context) string {
	return sessions.Default(c).Get("openid").(string)
}
