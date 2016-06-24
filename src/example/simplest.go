package example

import "sdk"

func main() {
	var config sdk.Config = sdk.NewConfig()
	var credential sdk.Credential = sdk.NewCredential()
	var client sdk.Client = sdk.NewClient()

	//init config
	config.SetAppVersionValue("1.0.1")
	config.SetAppMarketIdValue("678")

	//init credential
	credential.SetAppKey("1234567")
	credential.SetSecret("S#$%^&UJHVF")

	//init client
	client.RequestBuilder.SetConfig(config)
	client.RequestBuilder.SetCredential(credential)

	//init user
	var resp = client.CallApi("demo.time.get", "1.0", nil)
	resp.IsSuccess()
}
