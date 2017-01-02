package model

type Config struct {
	DisableSSL           bool //platform address
	GatewayHost          string
	GatewayPath          string
	GatewayPort          int
	AppKeyFieldName      string //System parameter name in HTTP header
	AppVersionFieldName  string
	AppMarketIdFieldName string
	DeviceIdFieldName    string
	SessionIdFieldName   string
	SignFieldName        string
	ApiFieldName         string //System parameter name in HTTP body
	ParamsFieldName      string
	TimestampFieldName   string
	VersionFieldName     string
	AppVersionValue      string //System parameter value
	AppMarketIdValue     string
}

func (config *Config) SetDefaultOption() *Config {
	config.AppKeyFieldName = "OA-App-Key"
	config.AppVersionFieldName = "OA-App-Version"
	config.AppMarketIdFieldName = "OA-App-Market-ID"
	config.DeviceIdFieldName = "OA-Device-Id"
	config.SessionIdFieldName = "OA-Session-Id"
	config.SignFieldName = "OA-Sign"
	config.ApiFieldName = "api"
	config.ParamsFieldName = "params"
	config.TimestampFieldName = "timestamp"
	config.VersionFieldName = "version"
	return config
}

func (config *Config) SetDisableSSL(disableSSL bool) *Config {
	config.DisableSSL = disableSSL
	return config
}

func (config *Config) SetGatewayHost(gatewayHost string) *Config {
	config.GatewayHost = gatewayHost
	return config
}

func (config *Config) SetGatewayPath(gatewayPath string) *Config {
	config.GatewayPath = gatewayPath
	return config
}

func (config *Config) SetGatewayPort(gatewayPort int) *Config {
	config.GatewayPort = gatewayPort
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
