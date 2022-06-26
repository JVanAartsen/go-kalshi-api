# UserGetReferralInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EligibleToRefer** | Pointer to **bool** |  | [optional] 
**NumContractsTraded** | Pointer to **int64** |  | [optional] 
**PeopleReferred** | Pointer to [**[]PeopleReferred**](PeopleReferred.md) |  | [optional] 
**ReferralCode** | Pointer to **string** |  | [optional] 
**ReferralId** | Pointer to **string** |  | [optional] 
**ReferralMoneyRewarded** | Pointer to **int64** |  | [optional] 
**StageInFunnel** | Pointer to **string** |  | [optional] 

## Methods

### NewUserGetReferralInfoResponse

`func NewUserGetReferralInfoResponse() *UserGetReferralInfoResponse`

NewUserGetReferralInfoResponse instantiates a new UserGetReferralInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetReferralInfoResponseWithDefaults

`func NewUserGetReferralInfoResponseWithDefaults() *UserGetReferralInfoResponse`

NewUserGetReferralInfoResponseWithDefaults instantiates a new UserGetReferralInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibleToRefer

`func (o *UserGetReferralInfoResponse) GetEligibleToRefer() bool`

GetEligibleToRefer returns the EligibleToRefer field if non-nil, zero value otherwise.

### GetEligibleToReferOk

`func (o *UserGetReferralInfoResponse) GetEligibleToReferOk() (*bool, bool)`

GetEligibleToReferOk returns a tuple with the EligibleToRefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleToRefer

`func (o *UserGetReferralInfoResponse) SetEligibleToRefer(v bool)`

SetEligibleToRefer sets EligibleToRefer field to given value.

### HasEligibleToRefer

`func (o *UserGetReferralInfoResponse) HasEligibleToRefer() bool`

HasEligibleToRefer returns a boolean if a field has been set.

### GetNumContractsTraded

`func (o *UserGetReferralInfoResponse) GetNumContractsTraded() int64`

GetNumContractsTraded returns the NumContractsTraded field if non-nil, zero value otherwise.

### GetNumContractsTradedOk

`func (o *UserGetReferralInfoResponse) GetNumContractsTradedOk() (*int64, bool)`

GetNumContractsTradedOk returns a tuple with the NumContractsTraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumContractsTraded

`func (o *UserGetReferralInfoResponse) SetNumContractsTraded(v int64)`

SetNumContractsTraded sets NumContractsTraded field to given value.

### HasNumContractsTraded

`func (o *UserGetReferralInfoResponse) HasNumContractsTraded() bool`

HasNumContractsTraded returns a boolean if a field has been set.

### GetPeopleReferred

`func (o *UserGetReferralInfoResponse) GetPeopleReferred() []PeopleReferred`

GetPeopleReferred returns the PeopleReferred field if non-nil, zero value otherwise.

### GetPeopleReferredOk

`func (o *UserGetReferralInfoResponse) GetPeopleReferredOk() (*[]PeopleReferred, bool)`

GetPeopleReferredOk returns a tuple with the PeopleReferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleReferred

`func (o *UserGetReferralInfoResponse) SetPeopleReferred(v []PeopleReferred)`

SetPeopleReferred sets PeopleReferred field to given value.

### HasPeopleReferred

`func (o *UserGetReferralInfoResponse) HasPeopleReferred() bool`

HasPeopleReferred returns a boolean if a field has been set.

### GetReferralCode

`func (o *UserGetReferralInfoResponse) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *UserGetReferralInfoResponse) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *UserGetReferralInfoResponse) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *UserGetReferralInfoResponse) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### GetReferralId

`func (o *UserGetReferralInfoResponse) GetReferralId() string`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *UserGetReferralInfoResponse) GetReferralIdOk() (*string, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *UserGetReferralInfoResponse) SetReferralId(v string)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *UserGetReferralInfoResponse) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetReferralMoneyRewarded

`func (o *UserGetReferralInfoResponse) GetReferralMoneyRewarded() int64`

GetReferralMoneyRewarded returns the ReferralMoneyRewarded field if non-nil, zero value otherwise.

### GetReferralMoneyRewardedOk

`func (o *UserGetReferralInfoResponse) GetReferralMoneyRewardedOk() (*int64, bool)`

GetReferralMoneyRewardedOk returns a tuple with the ReferralMoneyRewarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralMoneyRewarded

`func (o *UserGetReferralInfoResponse) SetReferralMoneyRewarded(v int64)`

SetReferralMoneyRewarded sets ReferralMoneyRewarded field to given value.

### HasReferralMoneyRewarded

`func (o *UserGetReferralInfoResponse) HasReferralMoneyRewarded() bool`

HasReferralMoneyRewarded returns a boolean if a field has been set.

### GetStageInFunnel

`func (o *UserGetReferralInfoResponse) GetStageInFunnel() string`

GetStageInFunnel returns the StageInFunnel field if non-nil, zero value otherwise.

### GetStageInFunnelOk

`func (o *UserGetReferralInfoResponse) GetStageInFunnelOk() (*string, bool)`

GetStageInFunnelOk returns a tuple with the StageInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageInFunnel

`func (o *UserGetReferralInfoResponse) SetStageInFunnel(v string)`

SetStageInFunnel sets StageInFunnel field to given value.

### HasStageInFunnel

`func (o *UserGetReferralInfoResponse) HasStageInFunnel() bool`

HasStageInFunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


