# UserOrderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Specifies how many contracts should be bought | 
**ExpirationUnixTs** | Pointer to **int64** | Specifies the expiration time of the order, in unix seconds.  If this is not supplied, or is 0, the order won&#39;t expire. | [optional] 
**MarketId** | **string** | Specifies the id of the market for this order | 
**MaxCostCents** | Pointer to **int64** |  | [optional] 
**Price** | **int64** |  | 
**SellPositionCapped** | Pointer to **bool** | Specifies whether the order place count should be capped by the members current position. | [optional] 
**Side** | **string** | Specifies if this is a &#39;yes&#39; or &#39;no&#39; order | 

## Methods

### NewUserOrderCreateRequest

`func NewUserOrderCreateRequest(count int32, marketId string, price int64, side string, ) *UserOrderCreateRequest`

NewUserOrderCreateRequest instantiates a new UserOrderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrderCreateRequestWithDefaults

`func NewUserOrderCreateRequestWithDefaults() *UserOrderCreateRequest`

NewUserOrderCreateRequestWithDefaults instantiates a new UserOrderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UserOrderCreateRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserOrderCreateRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserOrderCreateRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetExpirationUnixTs

`func (o *UserOrderCreateRequest) GetExpirationUnixTs() int64`

GetExpirationUnixTs returns the ExpirationUnixTs field if non-nil, zero value otherwise.

### GetExpirationUnixTsOk

`func (o *UserOrderCreateRequest) GetExpirationUnixTsOk() (*int64, bool)`

GetExpirationUnixTsOk returns a tuple with the ExpirationUnixTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationUnixTs

`func (o *UserOrderCreateRequest) SetExpirationUnixTs(v int64)`

SetExpirationUnixTs sets ExpirationUnixTs field to given value.

### HasExpirationUnixTs

`func (o *UserOrderCreateRequest) HasExpirationUnixTs() bool`

HasExpirationUnixTs returns a boolean if a field has been set.

### GetMarketId

`func (o *UserOrderCreateRequest) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *UserOrderCreateRequest) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *UserOrderCreateRequest) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.


### GetMaxCostCents

`func (o *UserOrderCreateRequest) GetMaxCostCents() int64`

GetMaxCostCents returns the MaxCostCents field if non-nil, zero value otherwise.

### GetMaxCostCentsOk

`func (o *UserOrderCreateRequest) GetMaxCostCentsOk() (*int64, bool)`

GetMaxCostCentsOk returns a tuple with the MaxCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCostCents

`func (o *UserOrderCreateRequest) SetMaxCostCents(v int64)`

SetMaxCostCents sets MaxCostCents field to given value.

### HasMaxCostCents

`func (o *UserOrderCreateRequest) HasMaxCostCents() bool`

HasMaxCostCents returns a boolean if a field has been set.

### GetPrice

`func (o *UserOrderCreateRequest) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UserOrderCreateRequest) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UserOrderCreateRequest) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetSellPositionCapped

`func (o *UserOrderCreateRequest) GetSellPositionCapped() bool`

GetSellPositionCapped returns the SellPositionCapped field if non-nil, zero value otherwise.

### GetSellPositionCappedOk

`func (o *UserOrderCreateRequest) GetSellPositionCappedOk() (*bool, bool)`

GetSellPositionCappedOk returns a tuple with the SellPositionCapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPositionCapped

`func (o *UserOrderCreateRequest) SetSellPositionCapped(v bool)`

SetSellPositionCapped sets SellPositionCapped field to given value.

### HasSellPositionCapped

`func (o *UserOrderCreateRequest) HasSellPositionCapped() bool`

HasSellPositionCapped returns a boolean if a field has been set.

### GetSide

`func (o *UserOrderCreateRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UserOrderCreateRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UserOrderCreateRequest) SetSide(v string)`

SetSide sets Side field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


