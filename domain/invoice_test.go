package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
	_"fmt"

)

func TestComposeCode2Session(t *testing.T) {
	//fmt.Print(getWechatURL())

}

func TestGetOpenID(t *testing.T) {
/*	session := new(WechatSession)
	err := session.getOpenID(getWechatURL())
	fmt.Printf("err message is %#v", err) */
}

func TestQueryAllInvoice(t *testing.T) {
	invoiceList, err := QueryAllInvoices("o5gGe4khB5GaEXO-Dn2waDD13zSs")
	assert.Empty(t, err)
	assert.NotEmpty(t,invoiceList)
	assert.NotEmpty(t, invoiceList[0])
}