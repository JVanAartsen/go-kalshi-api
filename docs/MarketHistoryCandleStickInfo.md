# MarketHistoryCandleStickInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTs** | Pointer to **int64** |  | [optional] 
**OpenInterest** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to [**OLHCM**](OLHCM.md) |  | [optional] 
**StartTs** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to **int64** |  | [optional] 
**YesAsk** | Pointer to [**OLHCM**](OLHCM.md) |  | [optional] 
**YesBid** | Pointer to [**OLHCM**](OLHCM.md) |  | [optional] 

## Methods

### NewMarketHistoryCandleStickInfo

`func NewMarketHistoryCandleStickInfo() *MarketHistoryCandleStickInfo`

NewMarketHistoryCandleStickInfo instantiates a new MarketHistoryCandleStickInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketHistoryCandleStickInfoWithDefaults

`func NewMarketHistoryCandleStickInfoWithDefaults() *MarketHistoryCandleStickInfo`

NewMarketHistoryCandleStickInfoWithDefaults instantiates a new MarketHistoryCandleStickInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTs

`func (o *MarketHistoryCandleStickInfo) GetEndTs() int64`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *MarketHistoryCandleStickInfo) GetEndTsOk() (*int64, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *MarketHistoryCandleStickInfo) SetEndTs(v int64)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *MarketHistoryCandleStickInfo) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOpenInterest

`func (o *MarketHistoryCandleStickInfo) GetOpenInterest() int64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *MarketHistoryCandleStickInfo) GetOpenInterestOk() (*int64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *MarketHistoryCandleStickInfo) SetOpenInterest(v int64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *MarketHistoryCandleStickInfo) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetPrice

`func (o *MarketHistoryCandleStickInfo) GetPrice() OLHCM`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketHistoryCandleStickInfo) GetPriceOk() (*OLHCM, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketHistoryCandleStickInfo) SetPrice(v OLHCM)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketHistoryCandleStickInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStartTs

`func (o *MarketHistoryCandleStickInfo) GetStartTs() int64`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *MarketHistoryCandleStickInfo) GetStartTsOk() (*int64, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *MarketHistoryCandleStickInfo) SetStartTs(v int64)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *MarketHistoryCandleStickInfo) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetVolume

`func (o *MarketHistoryCandleStickInfo) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketHistoryCandleStickInfo) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketHistoryCandleStickInfo) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketHistoryCandleStickInfo) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetYesAsk

`func (o *MarketHistoryCandleStickInfo) GetYesAsk() OLHCM`

GetYesAsk returns the YesAsk field if non-nil, zero value otherwise.

### GetYesAskOk

`func (o *MarketHistoryCandleStickInfo) GetYesAskOk() (*OLHCM, bool)`

GetYesAskOk returns a tuple with the YesAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesAsk

`func (o *MarketHistoryCandleStickInfo) SetYesAsk(v OLHCM)`

SetYesAsk sets YesAsk field to given value.

### HasYesAsk

`func (o *MarketHistoryCandleStickInfo) HasYesAsk() bool`

HasYesAsk returns a boolean if a field has been set.

### GetYesBid

`func (o *MarketHistoryCandleStickInfo) GetYesBid() OLHCM`

GetYesBid returns the YesBid field if non-nil, zero value otherwise.

### GetYesBidOk

`func (o *MarketHistoryCandleStickInfo) GetYesBidOk() (*OLHCM, bool)`

GetYesBidOk returns a tuple with the YesBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesBid

`func (o *MarketHistoryCandleStickInfo) SetYesBid(v OLHCM)`

SetYesBid sets YesBid field to given value.

### HasYesBid

`func (o *MarketHistoryCandleStickInfo) HasYesBid() bool`

HasYesBid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


