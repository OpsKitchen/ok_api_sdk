package sdk

import (
	"encoding/json"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	RequestBuilder *RequestBuilder
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

func (client *Client) CallApi(api string, version string, params interface{}, returnDataPointer interface{}) (*model.ApiResult, error) {
	var apiResult *model.ApiResult
	request, err := client.RequestBuilder.Build(api, version, params)
	if err != nil {
		return nil, err
	}

	response, err := client.HttpClient.Do(request)
	if err != nil {
		DefaultLogger.Error("Failed to do http communication: " + err.Error())
		return nil, err
	}

	defer response.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		DefaultLogger.Error("Failed to read response body: " + err.Error())
		return nil, err
	}
	DefaultLogger.Debug("Response body: " + string(responseBodyBytes))
	if err := json.Unmarshal(responseBodyBytes, &apiResult); err != nil {
		DefaultLogger.Error("Reponse body is not valid json.")
		return nil, err
	}

	//type casting
	if apiResult.Data != nil && returnDataPointer != nil {
		responseDataBytes, _ := json.Marshal(apiResult.Data)
		if err = json.Unmarshal(responseDataBytes, returnDataPointer); err != nil {
			DefaultLogger.Error("Failed to cast return data type")
		}
	}
	return apiResult, nil
}
