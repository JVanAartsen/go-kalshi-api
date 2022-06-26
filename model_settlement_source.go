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

// SettlementSource struct for SettlementSource
type SettlementSource struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewSettlementSource instantiates a new SettlementSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettlementSource() *SettlementSource {
	this := SettlementSource{}
	return &this
}

// NewSettlementSourceWithDefaults instantiates a new SettlementSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettlementSourceWithDefaults() *SettlementSource {
	this := SettlementSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SettlementSource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementSource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SettlementSource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SettlementSource) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SettlementSource) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementSource) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SettlementSource) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SettlementSource) SetUrl(v string) {
	o.Url = &v
}

func (o SettlementSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableSettlementSource struct {
	value *SettlementSource
	isSet bool
}

func (v NullableSettlementSource) Get() *SettlementSource {
	return v.value
}

func (v *NullableSettlementSource) Set(val *SettlementSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSettlementSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSettlementSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettlementSource(val *SettlementSource) *NullableSettlementSource {
	return &NullableSettlementSource{value: val, isSet: true}
}

func (v NullableSettlementSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettlementSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

