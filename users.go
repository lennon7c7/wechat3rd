package wechat3rd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/lennon7c7/wechat3rd/core"
	"net/url"
)

type Jscode2sessionResp struct {
	Openid string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	ErrCode int `json:"errcode,omitempty"`
	ErrMsg string `json:"errmsg,omitempty"`
}

func (s *Server) Jscode2session(appId ,jsCode string) (resp *Jscode2sessionResp,err error){
	var(
		req = make(url.Values)
		u = "https://api.weixin.qq.com/sns/component/jscode2session?"
		accessToken string
	)
	accessToken,err=s.Token()
	if err!=nil{
		return
	}
	resp = &Jscode2sessionResp{}
	req.Set("appid",appId)
	req.Set("js_code",jsCode)
	req.Set("grant_type","authorization_code")
	req.Set("component_appid",s.cfg.AppID)
	req.Set("component_access_token",accessToken)

	err=core.GetRequest(u,req,resp)
	return
}

// 解密用户信息
func (s *Server) AESCBCDecrypt(encryptData, key, iv string) (data []byte,err error) {
	var(
		encBytes []byte
		keyBytes []byte
		ivBytes []byte
		block cipher.Block
		mode cipher.BlockMode
	)
	encBytes, err = base64.StdEncoding.DecodeString(encryptData)
	if err!=nil{
		return
	}
	keyBytes, err = base64.StdEncoding.DecodeString(key)
	if err!=nil{
		return
	}
	ivBytes, err = base64.StdEncoding.DecodeString(iv)
	if err!=nil{
		return
	}

	block, err = aes.NewCipher(keyBytes)
	if err != nil {
		return
	}
	if len(encBytes) < block.BlockSize() {
		err=errors.New("ciphertext too short")
		return
	}
	if len(encBytes)%block.BlockSize() != 0 {
		err=errors.New("ciphertext is not a multiple of the block size")
		return
	}
	mode = cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(encBytes, encBytes)
	// 解填充
	encBytes = pkcs7UnPadding(encBytes)
	return encBytes, nil
}

//去除填充
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
