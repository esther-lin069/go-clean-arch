package connector

import (
	"context"
	"fmt"
	"go-clean-arch/app/websocketserver/controller"
	"go-clean-arch/app/websocketserver/route"
	_bookieRepo "go-clean-arch/domain/bookie/repository"
	_bookieUsecase "go-clean-arch/domain/bookie/usecase"
	_logger "go-clean-arch/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"gorm.io/gorm"
)

// connector ws server 進入點
func Setup(
	ctx context.Context,
	wsAuthKey string,
	m *melody.Melody,
	g *gin.Engine,
	logger _logger.Logger,
	bookieMasterDB *gorm.DB,
	bookieSlaveDB *gorm.DB,
) {
	// route group
	defaultGroup := g.Group("")

	// health check
	defaultGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	// new controller
	bookieUsecase := _bookieUsecase.NewBookieUsecase(_bookieRepo.NewBookieRepository(bookieMasterDB, bookieSlaveDB), logger)
	ctl := controller.NewController(logger, bookieUsecase)

	// 綁定 http 路徑
	BindWebsocket(wsAuthKey, m, defaultGroup, logger)

	// 綁定訊息處理方法
	BindMessageHandler(ctx, m, logger, ctl)

	// go TestTicker(ctx, m)

	return
}

// 綁定 ws server 路徑
func BindWebsocket(wsAuthKey string, m *melody.Melody, router *gin.RouterGroup, logger _logger.Logger) {
	// 設定 ws 連線路徑
	router.GET("/ws", func(c *gin.Context) {
		// 基礎的身份驗證
		key := c.GetHeader("auth-key")
		if key != wsAuthKey {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// 進入連線
		m.HandleRequest(c.Writer, c.Request)
	})
}

// 綁定 melody 訊息處理方法
func BindMessageHandler(ctx context.Context, m *melody.Melody, logger _logger.Logger, ctl *controller.Controller) {
	// 設定連線設定值，目前與套件預設值相同
	m.Config = &melody.Config{
		WriteWait:         10 * time.Second,
		PongWait:          60 * time.Second,
		PingPeriod:        (60 * time.Second * 9) / 10,
		MaxMessageSize:    512,
		MessageBufferSize: 256,
	}

	// 設定訊息接收後的處理方法
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("Received: %s \n", msg)

		route.MessageHandler(m, s, logger, ctl, msg)
	})
}

// 測試用
func TestTicker(ctx context.Context, m *melody.Melody) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("send message")

		case <-ctx.Done():
			fmt.Println("ticker done")
			return
		}
	}
}
