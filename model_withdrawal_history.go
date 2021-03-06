/*
Kalshi API.

This documentation describes Kalshi's REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kalshiAPI

import (
	"encoding/json"
	"time"
)

// WithdrawalHistory struct for WithdrawalHistory
type WithdrawalHistory struct {
	Amount *int64 `json:"amount,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Fee *int64 `json:"fee,omitempty"`
	ReturnedAmount *int64 `json:"returned_amount,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewWithdrawalHistory instantiates a new WithdrawalHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawalHistory() *WithdrawalHistory {
	this := WithdrawalHistory{}
	return &this
}

// NewWithdrawalHistoryWithDefaults instantiates a new WithdrawalHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawalHistoryWithDefaults() *WithdrawalHistory {
	this := WithdrawalHistory{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *WithdrawalHistory) SetAmount(v int64) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WithdrawalHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetFee() int64 {
	if o == nil || o.Fee == nil {
		var ret int64
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetFeeOk() (*int64, bool) {
	if o == nil || o.Fee == nil {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasFee() bool {
	if o != nil && o.Fee != nil {
		return true
	}

	return false
}

// SetFee gets a reference to the given int64 and assigns it to the Fee field.
func (o *WithdrawalHistory) SetFee(v int64) {
	o.Fee = &v
}

// GetReturnedAmount returns the ReturnedAmount field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetReturnedAmount() int64 {
	if o == nil || o.ReturnedAmount == nil {
		var ret int64
		return ret
	}
	return *o.ReturnedAmount
}

// GetReturnedAmountOk returns a tuple with the ReturnedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetReturnedAmountOk() (*int64, bool) {
	if o == nil || o.ReturnedAmount == nil {
		return nil, false
	}
	return o.ReturnedAmount, true
}

// HasReturnedAmount returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasReturnedAmount() bool {
	if o != nil && o.ReturnedAmount != nil {
		return true
	}

	return false
}

// SetReturnedAmount gets a reference to the given int64 and assigns it to the ReturnedAmount field.
func (o *WithdrawalHistory) SetReturnedAmount(v int64) {
	o.ReturnedAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WithdrawalHistory) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WithdrawalHistory) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WithdrawalHistory) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WithdrawalHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o WithdrawalHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Fee != nil {
		toSerialize["fee"] = o.Fee
	}
	if o.ReturnedAmount != nil {
		toSerialize["returned_amount"] = o.ReturnedAmount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawalHistory struct {
	value *WithdrawalHistory
	isSet bool
}

func (v NullableWithdrawalHistory) Get() *WithdrawalHistory {
	return v.value
}

func (v *NullableWithdrawalHistory) Set(val *WithdrawalHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawalHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawalHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawalHistory(val *WithdrawalHistory) *NullableWithdrawalHistory {
	return &NullableWithdrawalHistory{value: val, isSet: true}
}

func (v NullableWithdrawalHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawalHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


