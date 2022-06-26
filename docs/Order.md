# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseCancelCount** | Pointer to **int32** |  | [optional] 
**CreateTs** | Pointer to **time.Time** |  | [optional] 
**DecreaseCount** | Pointer to **int32** |  | [optional] 
**ExpirationTs** | Pointer to **time.Time** |  | [optional] 
**ExtraCost** | Pointer to **int64** |  | [optional] 
**ExtraCount** | Pointer to **int32** |  | [optional] 
**FccCancelCount** | Pointer to **int32** |  | [optional] 
**IsYes** | Pointer to **bool** |  | [optional] 
**LastUpdateOp** | Pointer to **string** |  | [optional] 
**MakerFillCount** | Pointer to **int32** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**PlaceCount** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **int64** |  | [optional] 
**QueuePosition** | Pointer to **int32** |  | [optional] 
**RemainingCount** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TakerFees** | Pointer to **int64** |  | [optional] 
**TakerFillCost** | Pointer to **int64** |  | [optional] 
**TakerFillCount** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseCancelCount

`func (o *Order) GetCloseCancelCount() int32`

GetCloseCancelCount returns the CloseCancelCount field if non-nil, zero value otherwise.

### GetCloseCancelCountOk

`func (o *Order) GetCloseCancelCountOk() (*int32, bool)`

GetCloseCancelCountOk returns a tuple with the CloseCancelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseCancelCount

`func (o *Order) SetCloseCancelCount(v int32)`

SetCloseCancelCount sets CloseCancelCount field to given value.

### HasCloseCancelCount

`func (o *Order) HasCloseCancelCount() bool`

HasCloseCancelCount returns a boolean if a field has been set.

### GetCreateTs

`func (o *Order) GetCreateTs() time.Time`

GetCreateTs returns the CreateTs field if non-nil, zero value otherwise.

### GetCreateTsOk

`func (o *Order) GetCreateTsOk() (*time.Time, bool)`

GetCreateTsOk returns a tuple with the CreateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTs

`func (o *Order) SetCreateTs(v time.Time)`

SetCreateTs sets CreateTs field to given value.

### HasCreateTs

`func (o *Order) HasCreateTs() bool`

HasCreateTs returns a boolean if a field has been set.

### GetDecreaseCount

`func (o *Order) GetDecreaseCount() int32`

GetDecreaseCount returns the DecreaseCount field if non-nil, zero value otherwise.

### GetDecreaseCountOk

`func (o *Order) GetDecreaseCountOk() (*int32, bool)`

GetDecreaseCountOk returns a tuple with the DecreaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecreaseCount

`func (o *Order) SetDecreaseCount(v int32)`

SetDecreaseCount sets DecreaseCount field to given value.

### HasDecreaseCount

`func (o *Order) HasDecreaseCount() bool`

HasDecreaseCount returns a boolean if a field has been set.

### GetExpirationTs

`func (o *Order) GetExpirationTs() time.Time`

GetExpirationTs returns the ExpirationTs field if non-nil, zero value otherwise.

### GetExpirationTsOk

`func (o *Order) GetExpirationTsOk() (*time.Time, bool)`

GetExpirationTsOk returns a tuple with the ExpirationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTs

`func (o *Order) SetExpirationTs(v time.Time)`

SetExpirationTs sets ExpirationTs field to given value.

### HasExpirationTs

`func (o *Order) HasExpirationTs() bool`

HasExpirationTs returns a boolean if a field has been set.

### GetExtraCost

`func (o *Order) GetExtraCost() int64`

GetExtraCost returns the ExtraCost field if non-nil, zero value otherwise.

### GetExtraCostOk

`func (o *Order) GetExtraCostOk() (*int64, bool)`

GetExtraCostOk returns a tuple with the ExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCost

`func (o *Order) SetExtraCost(v int64)`

SetExtraCost sets ExtraCost field to given value.

### HasExtraCost

`func (o *Order) HasExtraCost() bool`

HasExtraCost returns a boolean if a field has been set.

### GetExtraCount

`func (o *Order) GetExtraCount() int32`

GetExtraCount returns the ExtraCount field if non-nil, zero value otherwise.

### GetExtraCountOk

`func (o *Order) GetExtraCountOk() (*int32, bool)`

GetExtraCountOk returns a tuple with the ExtraCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCount

`func (o *Order) SetExtraCount(v int32)`

SetExtraCount sets ExtraCount field to given value.

### HasExtraCount

`func (o *Order) HasExtraCount() bool`

HasExtraCount returns a boolean if a field has been set.

### GetFccCancelCount

`func (o *Order) GetFccCancelCount() int32`

GetFccCancelCount returns the FccCancelCount field if non-nil, zero value otherwise.

### GetFccCancelCountOk

`func (o *Order) GetFccCancelCountOk() (*int32, bool)`

GetFccCancelCountOk returns a tuple with the FccCancelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFccCancelCount

`func (o *Order) SetFccCancelCount(v int32)`

SetFccCancelCount sets FccCancelCount field to given value.

### HasFccCancelCount

`func (o *Order) HasFccCancelCount() bool`

HasFccCancelCount returns a boolean if a field has been set.

### GetIsYes

`func (o *Order) GetIsYes() bool`

GetIsYes returns the IsYes field if non-nil, zero value otherwise.

### GetIsYesOk

`func (o *Order) GetIsYesOk() (*bool, bool)`

GetIsYesOk returns a tuple with the IsYes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYes

`func (o *Order) SetIsYes(v bool)`

SetIsYes sets IsYes field to given value.

### HasIsYes

`func (o *Order) HasIsYes() bool`

HasIsYes returns a boolean if a field has been set.

### GetLastUpdateOp

`func (o *Order) GetLastUpdateOp() string`

GetLastUpdateOp returns the LastUpdateOp field if non-nil, zero value otherwise.

### GetLastUpdateOpOk

`func (o *Order) GetLastUpdateOpOk() (*string, bool)`

GetLastUpdateOpOk returns a tuple with the LastUpdateOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateOp

`func (o *Order) SetLastUpdateOp(v string)`

SetLastUpdateOp sets LastUpdateOp field to given value.

### HasLastUpdateOp

`func (o *Order) HasLastUpdateOp() bool`

HasLastUpdateOp returns a boolean if a field has been set.

### GetMakerFillCount

`func (o *Order) GetMakerFillCount() int32`

GetMakerFillCount returns the MakerFillCount field if non-nil, zero value otherwise.

### GetMakerFillCountOk

`func (o *Order) GetMakerFillCountOk() (*int32, bool)`

GetMakerFillCountOk returns a tuple with the MakerFillCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFillCount

`func (o *Order) SetMakerFillCount(v int32)`

SetMakerFillCount sets MakerFillCount field to given value.

### HasMakerFillCount

`func (o *Order) HasMakerFillCount() bool`

HasMakerFillCount returns a boolean if a field has been set.

### GetMarketId

`func (o *Order) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *Order) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *Order) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *Order) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetOrderId

`func (o *Order) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Order) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPlaceCount

`func (o *Order) GetPlaceCount() int32`

GetPlaceCount returns the PlaceCount field if non-nil, zero value otherwise.

### GetPlaceCountOk

`func (o *Order) GetPlaceCountOk() (*int32, bool)`

GetPlaceCountOk returns a tuple with the PlaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceCount

`func (o *Order) SetPlaceCount(v int32)`

SetPlaceCount sets PlaceCount field to given value.

### HasPlaceCount

`func (o *Order) HasPlaceCount() bool`

HasPlaceCount returns a boolean if a field has been set.

### GetPrice

`func (o *Order) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Order) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Order) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Order) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQueuePosition

`func (o *Order) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *Order) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *Order) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.

### HasQueuePosition

`func (o *Order) HasQueuePosition() bool`

HasQueuePosition returns a boolean if a field has been set.

### GetRemainingCount

`func (o *Order) GetRemainingCount() int32`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *Order) GetRemainingCountOk() (*int32, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *Order) SetRemainingCount(v int32)`

SetRemainingCount sets RemainingCount field to given value.

### HasRemainingCount

`func (o *Order) HasRemainingCount() bool`

HasRemainingCount returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTakerFees

`func (o *Order) GetTakerFees() int64`

GetTakerFees returns the TakerFees field if non-nil, zero value otherwise.

### GetTakerFeesOk

`func (o *Order) GetTakerFeesOk() (*int64, bool)`

GetTakerFeesOk returns a tuple with the TakerFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFees

`func (o *Order) SetTakerFees(v int64)`

SetTakerFees sets TakerFees field to given value.

### HasTakerFees

`func (o *Order) HasTakerFees() bool`

HasTakerFees returns a boolean if a field has been set.

### GetTakerFillCost

`func (o *Order) GetTakerFillCost() int64`

GetTakerFillCost returns the TakerFillCost field if non-nil, zero value otherwise.

### GetTakerFillCostOk

`func (o *Order) GetTakerFillCostOk() (*int64, bool)`

GetTakerFillCostOk returns a tuple with the TakerFillCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFillCost

`func (o *Order) SetTakerFillCost(v int64)`

SetTakerFillCost sets TakerFillCost field to given value.

### HasTakerFillCost

`func (o *Order) HasTakerFillCost() bool`

HasTakerFillCost returns a boolean if a field has been set.

### GetTakerFillCount

`func (o *Order) GetTakerFillCount() int32`

GetTakerFillCount returns the TakerFillCount field if non-nil, zero value otherwise.

### GetTakerFillCountOk

`func (o *Order) GetTakerFillCountOk() (*int32, bool)`

GetTakerFillCountOk returns a tuple with the TakerFillCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFillCount

`func (o *Order) SetTakerFillCount(v int32)`

SetTakerFillCount sets TakerFillCount field to given value.

### HasTakerFillCount

`func (o *Order) HasTakerFillCount() bool`

HasTakerFillCount returns a boolean if a field has been set.

### GetUserId

`func (o *Order) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Order) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Order) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Order) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


