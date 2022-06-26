# SettlementHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeterminedTime** | Pointer to **time.Time** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**MarketResult** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**NoCount** | Pointer to **int64** |  | [optional] 
**NoTotalCost** | Pointer to **int64** |  | [optional] 
**Profit** | Pointer to **int64** |  | [optional] 
**SettledTime** | Pointer to **time.Time** |  | [optional] 
**YesCount** | Pointer to **int64** |  | [optional] 
**YesTotalCost** | Pointer to **int64** |  | [optional] 

## Methods

### NewSettlementHistory

`func NewSettlementHistory() *SettlementHistory`

NewSettlementHistory instantiates a new SettlementHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementHistoryWithDefaults

`func NewSettlementHistoryWithDefaults() *SettlementHistory`

NewSettlementHistoryWithDefaults instantiates a new SettlementHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeterminedTime

`func (o *SettlementHistory) GetDeterminedTime() time.Time`

GetDeterminedTime returns the DeterminedTime field if non-nil, zero value otherwise.

### GetDeterminedTimeOk

`func (o *SettlementHistory) GetDeterminedTimeOk() (*time.Time, bool)`

GetDeterminedTimeOk returns a tuple with the DeterminedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterminedTime

`func (o *SettlementHistory) SetDeterminedTime(v time.Time)`

SetDeterminedTime sets DeterminedTime field to given value.

### HasDeterminedTime

`func (o *SettlementHistory) HasDeterminedTime() bool`

HasDeterminedTime returns a boolean if a field has been set.

### GetMarketId

`func (o *SettlementHistory) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *SettlementHistory) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *SettlementHistory) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *SettlementHistory) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketResult

`func (o *SettlementHistory) GetMarketResult() string`

GetMarketResult returns the MarketResult field if non-nil, zero value otherwise.

### GetMarketResultOk

`func (o *SettlementHistory) GetMarketResultOk() (*string, bool)`

GetMarketResultOk returns a tuple with the MarketResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketResult

`func (o *SettlementHistory) SetMarketResult(v string)`

SetMarketResult sets MarketResult field to given value.

### HasMarketResult

`func (o *SettlementHistory) HasMarketResult() bool`

HasMarketResult returns a boolean if a field has been set.

### GetMarketTitle

`func (o *SettlementHistory) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *SettlementHistory) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *SettlementHistory) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *SettlementHistory) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetNoCount

`func (o *SettlementHistory) GetNoCount() int64`

GetNoCount returns the NoCount field if non-nil, zero value otherwise.

### GetNoCountOk

`func (o *SettlementHistory) GetNoCountOk() (*int64, bool)`

GetNoCountOk returns a tuple with the NoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCount

`func (o *SettlementHistory) SetNoCount(v int64)`

SetNoCount sets NoCount field to given value.

### HasNoCount

`func (o *SettlementHistory) HasNoCount() bool`

HasNoCount returns a boolean if a field has been set.

### GetNoTotalCost

`func (o *SettlementHistory) GetNoTotalCost() int64`

GetNoTotalCost returns the NoTotalCost field if non-nil, zero value otherwise.

### GetNoTotalCostOk

`func (o *SettlementHistory) GetNoTotalCostOk() (*int64, bool)`

GetNoTotalCostOk returns a tuple with the NoTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTotalCost

`func (o *SettlementHistory) SetNoTotalCost(v int64)`

SetNoTotalCost sets NoTotalCost field to given value.

### HasNoTotalCost

`func (o *SettlementHistory) HasNoTotalCost() bool`

HasNoTotalCost returns a boolean if a field has been set.

### GetProfit

`func (o *SettlementHistory) GetProfit() int64`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *SettlementHistory) GetProfitOk() (*int64, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *SettlementHistory) SetProfit(v int64)`

SetProfit sets Profit field to given value.

### HasProfit

`func (o *SettlementHistory) HasProfit() bool`

HasProfit returns a boolean if a field has been set.

### GetSettledTime

`func (o *SettlementHistory) GetSettledTime() time.Time`

GetSettledTime returns the SettledTime field if non-nil, zero value otherwise.

### GetSettledTimeOk

`func (o *SettlementHistory) GetSettledTimeOk() (*time.Time, bool)`

GetSettledTimeOk returns a tuple with the SettledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledTime

`func (o *SettlementHistory) SetSettledTime(v time.Time)`

SetSettledTime sets SettledTime field to given value.

### HasSettledTime

`func (o *SettlementHistory) HasSettledTime() bool`

HasSettledTime returns a boolean if a field has been set.

### GetYesCount

`func (o *SettlementHistory) GetYesCount() int64`

GetYesCount returns the YesCount field if non-nil, zero value otherwise.

### GetYesCountOk

`func (o *SettlementHistory) GetYesCountOk() (*int64, bool)`

GetYesCountOk returns a tuple with the YesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesCount

`func (o *SettlementHistory) SetYesCount(v int64)`

SetYesCount sets YesCount field to given value.

### HasYesCount

`func (o *SettlementHistory) HasYesCount() bool`

HasYesCount returns a boolean if a field has been set.

### GetYesTotalCost

`func (o *SettlementHistory) GetYesTotalCost() int64`

GetYesTotalCost returns the YesTotalCost field if non-nil, zero value otherwise.

### GetYesTotalCostOk

`func (o *SettlementHistory) GetYesTotalCostOk() (*int64, bool)`

GetYesTotalCostOk returns a tuple with the YesTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesTotalCost

`func (o *SettlementHistory) SetYesTotalCost(v int64)`

SetYesTotalCost sets YesTotalCost field to given value.

### HasYesTotalCost

`func (o *SettlementHistory) HasYesTotalCost() bool`

HasYesTotalCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


