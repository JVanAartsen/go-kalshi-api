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

// GetEventByTickerResponse struct for GetEventByTickerResponse
type GetEventByTickerResponse struct {
	Event EventData `json:"event"`
}

// NewGetEventByTickerResponse instantiates a new GetEventByTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventByTickerResponse(event EventData) *GetEventByTickerResponse {
	this := GetEventByTickerResponse{}
	this.Event = event
	return &this
}

// NewGetEventByTickerResponseWithDefaults instantiates a new GetEventByTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventByTickerResponseWithDefaults() *GetEventByTickerResponse {
	this := GetEventByTickerResponse{}
	return &this
}

// GetEvent returns the Event field value
func (o *GetEventByTickerResponse) GetEvent() EventData {
	if o == nil {
		var ret EventData
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *GetEventByTickerResponse) GetEventOk() (*EventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *GetEventByTickerResponse) SetEvent(v EventData) {
	o.Event = v
}

func (o GetEventByTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableGetEventByTickerResponse struct {
	value *GetEventByTickerResponse
	isSet bool
}

func (v NullableGetEventByTickerResponse) Get() *GetEventByTickerResponse {
	return v.value
}

func (v *NullableGetEventByTickerResponse) Set(val *GetEventByTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventByTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventByTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventByTickerResponse(val *GetEventByTickerResponse) *NullableGetEventByTickerResponse {
	return &NullableGetEventByTickerResponse{value: val, isSet: true}
}

func (v NullableGetEventByTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventByTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


