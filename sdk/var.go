package sdk

import "github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"

var DefaultLogger logger.LoggerInterface = &logger.Logger{
	Level: logger.InfoLevel,
}
