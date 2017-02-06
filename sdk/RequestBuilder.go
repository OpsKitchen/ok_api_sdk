package sdk

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
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
	DefaultLogger.Debug("[API SDK] Gateway url: " + gatewayUrl)
	DefaultLogger.Debug("[API SDK] Request body: " + requestBody)

	//init http request
	request, err := http.NewRequest(http.MethodPost, gatewayUrl, strings.NewReader(requestBody))
	if err != nil {
		errMsg := "sdk: can not create http request object: " + err.Error()
		DefaultLogger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	//set headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set(rb.Config.AppKeyFieldName, rb.Credential.AppKey)
	request.Header.Set(rb.Config.AppMarketIdFieldName, rb.Config.AppMarketIdValue)
	request.Header.Set(rb.Config.AppVersionFieldName, rb.Config.AppVersionValue)
	request.Header.Set(rb.Config.DeviceIdFieldName, deviceId)
	request.Header.Set(rb.Config.SessionIdFieldName, rb.Credential.SessionId)
	request.Header.Set(rb.Config.SignFieldName, rb.getSign(api, version, paramJson, timestamp))
	DefaultLogger.Debug("[API SDK] Request header:", request.Header)
	return request, nil
}

func (rb *RequestBuilder) getDeviceId() (string, error) {
	uuidFile := "/root/.ok_agent/uuid"
	_, err := os.Stat(uuidFile)
	if err != nil {
		instanceId := uuid.NewV4().String()
		err := ioutil.WriteFile(uuidFile, []byte(instanceId), 0644)
		if err != nil {
			errMsg := "sdk: failed to write uuid file [" + uuidFile + "]: " + err.Error()
			DefaultLogger.Error(errMsg)
			return "", errors.New(errMsg)
		}
		return instanceId, nil
	} else {
		instanceId, err := ioutil.ReadFile(uuidFile)
		if err != nil {
			errMsg := "sdk: failed to read uuid file [" + uuidFile + "]: " + err.Error()
			DefaultLogger.Error(errMsg)
			return "", errors.New(errMsg)
		}
		return string(instanceId), nil
	}
}

func (rb *RequestBuilder) getGatewayUrl() string {
	urlObj := url.URL{
		Path: rb.Config.GatewayPath,
	}
	if rb.Config.DisableSSL {
		urlObj.Scheme = "http"
	} else {
		urlObj.Scheme = "https"
	}
	if rb.Config.GatewayPort != 0 {//port number configured
		urlObj.Host = fmt.Sprintf("%s:%s", rb.Config.GatewayHost, strconv.Itoa(rb.Config.GatewayPort))
	} else {
		urlObj.Host = rb.Config.GatewayHost
	}
	return urlObj.String()
}

func (rb *RequestBuilder) getParamsJson(params interface{}) (string, error) {
	if params == nil {
		return "", nil
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		errMsg := "sdk: can not encode api parameter as json: " + err.Error()
		DefaultLogger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return string(jsonBytes), nil
}

func (rb *RequestBuilder) getPostBody(api string, version string, paramJson string, timestamp string) string {
	values := &url.Values{}
	values.Add(rb.Config.ApiFieldName, api)
	values.Add(rb.Config.VersionFieldName, version)
	values.Add(rb.Config.TimestampFieldName, timestamp)
	values.Add(rb.Config.ParamsFieldName, paramJson)
	return values.Encode()
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
