# UserWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int64** |  | 
**BankId** | **string** |  | 
**FeeCents** | **int64** |  | 

## Methods

### NewUserWithdrawalRequest

`func NewUserWithdrawalRequest(amountCents int64, bankId string, feeCents int64, ) *UserWithdrawalRequest`

NewUserWithdrawalRequest instantiates a new UserWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithdrawalRequestWithDefaults

`func NewUserWithdrawalRequestWithDefaults() *UserWithdrawalRequest`

NewUserWithdrawalRequestWithDefaults instantiates a new UserWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *UserWithdrawalRequest) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *UserWithdrawalRequest) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *UserWithdrawalRequest) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.


### GetBankId

`func (o *UserWithdrawalRequest) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *UserWithdrawalRequest) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *UserWithdrawalRequest) SetBankId(v string)`

SetBankId sets BankId field to given value.


### GetFeeCents

`func (o *UserWithdrawalRequest) GetFeeCents() int64`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *UserWithdrawalRequest) GetFeeCentsOk() (*int64, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *UserWithdrawalRequest) SetFeeCents(v int64)`

SetFeeCents sets FeeCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


