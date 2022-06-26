# Market

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCloseEarly** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CloseDate** | Pointer to **time.Time** |  | [optional] 
**CloseUnconfirmed** | Pointer to **bool** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**DescriptionCaseNo** | Pointer to **string** |  | [optional] 
**DescriptionCaseYes** | Pointer to **string** |  | [optional] 
**DescriptionContext** | Pointer to **string** |  | [optional] 
**DollarOpenInterest** | Pointer to **int64** |  | [optional] 
**DollarRecentVolume** | Pointer to **int64** |  | [optional] 
**DollarVolume** | Pointer to **int64** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationValue** | Pointer to **string** |  | [optional] 
**FrequencyInDays** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **int64** |  | [optional] 
**ListDate** | Pointer to **time.Time** |  | [optional] 
**MetricsTags** | Pointer to **[]string** |  | [optional] 
**MinTickSize** | Pointer to **string** |  | [optional] 
**MiniTitle** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **time.Time** |  | [optional] 
**OpenInterest** | Pointer to **int64** |  | [optional] 
**OptionType** | Pointer to **string** |  | [optional] 
**OriginalExpirationDate** | Pointer to **time.Time** |  | [optional] 
**PreviousPrice** | Pointer to **int64** |  | [optional] 
**PreviousYesAsk** | Pointer to **int64** |  | [optional] 
**PreviousYesBid** | Pointer to **int64** |  | [optional] 
**RangedGroupName** | Pointer to **string** |  | [optional] 
**RangedGroupTicker** | Pointer to **string** |  | [optional] 
**RecentVolume** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**RulebookVariables** | Pointer to **map[string]string** | You should not assume a fixed structure for this field. It is subject to change from market to market. | [optional] 
**SettleDetails** | Pointer to **string** |  | [optional] 
**SettlementSources** | Pointer to [**[]SettlementSource**](SettlementSource.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **float64** |  | [optional] 
**SubTitle** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TickerName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **int64** |  | [optional] 
**YesAsk** | Pointer to **int64** |  | [optional] 
**YesBid** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarket

`func NewMarket() *Market`

NewMarket instantiates a new Market object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketWithDefaults

`func NewMarketWithDefaults() *Market`

NewMarketWithDefaults instantiates a new Market object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCloseEarly

`func (o *Market) GetCanCloseEarly() bool`

GetCanCloseEarly returns the CanCloseEarly field if non-nil, zero value otherwise.

### GetCanCloseEarlyOk

`func (o *Market) GetCanCloseEarlyOk() (*bool, bool)`

GetCanCloseEarlyOk returns a tuple with the CanCloseEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCloseEarly

`func (o *Market) SetCanCloseEarly(v bool)`

SetCanCloseEarly sets CanCloseEarly field to given value.

### HasCanCloseEarly

`func (o *Market) HasCanCloseEarly() bool`

HasCanCloseEarly returns a boolean if a field has been set.

### GetCategory

`func (o *Market) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Market) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Market) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Market) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloseDate

`func (o *Market) GetCloseDate() time.Time`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *Market) GetCloseDateOk() (*time.Time, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *Market) SetCloseDate(v time.Time)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *Market) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCloseUnconfirmed

`func (o *Market) GetCloseUnconfirmed() bool`

GetCloseUnconfirmed returns the CloseUnconfirmed field if non-nil, zero value otherwise.

### GetCloseUnconfirmedOk

`func (o *Market) GetCloseUnconfirmedOk() (*bool, bool)`

GetCloseUnconfirmedOk returns a tuple with the CloseUnconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseUnconfirmed

`func (o *Market) SetCloseUnconfirmed(v bool)`

SetCloseUnconfirmed sets CloseUnconfirmed field to given value.

### HasCloseUnconfirmed

`func (o *Market) HasCloseUnconfirmed() bool`

HasCloseUnconfirmed returns a boolean if a field has been set.

### GetCreateDate

`func (o *Market) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Market) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Market) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Market) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDescriptionCaseNo

`func (o *Market) GetDescriptionCaseNo() string`

GetDescriptionCaseNo returns the DescriptionCaseNo field if non-nil, zero value otherwise.

### GetDescriptionCaseNoOk

`func (o *Market) GetDescriptionCaseNoOk() (*string, bool)`

GetDescriptionCaseNoOk returns a tuple with the DescriptionCaseNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionCaseNo

`func (o *Market) SetDescriptionCaseNo(v string)`

SetDescriptionCaseNo sets DescriptionCaseNo field to given value.

### HasDescriptionCaseNo

`func (o *Market) HasDescriptionCaseNo() bool`

HasDescriptionCaseNo returns a boolean if a field has been set.

### GetDescriptionCaseYes

`func (o *Market) GetDescriptionCaseYes() string`

GetDescriptionCaseYes returns the DescriptionCaseYes field if non-nil, zero value otherwise.

### GetDescriptionCaseYesOk

`func (o *Market) GetDescriptionCaseYesOk() (*string, bool)`

GetDescriptionCaseYesOk returns a tuple with the DescriptionCaseYes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionCaseYes

`func (o *Market) SetDescriptionCaseYes(v string)`

SetDescriptionCaseYes sets DescriptionCaseYes field to given value.

### HasDescriptionCaseYes

`func (o *Market) HasDescriptionCaseYes() bool`

HasDescriptionCaseYes returns a boolean if a field has been set.

### GetDescriptionContext

`func (o *Market) GetDescriptionContext() string`

GetDescriptionContext returns the DescriptionContext field if non-nil, zero value otherwise.

### GetDescriptionContextOk

`func (o *Market) GetDescriptionContextOk() (*string, bool)`

GetDescriptionContextOk returns a tuple with the DescriptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionContext

`func (o *Market) SetDescriptionContext(v string)`

SetDescriptionContext sets DescriptionContext field to given value.

### HasDescriptionContext

`func (o *Market) HasDescriptionContext() bool`

HasDescriptionContext returns a boolean if a field has been set.

### GetDollarOpenInterest

`func (o *Market) GetDollarOpenInterest() int64`

GetDollarOpenInterest returns the DollarOpenInterest field if non-nil, zero value otherwise.

### GetDollarOpenInterestOk

`func (o *Market) GetDollarOpenInterestOk() (*int64, bool)`

GetDollarOpenInterestOk returns a tuple with the DollarOpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarOpenInterest

`func (o *Market) SetDollarOpenInterest(v int64)`

SetDollarOpenInterest sets DollarOpenInterest field to given value.

### HasDollarOpenInterest

`func (o *Market) HasDollarOpenInterest() bool`

HasDollarOpenInterest returns a boolean if a field has been set.

### GetDollarRecentVolume

`func (o *Market) GetDollarRecentVolume() int64`

GetDollarRecentVolume returns the DollarRecentVolume field if non-nil, zero value otherwise.

### GetDollarRecentVolumeOk

`func (o *Market) GetDollarRecentVolumeOk() (*int64, bool)`

GetDollarRecentVolumeOk returns a tuple with the DollarRecentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarRecentVolume

`func (o *Market) SetDollarRecentVolume(v int64)`

SetDollarRecentVolume sets DollarRecentVolume field to given value.

### HasDollarRecentVolume

`func (o *Market) HasDollarRecentVolume() bool`

HasDollarRecentVolume returns a boolean if a field has been set.

### GetDollarVolume

`func (o *Market) GetDollarVolume() int64`

GetDollarVolume returns the DollarVolume field if non-nil, zero value otherwise.

### GetDollarVolumeOk

`func (o *Market) GetDollarVolumeOk() (*int64, bool)`

GetDollarVolumeOk returns a tuple with the DollarVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarVolume

`func (o *Market) SetDollarVolume(v int64)`

SetDollarVolume sets DollarVolume field to given value.

### HasDollarVolume

`func (o *Market) HasDollarVolume() bool`

HasDollarVolume returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Market) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Market) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Market) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Market) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpirationValue

`func (o *Market) GetExpirationValue() string`

GetExpirationValue returns the ExpirationValue field if non-nil, zero value otherwise.

### GetExpirationValueOk

`func (o *Market) GetExpirationValueOk() (*string, bool)`

GetExpirationValueOk returns a tuple with the ExpirationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationValue

`func (o *Market) SetExpirationValue(v string)`

SetExpirationValue sets ExpirationValue field to given value.

### HasExpirationValue

`func (o *Market) HasExpirationValue() bool`

HasExpirationValue returns a boolean if a field has been set.

### GetFrequencyInDays

`func (o *Market) GetFrequencyInDays() int32`

GetFrequencyInDays returns the FrequencyInDays field if non-nil, zero value otherwise.

### GetFrequencyInDaysOk

`func (o *Market) GetFrequencyInDaysOk() (*int32, bool)`

GetFrequencyInDaysOk returns a tuple with the FrequencyInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyInDays

`func (o *Market) SetFrequencyInDays(v int32)`

SetFrequencyInDays sets FrequencyInDays field to given value.

### HasFrequencyInDays

`func (o *Market) HasFrequencyInDays() bool`

HasFrequencyInDays returns a boolean if a field has been set.

### GetId

`func (o *Market) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Market) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Market) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Market) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *Market) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Market) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Market) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Market) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLastPrice

`func (o *Market) GetLastPrice() int64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *Market) GetLastPriceOk() (*int64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *Market) SetLastPrice(v int64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *Market) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLiquidity

`func (o *Market) GetLiquidity() int64`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *Market) GetLiquidityOk() (*int64, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *Market) SetLiquidity(v int64)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *Market) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetListDate

`func (o *Market) GetListDate() time.Time`

GetListDate returns the ListDate field if non-nil, zero value otherwise.

### GetListDateOk

`func (o *Market) GetListDateOk() (*time.Time, bool)`

GetListDateOk returns a tuple with the ListDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListDate

`func (o *Market) SetListDate(v time.Time)`

SetListDate sets ListDate field to given value.

### HasListDate

`func (o *Market) HasListDate() bool`

HasListDate returns a boolean if a field has been set.

### GetMetricsTags

`func (o *Market) GetMetricsTags() []string`

GetMetricsTags returns the MetricsTags field if non-nil, zero value otherwise.

### GetMetricsTagsOk

`func (o *Market) GetMetricsTagsOk() (*[]string, bool)`

GetMetricsTagsOk returns a tuple with the MetricsTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsTags

`func (o *Market) SetMetricsTags(v []string)`

SetMetricsTags sets MetricsTags field to given value.

### HasMetricsTags

`func (o *Market) HasMetricsTags() bool`

HasMetricsTags returns a boolean if a field has been set.

### GetMinTickSize

`func (o *Market) GetMinTickSize() string`

GetMinTickSize returns the MinTickSize field if non-nil, zero value otherwise.

### GetMinTickSizeOk

`func (o *Market) GetMinTickSizeOk() (*string, bool)`

GetMinTickSizeOk returns a tuple with the MinTickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTickSize

`func (o *Market) SetMinTickSize(v string)`

SetMinTickSize sets MinTickSize field to given value.

### HasMinTickSize

`func (o *Market) HasMinTickSize() bool`

HasMinTickSize returns a boolean if a field has been set.

### GetMiniTitle

`func (o *Market) GetMiniTitle() string`

GetMiniTitle returns the MiniTitle field if non-nil, zero value otherwise.

### GetMiniTitleOk

`func (o *Market) GetMiniTitleOk() (*string, bool)`

GetMiniTitleOk returns a tuple with the MiniTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniTitle

`func (o *Market) SetMiniTitle(v string)`

SetMiniTitle sets MiniTitle field to given value.

### HasMiniTitle

`func (o *Market) HasMiniTitle() bool`

HasMiniTitle returns a boolean if a field has been set.

### GetOpenDate

`func (o *Market) GetOpenDate() time.Time`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *Market) GetOpenDateOk() (*time.Time, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *Market) SetOpenDate(v time.Time)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *Market) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetOpenInterest

`func (o *Market) GetOpenInterest() int64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *Market) GetOpenInterestOk() (*int64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *Market) SetOpenInterest(v int64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *Market) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOptionType

`func (o *Market) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *Market) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *Market) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *Market) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetOriginalExpirationDate

`func (o *Market) GetOriginalExpirationDate() time.Time`

GetOriginalExpirationDate returns the OriginalExpirationDate field if non-nil, zero value otherwise.

### GetOriginalExpirationDateOk

`func (o *Market) GetOriginalExpirationDateOk() (*time.Time, bool)`

GetOriginalExpirationDateOk returns a tuple with the OriginalExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExpirationDate

`func (o *Market) SetOriginalExpirationDate(v time.Time)`

SetOriginalExpirationDate sets OriginalExpirationDate field to given value.

### HasOriginalExpirationDate

`func (o *Market) HasOriginalExpirationDate() bool`

HasOriginalExpirationDate returns a boolean if a field has been set.

### GetPreviousPrice

`func (o *Market) GetPreviousPrice() int64`

GetPreviousPrice returns the PreviousPrice field if non-nil, zero value otherwise.

### GetPreviousPriceOk

`func (o *Market) GetPreviousPriceOk() (*int64, bool)`

GetPreviousPriceOk returns a tuple with the PreviousPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPrice

`func (o *Market) SetPreviousPrice(v int64)`

SetPreviousPrice sets PreviousPrice field to given value.

### HasPreviousPrice

`func (o *Market) HasPreviousPrice() bool`

HasPreviousPrice returns a boolean if a field has been set.

### GetPreviousYesAsk

`func (o *Market) GetPreviousYesAsk() int64`

GetPreviousYesAsk returns the PreviousYesAsk field if non-nil, zero value otherwise.

### GetPreviousYesAskOk

`func (o *Market) GetPreviousYesAskOk() (*int64, bool)`

GetPreviousYesAskOk returns a tuple with the PreviousYesAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYesAsk

`func (o *Market) SetPreviousYesAsk(v int64)`

SetPreviousYesAsk sets PreviousYesAsk field to given value.

### HasPreviousYesAsk

`func (o *Market) HasPreviousYesAsk() bool`

HasPreviousYesAsk returns a boolean if a field has been set.

### GetPreviousYesBid

`func (o *Market) GetPreviousYesBid() int64`

GetPreviousYesBid returns the PreviousYesBid field if non-nil, zero value otherwise.

### GetPreviousYesBidOk

`func (o *Market) GetPreviousYesBidOk() (*int64, bool)`

GetPreviousYesBidOk returns a tuple with the PreviousYesBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYesBid

`func (o *Market) SetPreviousYesBid(v int64)`

SetPreviousYesBid sets PreviousYesBid field to given value.

### HasPreviousYesBid

`func (o *Market) HasPreviousYesBid() bool`

HasPreviousYesBid returns a boolean if a field has been set.

### GetRangedGroupName

`func (o *Market) GetRangedGroupName() string`

GetRangedGroupName returns the RangedGroupName field if non-nil, zero value otherwise.

### GetRangedGroupNameOk

`func (o *Market) GetRangedGroupNameOk() (*string, bool)`

GetRangedGroupNameOk returns a tuple with the RangedGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedGroupName

`func (o *Market) SetRangedGroupName(v string)`

SetRangedGroupName sets RangedGroupName field to given value.

### HasRangedGroupName

`func (o *Market) HasRangedGroupName() bool`

HasRangedGroupName returns a boolean if a field has been set.

### GetRangedGroupTicker

`func (o *Market) GetRangedGroupTicker() string`

GetRangedGroupTicker returns the RangedGroupTicker field if non-nil, zero value otherwise.

### GetRangedGroupTickerOk

`func (o *Market) GetRangedGroupTickerOk() (*string, bool)`

GetRangedGroupTickerOk returns a tuple with the RangedGroupTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangedGroupTicker

`func (o *Market) SetRangedGroupTicker(v string)`

SetRangedGroupTicker sets RangedGroupTicker field to given value.

### HasRangedGroupTicker

`func (o *Market) HasRangedGroupTicker() bool`

HasRangedGroupTicker returns a boolean if a field has been set.

### GetRecentVolume

`func (o *Market) GetRecentVolume() int64`

GetRecentVolume returns the RecentVolume field if non-nil, zero value otherwise.

### GetRecentVolumeOk

`func (o *Market) GetRecentVolumeOk() (*int64, bool)`

GetRecentVolumeOk returns a tuple with the RecentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVolume

`func (o *Market) SetRecentVolume(v int64)`

SetRecentVolume sets RecentVolume field to given value.

### HasRecentVolume

`func (o *Market) HasRecentVolume() bool`

HasRecentVolume returns a boolean if a field has been set.

### GetResult

`func (o *Market) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Market) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Market) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Market) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRulebookVariables

`func (o *Market) GetRulebookVariables() map[string]string`

GetRulebookVariables returns the RulebookVariables field if non-nil, zero value otherwise.

### GetRulebookVariablesOk

`func (o *Market) GetRulebookVariablesOk() (*map[string]string, bool)`

GetRulebookVariablesOk returns a tuple with the RulebookVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulebookVariables

`func (o *Market) SetRulebookVariables(v map[string]string)`

SetRulebookVariables sets RulebookVariables field to given value.

### HasRulebookVariables

`func (o *Market) HasRulebookVariables() bool`

HasRulebookVariables returns a boolean if a field has been set.

### GetSettleDetails

`func (o *Market) GetSettleDetails() string`

GetSettleDetails returns the SettleDetails field if non-nil, zero value otherwise.

### GetSettleDetailsOk

`func (o *Market) GetSettleDetailsOk() (*string, bool)`

GetSettleDetailsOk returns a tuple with the SettleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDetails

`func (o *Market) SetSettleDetails(v string)`

SetSettleDetails sets SettleDetails field to given value.

### HasSettleDetails

`func (o *Market) HasSettleDetails() bool`

HasSettleDetails returns a boolean if a field has been set.

### GetSettlementSources

`func (o *Market) GetSettlementSources() []SettlementSource`

GetSettlementSources returns the SettlementSources field if non-nil, zero value otherwise.

### GetSettlementSourcesOk

`func (o *Market) GetSettlementSourcesOk() (*[]SettlementSource, bool)`

GetSettlementSourcesOk returns a tuple with the SettlementSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementSources

`func (o *Market) SetSettlementSources(v []SettlementSource)`

SetSettlementSources sets SettlementSources field to given value.

### HasSettlementSources

`func (o *Market) HasSettlementSources() bool`

HasSettlementSources returns a boolean if a field has been set.

### GetStatus

`func (o *Market) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Market) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Market) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Market) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStrikePrice

`func (o *Market) GetStrikePrice() float64`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *Market) GetStrikePriceOk() (*float64, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *Market) SetStrikePrice(v float64)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *Market) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetSubTitle

`func (o *Market) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *Market) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *Market) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.

### HasSubTitle

`func (o *Market) HasSubTitle() bool`

HasSubTitle returns a boolean if a field has been set.

### GetTags

`func (o *Market) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Market) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Market) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Market) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTickerName

`func (o *Market) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *Market) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *Market) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *Market) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTitle

`func (o *Market) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Market) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Market) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Market) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnderlying

`func (o *Market) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *Market) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *Market) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *Market) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetVolume

`func (o *Market) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Market) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Market) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Market) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetYesAsk

`func (o *Market) GetYesAsk() int64`

GetYesAsk returns the YesAsk field if non-nil, zero value otherwise.

### GetYesAskOk

`func (o *Market) GetYesAskOk() (*int64, bool)`

GetYesAskOk returns a tuple with the YesAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesAsk

`func (o *Market) SetYesAsk(v int64)`

SetYesAsk sets YesAsk field to given value.

### HasYesAsk

`func (o *Market) HasYesAsk() bool`

HasYesAsk returns a boolean if a field has been set.

### GetYesBid

`func (o *Market) GetYesBid() int64`

GetYesBid returns the YesBid field if non-nil, zero value otherwise.

### GetYesBidOk

`func (o *Market) GetYesBidOk() (*int64, bool)`

GetYesBidOk returns a tuple with the YesBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesBid

`func (o *Market) SetYesBid(v int64)`

SetYesBid sets YesBid field to given value.

### HasYesBid

`func (o *Market) HasYesBid() bool`

HasYesBid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


