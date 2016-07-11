package sdk

import (
	"encoding/json"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"io/ioutil"
	"net/http"
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

func (client *Client) CallApi(api string, version string, params interface{}, returnDataPointer interface{}) (*model.ApiResult, error) {
	var apiResult *model.ApiResult
	var err error
	var request *http.Request
	var response *http.Response
	var responseBodyBytes []byte
	var responseBodyString string

	request, err = client.RequestBuilder.Build(api, version, params)
	if err != nil {
		return nil, err
	}

	response, err = client.HttpClient.Do(request)
	if err != nil {
		DefaultLogger.Error("Failed to do http communication: " + err.Error())
		return nil, err
	}

	defer response.Body.Close()
	responseBodyBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		DefaultLogger.Error("Failed to read response body: " + err.Error())
		return nil, err
	}
	responseBodyString = string(responseBodyBytes)
	DefaultLogger.Debug("Response body: " + responseBodyString)

	err = json.Unmarshal(responseBodyBytes, &apiResult)
	if err != nil {
		DefaultLogger.Error("Reponse body is not valid json.")
		return nil, err
	}

	if returnDataPointer != nil {
		responseBodyBytes, _ = json.Marshal(apiResult.Data)
		err = json.Unmarshal(responseBodyBytes, returnDataPointer)
		if err != nil {
			DefaultLogger.Error("Failed to cast return data type")
		}
	}

	return apiResult, nil
}
