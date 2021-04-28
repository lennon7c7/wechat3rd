package wechat3rd

import (
	"github.com/lennon7c7/wechat3rd/core"
)

type GetAccountBasicInfoResp struct {
	core.Error
	Appid          string `json:"appid"`
	AccountType    int8   `json:"account_type"`
	PrincipalType  int    `json:"principal_type"`
	PrincipalName  string `json:"principal_name"`
	RealnameStatus int8   `json:"realname_status"`
	WxVerifyInfo   struct {
		QualificationVerify   bool   `json:"qualification_verify"`
		NamingVerify          bool   `json:"naming_verify"`
		AnnualReview          *bool  `json:"annual_review,omitempty"`
		AnnualReviewBeginTime *int64 `json:"annual_review_begin_time,omitempty"`
		AnnualReviewEndTime   *int64 `json:"annual_review_end_time,omitempty"`
	} `json:"wx_verify_info"`
	SignatureInfo struct {
		Signature       string `json:"signature"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"signature_info"`
	HeadImageInfo struct {
		HeadImageUrl    string `json:"head_image_url"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"head_image_info"`
	NicknameInfo struct {
		Nickname        string `json:"nickname"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"nickname_info"`
	RegisteredCountry int `json:"registered_country"`
}

//获取基本信息
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html
func (s *Server) GetAccountBasicInfo(authToken string) (resp *GetAccountBasicInfoResp, err error) {
	var (
		u = CGIUrl + "/account/getaccountbasicinfo?"
	)
	resp = &GetAccountBasicInfoResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(authToken), resp)
	return
}

type ModifyDomainReq struct {
	Action          string   `json:"action"`
	Requestdomain   []string `json:"requestdomain"`
	Wsrequestdomain []string `json:"wsrequestdomain"`
	Uploaddomain    []string `json:"uploaddomain"`
	Downloaddomain  []string `json:"downloaddomain"`
}

type ModifyDomainResp struct {
	core.Error
	Requestdomain   []string `json:"requestdomain"`
	Wsrequestdomain []string `json:"wsrequestdomain"`
	Uploaddomain    []string `json:"uploaddomain"`
	Downloaddomain  []string `json:"downloaddomain"`
}

//设置服务器域名
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html
func (s *Server) ModifyDomain(authToken string, req *ModifyDomainReq) (resp *ModifyDomainResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxa/modify_domain?"
	)
	resp = &ModifyDomainResp{}
	err = core.PostJson(s.AuthToken2url(u, authToken), req, resp)
	return
}

type SetWebviewDomainReq struct {
	Action        *string  `json:"action,omitempty"`
	Webviewdomain []string `json:"webviewdomain,omitempty"`
}

//设置业务域名
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setwebviewdomain.html
func (s *Server) SetWebviewDomain(authToken string, req *SetWebviewDomainReq) (resp *core.Error, err error) {
	var (
		u = WECHAT_API_URL + "/wxa/setwebviewdomain?"
	)
	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, authToken), req, resp)
	return
}
