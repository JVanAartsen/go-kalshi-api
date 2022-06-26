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

// UserOrderDecreaseRequest struct for UserOrderDecreaseRequest
type UserOrderDecreaseRequest struct {
	Count *int32 `json:"count,omitempty"`
}

// NewUserOrderDecreaseRequest instantiates a new UserOrderDecreaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOrderDecreaseRequest() *UserOrderDecreaseRequest {
	this := UserOrderDecreaseRequest{}
	return &this
}

// NewUserOrderDecreaseRequestWithDefaults instantiates a new UserOrderDecreaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOrderDecreaseRequestWithDefaults() *UserOrderDecreaseRequest {
	this := UserOrderDecreaseRequest{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UserOrderDecreaseRequest) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOrderDecreaseRequest) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UserOrderDecreaseRequest) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UserOrderDecreaseRequest) SetCount(v int32) {
	o.Count = &v
}

func (o UserOrderDecreaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableUserOrderDecreaseRequest struct {
	value *UserOrderDecreaseRequest
	isSet bool
}

func (v NullableUserOrderDecreaseRequest) Get() *UserOrderDecreaseRequest {
	return v.value
}

func (v *NullableUserOrderDecreaseRequest) Set(val *UserOrderDecreaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOrderDecreaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOrderDecreaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOrderDecreaseRequest(val *UserOrderDecreaseRequest) *NullableUserOrderDecreaseRequest {
	return &NullableUserOrderDecreaseRequest{value: val, isSet: true}
}

func (v NullableUserOrderDecreaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOrderDecreaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


