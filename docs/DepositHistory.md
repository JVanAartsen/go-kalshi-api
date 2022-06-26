# DepositHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DepositType** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **int64** |  | [optional] 
**ImmediateAmount** | Pointer to **int64** |  | [optional] 
**ImmediateStatus** | Pointer to **string** |  | [optional] 
**ReturnedAmount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDepositHistory

`func NewDepositHistory() *DepositHistory`

NewDepositHistory instantiates a new DepositHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositHistoryWithDefaults

`func NewDepositHistoryWithDefaults() *DepositHistory`

NewDepositHistoryWithDefaults instantiates a new DepositHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DepositHistory) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositHistory) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositHistory) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DepositHistory) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DepositHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DepositHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DepositHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DepositHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepositType

`func (o *DepositHistory) GetDepositType() string`

GetDepositType returns the DepositType field if non-nil, zero value otherwise.

### GetDepositTypeOk

`func (o *DepositHistory) GetDepositTypeOk() (*string, bool)`

GetDepositTypeOk returns a tuple with the DepositType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositType

`func (o *DepositHistory) SetDepositType(v string)`

SetDepositType sets DepositType field to given value.

### HasDepositType

`func (o *DepositHistory) HasDepositType() bool`

HasDepositType returns a boolean if a field has been set.

### GetFee

`func (o *DepositHistory) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *DepositHistory) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *DepositHistory) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *DepositHistory) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetImmediateAmount

`func (o *DepositHistory) GetImmediateAmount() int64`

GetImmediateAmount returns the ImmediateAmount field if non-nil, zero value otherwise.

### GetImmediateAmountOk

`func (o *DepositHistory) GetImmediateAmountOk() (*int64, bool)`

GetImmediateAmountOk returns a tuple with the ImmediateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateAmount

`func (o *DepositHistory) SetImmediateAmount(v int64)`

SetImmediateAmount sets ImmediateAmount field to given value.

### HasImmediateAmount

`func (o *DepositHistory) HasImmediateAmount() bool`

HasImmediateAmount returns a boolean if a field has been set.

### GetImmediateStatus

`func (o *DepositHistory) GetImmediateStatus() string`

GetImmediateStatus returns the ImmediateStatus field if non-nil, zero value otherwise.

### GetImmediateStatusOk

`func (o *DepositHistory) GetImmediateStatusOk() (*string, bool)`

GetImmediateStatusOk returns a tuple with the ImmediateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateStatus

`func (o *DepositHistory) SetImmediateStatus(v string)`

SetImmediateStatus sets ImmediateStatus field to given value.

### HasImmediateStatus

`func (o *DepositHistory) HasImmediateStatus() bool`

HasImmediateStatus returns a boolean if a field has been set.

### GetReturnedAmount

`func (o *DepositHistory) GetReturnedAmount() int64`

GetReturnedAmount returns the ReturnedAmount field if non-nil, zero value otherwise.

### GetReturnedAmountOk

`func (o *DepositHistory) GetReturnedAmountOk() (*int64, bool)`

GetReturnedAmountOk returns a tuple with the ReturnedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedAmount

`func (o *DepositHistory) SetReturnedAmount(v int64)`

SetReturnedAmount sets ReturnedAmount field to given value.

### HasReturnedAmount

`func (o *DepositHistory) HasReturnedAmount() bool`

HasReturnedAmount returns a boolean if a field has been set.

### GetStatus

`func (o *DepositHistory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DepositHistory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DepositHistory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DepositHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DepositHistory) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DepositHistory) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DepositHistory) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DepositHistory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


