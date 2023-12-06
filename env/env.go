package env

import (
	_logger "go-clean-arch/pkg/logger"
	"strings"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type WebsocketClientEnv struct {
	Env       string `validate:"required"`
	Name      string `validate:"required"`
	Domain    string
	DB        DB              `validate:"required"`
	Websocket WebsocketClient `validate:"required"`
}

type WebsocketServerEnv struct {
	Env    string `validate:"required"`
	Port   string `validate:"required"`
	Name   string `validate:"required"`
	Domain string
	DB     DB `validate:"required"`
}

type DB struct {
	Mysql struct {
		// develop
		Develop struct {
			Database string
			Master   MysqlConfig
			Slave    MysqlConfig
		}
	}
}

type MysqlConfig struct {
	Host     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Port     string `validate:"required"`
	LifeTime int    `validate:"required"`
	MaxConn  int    `validate:"required"`
	Idle     int    `validate:"required"`
	Debug    bool
}

type WebsocketClient struct {
	Client struct {
		Origin  string
		URL     string
		AuthKey string
	}
}

func NewWebsocketClientEnv(logger _logger.Logger) (env *WebsocketClientEnv, err error) {
	env = &WebsocketClientEnv{}

	// env 檔案路徑
	path := "./env"
	configName := "websocketcilent"
	prefix := configName

	// 讀取設定
	viper, err := loadYaml(path, configName, prefix)
	if err != nil {
		logger.Error("讀取 websocketcilent.yaml 發生錯誤", err)
		return
	}

	// 解析設定
	err = viper.Unmarshal(env)
	if err != nil {
		logger.Error("解析 websocketcilent.yaml 發生錯誤", err)
		return
	}

	// 驗證環境變數是否有符合結構
	err = validator.New().Struct(env)
	if err != nil {
		logger.Error("環境設定驗證錯誤", err)
		return
	}

	logger.Info("讀取 env 完成", map[string]interface{}{
		"env": env,
	})

	return
}

func NewWebsocketServerEnv(logger _logger.Logger) (env *WebsocketServerEnv, err error) {
	env = &WebsocketServerEnv{}

	// env 檔案路徑
	path := "./env"
	configName := "websocketserver"
	prefix := configName

	// 讀取設定
	viper, err := loadYaml(path, configName, prefix)
	if err != nil {
		logger.Error("讀取 websocketserver.yaml 發生錯誤", err)
		return
	}

	// 解析設定
	err = viper.Unmarshal(env)
	if err != nil {
		logger.Error("解析 websocketserver.yaml 發生錯誤", err)
		return
	}

	// 驗證環境變數是否有符合結構
	err = validator.New().Struct(env)
	if err != nil {
		logger.Error("環境設定驗證錯誤", err)
		return
	}

	logger.Info("讀取 env 完成", map[string]interface{}{
		"env": env,
	})

	return
}

// 讀取 yaml 檔案
func loadYaml(path string, configName string, perfix string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	v.SetConfigName(configName)
	v.SetEnvPrefix(perfix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err = v.ReadInConfig()

	return
}
