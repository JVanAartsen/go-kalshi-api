# ConfirmPasswordResetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The new password. | 
**UserId** | **string** | UserUUID for your user. You can get this from the password reset link query parameter. | 

## Methods

### NewConfirmPasswordResetRequest

`func NewConfirmPasswordResetRequest(password string, userId string, ) *ConfirmPasswordResetRequest`

NewConfirmPasswordResetRequest instantiates a new ConfirmPasswordResetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmPasswordResetRequestWithDefaults

`func NewConfirmPasswordResetRequestWithDefaults() *ConfirmPasswordResetRequest`

NewConfirmPasswordResetRequestWithDefaults instantiates a new ConfirmPasswordResetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ConfirmPasswordResetRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConfirmPasswordResetRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConfirmPasswordResetRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUserId

`func (o *ConfirmPasswordResetRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConfirmPasswordResetRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConfirmPasswordResetRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


