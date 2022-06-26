# UserDepositResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposit** | Pointer to [**Deposit**](Deposit.md) |  | [optional] 
**DepositId** | **string** | Id for the deposit that was created. | 
**EstimatedAchTimeDays** | **int32** | The estimated number of days we believe the ach transfer will take | 

## Methods

### NewUserDepositResponse

`func NewUserDepositResponse(depositId string, estimatedAchTimeDays int32, ) *UserDepositResponse`

NewUserDepositResponse instantiates a new UserDepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDepositResponseWithDefaults

`func NewUserDepositResponseWithDefaults() *UserDepositResponse`

NewUserDepositResponseWithDefaults instantiates a new UserDepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposit

`func (o *UserDepositResponse) GetDeposit() Deposit`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *UserDepositResponse) GetDepositOk() (*Deposit, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *UserDepositResponse) SetDeposit(v Deposit)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *UserDepositResponse) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDepositId

`func (o *UserDepositResponse) GetDepositId() string`

GetDepositId returns the DepositId field if non-nil, zero value otherwise.

### GetDepositIdOk

`func (o *UserDepositResponse) GetDepositIdOk() (*string, bool)`

GetDepositIdOk returns a tuple with the DepositId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositId

`func (o *UserDepositResponse) SetDepositId(v string)`

SetDepositId sets DepositId field to given value.


### GetEstimatedAchTimeDays

`func (o *UserDepositResponse) GetEstimatedAchTimeDays() int32`

GetEstimatedAchTimeDays returns the EstimatedAchTimeDays field if non-nil, zero value otherwise.

### GetEstimatedAchTimeDaysOk

`func (o *UserDepositResponse) GetEstimatedAchTimeDaysOk() (*int32, bool)`

GetEstimatedAchTimeDaysOk returns a tuple with the EstimatedAchTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAchTimeDays

`func (o *UserDepositResponse) SetEstimatedAchTimeDays(v int32)`

SetEstimatedAchTimeDays sets EstimatedAchTimeDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


