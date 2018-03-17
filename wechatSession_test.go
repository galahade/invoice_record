package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/galahade/invoice_record/domain"
	"github.com/galahade/invoice_record/util"
)

func TestWechatSession(t *testing.T) {
	request := new(domain.WechatSessionRequest)
	sessionID, err := util.GenerateNewSessionID()
	session, err := request.GetWechatSession(sessionID)
	fmt.Printf("error is : %s", err)
	assert.Empty(t, err,)
	assert.NotEmpty(t,request.Appid)
	assert.NotEmpty(t, request.Secret)
	assert.NotEmpty(t, session.Openid)
	assert.NotEmpty(t, session.SessionKey)
	assert.NotEmpty(t, session.Unionid)
	fmt.Printf("session openid is %s", session.Openid)
}