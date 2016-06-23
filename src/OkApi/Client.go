package OkApi

import (
	"net/http"
	"log"
)

type Client struct {
	Config     *Config
	Credential *Credential
	HttpClient *http.Client
	Logger	   *log.Logger
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.DefaultClient,
	}
}

func (req *Client) SetCredential(credential Credential) *Client {
	req.Credential = &credential
	return req
}

func (req *Client) SetConfig(config Config) *Client {
	req.Config = &config
	return req
}

func (req *Client) SetHttpClient(httpClient *http.Client) *Client {
	req.HttpClient = &httpClient
	return req
}

func (req *Client) SetLogger(config *log.Logger) *Client {
	req.Logger = &config
	return req
}

func (req Client) Request(api string, version string, params interface{})  {

}