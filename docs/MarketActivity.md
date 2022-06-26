# MarketActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseDate** | Pointer to **time.Time** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **time.Time** |  | [optional] 
**PriceChange** | Pointer to **int64** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**VolumeIncrease** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarketActivity

`func NewMarketActivity() *MarketActivity`

NewMarketActivity instantiates a new MarketActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketActivityWithDefaults

`func NewMarketActivityWithDefaults() *MarketActivity`

NewMarketActivityWithDefaults instantiates a new MarketActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *MarketActivity) GetCloseDate() time.Time`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *MarketActivity) GetCloseDateOk() (*time.Time, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *MarketActivity) SetCloseDate(v time.Time)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *MarketActivity) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetMarketId

`func (o *MarketActivity) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *MarketActivity) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *MarketActivity) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *MarketActivity) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetOpenDate

`func (o *MarketActivity) GetOpenDate() time.Time`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *MarketActivity) GetOpenDateOk() (*time.Time, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *MarketActivity) SetOpenDate(v time.Time)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *MarketActivity) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetPriceChange

`func (o *MarketActivity) GetPriceChange() int64`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *MarketActivity) GetPriceChangeOk() (*int64, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *MarketActivity) SetPriceChange(v int64)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *MarketActivity) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetTicker

`func (o *MarketActivity) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *MarketActivity) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *MarketActivity) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *MarketActivity) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetVolumeIncrease

`func (o *MarketActivity) GetVolumeIncrease() int32`

GetVolumeIncrease returns the VolumeIncrease field if non-nil, zero value otherwise.

### GetVolumeIncreaseOk

`func (o *MarketActivity) GetVolumeIncreaseOk() (*int32, bool)`

GetVolumeIncreaseOk returns a tuple with the VolumeIncrease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIncrease

`func (o *MarketActivity) SetVolumeIncrease(v int32)`

SetVolumeIncrease sets VolumeIncrease field to given value.

### HasVolumeIncrease

`func (o *MarketActivity) HasVolumeIncrease() bool`

HasVolumeIncrease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


