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

// LoginRequest struct for LoginRequest
type LoginRequest struct {
	// Email should be used as login identification credentials.
	Email string `json:"email"`
	// Password defined in the first step of the sign-up.
	Password string `json:"password"`
}

// NewLoginRequest instantiates a new LoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginRequest(email string, password string) *LoginRequest {
	this := LoginRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewLoginRequestWithDefaults instantiates a new LoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginRequestWithDefaults() *LoginRequest {
	this := LoginRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *LoginRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *LoginRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *LoginRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *LoginRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginRequest) SetPassword(v string) {
	o.Password = v
}

func (o LoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableLoginRequest struct {
	value *LoginRequest
	isSet bool
}

func (v NullableLoginRequest) Get() *LoginRequest {
	return v.value
}

func (v *NullableLoginRequest) Set(val *LoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginRequest(val *LoginRequest) *NullableLoginRequest {
	return &NullableLoginRequest{value: val, isSet: true}
}

func (v NullableLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


