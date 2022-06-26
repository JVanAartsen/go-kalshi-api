# GetUserWithdrawalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Withdrawals** | [**[]Withdrawal**](Withdrawal.md) | List of previous withdrawals for the user | 

## Methods

### NewGetUserWithdrawalsResponse

`func NewGetUserWithdrawalsResponse(withdrawals []Withdrawal, ) *GetUserWithdrawalsResponse`

NewGetUserWithdrawalsResponse instantiates a new GetUserWithdrawalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserWithdrawalsResponseWithDefaults

`func NewGetUserWithdrawalsResponseWithDefaults() *GetUserWithdrawalsResponse`

NewGetUserWithdrawalsResponseWithDefaults instantiates a new GetUserWithdrawalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithdrawals

`func (o *GetUserWithdrawalsResponse) GetWithdrawals() []Withdrawal`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *GetUserWithdrawalsResponse) GetWithdrawalsOk() (*[]Withdrawal, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *GetUserWithdrawalsResponse) SetWithdrawals(v []Withdrawal)`

SetWithdrawals sets Withdrawals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


