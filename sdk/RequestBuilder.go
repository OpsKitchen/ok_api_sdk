package sdk

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RequestBuilder struct {
	Config     *model.Config
	Credential *model.Credential
}

func (rb *RequestBuilder) Build(api string, version string, params interface{}) (*http.Request, error) {
	paramJson, err := rb.getParamsJson(params)
	if err != nil {
		return nil, err
	}
	deviceId, err := rb.getDeviceId()
	if err != nil {
		return nil, err
	}

	timestamp := rb.getTimestamp()
	gatewayUrl := rb.getGatewayUrl()
	requestBody := rb.getPostBody(api, version, paramJson, timestamp)
	DefaultLogger.Debug("Gateway url: " + gatewayUrl)
	DefaultLogger.Debug("Request body: " + requestBody)

	//init http request
	request, err := http.NewRequest(rb.Config.HttpMethod, gatewayUrl, strings.NewReader(requestBody));
	if err != nil {
		DefaultLogger.Error("Failed to create http request object: " + err.Error())
		return nil, err
	}

	//set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set(rb.Config.AppKeyFieldName, rb.Credential.AppKey)
	request.Header.Set(rb.Config.AppMarketIdFieldName, rb.Config.AppMarketIdValue)
	request.Header.Set(rb.Config.AppVersionFieldName, rb.Config.AppVersionValue)
	request.Header.Set(rb.Config.DeviceIdFieldName, deviceId)
	request.Header.Set(rb.Config.SessionIdFieldName, rb.Credential.SessionId)
	request.Header.Set(rb.Config.SignFieldName, rb.getSign(api, version, paramJson, timestamp))
	DefaultLogger.Debug("Request header:", request.Header)
	return request, nil
}

func (rb *RequestBuilder) getDeviceId() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		DefaultLogger.Error("Failed to get net interface list: " + err.Error())
		return "", err
	}
	if len(interfaces) < 2 {
		errMsg := "Amount of net interfaces is less than 2"
		DefaultLogger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return interfaces[1].HardwareAddr.String(), nil
}

func (rb *RequestBuilder) getGatewayUrl() string {
	urlObj := url.URL{
		Host: rb.Config.GatewayHost,
		Path: rb.Config.GatewayPath,
	}
	if rb.Config.DisableSSL {
		urlObj.Scheme = "http"
	} else {
		urlObj.Scheme = "https"
	}
	return urlObj.String()
}

func (rb *RequestBuilder) getParamsJson(params interface{}) (string, error) {
	jsonBytes, err := json.Marshal(params)
	if  err != nil {
		DefaultLogger.Error("Api parameter can not encode as json. Json encoder said: " + err.Error())
		return "", err
	}
	return string(jsonBytes), nil
}

func (rb *RequestBuilder) getPostBody(api string, version string, paramJson string, timestamp string) string {
	str := fmt.Sprintf("%s&%s&%s", rb.Config.ApiFieldName+"="+api, rb.Config.VersionFieldName+"="+version,
		rb.Config.TimestampFieldName+"="+rb.getTimestamp())
	if paramJson != "null" {
		str += "&" + rb.Config.ParamsFieldName+"="+paramJson
	}
	return str
}

func (rb *RequestBuilder) getSign(api string, version string, paramJson string, timestamp string) string {
	hashObj := md5.New()
	stringToBeSign := rb.Credential.Secret + api + version + paramJson + timestamp
	io.WriteString(hashObj, stringToBeSign)
	return fmt.Sprintf("%x", hashObj.Sum(nil))
}

func (rb *RequestBuilder) getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
