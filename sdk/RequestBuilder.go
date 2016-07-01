package sdk

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"hash"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RequestBuilder struct {
	Config     *model.Config
	Credential *model.Credential
}

func (requestBuilder *RequestBuilder) Build(api string, version string, params interface{}) (*http.Request, error) {
	var err error
	var request *http.Request
	var deviceId, gatewayUrl, paramJson, requestBody, timestamp string

	paramJson, err = requestBuilder.getParamsJson(params)
	if err != nil {
		return nil, err
	}

	deviceId, err = requestBuilder.getDeviceId()
	if err != nil {
		return nil, err
	}

	timestamp = requestBuilder.getTimestamp()
	gatewayUrl = requestBuilder.getGatewayUrl()
	requestBody = requestBuilder.getPostBody(api, version, paramJson, timestamp)
	DefaultLogger.Debug("Gateway url:", gatewayUrl)
	DefaultLogger.Debug("Request body:", requestBody)

	//init http request
	request, err = http.NewRequest(requestBuilder.Config.HttpMethod, gatewayUrl, strings.NewReader(requestBody))
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
	DefaultLogger.Debug("Request header:", request.Header)
	return request, nil
}

func (requestBuilder *RequestBuilder) getDeviceId() (string, error) {
	var err error
	var interfaces []net.Interface
	interfaces, err = net.Interfaces()
	if err != nil {
		return "", err
	}
	return interfaces[1].HardwareAddr.String(), nil
}

func (requestBuilder *RequestBuilder) getGatewayUrl() string {
	var prefix string
	if requestBuilder.Config.DisableSSL {
		prefix = "http://"
	} else {
		prefix = "https://"
	}
	return prefix + requestBuilder.Config.GatewayHost + "/gw/json"
}

func (requestBuilder *RequestBuilder) getParamsJson(params interface{}) (string, error) {
	var err error
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(params)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (requestBuilder *RequestBuilder) getPostBody(api string, version string, paramJson string, timestamp string) string {
	return fmt.Sprintf("%s&%s&%s&%s",
		requestBuilder.Config.ApiFieldName+"="+api,
		requestBuilder.Config.VersionFieldName+"="+version,
		requestBuilder.Config.TimestampFieldName+"="+requestBuilder.getTimestamp(),
		requestBuilder.Config.ParamsFieldName+"="+paramJson)
}

func (requestBuilder *RequestBuilder) getSign(api string, version string, paramJson string, timestamp string) string {
	var stringToBeSign string = requestBuilder.Credential.Secret + api + version + paramJson + timestamp
	var hash hash.Hash = md5.New()
	io.WriteString(hash, stringToBeSign)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (requestBuilder *RequestBuilder) getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
