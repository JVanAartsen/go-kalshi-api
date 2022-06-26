# RangedMarketPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollateralReturnedCents** | Pointer to **int64** |  | [optional] 
**MarketPositions** | Pointer to [**[]MarketPosition**](MarketPosition.md) |  | [optional] 
**RangedMarketId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewRangedMarketPosition

`func NewRangedMarketPosition() *RangedMarketPosition`

NewRangedMarketPosition instantiates a new RangedMarketPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangedMarketPositionWithDefaults

`func NewRangedMarketPositionWithDefaults() *RangedMarketPosition`

NewRangedMarketPositionWithDefaults instantiates a new RangedMarketPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollateralReturnedCents

`func (o *RangedMarketPosition) GetCollateralReturnedCents() int64`

GetCollateralReturnedCents returns the CollateralReturnedCents field if non-nil, zero value otherwise.

### GetCollateralReturnedCentsOk

`func (o *RangedMarketPosition) GetCollateralReturnedCentsOk() (*int64, bool)`

GetCollateralReturnedCentsOk returns a tuple with the CollateralReturnedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralReturnedCents

`func (o *RangedMarketPosition) SetCollateralReturnedCents(v int64)`

SetCollateralReturnedCents sets CollateralReturnedCents field to given value.

### HasCollateralReturnedCents

`func (o *RangedMarketPosition) HasCollateralReturnedCents() bool`

HasCollateralReturnedCents returns a boolean if a field has been set.

### GetMarketPositions

`func (o *RangedMarketPosition) GetMarketPositions() []MarketPosition`

GetMarketPositions returns the MarketPositions field if non-nil, zero value otherwise.

### GetMarketPositionsOk

`func (o *RangedMarketPosition) GetMarketPositionsOk() (*[]MarketPosition, bool)`

GetMarketPositionsOk returns a tuple with the MarketPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketPositions

`func (o *RangedMarketPosition) SetMarketPositions(v []MarketPosition)`

SetMarketPositions sets MarketPositions field to given value.

### HasMarketPositions

`func (o *RangedMarketPosition) HasMarketPositions() bool`

HasMarketPositions returns a boolean if a field has been set.

### GetRangedMarketId

`func (o *RangedMarketPosition) GetRangedMarketId() string`

GetRangedMarketId returns the RangedMarketId field if non-nil, zero value otherwise.

### GetRangedMarketIdOk

`func (o *RangedMarketPosition) GetRangedMarketIdOk() (*string, bool)`

GetRangedMarketIdOk returns a tuple with the RangedMarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedMarketId

`func (o *RangedMarketPosition) SetRangedMarketId(v string)`

SetRangedMarketId sets RangedMarketId field to given value.

### HasRangedMarketId

`func (o *RangedMarketPosition) HasRangedMarketId() bool`

HasRangedMarketId returns a boolean if a field has been set.

### GetUserId

`func (o *RangedMarketPosition) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RangedMarketPosition) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RangedMarketPosition) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RangedMarketPosition) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


