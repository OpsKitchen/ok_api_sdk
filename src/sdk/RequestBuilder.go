package sdk

import (
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"net"
	"crypto/md5"
	"io"
	"fmt"
)

type RequestBuilder struct {
	Config     *Config
	Credential *Credential
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) Build(api string, version string, params interface{}) *http.Request {
	var req http.Request
	var paramJson string
	paramJson, err := json.Marshal(params)
	if err != nil {
		panic("param json encode failed: " + err.Error())
	}
	var timestamp string = rb.getTimestamp()

	//init http request
	req = http.NewRequest(rb.Config.HttpMethod, rb.getGatewayUrl(),
		rb.getPostBody(api, version, paramJson, timestamp))

	//set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(rb.Config.AppKeyFieldName, rb.Credential.AppKey)
	req.Header.Set(rb.Config.AppMarketIdFieldName, rb.Config.AppMarketIdValue)
	req.Header.Set(rb.Config.AppVersionFieldName, rb.Config.AppVersionValue)
	req.Header.Set(rb.Config.DeviceIdFieldName, rb.getDeviceId())
	req.Header.Set(rb.Config.SessionIdFieldName, rb.Credential.SessionId)
	req.Header.Set(rb.Config.SignFieldName, rb.getSign(api, version, paramJson, timestamp))

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

func (rb *RequestBuilder) getDeviceId() string {
	interfaces, err :=  net.Interfaces()
	if err != nil {
		panic("No net interface found: " + err.Error())
	}
	for _, netInterface := range interfaces {
		return  netInterface.HardwareAddr
	}
}

func (rb *RequestBuilder) getGatewayUrl() string {
	var prefix string;
	if Config.DisableSSL {
		prefix = "http://"
	} else {
		prefix = "https://"
	}

	return prefix + Config.GatewayHost + "/gw/json";
}

func (rb *RequestBuilder) getParamsJson(v interface{}) string {
	return json.Marshal(v)
}

func (rb *RequestBuilder) getPostBody(api string, version string, paramJson string, timestamp string) string {
	return fmt.Sprintf("%s&%s&%s&%s",
		rb.Config.ApiFieldName + "=" + api,
		rb.Config.VersionFieldName + "=" + version,
		rb.Config.TimestampFieldName + "=" + rb.getTimestamp(),
		rb.Config.ParamsFieldName + "=" + paramJson)
}

func (rb *RequestBuilder) getSign(api string, version string, paramJson string, timestamp string) string {
	var stringToBeSign string = rb.Credential.Secret + api + version + paramJson + timestamp

	hash := md5.New()
	io.WriteString(hash, stringToBeSign)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (rb *RequestBuilder) getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}