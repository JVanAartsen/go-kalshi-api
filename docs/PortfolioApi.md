# \PortfolioApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeprecatedUserGetPortfolioHistory**](PortfolioApi.md#DeprecatedUserGetPortfolioHistory) | **Get** /users/{user_id}/portfolio/ | DeprecatedUserGetPortfolioHistory
[**UserGetSampledPortfolioHistory**](PortfolioApi.md#UserGetSampledPortfolioHistory) | **Get** /users/{user_id}/sampled_portfolio/ | UserGetSampledPortfolioHistory



## DeprecatedUserGetPortfolioHistory

> DeprecatedUserGetPortfolioHistoryResponse DeprecatedUserGetPortfolioHistory(ctx, userId).DeprecatedUserGetPortfolioHistoryRequest(deprecatedUserGetPortfolioHistoryRequest).Execute()

DeprecatedUserGetPortfolioHistory



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
    userId := "userId_example" // string | This parameter should be filled with your user_id provided on log_in
    deprecatedUserGetPortfolioHistoryRequest := *openapiclient.NewDeprecatedUserGetPortfolioHistoryRequest() // DeprecatedUserGetPortfolioHistoryRequest | Order create input data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortfolioApi.DeprecatedUserGetPortfolioHistory(context.Background(), userId).DeprecatedUserGetPortfolioHistoryRequest(deprecatedUserGetPortfolioHistoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioApi.DeprecatedUserGetPortfolioHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedUserGetPortfolioHistory`: DeprecatedUserGetPortfolioHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `PortfolioApi.DeprecatedUserGetPortfolioHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedUserGetPortfolioHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deprecatedUserGetPortfolioHistoryRequest** | [**DeprecatedUserGetPortfolioHistoryRequest**](DeprecatedUserGetPortfolioHistoryRequest.md) | Order create input data | 

### Return type

[**DeprecatedUserGetPortfolioHistoryResponse**](DeprecatedUserGetPortfolioHistoryResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetSampledPortfolioHistory

> UserGetSampledPortfolioHistoryResponse UserGetSampledPortfolioHistory(ctx, userId).Execute()

UserGetSampledPortfolioHistory



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
    userId := "userId_example" // string | Should be filled with your user_id provided on log_in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortfolioApi.UserGetSampledPortfolioHistory(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioApi.UserGetSampledPortfolioHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetSampledPortfolioHistory`: UserGetSampledPortfolioHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `PortfolioApi.UserGetSampledPortfolioHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetSampledPortfolioHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetSampledPortfolioHistoryResponse**](UserGetSampledPortfolioHistoryResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

