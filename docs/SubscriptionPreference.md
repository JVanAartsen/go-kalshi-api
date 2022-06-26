# SubscriptionPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushPreferences** | Pointer to **[]string** |  | [optional] 
**SubscriptionLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionPreference

`func NewSubscriptionPreference() *SubscriptionPreference`

NewSubscriptionPreference instantiates a new SubscriptionPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPreferenceWithDefaults

`func NewSubscriptionPreferenceWithDefaults() *SubscriptionPreference`

NewSubscriptionPreferenceWithDefaults instantiates a new SubscriptionPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushPreferences

`func (o *SubscriptionPreference) GetPushPreferences() []string`

GetPushPreferences returns the PushPreferences field if non-nil, zero value otherwise.

### GetPushPreferencesOk

`func (o *SubscriptionPreference) GetPushPreferencesOk() (*[]string, bool)`

GetPushPreferencesOk returns a tuple with the PushPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushPreferences

`func (o *SubscriptionPreference) SetPushPreferences(v []string)`

SetPushPreferences sets PushPreferences field to given value.

### HasPushPreferences

`func (o *SubscriptionPreference) HasPushPreferences() bool`

HasPushPreferences returns a boolean if a field has been set.

### GetSubscriptionLevel

`func (o *SubscriptionPreference) GetSubscriptionLevel() string`

GetSubscriptionLevel returns the SubscriptionLevel field if non-nil, zero value otherwise.

### GetSubscriptionLevelOk

`func (o *SubscriptionPreference) GetSubscriptionLevelOk() (*string, bool)`

GetSubscriptionLevelOk returns a tuple with the SubscriptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLevel

`func (o *SubscriptionPreference) SetSubscriptionLevel(v string)`

SetSubscriptionLevel sets SubscriptionLevel field to given value.

### HasSubscriptionLevel

`func (o *SubscriptionPreference) HasSubscriptionLevel() bool`

HasSubscriptionLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


