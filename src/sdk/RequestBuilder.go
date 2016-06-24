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
	"strings"
)

type RequestBuilder struct {
	Config     *Config
	Credential *Credential
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) Build(api string, version string, params interface{}) *http.Request {
	var req *http.Request
	var paramJson string
	var timestamp string

	paramJson = rb.getParamsJson(params)
	timestamp = rb.getTimestamp()

	//init http request
	req, err := http.NewRequest(rb.Config.HttpMethod, rb.getGatewayUrl(),
		strings.NewReader(rb.getPostBody(api, version, paramJson, timestamp)))
	if err != nil {
		panic("Create http request failed: " + err.Error())
	}

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

func (rb *RequestBuilder) getDeviceId() string {
	return "test"
	interfaces, err :=  net.Interfaces()
	if err != nil {
		panic("No net interface found: " + err.Error())
	}
	return  interfaces[0].HardwareAddr.String()
}

func (rb *RequestBuilder) getGatewayUrl() string {
	var prefix string;
	if rb.Config.DisableSSL {
		prefix = "http://"
	} else {
		prefix = "https://"
	}
	return prefix + rb.Config.GatewayHost + "/gw/json";
}

func (rb *RequestBuilder) getParamsJson(v interface{}) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		panic("param json encode failed: " + err.Error())
	}
	return string(jsonBytes)
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