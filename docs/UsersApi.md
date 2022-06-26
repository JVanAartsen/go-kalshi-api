# \UsersApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserImmediateFunding**](UsersApi.md#GetUserImmediateFunding) | **Get** /users/{user_id}/immediate_funding | GetUserImmediateFunding
[**GetUserWithdrawalAvailableBalance**](UsersApi.md#GetUserWithdrawalAvailableBalance) | **Get** /users/{user_id}/withdrawals/available | GetUserWithdrawalAvailableBalance



## GetUserImmediateFunding

> GetUserImmediateFundingResponse GetUserImmediateFunding(ctx, userId).DepositAmountCents(depositAmountCents).Execute()

GetUserImmediateFunding



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
    depositAmountCents := int64(789) // int64 | Pass this parameter if you'd like to see how much of a deposit will be funded by immediate funding. If you don't need this information, pass 0 cents. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserImmediateFunding(context.Background(), userId).DepositAmountCents(depositAmountCents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserImmediateFunding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserImmediateFunding`: GetUserImmediateFundingResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserImmediateFunding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserImmediateFundingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositAmountCents** | **int64** | Pass this parameter if you&#39;d like to see how much of a deposit will be funded by immediate funding. If you don&#39;t need this information, pass 0 cents. | 

### Return type

[**GetUserImmediateFundingResponse**](GetUserImmediateFundingResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserWithdrawalAvailableBalance

> UserWithdrawalAvailableBalanceResponse GetUserWithdrawalAvailableBalance(ctx, userId).Execute()

GetUserWithdrawalAvailableBalance



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserWithdrawalAvailableBalance(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserWithdrawalAvailableBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWithdrawalAvailableBalance`: UserWithdrawalAvailableBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserWithdrawalAvailableBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWithdrawalAvailableBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserWithdrawalAvailableBalanceResponse**](UserWithdrawalAvailableBalanceResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

