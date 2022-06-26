# EventChildMarket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCloseEarly** | Pointer to **bool** |  | [optional] 
**CloseDate** | Pointer to **time.Time** |  | [optional] 
**CloseUnconfirmed** | Pointer to **bool** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**DollarOpenInterest** | Pointer to **int64** |  | [optional] 
**DollarRecentVolume** | Pointer to **int64** |  | [optional] 
**DollarVolume** | Pointer to **int64** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationValue** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **int64** |  | [optional] 
**ListDate** | Pointer to **time.Time** |  | [optional] 
**MiniTitle** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **time.Time** |  | [optional] 
**OpenInterest** | Pointer to **int64** |  | [optional] 
**PreviousPrice** | Pointer to **int64** |  | [optional] 
**PreviousYesAsk** | Pointer to **int64** |  | [optional] 
**PreviousYesBid** | Pointer to **int64** |  | [optional] 
**RecentVolume** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**RulebookVariables** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubTitle** | Pointer to **string** |  | [optional] 
**TickerName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **int64** |  | [optional] 
**YesAsk** | Pointer to **int64** |  | [optional] 
**YesBid** | Pointer to **int64** |  | [optional] 

## Methods

### NewEventChildMarket

`func NewEventChildMarket() *EventChildMarket`

NewEventChildMarket instantiates a new EventChildMarket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventChildMarketWithDefaults

`func NewEventChildMarketWithDefaults() *EventChildMarket`

NewEventChildMarketWithDefaults instantiates a new EventChildMarket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCloseEarly

`func (o *EventChildMarket) GetCanCloseEarly() bool`

GetCanCloseEarly returns the CanCloseEarly field if non-nil, zero value otherwise.

### GetCanCloseEarlyOk

`func (o *EventChildMarket) GetCanCloseEarlyOk() (*bool, bool)`

GetCanCloseEarlyOk returns a tuple with the CanCloseEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCloseEarly

`func (o *EventChildMarket) SetCanCloseEarly(v bool)`

SetCanCloseEarly sets CanCloseEarly field to given value.

### HasCanCloseEarly

`func (o *EventChildMarket) HasCanCloseEarly() bool`

HasCanCloseEarly returns a boolean if a field has been set.

### GetCloseDate

`func (o *EventChildMarket) GetCloseDate() time.Time`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *EventChildMarket) GetCloseDateOk() (*time.Time, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *EventChildMarket) SetCloseDate(v time.Time)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *EventChildMarket) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCloseUnconfirmed

`func (o *EventChildMarket) GetCloseUnconfirmed() bool`

GetCloseUnconfirmed returns the CloseUnconfirmed field if non-nil, zero value otherwise.

### GetCloseUnconfirmedOk

`func (o *EventChildMarket) GetCloseUnconfirmedOk() (*bool, bool)`

GetCloseUnconfirmedOk returns a tuple with the CloseUnconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseUnconfirmed

`func (o *EventChildMarket) SetCloseUnconfirmed(v bool)`

SetCloseUnconfirmed sets CloseUnconfirmed field to given value.

### HasCloseUnconfirmed

`func (o *EventChildMarket) HasCloseUnconfirmed() bool`

HasCloseUnconfirmed returns a boolean if a field has been set.

### GetCreateDate

`func (o *EventChildMarket) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *EventChildMarket) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *EventChildMarket) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *EventChildMarket) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDollarOpenInterest

`func (o *EventChildMarket) GetDollarOpenInterest() int64`

GetDollarOpenInterest returns the DollarOpenInterest field if non-nil, zero value otherwise.

### GetDollarOpenInterestOk

`func (o *EventChildMarket) GetDollarOpenInterestOk() (*int64, bool)`

GetDollarOpenInterestOk returns a tuple with the DollarOpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarOpenInterest

`func (o *EventChildMarket) SetDollarOpenInterest(v int64)`

SetDollarOpenInterest sets DollarOpenInterest field to given value.

### HasDollarOpenInterest

`func (o *EventChildMarket) HasDollarOpenInterest() bool`

HasDollarOpenInterest returns a boolean if a field has been set.

### GetDollarRecentVolume

`func (o *EventChildMarket) GetDollarRecentVolume() int64`

GetDollarRecentVolume returns the DollarRecentVolume field if non-nil, zero value otherwise.

### GetDollarRecentVolumeOk

`func (o *EventChildMarket) GetDollarRecentVolumeOk() (*int64, bool)`

GetDollarRecentVolumeOk returns a tuple with the DollarRecentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarRecentVolume

`func (o *EventChildMarket) SetDollarRecentVolume(v int64)`

SetDollarRecentVolume sets DollarRecentVolume field to given value.

### HasDollarRecentVolume

`func (o *EventChildMarket) HasDollarRecentVolume() bool`

HasDollarRecentVolume returns a boolean if a field has been set.

### GetDollarVolume

`func (o *EventChildMarket) GetDollarVolume() int64`

GetDollarVolume returns the DollarVolume field if non-nil, zero value otherwise.

### GetDollarVolumeOk

`func (o *EventChildMarket) GetDollarVolumeOk() (*int64, bool)`

GetDollarVolumeOk returns a tuple with the DollarVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarVolume

`func (o *EventChildMarket) SetDollarVolume(v int64)`

SetDollarVolume sets DollarVolume field to given value.

### HasDollarVolume

`func (o *EventChildMarket) HasDollarVolume() bool`

HasDollarVolume returns a boolean if a field has been set.

### GetExpirationDate

`func (o *EventChildMarket) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *EventChildMarket) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *EventChildMarket) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *EventChildMarket) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpirationValue

`func (o *EventChildMarket) GetExpirationValue() string`

GetExpirationValue returns the ExpirationValue field if non-nil, zero value otherwise.

### GetExpirationValueOk

`func (o *EventChildMarket) GetExpirationValueOk() (*string, bool)`

GetExpirationValueOk returns a tuple with the ExpirationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationValue

`func (o *EventChildMarket) SetExpirationValue(v string)`

SetExpirationValue sets ExpirationValue field to given value.

### HasExpirationValue

`func (o *EventChildMarket) HasExpirationValue() bool`

HasExpirationValue returns a boolean if a field has been set.

### GetId

`func (o *EventChildMarket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventChildMarket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventChildMarket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventChildMarket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPrice

`func (o *EventChildMarket) GetLastPrice() int64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *EventChildMarket) GetLastPriceOk() (*int64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *EventChildMarket) SetLastPrice(v int64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *EventChildMarket) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLiquidity

`func (o *EventChildMarket) GetLiquidity() int64`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *EventChildMarket) GetLiquidityOk() (*int64, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *EventChildMarket) SetLiquidity(v int64)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *EventChildMarket) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetListDate

`func (o *EventChildMarket) GetListDate() time.Time`

GetListDate returns the ListDate field if non-nil, zero value otherwise.

### GetListDateOk

`func (o *EventChildMarket) GetListDateOk() (*time.Time, bool)`

GetListDateOk returns a tuple with the ListDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListDate

`func (o *EventChildMarket) SetListDate(v time.Time)`

SetListDate sets ListDate field to given value.

### HasListDate

`func (o *EventChildMarket) HasListDate() bool`

HasListDate returns a boolean if a field has been set.

### GetMiniTitle

`func (o *EventChildMarket) GetMiniTitle() string`

GetMiniTitle returns the MiniTitle field if non-nil, zero value otherwise.

### GetMiniTitleOk

`func (o *EventChildMarket) GetMiniTitleOk() (*string, bool)`

GetMiniTitleOk returns a tuple with the MiniTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniTitle

`func (o *EventChildMarket) SetMiniTitle(v string)`

SetMiniTitle sets MiniTitle field to given value.

### HasMiniTitle

`func (o *EventChildMarket) HasMiniTitle() bool`

HasMiniTitle returns a boolean if a field has been set.

### GetOpenDate

`func (o *EventChildMarket) GetOpenDate() time.Time`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *EventChildMarket) GetOpenDateOk() (*time.Time, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *EventChildMarket) SetOpenDate(v time.Time)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *EventChildMarket) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetOpenInterest

`func (o *EventChildMarket) GetOpenInterest() int64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *EventChildMarket) GetOpenInterestOk() (*int64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *EventChildMarket) SetOpenInterest(v int64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *EventChildMarket) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetPreviousPrice

`func (o *EventChildMarket) GetPreviousPrice() int64`

GetPreviousPrice returns the PreviousPrice field if non-nil, zero value otherwise.

### GetPreviousPriceOk

`func (o *EventChildMarket) GetPreviousPriceOk() (*int64, bool)`

GetPreviousPriceOk returns a tuple with the PreviousPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPrice

`func (o *EventChildMarket) SetPreviousPrice(v int64)`

SetPreviousPrice sets PreviousPrice field to given value.

### HasPreviousPrice

`func (o *EventChildMarket) HasPreviousPrice() bool`

HasPreviousPrice returns a boolean if a field has been set.

### GetPreviousYesAsk

`func (o *EventChildMarket) GetPreviousYesAsk() int64`

GetPreviousYesAsk returns the PreviousYesAsk field if non-nil, zero value otherwise.

### GetPreviousYesAskOk

`func (o *EventChildMarket) GetPreviousYesAskOk() (*int64, bool)`

GetPreviousYesAskOk returns a tuple with the PreviousYesAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYesAsk

`func (o *EventChildMarket) SetPreviousYesAsk(v int64)`

SetPreviousYesAsk sets PreviousYesAsk field to given value.

### HasPreviousYesAsk

`func (o *EventChildMarket) HasPreviousYesAsk() bool`

HasPreviousYesAsk returns a boolean if a field has been set.

### GetPreviousYesBid

`func (o *EventChildMarket) GetPreviousYesBid() int64`

GetPreviousYesBid returns the PreviousYesBid field if non-nil, zero value otherwise.

### GetPreviousYesBidOk

`func (o *EventChildMarket) GetPreviousYesBidOk() (*int64, bool)`

GetPreviousYesBidOk returns a tuple with the PreviousYesBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYesBid

`func (o *EventChildMarket) SetPreviousYesBid(v int64)`

SetPreviousYesBid sets PreviousYesBid field to given value.

### HasPreviousYesBid

`func (o *EventChildMarket) HasPreviousYesBid() bool`

HasPreviousYesBid returns a boolean if a field has been set.

### GetRecentVolume

`func (o *EventChildMarket) GetRecentVolume() int64`

GetRecentVolume returns the RecentVolume field if non-nil, zero value otherwise.

### GetRecentVolumeOk

`func (o *EventChildMarket) GetRecentVolumeOk() (*int64, bool)`

GetRecentVolumeOk returns a tuple with the RecentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentVolume

`func (o *EventChildMarket) SetRecentVolume(v int64)`

SetRecentVolume sets RecentVolume field to given value.

### HasRecentVolume

`func (o *EventChildMarket) HasRecentVolume() bool`

HasRecentVolume returns a boolean if a field has been set.

### GetResult

`func (o *EventChildMarket) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EventChildMarket) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EventChildMarket) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *EventChildMarket) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRulebookVariables

`func (o *EventChildMarket) GetRulebookVariables() map[string]string`

GetRulebookVariables returns the RulebookVariables field if non-nil, zero value otherwise.

### GetRulebookVariablesOk

`func (o *EventChildMarket) GetRulebookVariablesOk() (*map[string]string, bool)`

GetRulebookVariablesOk returns a tuple with the RulebookVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulebookVariables

`func (o *EventChildMarket) SetRulebookVariables(v map[string]string)`

SetRulebookVariables sets RulebookVariables field to given value.

### HasRulebookVariables

`func (o *EventChildMarket) HasRulebookVariables() bool`

HasRulebookVariables returns a boolean if a field has been set.

### GetStatus

`func (o *EventChildMarket) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventChildMarket) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventChildMarket) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventChildMarket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubTitle

`func (o *EventChildMarket) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *EventChildMarket) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *EventChildMarket) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.

### HasSubTitle

`func (o *EventChildMarket) HasSubTitle() bool`

HasSubTitle returns a boolean if a field has been set.

### GetTickerName

`func (o *EventChildMarket) GetTickerName() string`

GetTickerName returns the TickerName field if non-nil, zero value otherwise.

### GetTickerNameOk

`func (o *EventChildMarket) GetTickerNameOk() (*string, bool)`

GetTickerNameOk returns a tuple with the TickerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerName

`func (o *EventChildMarket) SetTickerName(v string)`

SetTickerName sets TickerName field to given value.

### HasTickerName

`func (o *EventChildMarket) HasTickerName() bool`

HasTickerName returns a boolean if a field has been set.

### GetTitle

`func (o *EventChildMarket) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventChildMarket) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventChildMarket) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventChildMarket) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVolume

`func (o *EventChildMarket) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *EventChildMarket) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *EventChildMarket) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *EventChildMarket) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetYesAsk

`func (o *EventChildMarket) GetYesAsk() int64`

GetYesAsk returns the YesAsk field if non-nil, zero value otherwise.

### GetYesAskOk

`func (o *EventChildMarket) GetYesAskOk() (*int64, bool)`

GetYesAskOk returns a tuple with the YesAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesAsk

`func (o *EventChildMarket) SetYesAsk(v int64)`

SetYesAsk sets YesAsk field to given value.

### HasYesAsk

`func (o *EventChildMarket) HasYesAsk() bool`

HasYesAsk returns a boolean if a field has been set.

### GetYesBid

`func (o *EventChildMarket) GetYesBid() int64`

GetYesBid returns the YesBid field if non-nil, zero value otherwise.

### GetYesBidOk

`func (o *EventChildMarket) GetYesBidOk() (*int64, bool)`

GetYesBidOk returns a tuple with the YesBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesBid

`func (o *EventChildMarket) SetYesBid(v int64)`

SetYesBid sets YesBid field to given value.

### HasYesBid

`func (o *EventChildMarket) HasYesBid() bool`

HasYesBid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


