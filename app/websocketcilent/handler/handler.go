package handler

import (
	"context"
	"fmt"
	"go-clean-arch/app/websocketcilent/route"
	"go-clean-arch/conn/websocket"
	_logger "go-clean-arch/pkg/logger"
	"sync"
	"time"
)

func Setup(ctx context.Context, logger _logger.Logger, ws *websocket.WebsocketConn) (wg *sync.WaitGroup) {
	// ping
	WritePing(logger, ws)

	wg = new(sync.WaitGroup)
	wg.Add(1)

	go Read(ctx, logger, ws, wg)

	wg.Wait()
	return
}

// 讀取訊息
func Read(ctx context.Context, logger _logger.Logger, ws *websocket.WebsocketConn, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		msg = make([]byte, 1024)
		n   int
		err error
	)

	for {
		select {
		// 接收 ws 結束訊號
		case <-ctx.Done():
			return

		// read message
		default:
			if n, err = ws.Conn.Read(msg); err != nil {
				logger.Error("WS Read failed", err)
				ws.Stop()
			}

			if len(msg[:n]) != 0 {
				logger.Info("Read Received", map[string]interface{}{
					"message": string(msg[:n]),
				})

				route.MessageHandler(ws, msg[:n])
			}
		}
	}

}

// 定義請求格式
type Request struct {
	Type     string      `json:"type"`
	Params   interface{} `json:"params,omitempty"`
	UnixTime int64       `json:"time"`
}

// 寫入 ping
func WritePing(logger _logger.Logger, ws *websocket.WebsocketConn) (err error) {
	req := &Request{
		Type: "ping",
	}

	// 發送訊息
	err = sendMessageJSON(logger, ws, req)
	return
}

// 發送 JSON 訊息
func sendMessageJSON(logger _logger.Logger, ws *websocket.WebsocketConn, message *Request) (err error) {
	// 設定時間戳記
	message.UnixTime = time.Now().Unix()

	// 使用 websocket 提供之方法發送 JSON 訊息
	err = ws.JSON.Send(ws.Conn, message)
	if err != nil {
		logger.Error("sendMessageJSON failed", err, map[string]interface{}{
			"message": message,
		})
		return
	}

	fmt.Printf("SendJSON: %v \n", message)

	return
}
