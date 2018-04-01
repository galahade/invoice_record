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
	sessionID := util.GenerateNewSessionID()
	cfg := util.LoadYamlConfigFile(fmt.Sprintf("%s/%s", util.GetRootPath(), "config.yml"))
	pool := util.GetRedisPool(cfg)
	defer pool.Close()
	conn := pool.Get()
	defer conn.Close()
	session, err := request.GetWechatSession(sessionID, conn, cfg)
	fmt.Printf("error is : %s", err)
	assert.Empty(t, err,)
	assert.NotEmpty(t,request.Appid)
	assert.NotEmpty(t, request.Secret)
	assert.NotEmpty(t, session.Openid)
	assert.NotEmpty(t, session.SessionKey)
	assert.NotEmpty(t, session.Unionid)
	fmt.Printf("session openid is %s", session.Openid)
}