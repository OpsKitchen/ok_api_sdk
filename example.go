package main

import (
	"fmt"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/di/logger"
	"github.com/OpsKitchen/ok_api_sdk_go/sdk/model"
)

func main() {
	var client *sdk.Client = sdk.NewClient()
	var resp *model.ApiResult
	var err error

	//init config
	client.RequestBuilder.Config.SetAppVersionValue("1.0.1").SetAppMarketIdValue("678").SetGatewayHost("api.OpsKitchen.com").SetDisableSSL(true)

	//init credential
	client.RequestBuilder.Credential.SetAppKey("101").SetSecret("7INWkF/qSkkXrFwZ")

	//enable debug log
	sdk.SetDefaultLogger(&logger.Logger{
		Level: logger.DebugLevel,
	})

	//call api without parameter
	resp, err = client.CallApi("ops.meta.os.list", "1.0", nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	//call api with parameter
	param := make(map[string]string)
	param["osReleaseId"] = "3022"
	resp, err = client.CallApi("ops.meta.osImage.listByOsReleaseId", "1.0", param)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
}
