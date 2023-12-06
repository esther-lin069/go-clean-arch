package route

import (
	"go-clean-arch/app/websocketserver/controller"
	_logger "go-clean-arch/pkg/logger"

	"encoding/json"

	"github.com/olahol/melody"
)

// 定義客戶端請求類型
const (
	Ping              = "ping"
	SubscribeEvent    = "subscribeEvent"
	UnsubscribeEvent  = "unsubscribeEvent"
	SubscribeMarket   = "subscribeMarket"
	UnsubscribeMarket = "unsubscribeMarket"

	// 針對性取資料請求
	GetEventByID = "getEventByID"
	GetSportList = "getSportList"

	// 新增訊息事件請在此定義
)

// 定義客戶端請求格式
type request struct {
	Type     string          `json:"type"`
	Params   json.RawMessage `json:"params,omitempty"`
	UnixTime int64           `json:"time"`
}

/* 說明：設定頻道方式視需求而定，可自由運用 melody.Session 提供的 Keys 欄位放入任意值
 * 例如："channel": "event"、"channel_event": true、"channels": ["event", "market"] 等等
 */

// 訊息事件處理
func MessageHandler(
	m *melody.Melody,
	s *melody.Session,
	logger _logger.Logger,
	c *controller.Controller,
	msg []byte,
) (err error) {
	// 解析訊息事件 Event type
	req, err := requestParse(msg)
	if err != nil {
		logger.Error("解析事件錯誤, err: ", err)
		return
	}

	// 有需求的話，context with 統一timeout 做在這裡

	// 依照 Event type 定義處理邏輯
	switch req.Type {
	// ping
	case Ping:
		c.Pong(s)

	// 訂閱賽事頻道
	case SubscribeEvent:
		c.SubscribeEvent(s)

	// 取消訂閱賽事頻道
	case UnsubscribeEvent:
		c.UnsubscribeEvent(s)

	// 訂閱盤口頻道
	case SubscribeMarket:
		c.SubscribeMarket(s)

	// 取消訂閱盤口頻道
	case UnsubscribeMarket:
		c.UnsubscribeMarket(s)

	// 取得球種列表
	case GetSportList:
		c.GetSportList(s)

	// 如果有新增訊息事件，可在此處新增對應 case 處理邏輯

	// 其他
	default:
		// 廣播給所有 session
		_ = c.Broadcast(m, msg)
	}

	return

}

// 服務端解析請求
func requestParse(msg []byte) (request *request, err error) {
	err = json.Unmarshal(msg, &request)
	return
}
