package websocketcilent

import (
	"context"
	"go-clean-arch/app/websocketcilent/handler"
	"go-clean-arch/conn/websocket"
	"go-clean-arch/env"
	"time"

	_logger "go-clean-arch/pkg/logger"
)

func Start() {
	// 載入 logger
	logger := _logger.NewLogger()

	// 載入 env
	envConfig, err := env.NewWebsocketClientEnv(logger)
	if err != nil {
		logger.Error("載入 env 失敗", err)
		return
	}

	// TODO graceful shutdown 卡住程序
	ctx := context.Background()

	// 重連相關設定
	var (
		connTimeout time.Duration = time.Second * 5
		connCount   int
		connLimit   int = 5
	)

	// 連線服務端含重連邏輯，阻塞使總程序不會結束
	for {
		// 建立連線
		ws, wsCtx, err := websocket.NewWebsocketConn(
			ctx,
			logger,
			envConfig.Websocket.Client.URL,
			envConfig.Websocket.Client.Origin,
			envConfig.Websocket.Client.AuthKey,
		)
		if err != nil {
			// 連線失敗，重連次數 +1
			connCount++
			logger.Error("連線失敗", err)

			// 重連次數達上限，結束重連程序
			if connCount >= connLimit {
				logger.Info("重連次數達上限，結束重連程序")

				break
			}

			// 每次重連間隔為 connTimeout
			time.Sleep(connTimeout)

			continue
		}

		logger.Info("連線成功")

		// 連線成功，重置重連次數
		connCount = 0

		// 啟動連線
		handler.Setup(wsCtx, logger, ws)

		// Setup 會阻塞，直到 wsCtx.Done() 事件發生
		logger.Info("即將重連")
	}

	logger.Info("總程序結束")

}
