package sdk

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"strconv"
	"time"
)

type RequestBuilder struct {
	Config     *Config
	Credential *Credential
}

func (requestBuilder *RequestBuilder) Build(api string, version string, params interface{}) (*http.Request, error) {
	var err error
	var request *http.Request
	var deviceId string
	var paramJson string
	var timestamp string

	paramJson, err = requestBuilder.getParamsJson(params)
	if err != nil {
		return nil, err
	}

	deviceId, err = requestBuilder.getDeviceId()
	if err != nil {
		return nil, err
	}

	timestamp = requestBuilder.getTimestamp()

	//init http request
	request, err = http.NewRequest(requestBuilder.Config.HttpMethod, requestBuilder.getGatewayUrl(),
		strings.NewReader(requestBuilder.getPostBody(api, version, paramJson, timestamp)))
	if err != nil {
		return nil, err
	}

	//set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set(requestBuilder.Config.AppKeyFieldName, requestBuilder.Credential.AppKey)
	request.Header.Set(requestBuilder.Config.AppMarketIdFieldName, requestBuilder.Config.AppMarketIdValue)
	request.Header.Set(requestBuilder.Config.AppVersionFieldName, requestBuilder.Config.AppVersionValue)
	request.Header.Set(requestBuilder.Config.DeviceIdFieldName, deviceId)
	request.Header.Set(requestBuilder.Config.SessionIdFieldName, requestBuilder.Credential.SessionId)
	request.Header.Set(requestBuilder.Config.SignFieldName, requestBuilder.getSign(api, version, paramJson, timestamp))

	return request, nil
}

func (requestBuilder *RequestBuilder) getDeviceId() (string, error) {
	interfaces, err :=  net.Interfaces()
	if err != nil {
		return "", err
	}
	return  interfaces[1].HardwareAddr.String(), nil
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

func (requestBuilder *RequestBuilder) getParamsJson(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
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