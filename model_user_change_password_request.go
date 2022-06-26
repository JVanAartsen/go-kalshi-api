/*
Kalshi API.

This documentation describes Kalshi's REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kalshiAPI

import (
	"encoding/json"
)

// UserChangePasswordRequest struct for UserChangePasswordRequest
type UserChangePasswordRequest struct {
	// New password value.
	NewPassword string `json:"new_password"`
	// Old password should be passed as a validation parameter.
	OldPassword string `json:"old_password"`
}

// NewUserChangePasswordRequest instantiates a new UserChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChangePasswordRequest(newPassword string, oldPassword string) *UserChangePasswordRequest {
	this := UserChangePasswordRequest{}
	this.NewPassword = newPassword
	this.OldPassword = oldPassword
	return &this
}

// NewUserChangePasswordRequestWithDefaults instantiates a new UserChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChangePasswordRequestWithDefaults() *UserChangePasswordRequest {
	this := UserChangePasswordRequest{}
	return &this
}

// GetNewPassword returns the NewPassword field value
func (o *UserChangePasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UserChangePasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UserChangePasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetOldPassword returns the OldPassword field value
func (o *UserChangePasswordRequest) GetOldPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value
// and a boolean to check if the value has been set.
func (o *UserChangePasswordRequest) GetOldPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldPassword, true
}

// SetOldPassword sets field value
func (o *UserChangePasswordRequest) SetOldPassword(v string) {
	o.OldPassword = v
}

func (o UserChangePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["new_password"] = o.NewPassword
	}
	if true {
		toSerialize["old_password"] = o.OldPassword
	}
	return json.Marshal(toSerialize)
}

type NullableUserChangePasswordRequest struct {
	value *UserChangePasswordRequest
	isSet bool
}

func (v NullableUserChangePasswordRequest) Get() *UserChangePasswordRequest {
	return v.value
}

func (v *NullableUserChangePasswordRequest) Set(val *UserChangePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChangePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChangePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChangePasswordRequest(val *UserChangePasswordRequest) *NullableUserChangePasswordRequest {
	return &NullableUserChangePasswordRequest{value: val, isSet: true}
}

func (v NullableUserChangePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChangePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


