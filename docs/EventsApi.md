# \EventsApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventByTickerCached**](EventsApi.md#GetEventByTickerCached) | **Get** /events/{ticker} | GetEventByTickerCached



## GetEventByTickerCached

> GetEventByTickerResponse GetEventByTickerCached(ctx, ticker).Execute()

GetEventByTickerCached



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
    ticker := "ticker_example" // string | Should be the ticker of the event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventByTickerCached(context.Background(), ticker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventByTickerCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventByTickerCached`: GetEventByTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventByTickerCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | Should be the ticker of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventByTickerCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventByTickerResponse**](GetEventByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

