# \SeriesApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSeriesByTicker**](SeriesApi.md#GetSeriesByTicker) | **Get** /series/{series_ticker} | GetSeriesByTicker



## GetSeriesByTicker

> GetSeriesByTickerResponse GetSeriesByTicker(ctx, seriesTicker).Execute()

GetSeriesByTicker



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
    seriesTicker := "seriesTicker_example" // string | Should be the ticker of the series

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesApi.GetSeriesByTicker(context.Background(), seriesTicker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.GetSeriesByTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesByTicker`: GetSeriesByTickerResponse
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.GetSeriesByTicker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesTicker** | **string** | Should be the ticker of the series | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesByTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSeriesByTickerResponse**](GetSeriesByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

