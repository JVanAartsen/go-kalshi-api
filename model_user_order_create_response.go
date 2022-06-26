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

// UserOrderCreateResponse Response for submitting an order
type UserOrderCreateResponse struct {
	Order Order `json:"order"`
	// Status of the order submit operation
	Status string `json:"status"`
}

// NewUserOrderCreateResponse instantiates a new UserOrderCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOrderCreateResponse(order Order, status string) *UserOrderCreateResponse {
	this := UserOrderCreateResponse{}
	this.Order = order
	this.Status = status
	return &this
}

// NewUserOrderCreateResponseWithDefaults instantiates a new UserOrderCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOrderCreateResponseWithDefaults() *UserOrderCreateResponse {
	this := UserOrderCreateResponse{}
	return &this
}

// GetOrder returns the Order field value
func (o *UserOrderCreateResponse) GetOrder() Order {
	if o == nil {
		var ret Order
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *UserOrderCreateResponse) GetOrderOk() (*Order, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *UserOrderCreateResponse) SetOrder(v Order) {
	o.Order = v
}

// GetStatus returns the Status field value
func (o *UserOrderCreateResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserOrderCreateResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserOrderCreateResponse) SetStatus(v string) {
	o.Status = v
}

func (o UserOrderCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUserOrderCreateResponse struct {
	value *UserOrderCreateResponse
	isSet bool
}

func (v NullableUserOrderCreateResponse) Get() *UserOrderCreateResponse {
	return v.value
}

func (v *NullableUserOrderCreateResponse) Set(val *UserOrderCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOrderCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOrderCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOrderCreateResponse(val *UserOrderCreateResponse) *NullableUserOrderCreateResponse {
	return &NullableUserOrderCreateResponse{value: val, isSet: true}
}

func (v NullableUserOrderCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOrderCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

