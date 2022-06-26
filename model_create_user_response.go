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

// CreateUserResponse Response for submitting an order
type CreateUserResponse struct {
	// swagger: ignore
	Code *string `json:"code,omitempty"`
	// user_id for the created user.
	UserId string `json:"user_id"`
}

// NewCreateUserResponse instantiates a new CreateUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserResponse(userId string) *CreateUserResponse {
	this := CreateUserResponse{}
	this.UserId = userId
	return &this
}

// NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserResponseWithDefaults() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateUserResponse) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateUserResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateUserResponse) SetCode(v string) {
	o.Code = &v
}

// GetUserId returns the UserId field value
func (o *CreateUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateUserResponse) SetUserId(v string) {
	o.UserId = v
}

func (o CreateUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUserResponse struct {
	value *CreateUserResponse
	isSet bool
}

func (v NullableCreateUserResponse) Get() *CreateUserResponse {
	return v.value
}

func (v *NullableCreateUserResponse) Set(val *CreateUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponse(val *CreateUserResponse) *NullableCreateUserResponse {
	return &NullableCreateUserResponse{value: val, isSet: true}
}

func (v NullableCreateUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


