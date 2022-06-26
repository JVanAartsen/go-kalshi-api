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

// DeprecatedUserGetPortfolioHistoryResponse struct for DeprecatedUserGetPortfolioHistoryResponse
type DeprecatedUserGetPortfolioHistoryResponse struct {
	Values []DeprecatedPortfolioMeasurement `json:"values,omitempty"`
}

// NewDeprecatedUserGetPortfolioHistoryResponse instantiates a new DeprecatedUserGetPortfolioHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedUserGetPortfolioHistoryResponse() *DeprecatedUserGetPortfolioHistoryResponse {
	this := DeprecatedUserGetPortfolioHistoryResponse{}
	return &this
}

// NewDeprecatedUserGetPortfolioHistoryResponseWithDefaults instantiates a new DeprecatedUserGetPortfolioHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedUserGetPortfolioHistoryResponseWithDefaults() *DeprecatedUserGetPortfolioHistoryResponse {
	this := DeprecatedUserGetPortfolioHistoryResponse{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DeprecatedUserGetPortfolioHistoryResponse) GetValues() []DeprecatedPortfolioMeasurement {
	if o == nil || o.Values == nil {
		var ret []DeprecatedPortfolioMeasurement
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedUserGetPortfolioHistoryResponse) GetValuesOk() ([]DeprecatedPortfolioMeasurement, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DeprecatedUserGetPortfolioHistoryResponse) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []DeprecatedPortfolioMeasurement and assigns it to the Values field.
func (o *DeprecatedUserGetPortfolioHistoryResponse) SetValues(v []DeprecatedPortfolioMeasurement) {
	o.Values = v
}

func (o DeprecatedUserGetPortfolioHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedUserGetPortfolioHistoryResponse struct {
	value *DeprecatedUserGetPortfolioHistoryResponse
	isSet bool
}

func (v NullableDeprecatedUserGetPortfolioHistoryResponse) Get() *DeprecatedUserGetPortfolioHistoryResponse {
	return v.value
}

func (v *NullableDeprecatedUserGetPortfolioHistoryResponse) Set(val *DeprecatedUserGetPortfolioHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedUserGetPortfolioHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedUserGetPortfolioHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedUserGetPortfolioHistoryResponse(val *DeprecatedUserGetPortfolioHistoryResponse) *NullableDeprecatedUserGetPortfolioHistoryResponse {
	return &NullableDeprecatedUserGetPortfolioHistoryResponse{value: val, isSet: true}
}

func (v NullableDeprecatedUserGetPortfolioHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedUserGetPortfolioHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

