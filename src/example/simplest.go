package example

import "sdk"

func main() {
	var config *sdk.Config = sdk.NewConfig()
	var credential *sdk.Credential = sdk.NewCredential()
	var client *sdk.Client = sdk.NewClient()

	//init config
	config.AppVersionValue = "1.0.1"
	config.AppMarketIdValue = "678"

	//init credential
	credential.AppKey = "1234567"
	credential.Secret = "S#$%^&UJHVF"

	//init client
	client.RequestBuilder.Config = config
	client.RequestBuilder.Credential = credential

	//init user
	resp, err := client.CallApi("demo.time.get", "1.0", nil)
	resp.Data
}
