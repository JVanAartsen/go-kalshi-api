# MarketTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AveragePriceCents** | Pointer to **int64** |  | [optional] 
**FeeCents** | Pointer to **int64** |  | [optional] 
**IsSideYes** | Pointer to **bool** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**MarketTicker** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**NumContracts** | Pointer to **int64** |  | [optional] 
**RealizedCostOfContractsCents** | Pointer to **int64** |  | [optional] 
**RealizedFifoProfitCents** | Pointer to **int64** |  | [optional] 
**RealizedRevenueCents** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketTransaction

`func NewMarketTransaction() *MarketTransaction`

NewMarketTransaction instantiates a new MarketTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTransactionWithDefaults

`func NewMarketTransactionWithDefaults() *MarketTransaction`

NewMarketTransactionWithDefaults instantiates a new MarketTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAveragePriceCents

`func (o *MarketTransaction) GetAveragePriceCents() int64`

GetAveragePriceCents returns the AveragePriceCents field if non-nil, zero value otherwise.

### GetAveragePriceCentsOk

`func (o *MarketTransaction) GetAveragePriceCentsOk() (*int64, bool)`

GetAveragePriceCentsOk returns a tuple with the AveragePriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePriceCents

`func (o *MarketTransaction) SetAveragePriceCents(v int64)`

SetAveragePriceCents sets AveragePriceCents field to given value.

### HasAveragePriceCents

`func (o *MarketTransaction) HasAveragePriceCents() bool`

HasAveragePriceCents returns a boolean if a field has been set.

### GetFeeCents

`func (o *MarketTransaction) GetFeeCents() int64`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *MarketTransaction) GetFeeCentsOk() (*int64, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *MarketTransaction) SetFeeCents(v int64)`

SetFeeCents sets FeeCents field to given value.

### HasFeeCents

`func (o *MarketTransaction) HasFeeCents() bool`

HasFeeCents returns a boolean if a field has been set.

### GetIsSideYes

`func (o *MarketTransaction) GetIsSideYes() bool`

GetIsSideYes returns the IsSideYes field if non-nil, zero value otherwise.

### GetIsSideYesOk

`func (o *MarketTransaction) GetIsSideYesOk() (*bool, bool)`

GetIsSideYesOk returns a tuple with the IsSideYes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSideYes

`func (o *MarketTransaction) SetIsSideYes(v bool)`

SetIsSideYes sets IsSideYes field to given value.

### HasIsSideYes

`func (o *MarketTransaction) HasIsSideYes() bool`

HasIsSideYes returns a boolean if a field has been set.

### GetMarketId

`func (o *MarketTransaction) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *MarketTransaction) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *MarketTransaction) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *MarketTransaction) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTicker

`func (o *MarketTransaction) GetMarketTicker() string`

GetMarketTicker returns the MarketTicker field if non-nil, zero value otherwise.

### GetMarketTickerOk

`func (o *MarketTransaction) GetMarketTickerOk() (*string, bool)`

GetMarketTickerOk returns a tuple with the MarketTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTicker

`func (o *MarketTransaction) SetMarketTicker(v string)`

SetMarketTicker sets MarketTicker field to given value.

### HasMarketTicker

`func (o *MarketTransaction) HasMarketTicker() bool`

HasMarketTicker returns a boolean if a field has been set.

### GetMarketTitle

`func (o *MarketTransaction) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *MarketTransaction) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *MarketTransaction) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *MarketTransaction) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetNumContracts

`func (o *MarketTransaction) GetNumContracts() int64`

GetNumContracts returns the NumContracts field if non-nil, zero value otherwise.

### GetNumContractsOk

`func (o *MarketTransaction) GetNumContractsOk() (*int64, bool)`

GetNumContractsOk returns a tuple with the NumContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumContracts

`func (o *MarketTransaction) SetNumContracts(v int64)`

SetNumContracts sets NumContracts field to given value.

### HasNumContracts

`func (o *MarketTransaction) HasNumContracts() bool`

HasNumContracts returns a boolean if a field has been set.

### GetRealizedCostOfContractsCents

`func (o *MarketTransaction) GetRealizedCostOfContractsCents() int64`

GetRealizedCostOfContractsCents returns the RealizedCostOfContractsCents field if non-nil, zero value otherwise.

### GetRealizedCostOfContractsCentsOk

`func (o *MarketTransaction) GetRealizedCostOfContractsCentsOk() (*int64, bool)`

GetRealizedCostOfContractsCentsOk returns a tuple with the RealizedCostOfContractsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedCostOfContractsCents

`func (o *MarketTransaction) SetRealizedCostOfContractsCents(v int64)`

SetRealizedCostOfContractsCents sets RealizedCostOfContractsCents field to given value.

### HasRealizedCostOfContractsCents

`func (o *MarketTransaction) HasRealizedCostOfContractsCents() bool`

HasRealizedCostOfContractsCents returns a boolean if a field has been set.

### GetRealizedFifoProfitCents

`func (o *MarketTransaction) GetRealizedFifoProfitCents() int64`

GetRealizedFifoProfitCents returns the RealizedFifoProfitCents field if non-nil, zero value otherwise.

### GetRealizedFifoProfitCentsOk

`func (o *MarketTransaction) GetRealizedFifoProfitCentsOk() (*int64, bool)`

GetRealizedFifoProfitCentsOk returns a tuple with the RealizedFifoProfitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedFifoProfitCents

`func (o *MarketTransaction) SetRealizedFifoProfitCents(v int64)`

SetRealizedFifoProfitCents sets RealizedFifoProfitCents field to given value.

### HasRealizedFifoProfitCents

`func (o *MarketTransaction) HasRealizedFifoProfitCents() bool`

HasRealizedFifoProfitCents returns a boolean if a field has been set.

### GetRealizedRevenueCents

`func (o *MarketTransaction) GetRealizedRevenueCents() int64`

GetRealizedRevenueCents returns the RealizedRevenueCents field if non-nil, zero value otherwise.

### GetRealizedRevenueCentsOk

`func (o *MarketTransaction) GetRealizedRevenueCentsOk() (*int64, bool)`

GetRealizedRevenueCentsOk returns a tuple with the RealizedRevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedRevenueCents

`func (o *MarketTransaction) SetRealizedRevenueCents(v int64)`

SetRealizedRevenueCents sets RealizedRevenueCents field to given value.

### HasRealizedRevenueCents

`func (o *MarketTransaction) HasRealizedRevenueCents() bool`

HasRealizedRevenueCents returns a boolean if a field has been set.

### GetTime

`func (o *MarketTransaction) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarketTransaction) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarketTransaction) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarketTransaction) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *MarketTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarketTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


