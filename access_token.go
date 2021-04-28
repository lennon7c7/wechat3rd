package wechat3rd

import (
	"github.com/lennon7c7/wechat3rd/core"
	"github.com/lennon7c7/wechat3rd/util"
	"time"
)

const (
	componentAccessTokenUrl = WECHAT_API_URL + "/cgi-bin/component/api_component_token"
)

type AccessTokenServer interface {
	Token() (token string, err error)
}

type DefaultAccessTokenServer struct {
	TicketServer
	AppID     string
	AppSecret string

	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"` // 当前时间 + 过期时间
}

// token不使用不获取
func (d *DefaultAccessTokenServer) Token() (token string, err error) {
	var (
		ticket string
		resp   *AccessTokenResp
	)

	token, _ = util.CacheGetString("cachekey_of_wechat3rd_token")
	if token != "" {
		return
	}

	if d.ExpiresIn <= time.Now().Unix()-30 {
		ticket, err = d.GetTicket()
		if err != nil {
			return
		}
		resp, err = newAccessToken(&AccessTokenReq{
			ComponentAppid:        d.AppID,
			ComponentAppsecret:    d.AppSecret,
			ComponentVerifyTicket: ticket,
		})
		if err != nil {
			return
		}

		err = util.CacheSetString("cachekey_of_wechat3rd_token", resp.ComponentAccessToken, time.Hour*2)
		if err != nil {
			return
		}

		d.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
		d.ComponentAccessToken = resp.ComponentAccessToken
	}
	return d.ComponentAccessToken, nil
}

type AccessTokenResp struct {
	core.Error
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

type AccessTokenReq struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// 获取第三方应用token
func newAccessToken(r *AccessTokenReq) (*AccessTokenResp, error) {
	resp := &AccessTokenResp{}
	err := core.PostJson(componentAccessTokenUrl, r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
