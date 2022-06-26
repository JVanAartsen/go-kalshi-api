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

// UserBatchOrdersCreateRequest Request for submitting a batch of orders
type UserBatchOrdersCreateRequest struct {
	// An array of individual orders to place
	Orders []UserOrderCreateRequest `json:"orders"`
}

// NewUserBatchOrdersCreateRequest instantiates a new UserBatchOrdersCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBatchOrdersCreateRequest(orders []UserOrderCreateRequest) *UserBatchOrdersCreateRequest {
	this := UserBatchOrdersCreateRequest{}
	this.Orders = orders
	return &this
}

// NewUserBatchOrdersCreateRequestWithDefaults instantiates a new UserBatchOrdersCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBatchOrdersCreateRequestWithDefaults() *UserBatchOrdersCreateRequest {
	this := UserBatchOrdersCreateRequest{}
	return &this
}

// GetOrders returns the Orders field value
func (o *UserBatchOrdersCreateRequest) GetOrders() []UserOrderCreateRequest {
	if o == nil {
		var ret []UserOrderCreateRequest
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *UserBatchOrdersCreateRequest) GetOrdersOk() ([]UserOrderCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orders, true
}

// SetOrders sets field value
func (o *UserBatchOrdersCreateRequest) SetOrders(v []UserOrderCreateRequest) {
	o.Orders = v
}

func (o UserBatchOrdersCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableUserBatchOrdersCreateRequest struct {
	value *UserBatchOrdersCreateRequest
	isSet bool
}

func (v NullableUserBatchOrdersCreateRequest) Get() *UserBatchOrdersCreateRequest {
	return v.value
}

func (v *NullableUserBatchOrdersCreateRequest) Set(val *UserBatchOrdersCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBatchOrdersCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBatchOrdersCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBatchOrdersCreateRequest(val *UserBatchOrdersCreateRequest) *NullableUserBatchOrdersCreateRequest {
	return &NullableUserBatchOrdersCreateRequest{value: val, isSet: true}
}

func (v NullableUserBatchOrdersCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBatchOrdersCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

