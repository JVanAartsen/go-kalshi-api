# UserGetProfitsAndLossesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketTransactions** | Pointer to [**[]MarketTransaction**](MarketTransaction.md) | User market transactions (trades and settlements)  in: body | [optional] 
**PnlWithFeesCents** | Pointer to **int64** |  | [optional] 
**PnlWithoutFeesCents** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserGetProfitsAndLossesResponse

`func NewUserGetProfitsAndLossesResponse() *UserGetProfitsAndLossesResponse`

NewUserGetProfitsAndLossesResponse instantiates a new UserGetProfitsAndLossesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetProfitsAndLossesResponseWithDefaults

`func NewUserGetProfitsAndLossesResponseWithDefaults() *UserGetProfitsAndLossesResponse`

NewUserGetProfitsAndLossesResponseWithDefaults instantiates a new UserGetProfitsAndLossesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketTransactions

`func (o *UserGetProfitsAndLossesResponse) GetMarketTransactions() []MarketTransaction`

GetMarketTransactions returns the MarketTransactions field if non-nil, zero value otherwise.

### GetMarketTransactionsOk

`func (o *UserGetProfitsAndLossesResponse) GetMarketTransactionsOk() (*[]MarketTransaction, bool)`

GetMarketTransactionsOk returns a tuple with the MarketTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTransactions

`func (o *UserGetProfitsAndLossesResponse) SetMarketTransactions(v []MarketTransaction)`

SetMarketTransactions sets MarketTransactions field to given value.

### HasMarketTransactions

`func (o *UserGetProfitsAndLossesResponse) HasMarketTransactions() bool`

HasMarketTransactions returns a boolean if a field has been set.

### GetPnlWithFeesCents

`func (o *UserGetProfitsAndLossesResponse) GetPnlWithFeesCents() int64`

GetPnlWithFeesCents returns the PnlWithFeesCents field if non-nil, zero value otherwise.

### GetPnlWithFeesCentsOk

`func (o *UserGetProfitsAndLossesResponse) GetPnlWithFeesCentsOk() (*int64, bool)`

GetPnlWithFeesCentsOk returns a tuple with the PnlWithFeesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlWithFeesCents

`func (o *UserGetProfitsAndLossesResponse) SetPnlWithFeesCents(v int64)`

SetPnlWithFeesCents sets PnlWithFeesCents field to given value.

### HasPnlWithFeesCents

`func (o *UserGetProfitsAndLossesResponse) HasPnlWithFeesCents() bool`

HasPnlWithFeesCents returns a boolean if a field has been set.

### GetPnlWithoutFeesCents

`func (o *UserGetProfitsAndLossesResponse) GetPnlWithoutFeesCents() int64`

GetPnlWithoutFeesCents returns the PnlWithoutFeesCents field if non-nil, zero value otherwise.

### GetPnlWithoutFeesCentsOk

`func (o *UserGetProfitsAndLossesResponse) GetPnlWithoutFeesCentsOk() (*int64, bool)`

GetPnlWithoutFeesCentsOk returns a tuple with the PnlWithoutFeesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlWithoutFeesCents

`func (o *UserGetProfitsAndLossesResponse) SetPnlWithoutFeesCents(v int64)`

SetPnlWithoutFeesCents sets PnlWithoutFeesCents field to given value.

### HasPnlWithoutFeesCents

`func (o *UserGetProfitsAndLossesResponse) HasPnlWithoutFeesCents() bool`

HasPnlWithoutFeesCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


