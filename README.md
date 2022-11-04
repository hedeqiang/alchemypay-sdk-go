# alchemypay-sdk-go

# Installing

```go
go get -u github.com/hedeqiang/alchemypay
```



# Configuration

```go
config := sdk.Config{
    Scheme:       "https",
    Endpoint:     "xxx",
    MerchantCode: "xxx",
    PrivateKey:   "xxxx=",
}
client, _ := services.NewClient(&config)
```



# Usage

> 任意接口不需要传递 `merchantCode`、`sign` 字段。此 `SDK` 已实现签名功能.

### 数字货币支持列表接口

```go
req := services.NewPayRequest()
req.URL = "openApi/digitalCurrencyList"

resp, err := client.Pay(req)
if err != nil {
    fmt.Println("err:", err)
    return
}
fmt.Printf("get response body: `%s`\n", resp.String())
```



### 创建订单接口

```go
req := services.NewPayRequest()
req.URL = "openApi/createOrder"
req.Params = map[string]interface{}{
    "amount":           10.0,
    "fiatType":         "CNY",
    "callbackUrl":      "http://147.243.170.11:9091/transnotify",
    "merchantOrderNum": "testqwe1234567891035",
    "payMent":          "w1",
    "email":            "123456@qq.com",
}
resp, err := client.Pay(req)
if err != nil {
    fmt.Println("err:", err)
    return
}
fmt.Printf("get response body: `%s`\n", resp.String())
```



More

