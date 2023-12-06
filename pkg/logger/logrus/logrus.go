package logrus

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// 物件格式
type LogrusLogger struct {
	log *logrus.Logger
}

// NewLogger 物件
func NewLogger() (l *LogrusLogger) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(logrus.StandardLogger().Out)
	log.SetLevel(logrus.DebugLevel)

	l = &LogrusLogger{log: log}

	return l
}

// 開發或是除錯用訊息
func (l *LogrusLogger) Debug(s ...interface{}) {
	l.log.Debug(s...)
}

// 系統訊息
func (l *LogrusLogger) Info(message string, moreInfo ...interface{}) {
	moreInfoFields := getMoreInfo(moreInfo...)

	l.log.WithFields(moreInfoFields).Info(message)
}

// 錯誤訊息，moreInfo 必須為空或 map[string]interface{}
func (l *LogrusLogger) Error(message string, err error, moreInfo ...interface{}) {
	moreInfoFields := getMoreInfo(moreInfo...)
	moreInfoFields["err"] = err.Error()

	l.log.WithFields(moreInfoFields).Error(message)
}

// 不需被處理的警示
func (l *LogrusLogger) Warn(message string, err error, moreInfo ...interface{}) {
	moreInfoFields := getMoreInfo(moreInfo...)
	moreInfoFields["err"] = err.Error()

	l.log.WithFields(moreInfoFields).Warn(message)
}

// 組成更多資訊
func getMoreInfo(moreInfo ...interface{}) (moreInfoFields map[string]interface{}) {
	moreInfoFields = map[string]interface{}{}

	if len(moreInfo) > 0 {
		v, ok := moreInfo[0].(map[string]interface{})
		if ok {
			moreInfoFields = v
		}
	}

	moreInfoFields["caller"] = getCallerFunctionName()

	return
}

// 取得 func 呼叫者名稱
func getCallerFunctionName() (caller string) {
	pc, _, _, _ := runtime.Caller(3) // 上層 func 的名稱

	caller = runtime.FuncForPC(pc).Name()

	return
}
