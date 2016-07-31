package sdk

import (
	"encoding/json"
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
	return apiResult, nil
}
