package websocketserver

import (
	"context"
	"fmt"
	"go-clean-arch/app/websocketserver/connector"
	"go-clean-arch/conn/mysql"
	"go-clean-arch/env"
	_logger "go-clean-arch/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"gorm.io/gorm"
)

func Start() {
	// logger
	logger := _logger.NewLogger()

	// 載入 env
	envConfig, err := env.NewWebsocketServerEnv(logger)
	if err != nil {
		logger.Error("載入 env 失敗", err)
		return
	}

	// new melody
	m := melody.New()

	// new gin
	g := gin.Default()
	srv := &http.Server{
		Addr:    envConfig.Domain + envConfig.Port,
		Handler: g,
	}

	// ctx 監聽外部停止訊號
	ctx := withContext(srv, func() {
		// 定義服務關閉前的動作
		m.Close()
		logger.Info("m.Close(): melody 已關閉")

	})

	// 建立 DB 連線
	dbMaster, dbSlave, err := newMySQL(
		envConfig.DB.Mysql.Develop.Database,
		envConfig.DB.Mysql.Develop.Master,
		envConfig.DB.Mysql.Develop.Slave,
	)
	if err != nil {
		logger.Error("建立 bookie 資料庫連線錯誤，停止服務", err)
		return
	}

	// 啟動 ws server
	connector.Setup(ctx, "some-auth-key", m, g, logger, dbMaster, dbSlave)

	// 啟動 gin
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("srv.ListenAndServe()", err)
	}

}

// 監聽停止訊號
func withContext(srv *http.Server, fn func()) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 1)

		// 監聽系統關閉訊號
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		defer func() {
			signal.Stop(quit)

			if r := recover(); r != nil {
				fmt.Println("recovered from ", r)
			}
		}()

		// 堵塞直到監聽到訊號
		<-quit

		fmt.Println("shutting down server...")
		cancel()

		// 執行傳入的關閉動作
		fn()

		fmt.Println("shutting down http...")

		// http server shutdown，因為 ListenAndServe 會堵塞服務所以需最後執行
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("http server shutdown error:", err)
		}

		return

	}()

	return ctx
}

// 建立 MySQL 連線
func newMySQL(database string, masterEnv env.MysqlConfig, slaveEnv env.MysqlConfig) (masterDB *gorm.DB, slaveDB *gorm.DB, err error) {
	masterDB, err = mysql.NewMySQLConn(
		masterEnv.User,
		masterEnv.Password,
		masterEnv.Host,
		masterEnv.Port,
		database,
		masterEnv.Idle,
		masterEnv.MaxConn,
		masterEnv.LifeTime,
		masterEnv.Debug,
	)
	if err != nil {
		return
	}

	slaveDB, err = mysql.NewMySQLConn(
		slaveEnv.User,
		slaveEnv.Password,
		slaveEnv.Host,
		slaveEnv.Port,
		database,
		slaveEnv.Idle,
		slaveEnv.MaxConn,
		slaveEnv.LifeTime,
		slaveEnv.Debug,
	)
	if err != nil {
		return
	}

	return
}
