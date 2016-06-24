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