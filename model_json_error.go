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

// JSONError struct for JSONError
type JSONError struct {
	Code *string `json:"code,omitempty"`
	Details *string `json:"details,omitempty"`
	Message *string `json:"message,omitempty"`
	Service *string `json:"service,omitempty"`
}

// NewJSONError instantiates a new JSONError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONError() *JSONError {
	this := JSONError{}
	return &this
}

// NewJSONErrorWithDefaults instantiates a new JSONError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONErrorWithDefaults() *JSONError {
	this := JSONError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *JSONError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *JSONError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *JSONError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *JSONError) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONError) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *JSONError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *JSONError) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *JSONError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *JSONError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *JSONError) SetMessage(v string) {
	o.Message = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *JSONError) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONError) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *JSONError) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *JSONError) SetService(v string) {
	o.Service = &v
}

func (o JSONError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableJSONError struct {
	value *JSONError
	isSet bool
}

func (v NullableJSONError) Get() *JSONError {
	return v.value
}

func (v *NullableJSONError) Set(val *JSONError) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONError) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONError(val *JSONError) *NullableJSONError {
	return &NullableJSONError{value: val, isSet: true}
}

func (v NullableJSONError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

