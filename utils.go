package goaip

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func NowUTCInSec() int64 {
	return time.Now().Unix()
}

func HttpPostByJSON(url string, body []byte, params map[string]interface{}, timeoutInSec uint32) ([]byte, error) {
	var response *http.Response
	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.DefaultClient
	if timeoutInSec > 0 {
		client.Timeout = time.Duration(timeoutInSec) * time.Second
	}
	q := req.URL.Query()
	if params != nil {
		for k, v := range params {
			var kind = reflect.TypeOf(v).Kind()
			if kind == reflect.Int {
				q.Add(k, strconv.Itoa(int(reflect.ValueOf(v).Int())))
			} else if kind == reflect.Uint32 {
				q.Add(k, strconv.Itoa(int(reflect.ValueOf(v).Uint())))
			} else if kind == reflect.Bool {
				q.Add(k, strconv.FormatBool(reflect.ValueOf(v).Bool()))
			} else {
				q.Add(k, v.(string))
			}
		}
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", runtime.Version())
	req.Header.Set("Connection", "keep-alive")
	if response, err = client.Do(req); err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func HttpPostByForm(url string, body []byte, params map[string]interface{}, timeoutInSec uint32) ([]byte, error) {
	var response *http.Response
	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	client := http.DefaultClient
	if timeoutInSec > 0 {
		client.Timeout = time.Duration(timeoutInSec) * time.Second
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", runtime.Version())
	req.Header.Set("Connection", "keep-alive")
	if response, err = client.Do(req); err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func HttpGetByForm(url string, form map[string]interface{}, timeoutInSec uint32) (resp []byte, e error) {
	var req *http.Request
	var response *http.Response
	var ret []byte
	var err error
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return ret, nil
	}
	q := req.URL.Query()
	if form != nil {
		for k, v := range form {
			var kind = reflect.TypeOf(v).Kind()
			if kind == reflect.Int {
				q.Add(k, strconv.Itoa(int(reflect.ValueOf(v).Int())))
			} else if kind == reflect.Uint32 {
				q.Add(k, strconv.Itoa(int(reflect.ValueOf(v).Uint())))
			} else if kind == reflect.Bool {
				q.Add(k, strconv.FormatBool(reflect.ValueOf(v).Bool()))
			} else {
				q.Add(k, v.(string))
			}
		}
	}
	req.URL.RawQuery = q.Encode()
	client := http.DefaultClient
	if timeoutInSec > 0 {
		client.Timeout = time.Duration(timeoutInSec) * time.Second
	}
	if response, err = client.Do(req); err != nil {
		return ret, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
