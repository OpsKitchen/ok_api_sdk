package sdk

type Config struct {
	//platform address
	DisableSSL  bool
	GatewayHost string
	HttpMethod  string

	//System parameter name in HTTP header
	AppKeyFieldName      string
	AppVersionFieldName  string
	AppMarketIdFieldName string
	DeviceIdFieldName    string
	SessionIdFieldName   string
	SignFieldName        string

	//System parameter name in HTTP body
	ApiFieldName       string
	ParamsFieldName    string
	TimestampFieldName string
	VersionFieldName   string

	//System parameter value
	AppVersionValue  string
	AppMarketIdValue string
}

func NewConfig() *Config {
	return &Config {
		HttpMethod: "POST",

		//System parameter name in HTTP header
		AppKeyFieldName:      "OA-App-Key",
		AppVersionFieldName:  "OA-App-Version",
		AppMarketIdFieldName: "OA-App-Market-ID",
		DeviceIdFieldName:    "OA-Device-Id",
		SessionIdFieldName:   "OA-Session-Id",
		SignFieldName:        "OA-Sign",

		//System parameter name in HTTP body
		ApiFieldName:       "api",
		ParamsFieldName:    "params",
		TimestampFieldName: "timestamp",
		VersionFieldName:   "version",
	}
}

func (config *Config) SetDisableSSL(disableSSL bool) *Config {
	config.DisableSSL = disableSSL
	return config
}

func (config *Config) SetGatewayHost(gatewayHost string) *Config {
	config.GatewayHost = gatewayHost
	return config
}

func (config *Config) SetHttpMethod(httpMethod string) *Config {
	config.HttpMethod = httpMethod
	return config
}

func (config *Config) SetAppKeyFieldName(appKeyFieldName string) *Config {
	config.AppKeyFieldName = appKeyFieldName
	return config
}

func (config *Config) SetAppVersionFieldName(appVersionFieldName string) *Config {
	config.AppVersionFieldName = appVersionFieldName
	return config
}

func (config *Config) SetAppMarketIdFieldName(appMarketIdFieldName string) *Config {
	config.AppMarketIdFieldName = appMarketIdFieldName
	return config
}

func (config *Config) SetDeviceIdFieldName(deviceIdFieldName string) *Config {
	config.DeviceIdFieldName = deviceIdFieldName
	return config
}

func (config *Config) SetSessionIdFieldName(sessionIdFieldName string) *Config {
	config.SessionIdFieldName = sessionIdFieldName
	return config
}

func (config *Config) SetSignFieldName(signFieldName string) *Config {
	config.SignFieldName = signFieldName
	return config
}

func (config *Config) SetApiFieldName(apiFieldName string) *Config {
	config.ApiFieldName = apiFieldName
	return config
}

func (config *Config) SetParamsFieldName(paramsFieldName string) *Config {
	config.ParamsFieldName = paramsFieldName
	return config
}

func (config *Config) SetTimestampFieldName(timestampFieldName string) *Config {
	config.TimestampFieldName = timestampFieldName
	return config
}

func (config *Config) SetVersionFieldName(versionFieldName string) *Config {
	config.VersionFieldName = versionFieldName
	return config
}

func (config *Config) SetAppVersionValue(appVersion string) *Config {
	config.AppVersionValue = appVersion
	return config
}

func (config *Config) SetAppMarketIdValue(appMarketId string) *Config {
	config.AppMarketIdValue = appMarketId
	return config
}
