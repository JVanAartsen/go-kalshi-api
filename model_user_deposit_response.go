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

// UserDepositResponse struct for UserDepositResponse
type UserDepositResponse struct {
	Deposit *Deposit `json:"deposit,omitempty"`
	// Id for the deposit that was created.
	DepositId string `json:"deposit_id"`
	// The estimated number of days we believe the ach transfer will take
	EstimatedAchTimeDays int32 `json:"estimated_ach_time_days"`
}

// NewUserDepositResponse instantiates a new UserDepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDepositResponse(depositId string, estimatedAchTimeDays int32) *UserDepositResponse {
	this := UserDepositResponse{}
	this.DepositId = depositId
	this.EstimatedAchTimeDays = estimatedAchTimeDays
	return &this
}

// NewUserDepositResponseWithDefaults instantiates a new UserDepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDepositResponseWithDefaults() *UserDepositResponse {
	this := UserDepositResponse{}
	return &this
}

// GetDeposit returns the Deposit field value if set, zero value otherwise.
func (o *UserDepositResponse) GetDeposit() Deposit {
	if o == nil || o.Deposit == nil {
		var ret Deposit
		return ret
	}
	return *o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDepositResponse) GetDepositOk() (*Deposit, bool) {
	if o == nil || o.Deposit == nil {
		return nil, false
	}
	return o.Deposit, true
}

// HasDeposit returns a boolean if a field has been set.
func (o *UserDepositResponse) HasDeposit() bool {
	if o != nil && o.Deposit != nil {
		return true
	}

	return false
}

// SetDeposit gets a reference to the given Deposit and assigns it to the Deposit field.
func (o *UserDepositResponse) SetDeposit(v Deposit) {
	o.Deposit = &v
}

// GetDepositId returns the DepositId field value
func (o *UserDepositResponse) GetDepositId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositId
}

// GetDepositIdOk returns a tuple with the DepositId field value
// and a boolean to check if the value has been set.
func (o *UserDepositResponse) GetDepositIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositId, true
}

// SetDepositId sets field value
func (o *UserDepositResponse) SetDepositId(v string) {
	o.DepositId = v
}

// GetEstimatedAchTimeDays returns the EstimatedAchTimeDays field value
func (o *UserDepositResponse) GetEstimatedAchTimeDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedAchTimeDays
}

// GetEstimatedAchTimeDaysOk returns a tuple with the EstimatedAchTimeDays field value
// and a boolean to check if the value has been set.
func (o *UserDepositResponse) GetEstimatedAchTimeDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedAchTimeDays, true
}

// SetEstimatedAchTimeDays sets field value
func (o *UserDepositResponse) SetEstimatedAchTimeDays(v int32) {
	o.EstimatedAchTimeDays = v
}

func (o UserDepositResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deposit != nil {
		toSerialize["deposit"] = o.Deposit
	}
	if true {
		toSerialize["deposit_id"] = o.DepositId
	}
	if true {
		toSerialize["estimated_ach_time_days"] = o.EstimatedAchTimeDays
	}
	return json.Marshal(toSerialize)
}

type NullableUserDepositResponse struct {
	value *UserDepositResponse
	isSet bool
}

func (v NullableUserDepositResponse) Get() *UserDepositResponse {
	return v.value
}

func (v *NullableUserDepositResponse) Set(val *UserDepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDepositResponse(val *UserDepositResponse) *NullableUserDepositResponse {
	return &NullableUserDepositResponse{value: val, isSet: true}
}

func (v NullableUserDepositResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


