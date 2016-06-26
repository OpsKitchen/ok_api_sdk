package main

import (
	"fmt"
	"sdk"
	"sdk/di/logger"
)

func main() {
	var client *sdk.Client = sdk.NewClient()

	//init config
	client.RequestBuilder.Config.SetAppVersionValue("1.0.1").SetAppMarketIdValue("678").SetGatewayHost("api.OpsKitchen.com").SetDisableSSL(true)

	//init credential
	client.RequestBuilder.Credential.SetAppKey("101").SetSecret("7INWkF/qSkkXrFwZ")

	//enable debug log
	sdk.DefaultLogger.SetLevel(logger.DebugLevel)

	//call api
	resp, err := client.CallApi("ops.meta.os.list", "1.0", nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
}