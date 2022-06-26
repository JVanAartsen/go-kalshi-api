# WithdrawalHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Fee** | Pointer to **int64** |  | [optional] 
**ReturnedAmount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWithdrawalHistory

`func NewWithdrawalHistory() *WithdrawalHistory`

NewWithdrawalHistory instantiates a new WithdrawalHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalHistoryWithDefaults

`func NewWithdrawalHistoryWithDefaults() *WithdrawalHistory`

NewWithdrawalHistoryWithDefaults instantiates a new WithdrawalHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WithdrawalHistory) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawalHistory) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawalHistory) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WithdrawalHistory) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WithdrawalHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WithdrawalHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WithdrawalHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WithdrawalHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFee

`func (o *WithdrawalHistory) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *WithdrawalHistory) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *WithdrawalHistory) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *WithdrawalHistory) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetReturnedAmount

`func (o *WithdrawalHistory) GetReturnedAmount() int64`

GetReturnedAmount returns the ReturnedAmount field if non-nil, zero value otherwise.

### GetReturnedAmountOk

`func (o *WithdrawalHistory) GetReturnedAmountOk() (*int64, bool)`

GetReturnedAmountOk returns a tuple with the ReturnedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedAmount

`func (o *WithdrawalHistory) SetReturnedAmount(v int64)`

SetReturnedAmount sets ReturnedAmount field to given value.

### HasReturnedAmount

`func (o *WithdrawalHistory) HasReturnedAmount() bool`

HasReturnedAmount returns a boolean if a field has been set.

### GetStatus

`func (o *WithdrawalHistory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WithdrawalHistory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WithdrawalHistory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WithdrawalHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WithdrawalHistory) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WithdrawalHistory) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WithdrawalHistory) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WithdrawalHistory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


