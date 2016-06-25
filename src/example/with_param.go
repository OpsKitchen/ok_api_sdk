package main

import (
	"fmt"
	"sdk"
)

func main() {
	var client *sdk.Client = sdk.NewClient()

	//init config
	client.RequestBuilder.Config.SetAppVersionValue("1.0.1").SetAppMarketIdValue("678").SetGatewayHost("api.OpsKitchen.com").SetDisableSSL(true)

	//init credential
	client.RequestBuilder.Credential.SetAppKey("101").SetSecret("7INWkF/qSkkXrFwZ")

	//query param
	param := make(map[string]string)
	param["osReleaseId"] = "3022"

	//call api
	resp, err := client.CallApi("ops.meta.osImage.listByOsReleaseId", "1.0", param)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
}
