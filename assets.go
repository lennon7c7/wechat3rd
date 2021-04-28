package wechat3rd

import "github.com/lennon7c7/wechat3rd/core"

type MaterialItem struct {
	Title string `json:"title"`
	ThumbMediaId string `json:"thumbMediaId"`
	ShowCoverPic int8 `json:"showCoverPic"`
	Author string `json:"author"`
	Digest string `json:"digest"`
	Content string `json:"content"`
	Url string `json:"url"`
	ContentSourceUrl string `json:"contentSourceUrl"`
}

type GetMaterialResp struct {
	core.Error
	//图文内容响应结果
	NewsItem []*MaterialItem `json:"newsItem,omitempty"`

	//视频消息响应结果
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	DownUrl *string `json:"downUrl,omitempty"`
}

//获取已上传的代码的页面列表
//https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (s *Server) GetMaterial(accessToken string,mediaId string) (resp *GetMaterialResp, err error) {
	var (
		u = CGIUrl + "/material/get_material?"
		req = &struct {
			MediaId string `json:"mediaId"`
		}{MediaId: mediaId}
	)
	resp = &GetMaterialResp{}

	err = core.PostJson(s.AuthToken2url(u,accessToken),req , resp)
	return
}
