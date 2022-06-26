# GetUserDepositsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposits** | [**[]Deposit**](Deposit.md) | List of previous deposits for the user | 

## Methods

### NewGetUserDepositsResponse

`func NewGetUserDepositsResponse(deposits []Deposit, ) *GetUserDepositsResponse`

NewGetUserDepositsResponse instantiates a new GetUserDepositsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDepositsResponseWithDefaults

`func NewGetUserDepositsResponseWithDefaults() *GetUserDepositsResponse`

NewGetUserDepositsResponseWithDefaults instantiates a new GetUserDepositsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposits

`func (o *GetUserDepositsResponse) GetDeposits() []Deposit`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *GetUserDepositsResponse) GetDepositsOk() (*[]Deposit, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *GetUserDepositsResponse) SetDeposits(v []Deposit)`

SetDeposits sets Deposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


