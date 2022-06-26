# EventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**DescriptionContext** | Pointer to **string** |  | [optional] 
**Markets** | Pointer to [**[]EventChildMarket**](EventChildMarket.md) |  | [optional] 
**MetricsTags** | Pointer to **[]string** |  | [optional] 
**MinTickSize** | Pointer to **string** |  | [optional] 
**MiniTitle** | Pointer to **string** |  | [optional] 
**MutuallyExclusive** | Pointer to **bool** |  | [optional] 
**MutuallyExclusiveSide** | Pointer to **string** |  | [optional] 
**SeriesTicker** | Pointer to **string** |  | [optional] 
**SettleDetails** | Pointer to **string** |  | [optional] 
**SettlementSources** | Pointer to [**[]SettlementSource**](SettlementSource.md) |  | [optional] 
**SubTitle** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TargetDatetime** | Pointer to **time.Time** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 

## Methods

### NewEventData

`func NewEventData() *EventData`

NewEventData instantiates a new EventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataWithDefaults

`func NewEventDataWithDefaults() *EventData`

NewEventDataWithDefaults instantiates a new EventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *EventData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescriptionContext

`func (o *EventData) GetDescriptionContext() string`

GetDescriptionContext returns the DescriptionContext field if non-nil, zero value otherwise.

### GetDescriptionContextOk

`func (o *EventData) GetDescriptionContextOk() (*string, bool)`

GetDescriptionContextOk returns a tuple with the DescriptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionContext

`func (o *EventData) SetDescriptionContext(v string)`

SetDescriptionContext sets DescriptionContext field to given value.

### HasDescriptionContext

`func (o *EventData) HasDescriptionContext() bool`

HasDescriptionContext returns a boolean if a field has been set.

### GetMarkets

`func (o *EventData) GetMarkets() []EventChildMarket`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *EventData) GetMarketsOk() (*[]EventChildMarket, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *EventData) SetMarkets(v []EventChildMarket)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *EventData) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetMetricsTags

`func (o *EventData) GetMetricsTags() []string`

GetMetricsTags returns the MetricsTags field if non-nil, zero value otherwise.

### GetMetricsTagsOk

`func (o *EventData) GetMetricsTagsOk() (*[]string, bool)`

GetMetricsTagsOk returns a tuple with the MetricsTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsTags

`func (o *EventData) SetMetricsTags(v []string)`

SetMetricsTags sets MetricsTags field to given value.

### HasMetricsTags

`func (o *EventData) HasMetricsTags() bool`

HasMetricsTags returns a boolean if a field has been set.

### GetMinTickSize

`func (o *EventData) GetMinTickSize() string`

GetMinTickSize returns the MinTickSize field if non-nil, zero value otherwise.

### GetMinTickSizeOk

`func (o *EventData) GetMinTickSizeOk() (*string, bool)`

GetMinTickSizeOk returns a tuple with the MinTickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTickSize

`func (o *EventData) SetMinTickSize(v string)`

SetMinTickSize sets MinTickSize field to given value.

### HasMinTickSize

`func (o *EventData) HasMinTickSize() bool`

HasMinTickSize returns a boolean if a field has been set.

### GetMiniTitle

`func (o *EventData) GetMiniTitle() string`

GetMiniTitle returns the MiniTitle field if non-nil, zero value otherwise.

### GetMiniTitleOk

`func (o *EventData) GetMiniTitleOk() (*string, bool)`

GetMiniTitleOk returns a tuple with the MiniTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniTitle

`func (o *EventData) SetMiniTitle(v string)`

SetMiniTitle sets MiniTitle field to given value.

### HasMiniTitle

`func (o *EventData) HasMiniTitle() bool`

HasMiniTitle returns a boolean if a field has been set.

### GetMutuallyExclusive

`func (o *EventData) GetMutuallyExclusive() bool`

GetMutuallyExclusive returns the MutuallyExclusive field if non-nil, zero value otherwise.

### GetMutuallyExclusiveOk

`func (o *EventData) GetMutuallyExclusiveOk() (*bool, bool)`

GetMutuallyExclusiveOk returns a tuple with the MutuallyExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutuallyExclusive

`func (o *EventData) SetMutuallyExclusive(v bool)`

SetMutuallyExclusive sets MutuallyExclusive field to given value.

### HasMutuallyExclusive

`func (o *EventData) HasMutuallyExclusive() bool`

HasMutuallyExclusive returns a boolean if a field has been set.

### GetMutuallyExclusiveSide

`func (o *EventData) GetMutuallyExclusiveSide() string`

GetMutuallyExclusiveSide returns the MutuallyExclusiveSide field if non-nil, zero value otherwise.

### GetMutuallyExclusiveSideOk

`func (o *EventData) GetMutuallyExclusiveSideOk() (*string, bool)`

GetMutuallyExclusiveSideOk returns a tuple with the MutuallyExclusiveSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutuallyExclusiveSide

`func (o *EventData) SetMutuallyExclusiveSide(v string)`

SetMutuallyExclusiveSide sets MutuallyExclusiveSide field to given value.

### HasMutuallyExclusiveSide

`func (o *EventData) HasMutuallyExclusiveSide() bool`

HasMutuallyExclusiveSide returns a boolean if a field has been set.

### GetSeriesTicker

`func (o *EventData) GetSeriesTicker() string`

GetSeriesTicker returns the SeriesTicker field if non-nil, zero value otherwise.

### GetSeriesTickerOk

`func (o *EventData) GetSeriesTickerOk() (*string, bool)`

GetSeriesTickerOk returns a tuple with the SeriesTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTicker

`func (o *EventData) SetSeriesTicker(v string)`

SetSeriesTicker sets SeriesTicker field to given value.

### HasSeriesTicker

`func (o *EventData) HasSeriesTicker() bool`

HasSeriesTicker returns a boolean if a field has been set.

### GetSettleDetails

`func (o *EventData) GetSettleDetails() string`

GetSettleDetails returns the SettleDetails field if non-nil, zero value otherwise.

### GetSettleDetailsOk

`func (o *EventData) GetSettleDetailsOk() (*string, bool)`

GetSettleDetailsOk returns a tuple with the SettleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDetails

`func (o *EventData) SetSettleDetails(v string)`

SetSettleDetails sets SettleDetails field to given value.

### HasSettleDetails

`func (o *EventData) HasSettleDetails() bool`

HasSettleDetails returns a boolean if a field has been set.

### GetSettlementSources

`func (o *EventData) GetSettlementSources() []SettlementSource`

GetSettlementSources returns the SettlementSources field if non-nil, zero value otherwise.

### GetSettlementSourcesOk

`func (o *EventData) GetSettlementSourcesOk() (*[]SettlementSource, bool)`

GetSettlementSourcesOk returns a tuple with the SettlementSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementSources

`func (o *EventData) SetSettlementSources(v []SettlementSource)`

SetSettlementSources sets SettlementSources field to given value.

### HasSettlementSources

`func (o *EventData) HasSettlementSources() bool`

HasSettlementSources returns a boolean if a field has been set.

### GetSubTitle

`func (o *EventData) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *EventData) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *EventData) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.

### HasSubTitle

`func (o *EventData) HasSubTitle() bool`

HasSubTitle returns a boolean if a field has been set.

### GetTags

`func (o *EventData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetDatetime

`func (o *EventData) GetTargetDatetime() time.Time`

GetTargetDatetime returns the TargetDatetime field if non-nil, zero value otherwise.

### GetTargetDatetimeOk

`func (o *EventData) GetTargetDatetimeOk() (*time.Time, bool)`

GetTargetDatetimeOk returns a tuple with the TargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatetime

`func (o *EventData) SetTargetDatetime(v time.Time)`

SetTargetDatetime sets TargetDatetime field to given value.

### HasTargetDatetime

`func (o *EventData) HasTargetDatetime() bool`

HasTargetDatetime returns a boolean if a field has been set.

### GetTicker

`func (o *EventData) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *EventData) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *EventData) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *EventData) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTitle

`func (o *EventData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnderlying

`func (o *EventData) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *EventData) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *EventData) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *EventData) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


