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

// UserTradesGetResponse struct for UserTradesGetResponse
type UserTradesGetResponse struct {
	Trades []UserTrade `json:"trades"`
}

// NewUserTradesGetResponse instantiates a new UserTradesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTradesGetResponse(trades []UserTrade) *UserTradesGetResponse {
	this := UserTradesGetResponse{}
	this.Trades = trades
	return &this
}

// NewUserTradesGetResponseWithDefaults instantiates a new UserTradesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTradesGetResponseWithDefaults() *UserTradesGetResponse {
	this := UserTradesGetResponse{}
	return &this
}

// GetTrades returns the Trades field value
func (o *UserTradesGetResponse) GetTrades() []UserTrade {
	if o == nil {
		var ret []UserTrade
		return ret
	}

	return o.Trades
}

// GetTradesOk returns a tuple with the Trades field value
// and a boolean to check if the value has been set.
func (o *UserTradesGetResponse) GetTradesOk() ([]UserTrade, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trades, true
}

// SetTrades sets field value
func (o *UserTradesGetResponse) SetTrades(v []UserTrade) {
	o.Trades = v
}

func (o UserTradesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trades"] = o.Trades
	}
	return json.Marshal(toSerialize)
}

type NullableUserTradesGetResponse struct {
	value *UserTradesGetResponse
	isSet bool
}

func (v NullableUserTradesGetResponse) Get() *UserTradesGetResponse {
	return v.value
}

func (v *NullableUserTradesGetResponse) Set(val *UserTradesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTradesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTradesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTradesGetResponse(val *UserTradesGetResponse) *NullableUserTradesGetResponse {
	return &NullableUserTradesGetResponse{value: val, isSet: true}
}

func (v NullableUserTradesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTradesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

