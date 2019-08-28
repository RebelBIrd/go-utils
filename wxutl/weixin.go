package wxutl

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/xlstudio/wxbizdatacrypt"
	"io/ioutil"
	"net/http"
)

type RespWXSmall struct {
	OpenId     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionId    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	Errcode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errMsg"`      //错误信息
}

type WxConf struct {
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

func (conf *WxConf) LoginWXSmall(code string) (wxInfo RespWXSmall, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, conf.AppId, conf.AppSecret, code))
	if err != nil {
		return wxInfo, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &wxInfo)
	if err != nil {
		return wxInfo, err
	}
	if wxInfo.Errcode != 0 {
		return wxInfo, errors.New(fmt.Sprintf("code: %d, errmsg: %s", wxInfo.Errcode, wxInfo.ErrMsg))
	}
	return wxInfo, nil
}

func (conf *WxConf) DecryptWXOpenData(sessionKey, encryptData, iv string) (string, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: conf.AppId, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptData, iv, true)
	if err != nil {
		logutl.Error(err.Error())
		return "", err
	} else {
		decodeStr := result.(string)
		return decodeStr, nil
	}
}
