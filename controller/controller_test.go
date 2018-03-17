package controller

import (
	"testing"
)


func TestInvoiceModleSetValue(t *testing.T) {
	invoiceModle := new(InvoiceModle)
	invoiceModle.InvoiceCode = "code"
	invoiceModle.Status = "error"
}