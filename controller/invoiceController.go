package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/galahade/invoice_record/domain"
	"github.com/galahade/invoice_record/middleware"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
			invoiceM.Date = invoice.Date
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
	invoiceCodeNo := c.Param("invoice_code")
	invoiceInfos := strings.Split(invoiceCodeNo, ":")
	if len(invoiceInfos) == 2 {
		conn := c.MustGet(middleware.RedisConnKey).(redis.Conn)
		defer conn.Close()
		invoice, ok := domain.QueryByNo(invoiceInfos[0], invoiceInfos[1], getOpenID(c), conn)
		if ok {
			invoiceM := new(InvoiceResponseModel)
			invoiceM.InvoiceCode = invoice.Code
			invoiceM.Number = invoice.No
			invoiceM.Amount = invoice.Amount
			invoiceM.Date = invoice.Date
			invoiceM.CreateDate = domain.JsonDashTime(invoice.CreateDate)
			c.JSON(http.StatusOK, invoiceM)
		} else {
			c.JSON(http.StatusNotFound, "{}")
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  ErrorStatus,
			"message": "Request param must be in format of code:no",
		})
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
