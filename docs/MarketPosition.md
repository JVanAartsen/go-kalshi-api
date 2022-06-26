# MarketPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeesPaid** | Pointer to **int64** |  | [optional] 
**FinalPosition** | Pointer to **int32** | Settlement stats | [optional] 
**FinalPositionCost** | Pointer to **int64** |  | [optional] 
**IsSettled** | Pointer to **bool** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** | Current stats | [optional] 
**PositionCost** | Pointer to **int64** |  | [optional] 
**RealizedPnl** | Pointer to **int64** |  | [optional] 
**RestingOrdersCount** | Pointer to **int32** |  | [optional] 
**TotalCost** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarketPosition

`func NewMarketPosition() *MarketPosition`

NewMarketPosition instantiates a new MarketPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPositionWithDefaults

`func NewMarketPositionWithDefaults() *MarketPosition`

NewMarketPositionWithDefaults instantiates a new MarketPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeesPaid

`func (o *MarketPosition) GetFeesPaid() int64`

GetFeesPaid returns the FeesPaid field if non-nil, zero value otherwise.

### GetFeesPaidOk

`func (o *MarketPosition) GetFeesPaidOk() (*int64, bool)`

GetFeesPaidOk returns a tuple with the FeesPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesPaid

`func (o *MarketPosition) SetFeesPaid(v int64)`

SetFeesPaid sets FeesPaid field to given value.

### HasFeesPaid

`func (o *MarketPosition) HasFeesPaid() bool`

HasFeesPaid returns a boolean if a field has been set.

### GetFinalPosition

`func (o *MarketPosition) GetFinalPosition() int32`

GetFinalPosition returns the FinalPosition field if non-nil, zero value otherwise.

### GetFinalPositionOk

`func (o *MarketPosition) GetFinalPositionOk() (*int32, bool)`

GetFinalPositionOk returns a tuple with the FinalPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPosition

`func (o *MarketPosition) SetFinalPosition(v int32)`

SetFinalPosition sets FinalPosition field to given value.

### HasFinalPosition

`func (o *MarketPosition) HasFinalPosition() bool`

HasFinalPosition returns a boolean if a field has been set.

### GetFinalPositionCost

`func (o *MarketPosition) GetFinalPositionCost() int64`

GetFinalPositionCost returns the FinalPositionCost field if non-nil, zero value otherwise.

### GetFinalPositionCostOk

`func (o *MarketPosition) GetFinalPositionCostOk() (*int64, bool)`

GetFinalPositionCostOk returns a tuple with the FinalPositionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPositionCost

`func (o *MarketPosition) SetFinalPositionCost(v int64)`

SetFinalPositionCost sets FinalPositionCost field to given value.

### HasFinalPositionCost

`func (o *MarketPosition) HasFinalPositionCost() bool`

HasFinalPositionCost returns a boolean if a field has been set.

### GetIsSettled

`func (o *MarketPosition) GetIsSettled() bool`

GetIsSettled returns the IsSettled field if non-nil, zero value otherwise.

### GetIsSettledOk

`func (o *MarketPosition) GetIsSettledOk() (*bool, bool)`

GetIsSettledOk returns a tuple with the IsSettled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSettled

`func (o *MarketPosition) SetIsSettled(v bool)`

SetIsSettled sets IsSettled field to given value.

### HasIsSettled

`func (o *MarketPosition) HasIsSettled() bool`

HasIsSettled returns a boolean if a field has been set.

### GetMarketId

`func (o *MarketPosition) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *MarketPosition) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *MarketPosition) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *MarketPosition) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetPosition

`func (o *MarketPosition) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MarketPosition) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MarketPosition) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MarketPosition) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionCost

`func (o *MarketPosition) GetPositionCost() int64`

GetPositionCost returns the PositionCost field if non-nil, zero value otherwise.

### GetPositionCostOk

`func (o *MarketPosition) GetPositionCostOk() (*int64, bool)`

GetPositionCostOk returns a tuple with the PositionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCost

`func (o *MarketPosition) SetPositionCost(v int64)`

SetPositionCost sets PositionCost field to given value.

### HasPositionCost

`func (o *MarketPosition) HasPositionCost() bool`

HasPositionCost returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *MarketPosition) GetRealizedPnl() int64`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *MarketPosition) GetRealizedPnlOk() (*int64, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *MarketPosition) SetRealizedPnl(v int64)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *MarketPosition) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetRestingOrdersCount

`func (o *MarketPosition) GetRestingOrdersCount() int32`

GetRestingOrdersCount returns the RestingOrdersCount field if non-nil, zero value otherwise.

### GetRestingOrdersCountOk

`func (o *MarketPosition) GetRestingOrdersCountOk() (*int32, bool)`

GetRestingOrdersCountOk returns a tuple with the RestingOrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestingOrdersCount

`func (o *MarketPosition) SetRestingOrdersCount(v int32)`

SetRestingOrdersCount sets RestingOrdersCount field to given value.

### HasRestingOrdersCount

`func (o *MarketPosition) HasRestingOrdersCount() bool`

HasRestingOrdersCount returns a boolean if a field has been set.

### GetTotalCost

`func (o *MarketPosition) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *MarketPosition) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *MarketPosition) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *MarketPosition) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetUserId

`func (o *MarketPosition) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MarketPosition) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MarketPosition) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MarketPosition) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVolume

`func (o *MarketPosition) GetVolume() int32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketPosition) GetVolumeOk() (*int32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketPosition) SetVolume(v int32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketPosition) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


