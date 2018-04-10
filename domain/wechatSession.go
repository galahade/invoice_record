package domain

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/golang/glog"
	"github.com/garyburd/redigo/redis"
	"github.com/olebedev/config"
	"github.com/galahade/invoice_record/util"
)

type WechatSessionRequest struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	JsCode    string `json:"js_code"`
	GrantType string `json:"grant_type"`
}

type WechatSession struct {
	SessionID  string `json:"-"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMSG     string `json:"errmsg"`
}

func (we *WechatSessionRequest) GetWechatSession(conn redis.Conn, cfg config.Config ) (WechatSession, error) {
	var err error
	session := new(WechatSession)
	appid, _ := cfg.String("Wechat.appid") 
	secret, _ := cfg.String("Wechat.secret")
	we.Appid = appid
	we.Secret = secret
	url, err := cfg.String("Wechat.session.url")
	if err != nil {
		panic("Fail to load wechat auth config")
	}
	if err1 := session.setOpenID(we.ComposeCode2SessionURL(url)); err1 == nil {
		sessionID := util.GenerateNewSessionID()
		session.SessionID = sessionID
		if _, err = conn.Do("SET", sessionID, session.Openid); err == nil {
			conn.Do("EXPIRE", sessionID, 1800)
			var b []byte
			if b, err = json.Marshal(session); err == nil {
				glog.V(3).Infof("wechat store to db value is : %s", string(b))
				if _, err = conn.Do("SET", fmt.Sprintf("openid::%s", session.Openid), b); err == nil {
					return *session, nil
				}
			}
		}
	} else {
		err = err1
	}

	return *session, err
}

//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func (we WechatSessionRequest) ComposeCode2SessionURL(url string) string {
	return fmt.Sprintf("%sappid=%s&secret=%s&js_code=%s", url, we.Appid, we.Secret, we.JsCode)
}

func (session *WechatSession) setOpenID(url string) error {
	glog.V(3).Infof("wechat server url is : %s", url)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	statusCode := resp.StatusCode()
	glog.V(3).Infof("Wechat check session interface response code is %d.\n", statusCode)
	bodyBytes := resp.Body()
	json.Unmarshal(bodyBytes, session)
	glog.V(3).Infof("wechat check session interface raw string %s\n", string(bodyBytes))
	glog.V(2).Infof("wechat check session interface response is : %#v\n", session)
	if session.ErrCode != 0 {
		err = fmt.Errorf("wechat check session interface err. errCode is %d, errMessage is %s", session.ErrCode, session.ErrMSG)
	}
	return err
}
