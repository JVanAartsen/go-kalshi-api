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

// GetUserDepositsResponse struct for GetUserDepositsResponse
type GetUserDepositsResponse struct {
	// List of previous deposits for the user
	Deposits []Deposit `json:"deposits"`
}

// NewGetUserDepositsResponse instantiates a new GetUserDepositsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserDepositsResponse(deposits []Deposit) *GetUserDepositsResponse {
	this := GetUserDepositsResponse{}
	this.Deposits = deposits
	return &this
}

// NewGetUserDepositsResponseWithDefaults instantiates a new GetUserDepositsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserDepositsResponseWithDefaults() *GetUserDepositsResponse {
	this := GetUserDepositsResponse{}
	return &this
}

// GetDeposits returns the Deposits field value
func (o *GetUserDepositsResponse) GetDeposits() []Deposit {
	if o == nil {
		var ret []Deposit
		return ret
	}

	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value
// and a boolean to check if the value has been set.
func (o *GetUserDepositsResponse) GetDepositsOk() ([]Deposit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deposits, true
}

// SetDeposits sets field value
func (o *GetUserDepositsResponse) SetDeposits(v []Deposit) {
	o.Deposits = v
}

func (o GetUserDepositsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposits"] = o.Deposits
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserDepositsResponse struct {
	value *GetUserDepositsResponse
	isSet bool
}

func (v NullableGetUserDepositsResponse) Get() *GetUserDepositsResponse {
	return v.value
}

func (v *NullableGetUserDepositsResponse) Set(val *GetUserDepositsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserDepositsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserDepositsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserDepositsResponse(val *GetUserDepositsResponse) *NullableGetUserDepositsResponse {
	return &NullableGetUserDepositsResponse{value: val, isSet: true}
}

func (v NullableGetUserDepositsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserDepositsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


