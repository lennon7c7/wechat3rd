package wechat3rd_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/wechat3rd"
	"log"
	"testing"
)

const (
	AppID     = "wx88773bff1cd12c8c"                          //第三方平台appid
	AppSecret = "d08563f46dc21b9eaecb2067f140b073"            //第三方平台app_secret
	AESKey    = "CKUE1iJDu83kgbZ4NfL6cYhz2xs7mGo9QnTAp5yRH0e" //第三方平台消息加解密Key
	Token     = "UklM5J3G2BpZXnLoq0Namx84rjtihYSF"            //消息校验Token

	Ticket = "ticket@@@bw-1kI2L4e-R7B78byTo1RfGAaXIjq9t42W7C2TmnyMflZvdjBPVtV0vC0gvjkjYqlqzuT1mFarj8dwHfZ2XKg" // 票据
)

const (
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = "root"
	RedisDB       = 0
)

func TestGetToken(t *testing.T) {
	// 除Config外的其它参数传nil则使用默认配置.  该处代码你应该使用单例模式或服务池方式来管理
	service, err := wechat3rd.NewService(wechat3rd.Config{
		AppID:     AppID,
		AppSecret: AppSecret,
		AESKey:    AESKey,
		Token:     Token,
	}, redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       RedisDB,
	}), nil, nil, nil)
	if err != nil {
		t.Error("NewService err：" + err.Error())
	}

	err = service.SetTicket(Ticket)
	if err != nil {
		t.Error("SetTicket err：" + err.Error())
	}

	token, err := service.Token()
	if err != nil {
		t.Error("Token err：" + err.Error())
	}
	log.Println(token)
}
