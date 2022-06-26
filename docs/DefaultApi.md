# \DefaultApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsCached**](DefaultApi.md#GetEventsCached) | **Get** /events/ | GetEventsCached
[**GetSeriesList**](DefaultApi.md#GetSeriesList) | **Get** /series/ | GetSeriesList
[**GetSeriesListCached**](DefaultApi.md#GetSeriesListCached) | **Get** /cached/series | GetSeriesListCached
[**GetTrades**](DefaultApi.md#GetTrades) | **Get** /trades | GetTrades



## GetEventsCached

> GetEventsResponse GetEventsCached(ctx).Execute()

GetEventsCached



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEventsCached(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEventsCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsCached`: GetEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEventsCached`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsCachedRequest struct via the builder pattern


### Return type

[**GetEventsResponse**](GetEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesList

> GetSeriesListResponse GetSeriesList(ctx).Execute()

GetSeriesList



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSeriesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSeriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesList`: GetSeriesListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSeriesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesListRequest struct via the builder pattern


### Return type

[**GetSeriesListResponse**](GetSeriesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesListCached

> GetSeriesListResponse GetSeriesListCached(ctx).Execute()

GetSeriesListCached



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSeriesListCached(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSeriesListCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesListCached`: GetSeriesListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSeriesListCached`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesListCachedRequest struct via the builder pattern


### Return type

[**GetSeriesListResponse**](GetSeriesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrades

> TradesGetResponse GetTrades(ctx).TradesDate(tradesDate).PageSize(pageSize).PageNumber(pageNumber).MarketId(marketId).Execute()

GetTrades



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
    tradesDate := "tradesDate_example" // string | Restricts the response to trades during a certain day (give trades_date in ET with format: YYYY-MM-DD). Dates returned will be UTC (optional)
    pageSize := int64(789) // int64 | Parameter to specify the number of results per page (optional)
    pageNumber := int64(789) // int64 | Parameter to specify which page of the results should be retrieved (optional)
    marketId := "marketId_example" // string | Parameter to specify a specific marketId to get trades from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTrades(context.Background()).TradesDate(tradesDate).PageSize(pageSize).PageNumber(pageNumber).MarketId(marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrades`: TradesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradesDate** | **string** | Restricts the response to trades during a certain day (give trades_date in ET with format: YYYY-MM-DD). Dates returned will be UTC | 
 **pageSize** | **int64** | Parameter to specify the number of results per page | 
 **pageNumber** | **int64** | Parameter to specify which page of the results should be retrieved | 
 **marketId** | **string** | Parameter to specify a specific marketId to get trades from | 

### Return type

[**TradesGetResponse**](TradesGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

