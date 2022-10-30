# alchemypay-sdk-go

# Usage
```go
config := sdk.Config{
    Scheme:       "https",
    Endpoint:     "xxx",
    MerchantCode: "xxx",
    PrivateKey:   "xxxx=",
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
```
