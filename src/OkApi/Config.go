package OkApi

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
	return &Config{
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

func (c *Config) SetDisableSSL(disableSSL bool) *Config {
	c.DisableSSL = disableSSL
	return c
}

func (c *Config) SetGatewayHost(gatewayHost string) *Config {
	c.GatewayHost = gatewayHost
	return c
}

func (c *Config) SetHttpMethod(httpMethod string) *Config {
	c.HttpMethod = httpMethod
	return c
}

func (c *Config) SetAppKeyFieldName(appkeyFieldName string) *Config {
	c.AppVersionFieldName = appkeyFieldName
	return c
}

func (c *Config) SetAppVersionFieldName(appVersionFieldName string) *Config {
	c.AppVersionFieldName = appVersionFieldName
	return c
}

func (c *Config) SetAppMarketIdFieldName(appMarketIdFieldName string) *Config {
	c.AppMarketIdFieldName = appMarketIdFieldName
	return c
}

func (c *Config) SetDeviceIdFieldName(deviceIdFieldName string) *Config {
	c.DeviceIdFieldName = deviceIdFieldName
	return c
}

func (c *Config) SetSessionIdFieldName(sessionIdFieldName string) *Config {
	c.AppVersionFieldName = sessionIdFieldName
	return c
}

func (c *Config) SetSignFieldName(signFieldName string) *Config {
	c.SignFieldName = signFieldName
	return c
}

func (c *Config) SetApiFieldName(apiFieldName string) *Config {
	c.ApiFieldName = apiFieldName
	return c
}

func (c *Config) SetParamsFieldName(paramsFieldName string) *Config {
	c.ParamsFieldName = paramsFieldName
	return c
}

func (c *Config) SetTimestampFieldName(timestampFieldName string) *Config {
	c.TimestampFieldName = timestampFieldName
	return c
}

func (c *Config) SetVersionFieldName(versionFieldName string) *Config {
	c.VersionFieldName = versionFieldName
	return c
}

func (c *Config) SetAppVersionValue(appVersion string) *Config {
	c.AppVersionValue = appVersion
	return c
}

func (c *Config) SetAppMarketIdValue(appMarketId string) *Config {
	c.AppMarketIdValue = appMarketId
	return c
}
