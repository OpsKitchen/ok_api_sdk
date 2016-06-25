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

func (requestBuilder *RequestBuilder) Build(api string, version string, params interface{}) *http.Request {
	var request *http.Request
	var paramJson string
	var timestamp string

	paramJson = requestBuilder.getParamsJson(params)
	timestamp = requestBuilder.getTimestamp()

	//init http request
	request, err := http.NewRequest(requestBuilder.Config.HttpMethod, requestBuilder.getGatewayUrl(),
		strings.NewReader(requestBuilder.getPostBody(api, version, paramJson, timestamp)))
	if err != nil {
		panic("Create http request failed: " + err.Error())
	}

	//set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set(requestBuilder.Config.AppKeyFieldName, requestBuilder.Credential.AppKey)
	request.Header.Set(requestBuilder.Config.AppMarketIdFieldName, requestBuilder.Config.AppMarketIdValue)
	request.Header.Set(requestBuilder.Config.AppVersionFieldName, requestBuilder.Config.AppVersionValue)
	request.Header.Set(requestBuilder.Config.DeviceIdFieldName, requestBuilder.getDeviceId())
	request.Header.Set(requestBuilder.Config.SessionIdFieldName, requestBuilder.Credential.SessionId)
	request.Header.Set(requestBuilder.Config.SignFieldName, requestBuilder.getSign(api, version, paramJson, timestamp))

	return request
}

func (requestBuilder *RequestBuilder) getDeviceId() string {
	interfaces, err :=  net.Interfaces()
	if err != nil {
		panic("No net interface found: " + err.Error())
	}
	return  interfaces[1].HardwareAddr.String()
}

func (requestBuilder *RequestBuilder) getGatewayUrl() string {
	var prefix string;
	if requestBuilder.Config.DisableSSL {
		prefix = "http://"
	} else {
		prefix = "https://"
	}
	return prefix + requestBuilder.Config.GatewayHost + "/gw/json";
}

func (requestBuilder *RequestBuilder) getParamsJson(v interface{}) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		panic("param json encode failed: " + err.Error())
	}
	return string(jsonBytes)
}

func (requestBuilder *RequestBuilder) getPostBody(api string, version string, paramJson string, timestamp string) string {
	return fmt.Sprintf("%s&%s&%s&%s",
		requestBuilder.Config.ApiFieldName + "=" + api,
		requestBuilder.Config.VersionFieldName + "=" + version,
		requestBuilder.Config.TimestampFieldName + "=" + requestBuilder.getTimestamp(),
		requestBuilder.Config.ParamsFieldName + "=" + paramJson)
}

func (requestBuilder *RequestBuilder) getSign(api string, version string, paramJson string, timestamp string) string {
	var stringToBeSign string = requestBuilder.Credential.Secret + api + version + paramJson + timestamp

	hash := md5.New()
	io.WriteString(hash, stringToBeSign)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (requestBuilder *RequestBuilder) getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}