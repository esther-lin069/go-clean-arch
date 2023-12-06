package responder

import (
	"encoding/json"
	"strings"
	"time"
)

// 頻道名稱定義
const (
	ChannelEvent  = "channelEvent"
	ChannelMarket = "channelMarket"
)

// 定義回應類型
const (
	// 回應客戶端，原則上與請求的 Type 一致
	Pong              = "pong"
	SubscribeEvent    = "subscribeEvent"    // 訂閱賽事
	UnsubscribeEvent  = "unsubscribeEvent"  // 取消訂閱賽事
	SubscribeMarket   = "subscribeMarket"   // 訂閱盤口
	UnsubscribeMarket = "unsubscribeMarket" // 取消訂閱盤口
	GetEventByID      = "getEventByID"      // 取得賽事資訊 By ID
	GetSportList      = "getSportList"      // 取得球種列表

	// 主動推送
	UpdateEvent  = "updateEvent"  // 更新賽事
	CreateEvent  = "createEvent"  // 新增賽事
	DeleteEvent  = "deleteEvent"  // 刪除賽事
	UpdateMarket = "updateMarket" // 更新盤口
	CreateMarket = "createMarket" // 新增盤口
	DeleteMarket = "deleteMarket" // 刪除盤口
)

// 定義回應格式
type Response struct {
	Type string      `json:"type"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"` // unix time nano
}

// 錯誤回傳
func ErrorResp(msgType, errorCode string, data ...string) (errMessage []byte, err error) {
	resp := &Response{
		Type: msgType,
		Code: errorCode,
		Ts:   time.Now().UnixNano(),
	}

	if len(data) > 0 {
		resp.Data = strings.Join(data, ", ")
	}

	errMessage, err = json.Marshal(resp)

	return
}

// 成功回傳
func SuccessResp(msgType string, data interface{}) (resp []byte, err error) {
	resp, err = json.Marshal(&Response{
		Type: msgType,
		Code: "0",
		Data: data,
		Ts:   time.Now().UnixNano(),
	})

	return
}
