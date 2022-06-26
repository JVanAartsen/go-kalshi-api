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

// LoginResponse Response for login request
type LoginResponse struct {
	// Access level, to limit the access that applicants have
	AccessLevel string `json:"access_level"`
	// Access token for an member role session in the api
	Token string `json:"token"`
	// Your user_id, this will be required in all requests under the /users prefix
	UserId string `json:"user_id"`
}

// NewLoginResponse instantiates a new LoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponse(accessLevel string, token string, userId string) *LoginResponse {
	this := LoginResponse{}
	this.AccessLevel = accessLevel
	this.Token = token
	this.UserId = userId
	return &this
}

// NewLoginResponseWithDefaults instantiates a new LoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseWithDefaults() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// GetAccessLevel returns the AccessLevel field value
func (o *LoginResponse) GetAccessLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetAccessLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *LoginResponse) SetAccessLevel(v string) {
	o.AccessLevel = v
}

// GetToken returns the Token field value
func (o *LoginResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *LoginResponse) SetToken(v string) {
	o.Token = v
}

// GetUserId returns the UserId field value
func (o *LoginResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *LoginResponse) SetUserId(v string) {
	o.UserId = v
}

func (o LoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_level"] = o.AccessLevel
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableLoginResponse struct {
	value *LoginResponse
	isSet bool
}

func (v NullableLoginResponse) Get() *LoginResponse {
	return v.value
}

func (v *NullableLoginResponse) Set(val *LoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponse(val *LoginResponse) *NullableLoginResponse {
	return &NullableLoginResponse{value: val, isSet: true}
}

func (v NullableLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

