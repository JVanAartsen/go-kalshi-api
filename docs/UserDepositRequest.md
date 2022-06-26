# UserDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int64** |  | 
**BankId** | **string** |  | 
**FeeCents** | **int64** |  | 

## Methods

### NewUserDepositRequest

`func NewUserDepositRequest(amountCents int64, bankId string, feeCents int64, ) *UserDepositRequest`

NewUserDepositRequest instantiates a new UserDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDepositRequestWithDefaults

`func NewUserDepositRequestWithDefaults() *UserDepositRequest`

NewUserDepositRequestWithDefaults instantiates a new UserDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *UserDepositRequest) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *UserDepositRequest) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *UserDepositRequest) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.


### GetBankId

`func (o *UserDepositRequest) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *UserDepositRequest) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *UserDepositRequest) SetBankId(v string)`

SetBankId sets BankId field to given value.


### GetFeeCents

`func (o *UserDepositRequest) GetFeeCents() int64`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *UserDepositRequest) GetFeeCentsOk() (*int64, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *UserDepositRequest) SetFeeCents(v int64)`

SetFeeCents sets FeeCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


