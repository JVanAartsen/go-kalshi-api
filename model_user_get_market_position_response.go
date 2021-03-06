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

// UserGetMarketPositionResponse struct for UserGetMarketPositionResponse
type UserGetMarketPositionResponse struct {
	MarketPosition MarketPosition `json:"market_position"`
}

// NewUserGetMarketPositionResponse instantiates a new UserGetMarketPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetMarketPositionResponse(marketPosition MarketPosition) *UserGetMarketPositionResponse {
	this := UserGetMarketPositionResponse{}
	this.MarketPosition = marketPosition
	return &this
}

// NewUserGetMarketPositionResponseWithDefaults instantiates a new UserGetMarketPositionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetMarketPositionResponseWithDefaults() *UserGetMarketPositionResponse {
	this := UserGetMarketPositionResponse{}
	return &this
}

// GetMarketPosition returns the MarketPosition field value
func (o *UserGetMarketPositionResponse) GetMarketPosition() MarketPosition {
	if o == nil {
		var ret MarketPosition
		return ret
	}

	return o.MarketPosition
}

// GetMarketPositionOk returns a tuple with the MarketPosition field value
// and a boolean to check if the value has been set.
func (o *UserGetMarketPositionResponse) GetMarketPositionOk() (*MarketPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketPosition, true
}

// SetMarketPosition sets field value
func (o *UserGetMarketPositionResponse) SetMarketPosition(v MarketPosition) {
	o.MarketPosition = v
}

func (o UserGetMarketPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["market_position"] = o.MarketPosition
	}
	return json.Marshal(toSerialize)
}

type NullableUserGetMarketPositionResponse struct {
	value *UserGetMarketPositionResponse
	isSet bool
}

func (v NullableUserGetMarketPositionResponse) Get() *UserGetMarketPositionResponse {
	return v.value
}

func (v *NullableUserGetMarketPositionResponse) Set(val *UserGetMarketPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetMarketPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetMarketPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetMarketPositionResponse(val *UserGetMarketPositionResponse) *NullableUserGetMarketPositionResponse {
	return &NullableUserGetMarketPositionResponse{value: val, isSet: true}
}

func (v NullableUserGetMarketPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetMarketPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


