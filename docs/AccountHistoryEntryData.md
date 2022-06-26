# AccountHistoryEntryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | Pointer to [**CreditHistory**](CreditHistory.md) |  | [optional] 
**Deposit** | Pointer to [**DepositHistory**](DepositHistory.md) |  | [optional] 
**Order** | Pointer to [**OrderHistory**](OrderHistory.md) |  | [optional] 
**Settlement** | Pointer to [**SettlementHistory**](SettlementHistory.md) |  | [optional] 
**Trade** | Pointer to [**TradeHistory**](TradeHistory.md) |  | [optional] 
**Withdrawal** | Pointer to [**WithdrawalHistory**](WithdrawalHistory.md) |  | [optional] 

## Methods

### NewAccountHistoryEntryData

`func NewAccountHistoryEntryData() *AccountHistoryEntryData`

NewAccountHistoryEntryData instantiates a new AccountHistoryEntryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHistoryEntryDataWithDefaults

`func NewAccountHistoryEntryDataWithDefaults() *AccountHistoryEntryData`

NewAccountHistoryEntryDataWithDefaults instantiates a new AccountHistoryEntryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *AccountHistoryEntryData) GetCredit() CreditHistory`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AccountHistoryEntryData) GetCreditOk() (*CreditHistory, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AccountHistoryEntryData) SetCredit(v CreditHistory)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AccountHistoryEntryData) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDeposit

`func (o *AccountHistoryEntryData) GetDeposit() DepositHistory`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *AccountHistoryEntryData) GetDepositOk() (*DepositHistory, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *AccountHistoryEntryData) SetDeposit(v DepositHistory)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *AccountHistoryEntryData) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetOrder

`func (o *AccountHistoryEntryData) GetOrder() OrderHistory`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AccountHistoryEntryData) GetOrderOk() (*OrderHistory, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AccountHistoryEntryData) SetOrder(v OrderHistory)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AccountHistoryEntryData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSettlement

`func (o *AccountHistoryEntryData) GetSettlement() SettlementHistory`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *AccountHistoryEntryData) GetSettlementOk() (*SettlementHistory, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *AccountHistoryEntryData) SetSettlement(v SettlementHistory)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *AccountHistoryEntryData) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### GetTrade

`func (o *AccountHistoryEntryData) GetTrade() TradeHistory`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *AccountHistoryEntryData) GetTradeOk() (*TradeHistory, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *AccountHistoryEntryData) SetTrade(v TradeHistory)`

SetTrade sets Trade field to given value.

### HasTrade

`func (o *AccountHistoryEntryData) HasTrade() bool`

HasTrade returns a boolean if a field has been set.

### GetWithdrawal

`func (o *AccountHistoryEntryData) GetWithdrawal() WithdrawalHistory`

GetWithdrawal returns the Withdrawal field if non-nil, zero value otherwise.

### GetWithdrawalOk

`func (o *AccountHistoryEntryData) GetWithdrawalOk() (*WithdrawalHistory, bool)`

GetWithdrawalOk returns a tuple with the Withdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawal

`func (o *AccountHistoryEntryData) SetWithdrawal(v WithdrawalHistory)`

SetWithdrawal sets Withdrawal field to given value.

### HasWithdrawal

`func (o *AccountHistoryEntryData) HasWithdrawal() bool`

HasWithdrawal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


