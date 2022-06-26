# UserWithdrawalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Withdrawal** | Pointer to [**Withdrawal**](Withdrawal.md) |  | [optional] 
**WithdrawalId** | **string** | Id for the withdrawal that was created. | 

## Methods

### NewUserWithdrawalResponse

`func NewUserWithdrawalResponse(withdrawalId string, ) *UserWithdrawalResponse`

NewUserWithdrawalResponse instantiates a new UserWithdrawalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithdrawalResponseWithDefaults

`func NewUserWithdrawalResponseWithDefaults() *UserWithdrawalResponse`

NewUserWithdrawalResponseWithDefaults instantiates a new UserWithdrawalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithdrawal

`func (o *UserWithdrawalResponse) GetWithdrawal() Withdrawal`

GetWithdrawal returns the Withdrawal field if non-nil, zero value otherwise.

### GetWithdrawalOk

`func (o *UserWithdrawalResponse) GetWithdrawalOk() (*Withdrawal, bool)`

GetWithdrawalOk returns a tuple with the Withdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawal

`func (o *UserWithdrawalResponse) SetWithdrawal(v Withdrawal)`

SetWithdrawal sets Withdrawal field to given value.

### HasWithdrawal

`func (o *UserWithdrawalResponse) HasWithdrawal() bool`

HasWithdrawal returns a boolean if a field has been set.

### GetWithdrawalId

`func (o *UserWithdrawalResponse) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *UserWithdrawalResponse) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *UserWithdrawalResponse) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


