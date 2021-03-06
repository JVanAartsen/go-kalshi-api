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

// Withdrawal Represents a withdrawal.
type Withdrawal struct {
	AmountCents *int64 `json:"amount_cents,omitempty"`
	BankId *string `json:"bank_id,omitempty"`
	CreatedTs *time.Time `json:"created_ts,omitempty"`
	Id *string `json:"id,omitempty"`
	ReturnCode *string `json:"return_code,omitempty"`
	ReturnReason *string `json:"return_reason,omitempty"`
	Status *string `json:"status,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// NewWithdrawal instantiates a new Withdrawal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawal() *Withdrawal {
	this := Withdrawal{}
	return &this
}

// NewWithdrawalWithDefaults instantiates a new Withdrawal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawalWithDefaults() *Withdrawal {
	this := Withdrawal{}
	return &this
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *Withdrawal) GetAmountCents() int64 {
	if o == nil || o.AmountCents == nil {
		var ret int64
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetAmountCentsOk() (*int64, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *Withdrawal) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int64 and assigns it to the AmountCents field.
func (o *Withdrawal) SetAmountCents(v int64) {
	o.AmountCents = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *Withdrawal) GetBankId() string {
	if o == nil || o.BankId == nil {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetBankIdOk() (*string, bool) {
	if o == nil || o.BankId == nil {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *Withdrawal) HasBankId() bool {
	if o != nil && o.BankId != nil {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *Withdrawal) SetBankId(v string) {
	o.BankId = &v
}

// GetCreatedTs returns the CreatedTs field value if set, zero value otherwise.
func (o *Withdrawal) GetCreatedTs() time.Time {
	if o == nil || o.CreatedTs == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTs
}

// GetCreatedTsOk returns a tuple with the CreatedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetCreatedTsOk() (*time.Time, bool) {
	if o == nil || o.CreatedTs == nil {
		return nil, false
	}
	return o.CreatedTs, true
}

// HasCreatedTs returns a boolean if a field has been set.
func (o *Withdrawal) HasCreatedTs() bool {
	if o != nil && o.CreatedTs != nil {
		return true
	}

	return false
}

// SetCreatedTs gets a reference to the given time.Time and assigns it to the CreatedTs field.
func (o *Withdrawal) SetCreatedTs(v time.Time) {
	o.CreatedTs = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Withdrawal) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Withdrawal) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Withdrawal) SetId(v string) {
	o.Id = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *Withdrawal) GetReturnCode() string {
	if o == nil || o.ReturnCode == nil {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetReturnCodeOk() (*string, bool) {
	if o == nil || o.ReturnCode == nil {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *Withdrawal) HasReturnCode() bool {
	if o != nil && o.ReturnCode != nil {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *Withdrawal) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *Withdrawal) GetReturnReason() string {
	if o == nil || o.ReturnReason == nil {
		var ret string
		return ret
	}
	return *o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetReturnReasonOk() (*string, bool) {
	if o == nil || o.ReturnReason == nil {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *Withdrawal) HasReturnReason() bool {
	if o != nil && o.ReturnReason != nil {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given string and assigns it to the ReturnReason field.
func (o *Withdrawal) SetReturnReason(v string) {
	o.ReturnReason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Withdrawal) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Withdrawal) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Withdrawal) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Withdrawal) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Withdrawal) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Withdrawal) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Withdrawal) SetUserId(v string) {
	o.UserId = &v
}

func (o Withdrawal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.BankId != nil {
		toSerialize["bank_id"] = o.BankId
	}
	if o.CreatedTs != nil {
		toSerialize["created_ts"] = o.CreatedTs
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReturnCode != nil {
		toSerialize["return_code"] = o.ReturnCode
	}
	if o.ReturnReason != nil {
		toSerialize["return_reason"] = o.ReturnReason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawal struct {
	value *Withdrawal
	isSet bool
}

func (v NullableWithdrawal) Get() *Withdrawal {
	return v.value
}

func (v *NullableWithdrawal) Set(val *Withdrawal) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawal) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawal(val *Withdrawal) *NullableWithdrawal {
	return &NullableWithdrawal{value: val, isSet: true}
}

func (v NullableWithdrawal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


