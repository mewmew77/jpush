package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	JPushDevicesURL = "https://device.jpush.cn/v3/devices"
)

const (
	JPushHeaderCharset     = "UTF-8"
	JPushHeaderContentType = "application/json"
)

const (
	JPushURL     = "https://api.jpush.cn/v3/push"
	BASE64_TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type JPushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *JPushClient {
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &JPushClient{secret, appKey, auth, JPushURL}
	return pusher
}

func JPushPost(url string, header map[string]string, params interface{}) (*http.Response, error) {
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramsBytes))
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("fail")
	}
	fmt.Println(string(r))
	return resp, nil
}

func (c *JPushClient) JPushSetDevicesAlias(userIdAlias string, registrationId string) error {

	url := strings.Join([]string{JPushDevicesURL, registrationId}, "/")
	header := map[string]string{
		"Authorization": c.AuthCode,
		"Charset":       JPushHeaderCharset,
		"Content-Type":  JPushHeaderContentType,
	}
	params := map[string]string{
		"alias": userIdAlias,
	}
	resp, err := JPushPost(url, header, params)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("设置别名失败")
	}

	return nil
}

func (c *JPushClient) JPushSendNotice(alias []string, message interface{}) error {
	//推送平台
	var pf Platform
	pf.Add(ANDROID)

	//推送目标设备
	var ad Audience
	//ad.SetAlias(alias)
	ad.SetID([]string{"120c83f7613a041de4d"})

	//推送内容
	var notice Notice
	notice.SetAlert("通知内容")

	//todo 自定义消息结构
	extras := map[string]interface{}{
		"id":   2,
		"name": "testdevice1",
	}
	notice.SetAndroidNotice(&AndroidNotice{Alert: "通知内容", Extras: extras})

	//var msg Message
	//msg.Title = "自定义消息"
	//msg.Content = "app自行处理自定义消息，不展示在通知中"
	//msg.Extras = map[string]interface{}{
	//	"device_id": 1,
	//}

	payload := NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)
	//payload.SetMessage(&msg)
	header := map[string]string{
		"Authorization": c.AuthCode,
		"Charset":       JPushHeaderCharset,
		"Content-Type":  JPushHeaderContentType,
	}
	resp, err := JPushPost(c.BaseUrl, header, payload)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("推送消息失败")
	}
	return nil
}
