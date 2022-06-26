# OrderHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledCount** | Pointer to **int64** |  | [optional] 
**CloseCancelCount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FccCanceledCount** | Pointer to **int64** |  | [optional] 
**FilledCount** | Pointer to **int64** |  | [optional] 
**IsYes** | Pointer to **bool** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**OriginalCount** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **int64** |  | [optional] 
**RemainingCount** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrderHistory

`func NewOrderHistory() *OrderHistory`

NewOrderHistory instantiates a new OrderHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderHistoryWithDefaults

`func NewOrderHistoryWithDefaults() *OrderHistory`

NewOrderHistoryWithDefaults instantiates a new OrderHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledCount

`func (o *OrderHistory) GetCanceledCount() int64`

GetCanceledCount returns the CanceledCount field if non-nil, zero value otherwise.

### GetCanceledCountOk

`func (o *OrderHistory) GetCanceledCountOk() (*int64, bool)`

GetCanceledCountOk returns a tuple with the CanceledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledCount

`func (o *OrderHistory) SetCanceledCount(v int64)`

SetCanceledCount sets CanceledCount field to given value.

### HasCanceledCount

`func (o *OrderHistory) HasCanceledCount() bool`

HasCanceledCount returns a boolean if a field has been set.

### GetCloseCancelCount

`func (o *OrderHistory) GetCloseCancelCount() int64`

GetCloseCancelCount returns the CloseCancelCount field if non-nil, zero value otherwise.

### GetCloseCancelCountOk

`func (o *OrderHistory) GetCloseCancelCountOk() (*int64, bool)`

GetCloseCancelCountOk returns a tuple with the CloseCancelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseCancelCount

`func (o *OrderHistory) SetCloseCancelCount(v int64)`

SetCloseCancelCount sets CloseCancelCount field to given value.

### HasCloseCancelCount

`func (o *OrderHistory) HasCloseCancelCount() bool`

HasCloseCancelCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFccCanceledCount

`func (o *OrderHistory) GetFccCanceledCount() int64`

GetFccCanceledCount returns the FccCanceledCount field if non-nil, zero value otherwise.

### GetFccCanceledCountOk

`func (o *OrderHistory) GetFccCanceledCountOk() (*int64, bool)`

GetFccCanceledCountOk returns a tuple with the FccCanceledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFccCanceledCount

`func (o *OrderHistory) SetFccCanceledCount(v int64)`

SetFccCanceledCount sets FccCanceledCount field to given value.

### HasFccCanceledCount

`func (o *OrderHistory) HasFccCanceledCount() bool`

HasFccCanceledCount returns a boolean if a field has been set.

### GetFilledCount

`func (o *OrderHistory) GetFilledCount() int64`

GetFilledCount returns the FilledCount field if non-nil, zero value otherwise.

### GetFilledCountOk

`func (o *OrderHistory) GetFilledCountOk() (*int64, bool)`

GetFilledCountOk returns a tuple with the FilledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCount

`func (o *OrderHistory) SetFilledCount(v int64)`

SetFilledCount sets FilledCount field to given value.

### HasFilledCount

`func (o *OrderHistory) HasFilledCount() bool`

HasFilledCount returns a boolean if a field has been set.

### GetIsYes

`func (o *OrderHistory) GetIsYes() bool`

GetIsYes returns the IsYes field if non-nil, zero value otherwise.

### GetIsYesOk

`func (o *OrderHistory) GetIsYesOk() (*bool, bool)`

GetIsYesOk returns a tuple with the IsYes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYes

`func (o *OrderHistory) SetIsYes(v bool)`

SetIsYes sets IsYes field to given value.

### HasIsYes

`func (o *OrderHistory) HasIsYes() bool`

HasIsYes returns a boolean if a field has been set.

### GetMarketId

`func (o *OrderHistory) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *OrderHistory) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *OrderHistory) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *OrderHistory) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTitle

`func (o *OrderHistory) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *OrderHistory) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *OrderHistory) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *OrderHistory) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetOriginalCount

`func (o *OrderHistory) GetOriginalCount() int64`

GetOriginalCount returns the OriginalCount field if non-nil, zero value otherwise.

### GetOriginalCountOk

`func (o *OrderHistory) GetOriginalCountOk() (*int64, bool)`

GetOriginalCountOk returns a tuple with the OriginalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCount

`func (o *OrderHistory) SetOriginalCount(v int64)`

SetOriginalCount sets OriginalCount field to given value.

### HasOriginalCount

`func (o *OrderHistory) HasOriginalCount() bool`

HasOriginalCount returns a boolean if a field has been set.

### GetPrice

`func (o *OrderHistory) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderHistory) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderHistory) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderHistory) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRemainingCount

`func (o *OrderHistory) GetRemainingCount() int64`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *OrderHistory) GetRemainingCountOk() (*int64, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *OrderHistory) SetRemainingCount(v int64)`

SetRemainingCount sets RemainingCount field to given value.

### HasRemainingCount

`func (o *OrderHistory) HasRemainingCount() bool`

HasRemainingCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderHistory) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderHistory) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderHistory) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderHistory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


