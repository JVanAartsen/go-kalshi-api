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

// DepositHistory Represents a deposit account history item
type DepositHistory struct {
	Amount *int64 `json:"amount,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DepositType *string `json:"deposit_type,omitempty"`
	Fee *int64 `json:"fee,omitempty"`
	ImmediateAmount *int64 `json:"immediate_amount,omitempty"`
	ImmediateStatus *string `json:"immediate_status,omitempty"`
	ReturnedAmount *int64 `json:"returned_amount,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDepositHistory instantiates a new DepositHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistory() *DepositHistory {
	this := DepositHistory{}
	return &this
}

// NewDepositHistoryWithDefaults instantiates a new DepositHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryWithDefaults() *DepositHistory {
	this := DepositHistory{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DepositHistory) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DepositHistory) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *DepositHistory) SetAmount(v int64) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DepositHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DepositHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DepositHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDepositType returns the DepositType field value if set, zero value otherwise.
func (o *DepositHistory) GetDepositType() string {
	if o == nil || o.DepositType == nil {
		var ret string
		return ret
	}
	return *o.DepositType
}

// GetDepositTypeOk returns a tuple with the DepositType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetDepositTypeOk() (*string, bool) {
	if o == nil || o.DepositType == nil {
		return nil, false
	}
	return o.DepositType, true
}

// HasDepositType returns a boolean if a field has been set.
func (o *DepositHistory) HasDepositType() bool {
	if o != nil && o.DepositType != nil {
		return true
	}

	return false
}

// SetDepositType gets a reference to the given string and assigns it to the DepositType field.
func (o *DepositHistory) SetDepositType(v string) {
	o.DepositType = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *DepositHistory) GetFee() int64 {
	if o == nil || o.Fee == nil {
		var ret int64
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetFeeOk() (*int64, bool) {
	if o == nil || o.Fee == nil {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *DepositHistory) HasFee() bool {
	if o != nil && o.Fee != nil {
		return true
	}

	return false
}

// SetFee gets a reference to the given int64 and assigns it to the Fee field.
func (o *DepositHistory) SetFee(v int64) {
	o.Fee = &v
}

// GetImmediateAmount returns the ImmediateAmount field value if set, zero value otherwise.
func (o *DepositHistory) GetImmediateAmount() int64 {
	if o == nil || o.ImmediateAmount == nil {
		var ret int64
		return ret
	}
	return *o.ImmediateAmount
}

// GetImmediateAmountOk returns a tuple with the ImmediateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetImmediateAmountOk() (*int64, bool) {
	if o == nil || o.ImmediateAmount == nil {
		return nil, false
	}
	return o.ImmediateAmount, true
}

// HasImmediateAmount returns a boolean if a field has been set.
func (o *DepositHistory) HasImmediateAmount() bool {
	if o != nil && o.ImmediateAmount != nil {
		return true
	}

	return false
}

// SetImmediateAmount gets a reference to the given int64 and assigns it to the ImmediateAmount field.
func (o *DepositHistory) SetImmediateAmount(v int64) {
	o.ImmediateAmount = &v
}

// GetImmediateStatus returns the ImmediateStatus field value if set, zero value otherwise.
func (o *DepositHistory) GetImmediateStatus() string {
	if o == nil || o.ImmediateStatus == nil {
		var ret string
		return ret
	}
	return *o.ImmediateStatus
}

// GetImmediateStatusOk returns a tuple with the ImmediateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetImmediateStatusOk() (*string, bool) {
	if o == nil || o.ImmediateStatus == nil {
		return nil, false
	}
	return o.ImmediateStatus, true
}

// HasImmediateStatus returns a boolean if a field has been set.
func (o *DepositHistory) HasImmediateStatus() bool {
	if o != nil && o.ImmediateStatus != nil {
		return true
	}

	return false
}

// SetImmediateStatus gets a reference to the given string and assigns it to the ImmediateStatus field.
func (o *DepositHistory) SetImmediateStatus(v string) {
	o.ImmediateStatus = &v
}

// GetReturnedAmount returns the ReturnedAmount field value if set, zero value otherwise.
func (o *DepositHistory) GetReturnedAmount() int64 {
	if o == nil || o.ReturnedAmount == nil {
		var ret int64
		return ret
	}
	return *o.ReturnedAmount
}

// GetReturnedAmountOk returns a tuple with the ReturnedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetReturnedAmountOk() (*int64, bool) {
	if o == nil || o.ReturnedAmount == nil {
		return nil, false
	}
	return o.ReturnedAmount, true
}

// HasReturnedAmount returns a boolean if a field has been set.
func (o *DepositHistory) HasReturnedAmount() bool {
	if o != nil && o.ReturnedAmount != nil {
		return true
	}

	return false
}

// SetReturnedAmount gets a reference to the given int64 and assigns it to the ReturnedAmount field.
func (o *DepositHistory) SetReturnedAmount(v int64) {
	o.ReturnedAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DepositHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DepositHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DepositHistory) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DepositHistory) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DepositHistory) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DepositHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DepositHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DepositType != nil {
		toSerialize["deposit_type"] = o.DepositType
	}
	if o.Fee != nil {
		toSerialize["fee"] = o.Fee
	}
	if o.ImmediateAmount != nil {
		toSerialize["immediate_amount"] = o.ImmediateAmount
	}
	if o.ImmediateStatus != nil {
		toSerialize["immediate_status"] = o.ImmediateStatus
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

type NullableDepositHistory struct {
	value *DepositHistory
	isSet bool
}

func (v NullableDepositHistory) Get() *DepositHistory {
	return v.value
}

func (v *NullableDepositHistory) Set(val *DepositHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositHistory(val *DepositHistory) *NullableDepositHistory {
	return &NullableDepositHistory{value: val, isSet: true}
}

func (v NullableDepositHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

