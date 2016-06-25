package sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	Logger         *log.Logger
	RequestBuilder *RequestBuilder
}

func NewClient() *Client {
	return &Client {
		HttpClient: http.DefaultClient,
		RequestBuilder: &RequestBuilder {
			Config: NewConfig(),
			Credential: &Credential {},
		},
	}
}

func (client *Client) CallApi(api string, version string, params interface{}) (*ApiResult, error)  {
	var err error
	var apiResult *ApiResult
	var request *http.Request
	var response *http.Response

	request, err = client.RequestBuilder.Build(api, version, params)
	if err != nil {
		//client.Logger.Fatal("Build request failed")
		return nil, err
	}

	response, err = client.HttpClient.Do(request)
	if err != nil {
		//client.Logger.Fatal("Do http request failed")
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//client.Logger.Fatal("Read response body failed")
		return nil, err
	}

	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		//client.Logger.Fatal("Json decode failed")
		return nil, err
	}
	return apiResult, nil
}