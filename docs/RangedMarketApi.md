# \RangedMarketApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRangedMarket**](RangedMarketApi.md#GetRangedMarket) | **Get** /ranged_markets/{ranged_market_id} | GetRangedMarket



## GetRangedMarket

> GetRangedMarketResponse GetRangedMarket(ctx, rangedMarketId).Execute()

GetRangedMarket



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
    rangedMarketId := "rangedMarketId_example" // string | Should be filled in with a ranged market id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RangedMarketApi.GetRangedMarket(context.Background(), rangedMarketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RangedMarketApi.GetRangedMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRangedMarket`: GetRangedMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `RangedMarketApi.GetRangedMarket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangedMarketId** | **string** | Should be filled in with a ranged market id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRangedMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRangedMarketResponse**](GetRangedMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

