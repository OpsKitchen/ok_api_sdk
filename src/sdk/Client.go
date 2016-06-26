package sdk

import (
	//go builtin pkg
	"encoding/json"
	"io/ioutil"
	"net/http"

	//local pkg
	"sdk/di"
	"sdk/di/defaultimpl"
	"sdk/model"
)

var DefaultLogger di.Logger = &defaultimpl.Logger{}

type Client struct {
	HttpClient     *http.Client
	RequestBuilder *RequestBuilder
}

func NewClient() *Client {
	return &Client {
		HttpClient: http.DefaultClient,
		RequestBuilder: &RequestBuilder {
			Config: model.NewConfig(),
			Credential: &model.Credential {},
		},
	}
}

func SetDefaultLogger(logger di.Logger) {
	DefaultLogger = logger
}

func (client *Client) CallApi(api string, version string, params interface{}) (*model.ApiResult, error)  {
	var err error
	var apiResult *model.ApiResult
	var request *http.Request
	var response *http.Response

	request, err = client.RequestBuilder.Build(api, version, params)
	if err != nil {
		DefaultLogger.Fatal("Build request failed")
		return nil, err
	}

	response, err = client.HttpClient.Do(request)
	if err != nil {
		DefaultLogger.Fatal("Do http request failed")
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		DefaultLogger.Fatal("Read response body failed")
		return nil, err
	}
	DefaultLogger.Debug(string(body))

	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		DefaultLogger.Fatal("Json decode failed")
		return nil, err
	}
	return apiResult, nil
}