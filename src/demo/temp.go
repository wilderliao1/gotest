package demo

import (
	"fmt"
	"time"
	"encoding/base64"
	"net/url"
	"errors"
	"strconv"
)

// GenClassid 生成classid
func GenClassid(sdkappid int, avRoomId int) string {
	return genCkvKey("classid", sdkappid, avRoomId)
}

// GenTaskid 生成taskid
func GenTaskid(sdkappid int, avRoomId int) string {
	keyStr := fmt.Sprintf("flag=taskid&sdkappid=%d&avroomid=%d&time=%d", sdkappid, avRoomId,time.Now().Unix())
	key := base64.StdEncoding.EncodeToString([]byte(keyStr))
	return key
}

// "课堂正在进行时的录制回调信息"保存在一个buff中，生成这个buff保存时所需的key
func GenClassCBID(sdkappid int, avRoomId int) string {
	return genCkvKey("class_callback", sdkappid, avRoomId)
}

//// 生成"课堂结束，录制工具正在进行时的录制回调信息"保存时所需的key
//func GenRecordCBID(sdkappid int, avRoomId int) string {
//	return genCkvKey("record_callback", sdkappid, avRoomId)
//}

// "可以发起录制的课堂信息"保存在一个buff中，生成这个buff保存时所需的key
func GenRecordBuffid() string {
	return "will_record_buffid"
}

// 生成"管理员账号信息"保存时所需的key
func GenAdminInfoid(sdkappid int) string {
	keyStr := fmt.Sprintf("flag=admin&sdkappid=%d", sdkappid)
	key := base64.StdEncoding.EncodeToString([]byte(keyStr))
	return key
}

func genCkvKey(flag string, sdkappid int, avRoomId int) string {
	keyStr := fmt.Sprintf("flag=%s&sdkappid=%d&avroomid=%d", flag, sdkappid, avRoomId)
	key := base64.StdEncoding.EncodeToString([]byte(keyStr))
	return key
}

func TaskidToClassid(taskid string) string {
	keyMap, err := ParseCkvKey(taskid)
	if err != nil {
		return ""
	}
	if keyMap["flag"] != "taskid" {
		return ""
	}
	return GenClassid(keyMap["sdkappid"].(int), keyMap["avroomid"].(int))
}

func GetAvRoomIdFromTaskid(taskid string) (avRoomId int) {
	keyMap, err := ParseCkvKey(taskid)
	if err != nil {
		return -1
	}
	if keyMap["flag"] != "taskid" {
		return -1
	}
	return keyMap["avroomid"].(int)
}

func ParseCkvKey(key string) (map[string]interface{}, error) {
	keyByte, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	keyStr := string(keyByte)
	keyMap, err := url.ParseQuery(keyStr)
	if err != nil {
		return nil, err
	}
	if len(keyMap["flag"]) == 0 || len(keyMap["sdkappid"]) == 0 || len(keyMap["avroomid"]) == 0 {
		return nil, errors.New("parse key fail1")
	}

	result := make(map[string]interface{})
	result["flag"] = keyMap["flag"][0]

	sdkappid, err := strconv.Atoi(keyMap["sdkappid"][0])
	if err != nil {
		return nil, errors.New("parse key fail2")
	}
	result["sdkappid"] = sdkappid

	avRoomId, err := strconv.Atoi(keyMap["avroomid"][0])
	if err != nil {
		return nil, errors.New("parse key fail3")
	}
	result["avroomid"] = avRoomId

	return result, nil
}
