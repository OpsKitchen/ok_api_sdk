package sdk

import (
	"encoding/json"
	"errors"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	RequestBuilder *RequestBuilder
}

func (client *Client) CallApi(api string, version string, params interface{}) (*model.ApiResult, error) {
	var apiResult *model.ApiResult
	request, err := client.RequestBuilder.Build(api, version, params)
	if err != nil {
		return nil, err
	}

	response, err := client.HttpClient.Do(request)
	if err != nil {
		errMsg := "sdk: http communication error: " + err.Error()
		DefaultLogger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	defer response.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errMsg := "sdk: can not read response body: " + err.Error()
		DefaultLogger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	DefaultLogger.Debug("[API SDK] Response body: " + string(responseBodyBytes))
	if err := json.Unmarshal(responseBodyBytes, &apiResult); err != nil {
		errMsg := "sdk: can not parse response body, not a valid json."
		DefaultLogger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return apiResult, nil
}
