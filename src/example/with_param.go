package main

import (
	"sdk"
	"fmt"
)

func main() {
	var config *sdk.Config = sdk.NewConfig()
	var credential *sdk.Credential = sdk.NewCredential()
	var client *sdk.Client = sdk.NewClient()

	//init config
	config.AppVersionValue = "1.0.1"
	config.AppMarketIdValue = "678"
	config.DisableSSL = true
	config.GatewayHost = "api.OpsKitchen.com"

	//init credential
	credential.AppKey = "101"
	credential.Secret = "your secret"

	//init client
	client.RequestBuilder.Config = config
	client.RequestBuilder.Credential = credential

	//query param
	param := make(map[string]string)
	param["osReleaseId"] = "3022"

	//call api
	resp, err := client.CallApi("ops.meta.osImage.listByOsReleaseId", "1.0", param)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Data)
}
