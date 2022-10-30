package main

import (
	"fmt"
	"github.com/hedeqiang/alchemypay-sdk-go/sdk"
	"github.com/hedeqiang/alchemypay-sdk-go/services"
)

func main() {
	config := sdk.Config{
		Scheme:       "https",
		Endpoint:     "xxx",
		MerchantCode: "xxx",
		PrivateKey:   "xxx",
	}
	client, _ := services.NewClient(&config)

	req := services.NewPayRequest()
	req.URL = "openApi/digitalCurrencyList"

	resp, err := client.Pay(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("get response body: `%s`\n", resp.String())

	fmt.Printf("get response body: `%s`\n", resp.GetResponseBody())
}
