package sdk

import (
	//go builtin pkg
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

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

func (client *Client) CallApi(api string, version string, params interface{}, returnDataPointer interface{}) (*model.ApiResult, error) {
	var apiResult *model.ApiResult
	var byteArray []byte
	var err error
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
	byteArray, err = ioutil.ReadAll(response.Body)
	if err != nil {
		DefaultLogger.Error("Read response body failed: ", err.Error())
		return nil, err
	}
	DefaultLogger.Debug("Response body: " + string(byteArray))

	err = json.Unmarshal(byteArray, &apiResult)
	if err != nil {
		DefaultLogger.Error("Reponse body json decode failed: ", err.Error())
		return nil, err
	}

	if returnDataPointer != nil {
		var returnDataType reflect.Type = reflect.TypeOf(returnDataPointer)
		DefaultLogger.Debug("Return data type: ", returnDataType)
		byteArray, _ = json.Marshal(apiResult.Data)
		err = json.Unmarshal(byteArray, returnDataPointer)
		if err != nil {
			DefaultLogger.Error("Can not cast return data to type: ", returnDataType)
		}
	}

	return apiResult, nil
}
