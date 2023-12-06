package websocket

import (
	"context"
	"strings"

	_logger "go-clean-arch/pkg/logger"

	"golang.org/x/net/websocket"
)

type WebsocketConn struct {
	Conn   *websocket.Conn
	JSON   *websocket.Codec
	logger _logger.Logger
	url    string
	origin string
	cancel context.CancelFunc
}

// 建立 websocket 連線
func NewWebsocketConn(ctx context.Context, logger _logger.Logger, url, origin, authKey string) (wsConn *WebsocketConn, wsCtx context.Context, err error) {
	connConfig, err := websocket.NewConfig(url, origin)
	if err != nil {
		logger.Error("建立websocket config 失敗", err)
	}
	connConfig.Header.Add("auth-key", authKey)

	ws, err := websocket.DialConfig(connConfig)
	if err != nil {
		logger.Error("建立websocket連線失敗", err)
	}

	wsConn = &WebsocketConn{
		Conn:   ws,
		JSON:   &websocket.JSON,
		url:    url,
		origin: origin,
		logger: logger,
	}

	// 設定 context
	wsCtx, wsConn.cancel = context.WithCancel(ctx)

	return
}

// 停止連線物件運作，並觸發重連
func (ws *WebsocketConn) Stop() {
	ws.logger.Error("停止連線物件運作，並觸發重連", nil)
	ws.cancel()

	ws.Conn.Close()
}

// 判斷是否屬於來自服務端的連線錯誤
func (ws *WebsocketConn) IsDialError(err error) bool {
	return strings.Contains(err.Error(), "websocket.Dial")
}
