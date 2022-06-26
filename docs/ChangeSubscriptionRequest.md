# ChangeSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushPreferences** | Pointer to **bool** |  | [optional] 
**SubscriptionLevel** | **string** | Specifies the subscription level for email notifications its values can be: \&quot;none\&quot;, \&quot;basic\&quot; or \&quot;all\&quot; | 

## Methods

### NewChangeSubscriptionRequest

`func NewChangeSubscriptionRequest(subscriptionLevel string, ) *ChangeSubscriptionRequest`

NewChangeSubscriptionRequest instantiates a new ChangeSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSubscriptionRequestWithDefaults

`func NewChangeSubscriptionRequestWithDefaults() *ChangeSubscriptionRequest`

NewChangeSubscriptionRequestWithDefaults instantiates a new ChangeSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushPreferences

`func (o *ChangeSubscriptionRequest) GetPushPreferences() bool`

GetPushPreferences returns the PushPreferences field if non-nil, zero value otherwise.

### GetPushPreferencesOk

`func (o *ChangeSubscriptionRequest) GetPushPreferencesOk() (*bool, bool)`

GetPushPreferencesOk returns a tuple with the PushPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushPreferences

`func (o *ChangeSubscriptionRequest) SetPushPreferences(v bool)`

SetPushPreferences sets PushPreferences field to given value.

### HasPushPreferences

`func (o *ChangeSubscriptionRequest) HasPushPreferences() bool`

HasPushPreferences returns a boolean if a field has been set.

### GetSubscriptionLevel

`func (o *ChangeSubscriptionRequest) GetSubscriptionLevel() string`

GetSubscriptionLevel returns the SubscriptionLevel field if non-nil, zero value otherwise.

### GetSubscriptionLevelOk

`func (o *ChangeSubscriptionRequest) GetSubscriptionLevelOk() (*string, bool)`

GetSubscriptionLevelOk returns a tuple with the SubscriptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLevel

`func (o *ChangeSubscriptionRequest) SetSubscriptionLevel(v string)`

SetSubscriptionLevel sets SubscriptionLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


