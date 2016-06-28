package sdk

import (
	//go builtin pkg
	"encoding/json"
	"io/ioutil"
	"net/http"

	//local pkg
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
)

var DefaultLogger logger.LoggerInterface = &logger.Logger{
	Level: logger.InfoLevel,
}

type Client struct {
	HttpClient     *http.Client
	RequestBuilder *RequestBuilder
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		RequestBuilder: &RequestBuilder{
			Config:     model.NewConfig(),
			Credential: &model.Credential{},
		},
	}
}

func SetDefaultLogger(logger logger.LoggerInterface) {
	DefaultLogger = logger
}

func (client *Client) CallApi(api string, version string, params interface{}) (*model.ApiResult, error) {
	var err error
	var apiResult *model.ApiResult
	var request *http.Request
	var response *http.Response

	request, err = client.RequestBuilder.Build(api, version, params)
	if err != nil {
		DefaultLogger.Error("Build request failed: ", err.Error())
		return nil, err
	}

	response, err = client.HttpClient.Do(request)
	if err != nil {
		DefaultLogger.Error("Do http request failed: ", err.Error())
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		DefaultLogger.Error("Read response body failed: ", err.Error())
		return nil, err
	}
	DefaultLogger.Debug("Response body: " + string(body))

	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		DefaultLogger.Error("Reponse body json decode failed: ", err.Error())
		return nil, err
	}

	apiResult.DataBytes, _ = json.Marshal(apiResult.Data)
	return apiResult, nil
}
