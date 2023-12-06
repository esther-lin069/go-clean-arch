package logger

import _logrus "go-clean-arch/pkg/logger/logrus"

//go:generate mockgen -destination=mock/logger.go -source=logger.go

/*
僅使用以下四個等級：
Debug: 工程師開發或是除錯時使用，會跟 debug mode 連動，開啟時才會印出
Info: 系統資訊，固定會紀錄，例如服務啟動、服務結束
Warn: 不需被處理的錯誤，僅需要紀錄，供後續追蹤
Error: 需要被處理的錯誤，包含服務啟動初始化異常、對外請求錯誤

不使用：Trace、Fatal、Panic
*/
type Logger interface {
	// 開發過程或是 debug 模式會紀錄的
	Debug(s ...interface{})
	// 系統訊息，例如啟動訊息、結束訊息
	Info(message string, moreInfo ...interface{})
	// 不需被處理的警示，例如流程中要刪除快取資料，但是快取資料不存在或是發生錯誤(不影響後續操作，或是不需進行修正)
	Warn(message string, err error, moreInfo ...interface{})
	// 需要被處理的問題，但是服務不會自動停止，可能是外部資源錯誤或是 DB 請求異常
	Error(message string, err error, moreInfo ...interface{})
}

// new logger 物件，目前使用 datadog
func NewLogger() Logger {
	return _logrus.NewLogger()
}
