package domain

import (
	"fmt"

	"github.com/galahade/invoice_record/util"
	"github.com/valyala/fasthttp"
)

type WechatSessionRequest struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	JsCode    string `json:"js_code"`
	GrantType string `json:"grant_type"`
}

type WechatSession struct {
	SessionID  string `json:"sessionid"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

func (we *WechatSessionRequest) GetWechatSession(sessionID string) (WechatSession, error) {
	var err error
	session := new(WechatSession)
	appid, _ := util.Config.String("Wechat.appid")
	secret, _ := util.Config.String("Wechat.secret")
	we.Appid = appid
	we.Secret = secret
	url, err := util.Config.String("Wechat.session.url")
	if err != nil {
		panic("Fail to load wechat auth config")
	}
	if session.getOpenID(url); err == nil {
		//TO-DO this need to implement by http request.
		session.SessionID = sessionID
		session.Openid = "o5gGe4khB5GaEXO-Dn2waDD13zSs"
		session.SessionKey = "egvKJxQKzT2IamEYHQRLQA=="
		session.Unionid = "unionid"
		redisClient := util.GetRedisClient(util.Config)
		conn := redisClient.Get()
		defer conn.Close()
		if _, err = conn.Do("SET", sessionID, session.Openid); err == nil {
			conn.Do("EXPIRE", sessionID, 1800)
			var b []byte
			if b, err = json.Marshal(session); err == nil {
				if _, err = conn.Do("SETNX", session.Openid, b); err == nil {
					return *session, nil
				}
			}
		}
	}
	//TODO also need to store these value to

	return *session, err
}

//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func (we WechatSessionRequest) ComposeCode2SessionURL(url string) string {
	return fmt.Sprintf("%sappid=%s&secret=%s&js_code=%s", url, we.Appid, we.Secret, we.JsCode)
}

func (session *WechatSession) getOpenID(url string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	//  bodyBytes := resp.Body()
	//    println(string(bodyBytes))
	return err
}
