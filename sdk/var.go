package sdk

import (
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"net/http"
)

var DefaultLogger logger.LoggerInterface = &logger.Logger{
	Level: logger.InfoLevel,
}

func NewClient() *Client {
	client := &Client{
		HttpClient: http.DefaultClient,
		RequestBuilder: &RequestBuilder{
			Config:     &model.Config{},
			Credential: &model.Credential{},
		},
	}
	client.RequestBuilder.Config.SetDefaultOption()
	return client
}

func SetDefaultLogger(logger logger.LoggerInterface) {
	DefaultLogger = logger
}
