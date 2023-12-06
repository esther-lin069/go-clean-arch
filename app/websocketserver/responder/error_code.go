package responder

// 定義錯誤代碼
var codeMap = map[string]string{
	"not_found":       "2",
	"params_error":    "3",
	"bookie_db_error": "4",
}

// 取得錯誤代碼
func GetErrorCode(message string) (code string) {
	code, exist := codeMap[message]

	if !exist {
		code = codeMap["not_found"]
	}

	return
}
