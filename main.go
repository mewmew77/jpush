package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	appKey         = "xxx"
	secret         = "xxxxx"
	registrationId = "xxxxxx"
)

//func main() {
//
//	//Platform
//	var pf jpushclient.Platform
//	pf.Add(jpushclient.ANDROID)
//	//pf.Add(jpushclient.IOS)
//	//pf.Add(jpushclient.WINPHONE)
//	//pf.All()
//
//	//Audience
//	var ad jpushclient.Audience
//	s := []string{registrationId}
//	//ad.SetTag(s)
//	//ad.SetAlias(s)
//	ad.SetID(s)
//	//ad.All()
//
//	//Notice
//	var notice jpushclient.Notice
//	notice.SetAlert("你好啊")
//	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "你好啊"})
//	//notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
//	//notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})
//
//	var msg jpushclient.Message
//	msg.Title = "你好啊"
//	msg.Content = "你好啊"
//
//	payload := jpushclient.NewPushPayLoad()
//	payload.SetPlatform(&pf)
//	payload.SetAudience(&ad)
//	//payload.SetMessage(&msg)
//	payload.SetNotice(&notice)
//
//	bytes, _ := payload.ToBytes()
//	fmt.Printf("%s\r\n", string(bytes))
//
//	//push
//	c := jpushclient.NewPushClient(secret, appKey)
//	str, err := c.Send(bytes)
//	if err != nil {
//		fmt.Printf("err:%s", err.Error())
//	} else {
//		fmt.Printf("ok:%s", str)
//	}
//
//}

func main() {
	c := NewPushClient(secret, appKey)
	userId := 1
	alias := HashString(string(userId))
	fmt.Println(alias)

	//设置别名，可以绑定用户，指定推送
	//err := c.JPushSetDevicesAlias("c4ca4238a0b923820dcc509a6f75849b", registrationId)
	//if err != nil {
	//	fmt.Println("设置别名失败")
	//	return
	//}

	message := "信息通知内容"
	if err := c.JPushSendNotice([]string{alias}, message); err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}

	//退出登录后将别名置为空
	//err = c.JPushSetDevicesAlias("", registrationId)
	//if err != nil {
	//	fmt.Println("置空别名失败")
	//	return
	//}
}

func HashString(a string) string {
	h := md5.New()
	h.Write([]byte(registrationId + a))
	data := h.Sum(nil)
	return hex.EncodeToString(data)
}
