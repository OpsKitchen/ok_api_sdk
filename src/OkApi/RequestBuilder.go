package OkApi

import (
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

type RequestBuilder struct {
	Config     *Config
	Credential *Credential
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) Build(api string, version string, params interface{}) *http.Request {
	var query *string = rb.Config.ApiFieldName + "=" + api
	+ "&" + rb.Config.VersionFieldName + "=" + version
	+ "&" + rb.Config.TimestampFieldName + "=" + rb.getTime()
	+ "&" + rb.Config.ParamsFieldName + json.Marshal(params)

	req, err = http.NewRequest(rb.Config.HttpMethod, rb.getGatewayUrl(), query)
	req.Header.set("Content-Type", config.BaseConfig.CONTENT_TYPE)

	return req
}

func (rb *RequestBuilder) SetCredential(credential Credential) *RequestBuilder {
	rb.Credential = &credential
	return rb
}

func (rb *RequestBuilder) SetConfig(config Config) *RequestBuilder {
	rb.Config = &config
	return rb
}


func (rb *RequestBuilder) getGatewayUrl() *string {
	var prefix string;
	if Config.DisableSSL {
		prefix = "http://"
	} else {
		prefix = "https://"
	}

	return prefix + Config.GatewayHost + "/gw/json";
}

func (rb *RequestBuilder) getParamsJson(v interface{}) *string {
	return json.Marshal(v)
}

func (rb *RequestBuilder) getPostBody()

func (rb *RequestBuilder) getTime() *string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}