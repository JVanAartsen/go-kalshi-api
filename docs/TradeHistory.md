# TradeHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Fee** | Pointer to **int64** |  | [optional] 
**IsYes** | Pointer to **bool** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **int64** |  | [optional] 

## Methods

### NewTradeHistory

`func NewTradeHistory() *TradeHistory`

NewTradeHistory instantiates a new TradeHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeHistoryWithDefaults

`func NewTradeHistoryWithDefaults() *TradeHistory`

NewTradeHistoryWithDefaults instantiates a new TradeHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TradeHistory) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TradeHistory) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TradeHistory) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TradeHistory) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TradeHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TradeHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TradeHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TradeHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFee

`func (o *TradeHistory) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TradeHistory) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TradeHistory) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TradeHistory) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetIsYes

`func (o *TradeHistory) GetIsYes() bool`

GetIsYes returns the IsYes field if non-nil, zero value otherwise.

### GetIsYesOk

`func (o *TradeHistory) GetIsYesOk() (*bool, bool)`

GetIsYesOk returns a tuple with the IsYes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYes

`func (o *TradeHistory) SetIsYes(v bool)`

SetIsYes sets IsYes field to given value.

### HasIsYes

`func (o *TradeHistory) HasIsYes() bool`

HasIsYes returns a boolean if a field has been set.

### GetMarketId

`func (o *TradeHistory) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *TradeHistory) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *TradeHistory) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *TradeHistory) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTitle

`func (o *TradeHistory) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *TradeHistory) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *TradeHistory) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *TradeHistory) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetPrice

`func (o *TradeHistory) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TradeHistory) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TradeHistory) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TradeHistory) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


