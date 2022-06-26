# BankAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **string** |  | 
**Mask** | **string** |  | 
**Name** | **string** |  | 
**PlaidItemNeedsRelink** | **bool** |  | 
**RoutingNumber** | Pointer to **string** |  | [optional] 
**Subtype** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewBankAccountDetails

`func NewBankAccountDetails(bankId string, mask string, name string, plaidItemNeedsRelink bool, subtype string, type_ string, ) *BankAccountDetails`

NewBankAccountDetails instantiates a new BankAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountDetailsWithDefaults

`func NewBankAccountDetailsWithDefaults() *BankAccountDetails`

NewBankAccountDetailsWithDefaults instantiates a new BankAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *BankAccountDetails) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BankAccountDetails) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BankAccountDetails) SetBankId(v string)`

SetBankId sets BankId field to given value.


### GetMask

`func (o *BankAccountDetails) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *BankAccountDetails) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *BankAccountDetails) SetMask(v string)`

SetMask sets Mask field to given value.


### GetName

`func (o *BankAccountDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankAccountDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankAccountDetails) SetName(v string)`

SetName sets Name field to given value.


### GetPlaidItemNeedsRelink

`func (o *BankAccountDetails) GetPlaidItemNeedsRelink() bool`

GetPlaidItemNeedsRelink returns the PlaidItemNeedsRelink field if non-nil, zero value otherwise.

### GetPlaidItemNeedsRelinkOk

`func (o *BankAccountDetails) GetPlaidItemNeedsRelinkOk() (*bool, bool)`

GetPlaidItemNeedsRelinkOk returns a tuple with the PlaidItemNeedsRelink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaidItemNeedsRelink

`func (o *BankAccountDetails) SetPlaidItemNeedsRelink(v bool)`

SetPlaidItemNeedsRelink sets PlaidItemNeedsRelink field to given value.


### GetRoutingNumber

`func (o *BankAccountDetails) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountDetails) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountDetails) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *BankAccountDetails) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetSubtype

`func (o *BankAccountDetails) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BankAccountDetails) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BankAccountDetails) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetType

`func (o *BankAccountDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankAccountDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankAccountDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


