# \UserApi

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAddFavoritedSeries**](UserApi.md#UserAddFavoritedSeries) | **Put** /users/{user_id}/favorited_series/{series_ticker} | UserAddFavoritedSeries
[**UserAddWatchlist**](UserApi.md#UserAddWatchlist) | **Put** /users/{user_id}/watchlist/{market_id} | UserAddWatchlist
[**UserBatchOrdersCancel**](UserApi.md#UserBatchOrdersCancel) | **Delete** /users/{user_id}/batch_orders | UserBatchOrdersCancel
[**UserBatchOrdersCreate**](UserApi.md#UserBatchOrdersCreate) | **Post** /users/{user_id}/batch_orders | UserBatchOrdersCreate
[**UserChangePassword**](UserApi.md#UserChangePassword) | **Put** /users/{user_id}/password | UserChangePassword
[**UserDeactivate**](UserApi.md#UserDeactivate) | **Delete** /users | UserDeactivate
[**UserGetBalance**](UserApi.md#UserGetBalance) | **Get** /users/{user_id}/balance | UserGetBalance
[**UserGetFavoritedSeries**](UserApi.md#UserGetFavoritedSeries) | **Get** /users/{user_id}/favorited_series | UserGetFavoritedSeries
[**UserGetMarketPosition**](UserApi.md#UserGetMarketPosition) | **Get** /users/{user_id}/positions/{market_id} | UserGetMarketPosition
[**UserGetMarketPositions**](UserApi.md#UserGetMarketPositions) | **Get** /users/{user_id}/positions | UserGetMarketPositions
[**UserGetProfile**](UserApi.md#UserGetProfile) | **Get** /users/{user_id} | UserGetProfile
[**UserGetReferralInfo**](UserApi.md#UserGetReferralInfo) | **Get** /users/{user_id}/referrals | UserGetReferralInfo
[**UserGetWatchlist**](UserApi.md#UserGetWatchlist) | **Get** /users/{user_id}/watchlist | UserGetWatchlist
[**UserOrderCancel**](UserApi.md#UserOrderCancel) | **Delete** /users/{user_id}/orders/{order_id} | UserOrderCancel
[**UserOrderCreate**](UserApi.md#UserOrderCreate) | **Post** /users/{user_id}/orders | UserOrderCreate
[**UserOrderDecrease**](UserApi.md#UserOrderDecrease) | **Post** /users/{user_id}/orders/{order_id}/decrease | UserOrderDecrease
[**UserOrdersGet**](UserApi.md#UserOrdersGet) | **Get** /users/{user_id}/orders | UserOrdersGet
[**UserRemoveFavoritedSeries**](UserApi.md#UserRemoveFavoritedSeries) | **Delete** /users/{user_id}/favorited_series/{series_ticker} | UserRemoveFavoritedSeries
[**UserRemoveWatchlist**](UserApi.md#UserRemoveWatchlist) | **Delete** /users/{user_id}/watchlist/{market_id} | UserRemoveWatchlist
[**UserTradesGet**](UserApi.md#UserTradesGet) | **Get** /users/{user_id}/trades | UserTradesGet



## UserAddFavoritedSeries

> UserAddFavoritedSeries(ctx, userId, seriesTicker).Execute()

UserAddFavoritedSeries



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
    seriesTicker := "seriesTicker_example" // string | series_ticker should be filled with the ticker of the series to be added to the list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserAddFavoritedSeries(context.Background(), userId, seriesTicker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserAddFavoritedSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id should be filled with your user_id provided on log_in | 
**seriesTicker** | **string** | series_ticker should be filled with the ticker of the series to be added to the list | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAddFavoritedSeriesRequest struct via the builder pattern


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


## UserAddWatchlist

> UserAddWatchlist(ctx, userId, marketId).Execute()

UserAddWatchlist



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
    marketId := "marketId_example" // string | market_id should be filled with the id of the market to be added to the watchlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserAddWatchlist(context.Background(), userId, marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserAddWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id should be filled with your user_id provided on log_in | 
**marketId** | **string** | market_id should be filled with the id of the market to be added to the watchlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAddWatchlistRequest struct via the builder pattern


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


## UserBatchOrdersCancel

> UserBatchOrdersCancelResponse UserBatchOrdersCancel(ctx, userId).UserBatchOrdersCancelRequest(userBatchOrdersCancelRequest).Execute()

UserBatchOrdersCancel



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
    userBatchOrdersCancelRequest := *openapiclient.NewUserBatchOrdersCancelRequest([]string{"Ids_example"}) // UserBatchOrdersCancelRequest | Orders cancel input data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserBatchOrdersCancel(context.Background(), userId).UserBatchOrdersCancelRequest(userBatchOrdersCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserBatchOrdersCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserBatchOrdersCancel`: UserBatchOrdersCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserBatchOrdersCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserBatchOrdersCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userBatchOrdersCancelRequest** | [**UserBatchOrdersCancelRequest**](UserBatchOrdersCancelRequest.md) | Orders cancel input data | 

### Return type

[**UserBatchOrdersCancelResponse**](UserBatchOrdersCancelResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserBatchOrdersCreate

> UserBatchOrdersCreateResponse UserBatchOrdersCreate(ctx, userId).UserBatchOrdersCreateRequest(userBatchOrdersCreateRequest).Execute()

UserBatchOrdersCreate



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
    userBatchOrdersCreateRequest := *openapiclient.NewUserBatchOrdersCreateRequest([]openapiclient.UserOrderCreateRequest{*openapiclient.NewUserOrderCreateRequest(int32(123), "MarketId_example", int64(123), "Side_example")}) // UserBatchOrdersCreateRequest | Order create input data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserBatchOrdersCreate(context.Background(), userId).UserBatchOrdersCreateRequest(userBatchOrdersCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserBatchOrdersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserBatchOrdersCreate`: UserBatchOrdersCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserBatchOrdersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserBatchOrdersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userBatchOrdersCreateRequest** | [**UserBatchOrdersCreateRequest**](UserBatchOrdersCreateRequest.md) | Order create input data | 

### Return type

[**UserBatchOrdersCreateResponse**](UserBatchOrdersCreateResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChangePassword

> UserChangePassword(ctx, userId).UserChangePasswordRequest(userChangePasswordRequest).Execute()

UserChangePassword



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
    userChangePasswordRequest := *openapiclient.NewUserChangePasswordRequest("NewPassword_example", "OldPassword_example") // UserChangePasswordRequest | Change password input fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserChangePassword(context.Background(), userId).UserChangePasswordRequest(userChangePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userChangePasswordRequest** | [**UserChangePasswordRequest**](UserChangePasswordRequest.md) | Change password input fields. | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeactivate

> CreateUserResponse UserDeactivate(ctx).Execute()

UserDeactivate



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
    resp, r, err := apiClient.UserApi.UserDeactivate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDeactivate`: CreateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserDeactivate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeactivateRequest struct via the builder pattern


### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetBalance

> UserGetBalanceResponse UserGetBalance(ctx, userId).Execute()

UserGetBalance



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
    resp, r, err := apiClient.UserApi.UserGetBalance(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetBalance`: UserGetBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetBalanceResponse**](UserGetBalanceResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetFavoritedSeries

> UserGetFavoritedSeriesResponse UserGetFavoritedSeries(ctx, userId).Execute()

UserGetFavoritedSeries



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserGetFavoritedSeries(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetFavoritedSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetFavoritedSeries`: UserGetFavoritedSeriesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetFavoritedSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetFavoritedSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetFavoritedSeriesResponse**](UserGetFavoritedSeriesResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetMarketPosition

> UserGetMarketPositionResponse UserGetMarketPosition(ctx, userId, marketId).Execute()

UserGetMarketPosition



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserGetMarketPosition(context.Background(), userId, marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetMarketPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetMarketPosition`: UserGetMarketPositionResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetMarketPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetMarketPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserGetMarketPositionResponse**](UserGetMarketPositionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetMarketPositions

> UserGetMarketPositionsResponse UserGetMarketPositions(ctx, userId).Execute()

UserGetMarketPositions



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
    resp, r, err := apiClient.UserApi.UserGetMarketPositions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetMarketPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetMarketPositions`: UserGetMarketPositionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetMarketPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetMarketPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetMarketPositionsResponse**](UserGetMarketPositionsResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetProfile

> UserGetProfileResponse UserGetProfile(ctx, userId).Execute()

UserGetProfile



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
    resp, r, err := apiClient.UserApi.UserGetProfile(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetProfile`: UserGetProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetProfileResponse**](UserGetProfileResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetReferralInfo

> UserGetReferralInfoResponse UserGetReferralInfo(ctx, userId).Execute()

UserGetReferralInfo



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
    resp, r, err := apiClient.UserApi.UserGetReferralInfo(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetReferralInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetReferralInfo`: UserGetReferralInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetReferralInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetReferralInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetReferralInfoResponse**](UserGetReferralInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetWatchlist

> UserGetWatchlistResponse UserGetWatchlist(ctx, userId).Execute()

UserGetWatchlist



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
    resp, r, err := apiClient.UserApi.UserGetWatchlist(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetWatchlist`: UserGetWatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetWatchlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetWatchlistResponse**](UserGetWatchlistResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOrderCancel

> UserOrderDecreaseResponse UserOrderCancel(ctx, userId, orderId).Execute()

UserOrderCancel



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
    orderId := "orderId_example" // string | This order_id should be filled with the id of the order to be decrease

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserOrderCancel(context.Background(), userId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserOrderCancel`: UserOrderDecreaseResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 
**orderId** | **string** | This order_id should be filled with the id of the order to be decrease | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserOrderDecreaseResponse**](UserOrderDecreaseResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOrderCreate

> UserOrderCreateResponse UserOrderCreate(ctx, userId).UserOrderCreateRequest(userOrderCreateRequest).Execute()

UserOrderCreate



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
    userOrderCreateRequest := *openapiclient.NewUserOrderCreateRequest(int32(123), "MarketId_example", int64(123), "Side_example") // UserOrderCreateRequest | Order create input data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserOrderCreate(context.Background(), userId).UserOrderCreateRequest(userOrderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserOrderCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserOrderCreate`: UserOrderCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserOrderCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userOrderCreateRequest** | [**UserOrderCreateRequest**](UserOrderCreateRequest.md) | Order create input data | 

### Return type

[**UserOrderCreateResponse**](UserOrderCreateResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOrderDecrease

> UserOrderDecreaseResponse UserOrderDecrease(ctx, userId, orderId).UserOrderDecreaseRequest(userOrderDecreaseRequest).Execute()

UserOrderDecrease



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
    orderId := "orderId_example" // string | This order_id should be filled with the id of the order to be decrease
    userOrderDecreaseRequest := *openapiclient.NewUserOrderDecreaseRequest() // UserOrderDecreaseRequest | Order data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserOrderDecrease(context.Background(), userId, orderId).UserOrderDecreaseRequest(userOrderDecreaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserOrderDecrease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserOrderDecrease`: UserOrderDecreaseResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserOrderDecrease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 
**orderId** | **string** | This order_id should be filled with the id of the order to be decrease | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOrderDecreaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userOrderDecreaseRequest** | [**UserOrderDecreaseRequest**](UserOrderDecreaseRequest.md) | Order data | 

### Return type

[**UserOrderDecreaseResponse**](UserOrderDecreaseResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOrdersGet

> UserOrdersGetResponse UserOrdersGet(ctx, userId).MarketId(marketId).IsYes(isYes).MinPrice(minPrice).MaxPrice(maxPrice).MinPlaceCount(minPlaceCount).MaxPlaceCount(maxPlaceCount).MinInitialCount(minInitialCount).MaxInitialCount(maxInitialCount).MinRemainingCount(minRemainingCount).MaxRemainingCount(maxRemainingCount).MinDate(minDate).MaxDate(maxDate).GetQueuePosition(getQueuePosition).Status(status).PageSize(pageSize).PageNumber(pageNumber).Execute()

UserOrdersGet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | This parameter should be filled with your user_id provided on log_in
    marketId := "marketId_example" // string | Restricts the response to orders in a single market (optional)
    isYes := true // bool | Restricts the response to orders in a single direction (yes or no) (optional)
    minPrice := int64(789) // int64 | Restricts the response to orders within a minimum price (optional)
    maxPrice := int64(789) // int64 | Restricts the response to orders within a maximum price (optional)
    minPlaceCount := int32(56) // int32 | Restricts the response to orders within a minimum place count (optional)
    maxPlaceCount := int32(56) // int32 | Restricts the response to orders within a maximum place count (optional)
    minInitialCount := int32(56) // int32 | Restricts the response to orders within a minimum initial count (optional)
    maxInitialCount := int32(56) // int32 | Restricts the response to orders within a maximum initial count (optional)
    minRemainingCount := int32(56) // int32 | Restricts the response to orders within a minimum remaining resting contracts count (optional)
    maxRemainingCount := int32(56) // int32 | Restricts the response to orders within a maximum remaining resting contracts count (optional)
    minDate := time.Now() // time.Time | Restricts the response to orders after a timestamp (optional)
    maxDate := time.Now() // time.Time | Restricts the response to orders before a timestamp (optional)
    getQueuePosition := true // bool | If true, gets the queue placement for every resting order returned (optional)
    status := "status_example" // string | Restricts the response to orders that have a certain status: resting, canceled, or executed (optional)
    pageSize := int64(789) // int64 | Parameter to specify the number of results per page (optional)
    pageNumber := int64(789) // int64 | Parameter to specify which page of the results should be retrieved (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserOrdersGet(context.Background(), userId).MarketId(marketId).IsYes(isYes).MinPrice(minPrice).MaxPrice(maxPrice).MinPlaceCount(minPlaceCount).MaxPlaceCount(maxPlaceCount).MinInitialCount(minInitialCount).MaxInitialCount(maxInitialCount).MinRemainingCount(minRemainingCount).MaxRemainingCount(maxRemainingCount).MinDate(minDate).MaxDate(maxDate).GetQueuePosition(getQueuePosition).Status(status).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserOrdersGet`: UserOrdersGetResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserOrdersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketId** | **string** | Restricts the response to orders in a single market | 
 **isYes** | **bool** | Restricts the response to orders in a single direction (yes or no) | 
 **minPrice** | **int64** | Restricts the response to orders within a minimum price | 
 **maxPrice** | **int64** | Restricts the response to orders within a maximum price | 
 **minPlaceCount** | **int32** | Restricts the response to orders within a minimum place count | 
 **maxPlaceCount** | **int32** | Restricts the response to orders within a maximum place count | 
 **minInitialCount** | **int32** | Restricts the response to orders within a minimum initial count | 
 **maxInitialCount** | **int32** | Restricts the response to orders within a maximum initial count | 
 **minRemainingCount** | **int32** | Restricts the response to orders within a minimum remaining resting contracts count | 
 **maxRemainingCount** | **int32** | Restricts the response to orders within a maximum remaining resting contracts count | 
 **minDate** | **time.Time** | Restricts the response to orders after a timestamp | 
 **maxDate** | **time.Time** | Restricts the response to orders before a timestamp | 
 **getQueuePosition** | **bool** | If true, gets the queue placement for every resting order returned | 
 **status** | **string** | Restricts the response to orders that have a certain status: resting, canceled, or executed | 
 **pageSize** | **int64** | Parameter to specify the number of results per page | 
 **pageNumber** | **int64** | Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserOrdersGetResponse**](UserOrdersGetResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRemoveFavoritedSeries

> UserRemoveFavoritedSeries(ctx, userId, seriesTicker).Execute()

UserRemoveFavoritedSeries



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
    seriesTicker := "seriesTicker_example" // string | series_ticker should be filled with the ticker of the series to be removed from the list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserRemoveFavoritedSeries(context.Background(), userId, seriesTicker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserRemoveFavoritedSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id should be filled with your user_id provided on log_in | 
**seriesTicker** | **string** | series_ticker should be filled with the ticker of the series to be removed from the list | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserRemoveFavoritedSeriesRequest struct via the builder pattern


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


## UserRemoveWatchlist

> UserRemoveWatchlist(ctx, userId, marketId).Execute()

UserRemoveWatchlist



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
    marketId := "marketId_example" // string | Should be filled with the id of the target market

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserRemoveWatchlist(context.Background(), userId, marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserRemoveWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Should be filled with your user_id provided on log_in | 
**marketId** | **string** | Should be filled with the id of the target market | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserRemoveWatchlistRequest struct via the builder pattern


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


## UserTradesGet

> UserTradesGetResponse UserTradesGet(ctx, userId).MarketId(marketId).OrderId(orderId).MinPrice(minPrice).MaxPrice(maxPrice).MinCount(minCount).MaxCount(maxCount).MinDate(minDate).MaxDate(maxDate).PageSize(pageSize).PageNumber(pageNumber).Execute()

UserTradesGet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | This parameter should be filled with your user_id provided on log_in
    marketId := "marketId_example" // string | Restricts the response to trades in a specific market. (optional)
    orderId := "orderId_example" // string | Restricts the response to trades related to a specific order. (optional)
    minPrice := int64(789) // int64 | Restricts the response to trades within a minimum price. (optional)
    maxPrice := int64(789) // int64 | Restricts the response to trades within a maximum price. (optional)
    minCount := int32(56) // int32 | Restricts the response to trades within a minimum contracts count. (optional)
    maxCount := int32(56) // int32 | Restricts the response to trades within a maximum contracts count. (optional)
    minDate := time.Now() // time.Time | Restricts the response to trades after a timestamp. (optional)
    maxDate := time.Now() // time.Time | Restricts the response to trades before a timestamp. (optional)
    pageSize := int64(789) // int64 | Parameter to specify the number of results per page (optional)
    pageNumber := int64(789) // int64 | Parameter to specify which page of the results should be retrieved (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserTradesGet(context.Background(), userId).MarketId(marketId).OrderId(orderId).MinPrice(minPrice).MaxPrice(maxPrice).MinCount(minCount).MaxCount(maxCount).MinDate(minDate).MaxDate(maxDate).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTradesGet`: UserTradesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTradesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | This parameter should be filled with your user_id provided on log_in | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketId** | **string** | Restricts the response to trades in a specific market. | 
 **orderId** | **string** | Restricts the response to trades related to a specific order. | 
 **minPrice** | **int64** | Restricts the response to trades within a minimum price. | 
 **maxPrice** | **int64** | Restricts the response to trades within a maximum price. | 
 **minCount** | **int32** | Restricts the response to trades within a minimum contracts count. | 
 **maxCount** | **int32** | Restricts the response to trades within a maximum contracts count. | 
 **minDate** | **time.Time** | Restricts the response to trades after a timestamp. | 
 **maxDate** | **time.Time** | Restricts the response to trades before a timestamp. | 
 **pageSize** | **int64** | Parameter to specify the number of results per page | 
 **pageNumber** | **int64** | Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserTradesGetResponse**](UserTradesGetResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

