# \AccountApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeSubscription**](AccountApi.md#ChangeSubscription) | **Patch** /users/{user_id}/subscribe | ChangeSubscription
[**GetNotificationPreferences**](AccountApi.md#GetNotificationPreferences) | **Get** /users/{user_id}/notifications/preferences | GetNotificationPreferences
[**NotificationMarkRead**](AccountApi.md#NotificationMarkRead) | **Put** /users/{user_id}/notifications/{notification_id}/read | NotificationMarkRead
[**UserGetAccountHistory**](AccountApi.md#UserGetAccountHistory) | **Get** /users/{user_id}/account/history | UserGetAccountHistory
[**UserGetNotifications**](AccountApi.md#UserGetNotifications) | **Get** /users/{user_id}/notifications | UserGetNotifications
[**UserGetProfitsAndLosses**](AccountApi.md#UserGetProfitsAndLosses) | **Get** /users/{user_id}/account/pnl | UserGetProfitsAndLosses



## ChangeSubscription

> ChangeSubscriptionResponse ChangeSubscription(ctx, userId).ChangeSubscriptionRequest(changeSubscriptionRequest).Execute()

ChangeSubscription



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
    changeSubscriptionRequest := *openapiclient.NewChangeSubscriptionRequest("SubscriptionLevel_example") // ChangeSubscriptionRequest | Change subscription data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.ChangeSubscription(context.Background(), userId).ChangeSubscriptionRequest(changeSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ChangeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeSubscription`: ChangeSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ChangeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeSubscriptionRequest** | [**ChangeSubscriptionRequest**](ChangeSubscriptionRequest.md) | Change subscription data | 

### Return type

[**ChangeSubscriptionResponse**](ChangeSubscriptionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationPreferences

> GetNotificationPreferencesResponse GetNotificationPreferences(ctx, userId).Execute()

GetNotificationPreferences



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
    resp, r, err := apiClient.AccountApi.GetNotificationPreferences(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetNotificationPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPreferences`: GetNotificationPreferencesResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetNotificationPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationPreferencesResponse**](GetNotificationPreferencesResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationMarkRead

> NotificationMarkRead(ctx, userId, notificationId).Execute()

NotificationMarkRead



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
    userId := "userId_example" // string | user_id should be filled with your user_id provided on log_in
    notificationId := "notificationId_example" // string | notification_id should be filled with the id of the notification to be mark as read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.NotificationMarkRead(context.Background(), userId, notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.NotificationMarkRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id should be filled with your user_id provided on log_in | 
**notificationId** | **string** | notification_id should be filled with the id of the notification to be mark as read | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationMarkReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetAccountHistory

> UserGetAccountHistoryResponse UserGetAccountHistory(ctx, userId).Deposits(deposits).Withdrawals(withdrawals).Orders(orders).Settlements(settlements).Trades(trades).Credits(credits).Limit(limit).PageSize(pageSize).PageNumber(pageNumber).Execute()

UserGetAccountHistory



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
    deposits := true // bool | If true the response should include deposit entries (optional)
    withdrawals := true // bool | If true the response should include withdrawal entries (optional)
    orders := true // bool | If true the response should include order entries (optional)
    settlements := true // bool | If true the response should include settlement entries (optional)
    trades := true // bool | If true the response should include trade entries (optional)
    credits := true // bool | If true the response should include credit entries (optional)
    limit := int64(789) // int64 | Restricts the response to a return the first \"limit\" amount of acct history items. Note if you specify a limit, you cannot specify a PageSize or PageNumber (optional)
    pageSize := int64(789) // int64 | Parameter to specify the number of results per page. (optional)
    pageNumber := int64(789) // int64 | Parameter to specify which page of the results should be retrieved (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.UserGetAccountHistory(context.Background(), userId).Deposits(deposits).Withdrawals(withdrawals).Orders(orders).Settlements(settlements).Trades(trades).Credits(credits).Limit(limit).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UserGetAccountHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetAccountHistory`: UserGetAccountHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UserGetAccountHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetAccountHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deposits** | **bool** | If true the response should include deposit entries | 
 **withdrawals** | **bool** | If true the response should include withdrawal entries | 
 **orders** | **bool** | If true the response should include order entries | 
 **settlements** | **bool** | If true the response should include settlement entries | 
 **trades** | **bool** | If true the response should include trade entries | 
 **credits** | **bool** | If true the response should include credit entries | 
 **limit** | **int64** | Restricts the response to a return the first \&quot;limit\&quot; amount of acct history items. Note if you specify a limit, you cannot specify a PageSize or PageNumber | 
 **pageSize** | **int64** | Parameter to specify the number of results per page. | 
 **pageNumber** | **int64** | Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserGetAccountHistoryResponse**](UserGetAccountHistoryResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetNotifications

> UserGetNotificationsResponse UserGetNotifications(ctx, userId).PageSize(pageSize).PageNumber(pageNumber).Execute()

UserGetNotifications



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
    pageSize := int64(789) // int64 | Number of results per page
    pageNumber := int64(789) // int64 | Page of the results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.UserGetNotifications(context.Background(), userId).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UserGetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetNotifications`: UserGetNotificationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UserGetNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | Number of results per page | 
 **pageNumber** | **int64** | Page of the results | 

### Return type

[**UserGetNotificationsResponse**](UserGetNotificationsResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetProfitsAndLosses

> UserGetProfitsAndLossesResponse UserGetProfitsAndLosses(ctx, userId).StartTs(startTs).EndTs(endTs).Execute()

UserGetProfitsAndLosses



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
    startTs := int64(789) // int64 | Start time of pnl period, represented as the number of seconds since Unix epoch (optional)
    endTs := int64(789) // int64 | End time of pnl period, represented as the number of seconds since Unix epoch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.UserGetProfitsAndLosses(context.Background(), userId).StartTs(startTs).EndTs(endTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UserGetProfitsAndLosses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetProfitsAndLosses`: UserGetProfitsAndLossesResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UserGetProfitsAndLosses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetProfitsAndLossesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTs** | **int64** | Start time of pnl period, represented as the number of seconds since Unix epoch | 
 **endTs** | **int64** | End time of pnl period, represented as the number of seconds since Unix epoch | 

### Return type

[**UserGetProfitsAndLossesResponse**](UserGetProfitsAndLossesResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

