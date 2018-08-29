package goaip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
)

type AuthObject struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

type AIPBase struct {
	AuthObject
	AppID        string
	Key          string
	Secret       string
	IsCloudUser  bool
	TimeoutInSec uint32
	Version      string
	Proxy        string
	mutex        sync.Mutex
}

func (ab *AIPBase) Init() {
	ab.TimeoutInSec = 10
	ab.Version = "2_2_0"
}

func (ab *AIPBase) SetAppID(AppID string) *AIPBase {
	ab.AppID = AppID
	return ab
}

func (ab *AIPBase) SetKey(Key string) *AIPBase {
	ab.Key = Key
	return ab
}

func (ab *AIPBase) SetSecret(Secret string) *AIPBase {
	ab.Secret = Secret
	return ab
}

func (ab *AIPBase) SetTimeoutInMillis(TimeoutInMillis uint32) *AIPBase {
	ab.TimeoutInSec = TimeoutInMillis
	return ab
}

func (ab *AIPBase) SetVersion(Version string) *AIPBase {
	ab.Version = Version
	return ab
}

func (ab *AIPBase) SetProxy(proxy string) *AIPBase {
	if _, err := url.Parse(proxy); err != nil {
		msg := fmt.Sprintf("SetProxy 无效的代理地址 %s", proxy)
		panic(msg)
	}
	ab.Proxy = proxy
	return ab
}

func (ab *AIPBase) auth(refresh bool) error {
	var object AuthObject
	if !refresh {
		if ab.ExpiresIn > NowUTCInSec()+30 {
			return nil
		}
	}
	var param = make(map[string]interface{})
	var resp []byte
	var err error
	param["grant_type"] = "client_credentials"
	param["client_id"] = ab.Key
	param["client_secret"] = ab.Secret
	if resp, err = HttpGetByForm(ACCESS_TOKEN_URL, param, 60*1000); err != nil {
		return err
	}
	if err = json.Unmarshal(resp, &object); err != nil {
		return err
	}
	ab.AccessToken = object.AccessToken
	ab.ExpiresIn = object.ExpiresIn + NowUTCInSec()
	ab.RefreshToken = object.RefreshToken
	ab.Scope = object.Scope
	ab.SessionKey = object.SessionKey
	ab.SessionSecret = object.SessionSecret
	ab.IsCloudUser = !ab.checkPermission()
	return nil
}

func (ab *AIPBase) checkPermission() bool {
	var target = BRAIN_ALL_SCOPE
	for _, v := range strings.Split(ab.Scope, " ") {
		if target == v {
			return true
		}
	}
	return false
}

func (ab *AIPBase) getParam() (param map[string]interface{}, err error) {
	var ret = make(map[string]interface{})
	if err := ab.auth(false); err != nil {
		return ret, err
	}
	if !ab.IsCloudUser {
		ret["access_token"] = ab.AccessToken
	}
	return ret, nil
}

func (ab *AIPBase) Request(url string, form map[string]interface{}, timeoutInSec uint32) (string, error) {
	var reqBody []byte
	var respBody []byte
	var err error
	var params map[string]interface{}
	if reqBody, err = json.Marshal(form); err != nil {
		return "", err
	}
	if !ab.validate(url, reqBody) {
		return "", errors.New("AIPBase Request validate is false")
	}
	var gbk = Utf8ToGBK(string(reqBody))
	if params, err = ab.getParam(); err != nil {
		return "", err
	}
	if respBody, err = HttpPostByForm(url, gbk, params, timeoutInSec); err != nil {
		return "", err
	} else {
		return GBKToUtf8(string(respBody)), nil
	}
}

func (ab *AIPBase) validate(url string, data interface{}) bool {
	return true
}
