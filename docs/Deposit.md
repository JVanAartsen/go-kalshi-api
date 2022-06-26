# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** |  | [optional] 
**BankId** | Pointer to **string** |  | [optional] 
**CreatedTs** | Pointer to **time.Time** |  | [optional] 
**DepositType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImmediateAmount** | Pointer to **int64** |  | [optional] 
**ImmediateStatus** | Pointer to **string** |  | [optional] 
**ReturnCode** | Pointer to **string** |  | [optional] 
**ReturnReason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewDeposit

`func NewDeposit() *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Deposit) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Deposit) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Deposit) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Deposit) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetBankId

`func (o *Deposit) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Deposit) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Deposit) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *Deposit) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetCreatedTs

`func (o *Deposit) GetCreatedTs() time.Time`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *Deposit) GetCreatedTsOk() (*time.Time, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *Deposit) SetCreatedTs(v time.Time)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *Deposit) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetDepositType

`func (o *Deposit) GetDepositType() string`

GetDepositType returns the DepositType field if non-nil, zero value otherwise.

### GetDepositTypeOk

`func (o *Deposit) GetDepositTypeOk() (*string, bool)`

GetDepositTypeOk returns a tuple with the DepositType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositType

`func (o *Deposit) SetDepositType(v string)`

SetDepositType sets DepositType field to given value.

### HasDepositType

`func (o *Deposit) HasDepositType() bool`

HasDepositType returns a boolean if a field has been set.

### GetId

`func (o *Deposit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deposit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deposit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deposit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImmediateAmount

`func (o *Deposit) GetImmediateAmount() int64`

GetImmediateAmount returns the ImmediateAmount field if non-nil, zero value otherwise.

### GetImmediateAmountOk

`func (o *Deposit) GetImmediateAmountOk() (*int64, bool)`

GetImmediateAmountOk returns a tuple with the ImmediateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateAmount

`func (o *Deposit) SetImmediateAmount(v int64)`

SetImmediateAmount sets ImmediateAmount field to given value.

### HasImmediateAmount

`func (o *Deposit) HasImmediateAmount() bool`

HasImmediateAmount returns a boolean if a field has been set.

### GetImmediateStatus

`func (o *Deposit) GetImmediateStatus() string`

GetImmediateStatus returns the ImmediateStatus field if non-nil, zero value otherwise.

### GetImmediateStatusOk

`func (o *Deposit) GetImmediateStatusOk() (*string, bool)`

GetImmediateStatusOk returns a tuple with the ImmediateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateStatus

`func (o *Deposit) SetImmediateStatus(v string)`

SetImmediateStatus sets ImmediateStatus field to given value.

### HasImmediateStatus

`func (o *Deposit) HasImmediateStatus() bool`

HasImmediateStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *Deposit) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *Deposit) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *Deposit) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *Deposit) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnReason

`func (o *Deposit) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *Deposit) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *Deposit) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *Deposit) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetStatus

`func (o *Deposit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deposit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deposit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deposit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *Deposit) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Deposit) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Deposit) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Deposit) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


