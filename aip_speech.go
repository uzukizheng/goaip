package goaip

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type AIPSpeech struct {
	AIPBase
}

func MakeAIPSpeech(appID string, key string, secret string) *AIPSpeech {
	as := new(AIPSpeech)
	as.AppID = appID
	as.Key = key
	as.Secret = secret
	return as
}

func (as *AIPSpeech) makeRequestData(url string,
	params map[string]interface{},
	data map[string]interface{}) map[string]interface{} {
	token := as.AccessToken
	if _, ok := params["cuid"]; !ok {
		data["cuid"] = MD5(token)
	}
	if url == ASR_URL {
		data["token"] = token
	}
	return data
}

func (as *AIPSpeech) ASR(stream []byte) (string, error) {
	if len(stream) == 0 {
		return "", errors.New("no stream data")
	}
	var reqData = make(map[string]interface{})
	reqData["speech"] = base64.StdEncoding.EncodeToString(stream)
	reqData["len"] = len(stream)
	reqData["channel"] = 1
	reqData["format"] = "pcm"
	reqData["rate"] = 16000
	return as.Request(ASR_URL, reqData, 10)
}

func (as *AIPSpeech) Request(url string, form map[string]interface{}, timeoutInSec uint32) (string, error) {
	var reqBody []byte
	var respBody []byte
	var err error
	var params map[string]interface{}
	if params, err = as.getParam(); err != nil {
		return "", err
	}
	params = as.makeRequestData(url, params, form)
	if reqBody, err = json.Marshal(params); err != nil {
		return "", err
	}
	if respBody, err = HttpPostByJSON(url, reqBody, nil, timeoutInSec); err != nil {
		return "", err
	} else {
		return string(respBody), nil
	}
}
