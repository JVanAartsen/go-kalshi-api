# \RangedMarketsApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRangedMarketByTicker**](RangedMarketsApi.md#GetRangedMarketByTicker) | **Get** /ranged_markets_by_ticker/{ticker} | GetRangedMarketByTicker
[**UserGetRangedMarketPosition**](RangedMarketsApi.md#UserGetRangedMarketPosition) | **Get** /users/{user_id}/ranged_positions/{ranged_market_id} | UserGetRangedMarketPosition



## GetRangedMarketByTicker

> GetRangedMarketByTickerResponse GetRangedMarketByTicker(ctx, ticker).Execute()

GetRangedMarketByTicker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | Should be the ticker of the ranged market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RangedMarketsApi.GetRangedMarketByTicker(context.Background(), ticker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RangedMarketsApi.GetRangedMarketByTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRangedMarketByTicker`: GetRangedMarketByTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `RangedMarketsApi.GetRangedMarketByTicker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | Should be the ticker of the ranged market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRangedMarketByTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRangedMarketByTickerResponse**](GetRangedMarketByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetRangedMarketPosition

> UserGetRangedMarketPositionResponse UserGetRangedMarketPosition(ctx, userId, rangedMarketId).Execute()

UserGetRangedMarketPosition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | Should be filled in with your user_id provided on log_in
    rangedMarketId := "rangedMarketId_example" // string | Should be filled with the id of the target ranged market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RangedMarketsApi.UserGetRangedMarketPosition(context.Background(), userId, rangedMarketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RangedMarketsApi.UserGetRangedMarketPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetRangedMarketPosition`: UserGetRangedMarketPositionResponse
    fmt.Fprintf(os.Stdout, "Response from `RangedMarketsApi.UserGetRangedMarketPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled in with your user_id provided on log_in | 
**rangedMarketId** | **string** | Should be filled with the id of the target ranged market | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRangedMarketPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserGetRangedMarketPositionResponse**](UserGetRangedMarketPositionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

