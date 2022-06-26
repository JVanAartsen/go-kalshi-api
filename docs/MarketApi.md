# \MarketApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveMarkets**](MarketApi.md#GetActiveMarkets) | **Get** /active_markets | GetActiveMarkets
[**GetCandlestickMarketHistory**](MarketApi.md#GetCandlestickMarketHistory) | **Get** /markets/{market_id}/candlestick | GetCandlestickMarketHistory
[**GetCandlestickMarketHistoryCached**](MarketApi.md#GetCandlestickMarketHistoryCached) | **Get** /cached/markets/{market_id}/candlestick | GetCandlestickMarketHistoryCached
[**GetMarket**](MarketApi.md#GetMarket) | **Get** /markets/{market_id} | GetMarket
[**GetMarketByTicker**](MarketApi.md#GetMarketByTicker) | **Get** /markets_by_ticker/{ticker_name} | GetMarketByTicker
[**GetMarketByTickerCached**](MarketApi.md#GetMarketByTickerCached) | **Get** /cached/markets_by_ticker/{ticker_name} | GetMarketByTickerCached
[**GetMarketCached**](MarketApi.md#GetMarketCached) | **Get** /cached/markets/{market_id} | GetMarketCached
[**GetMarketHistory**](MarketApi.md#GetMarketHistory) | **Get** /markets/{market_id}/stats_history | GetMarketHistory
[**GetMarketHistoryCached**](MarketApi.md#GetMarketHistoryCached) | **Get** /cached/markets/{market_id}/stats_history | GetMarketHistoryCached
[**GetMarketOrderBook**](MarketApi.md#GetMarketOrderBook) | **Get** /markets/{market_id}/order_book | GetMarketOrderBook
[**GetMarketOrderBookCached**](MarketApi.md#GetMarketOrderBookCached) | **Get** /cached/markets/{market_id}/order_book | GetMarketOrderBookCached
[**GetMarkets**](MarketApi.md#GetMarkets) | **Get** /markets | GetMarkets
[**GetMarketsCached**](MarketApi.md#GetMarketsCached) | **Get** /cached/markets | GetMarketsCached



## GetActiveMarkets

> GetActiveMarketsResponse GetActiveMarkets(ctx).Limit(limit).MinDate(minDate).MaxDate(maxDate).Execute()

GetActiveMarkets



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
    limit := int64(789) // int64 | The maximum number of markets returned, this is capped at 20 (optional)
    minDate := int64(789) // int64 | The lower bound on the date searched through when looking for activity (optional)
    maxDate := int64(789) // int64 | The upper bound on the date searched through when looking for activity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetActiveMarkets(context.Background()).Limit(limit).MinDate(minDate).MaxDate(maxDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetActiveMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveMarkets`: GetActiveMarketsResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetActiveMarkets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveMarketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The maximum number of markets returned, this is capped at 20 | 
 **minDate** | **int64** | The lower bound on the date searched through when looking for activity | 
 **maxDate** | **int64** | The upper bound on the date searched through when looking for activity | 

### Return type

[**GetActiveMarketsResponse**](GetActiveMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCandlestickMarketHistory

> GetCandlestickMarketHistoryResponse GetCandlestickMarketHistory(ctx, marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()

GetCandlestickMarketHistory



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market
    lastSeenTs := int64(789) // int64 | If provided, restricts history to trades starting from lastSeenTs (optional)
    numBuckets := int32(56) // int32 | If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetCandlestickMarketHistory(context.Background(), marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetCandlestickMarketHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCandlestickMarketHistory`: GetCandlestickMarketHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetCandlestickMarketHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlestickMarketHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **int64** | If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **int32** | If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. | 

### Return type

[**GetCandlestickMarketHistoryResponse**](GetCandlestickMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCandlestickMarketHistoryCached

> GetCandlestickMarketHistoryResponse GetCandlestickMarketHistoryCached(ctx, marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()

GetCandlestickMarketHistoryCached



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market
    lastSeenTs := int64(789) // int64 | If provided, restricts history to trades starting from lastSeenTs (optional)
    numBuckets := int32(56) // int32 | If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetCandlestickMarketHistoryCached(context.Background(), marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetCandlestickMarketHistoryCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCandlestickMarketHistoryCached`: GetCandlestickMarketHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetCandlestickMarketHistoryCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlestickMarketHistoryCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **int64** | If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **int32** | If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. | 

### Return type

[**GetCandlestickMarketHistoryResponse**](GetCandlestickMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarket

> UserGetMarketResponse GetMarket(ctx, marketId).Execute()

GetMarket



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarket(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarket`: UserGetMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketByTicker

> UserGetMarketResponse GetMarketByTicker(ctx, tickerName).Execute()

GetMarketByTicker



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
    tickerName := "tickerName_example" // string | Should be filled with the ticker name of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketByTicker(context.Background(), tickerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketByTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketByTicker`: UserGetMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketByTicker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tickerName** | **string** | Should be filled with the ticker name of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketByTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketByTickerCached

> UserGetMarketResponse GetMarketByTickerCached(ctx, tickerName).Execute()

GetMarketByTickerCached



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
    tickerName := "tickerName_example" // string | Should be filled with the ticker name of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketByTickerCached(context.Background(), tickerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketByTickerCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketByTickerCached`: UserGetMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketByTickerCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tickerName** | **string** | Should be filled with the ticker name of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketByTickerCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCached

> UserGetMarketResponse GetMarketCached(ctx, marketId).Execute()

GetMarketCached



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketCached(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketCached`: UserGetMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHistory

> GetMarketHistoryResponse GetMarketHistory(ctx, marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()

GetMarketHistory



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market
    lastSeenTs := int64(789) // int64 | If provided, restricts history to trades starting from lastSeenTs (optional)
    numBuckets := int32(56) // int32 | If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketHistory(context.Background(), marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketHistory`: GetMarketHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **int64** | If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **int32** | If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) | 

### Return type

[**GetMarketHistoryResponse**](GetMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHistoryCached

> GetMarketHistoryResponse GetMarketHistoryCached(ctx, marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()

GetMarketHistoryCached



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market
    lastSeenTs := int64(789) // int64 | If provided, restricts history to trades starting from lastSeenTs (optional)
    numBuckets := int32(56) // int32 | If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketHistoryCached(context.Background(), marketId).LastSeenTs(lastSeenTs).NumBuckets(numBuckets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketHistoryCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketHistoryCached`: GetMarketHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketHistoryCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHistoryCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **int64** | If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **int32** | If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) | 

### Return type

[**GetMarketHistoryResponse**](GetMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketOrderBook

> GetMarketOrderBookResponse GetMarketOrderBook(ctx, marketId).Execute()

GetMarketOrderBook



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketOrderBook(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketOrderBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketOrderBook`: GetMarketOrderBookResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketOrderBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketOrderBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMarketOrderBookResponse**](GetMarketOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketOrderBookCached

> GetMarketOrderBookResponse GetMarketOrderBookCached(ctx, marketId).Execute()

GetMarketOrderBookCached



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketApi.GetMarketOrderBookCached(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketOrderBookCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketOrderBookCached`: GetMarketOrderBookResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketOrderBookCached`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketOrderBookCachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMarketOrderBookResponse**](GetMarketOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarkets

> UserGetMarketsResponse GetMarkets(ctx).Execute()

GetMarkets



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
    resp, r, err := apiClient.MarketApi.GetMarkets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarkets`: UserGetMarketsResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarkets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsRequest struct via the builder pattern


### Return type

[**UserGetMarketsResponse**](UserGetMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketsCached

> UserGetMarketsResponse GetMarketsCached(ctx).Execute()

GetMarketsCached



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
    resp, r, err := apiClient.MarketApi.GetMarketsCached(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetMarketsCached``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketsCached`: UserGetMarketsResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetMarketsCached`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsCachedRequest struct via the builder pattern


### Return type

[**UserGetMarketsResponse**](UserGetMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

