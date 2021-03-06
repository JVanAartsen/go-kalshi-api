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

// MarketStatsPoint struct for MarketStatsPoint
type MarketStatsPoint struct {
	DollarOpenInterest *int64 `json:"dollar_open_interest,omitempty"`
	DollarVolume *int64 `json:"dollar_volume,omitempty"`
	OpenInterest *int64 `json:"open_interest,omitempty"`
	Price *int64 `json:"price,omitempty"`
	Ts *int64 `json:"ts,omitempty"`
	Volume *int64 `json:"volume,omitempty"`
	YesAsk *int64 `json:"yes_ask,omitempty"`
	YesBid *int64 `json:"yes_bid,omitempty"`
}

// NewMarketStatsPoint instantiates a new MarketStatsPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketStatsPoint() *MarketStatsPoint {
	this := MarketStatsPoint{}
	return &this
}

// NewMarketStatsPointWithDefaults instantiates a new MarketStatsPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketStatsPointWithDefaults() *MarketStatsPoint {
	this := MarketStatsPoint{}
	return &this
}

// GetDollarOpenInterest returns the DollarOpenInterest field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetDollarOpenInterest() int64 {
	if o == nil || o.DollarOpenInterest == nil {
		var ret int64
		return ret
	}
	return *o.DollarOpenInterest
}

// GetDollarOpenInterestOk returns a tuple with the DollarOpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetDollarOpenInterestOk() (*int64, bool) {
	if o == nil || o.DollarOpenInterest == nil {
		return nil, false
	}
	return o.DollarOpenInterest, true
}

// HasDollarOpenInterest returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasDollarOpenInterest() bool {
	if o != nil && o.DollarOpenInterest != nil {
		return true
	}

	return false
}

// SetDollarOpenInterest gets a reference to the given int64 and assigns it to the DollarOpenInterest field.
func (o *MarketStatsPoint) SetDollarOpenInterest(v int64) {
	o.DollarOpenInterest = &v
}

// GetDollarVolume returns the DollarVolume field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetDollarVolume() int64 {
	if o == nil || o.DollarVolume == nil {
		var ret int64
		return ret
	}
	return *o.DollarVolume
}

// GetDollarVolumeOk returns a tuple with the DollarVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetDollarVolumeOk() (*int64, bool) {
	if o == nil || o.DollarVolume == nil {
		return nil, false
	}
	return o.DollarVolume, true
}

// HasDollarVolume returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasDollarVolume() bool {
	if o != nil && o.DollarVolume != nil {
		return true
	}

	return false
}

// SetDollarVolume gets a reference to the given int64 and assigns it to the DollarVolume field.
func (o *MarketStatsPoint) SetDollarVolume(v int64) {
	o.DollarVolume = &v
}

// GetOpenInterest returns the OpenInterest field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetOpenInterest() int64 {
	if o == nil || o.OpenInterest == nil {
		var ret int64
		return ret
	}
	return *o.OpenInterest
}

// GetOpenInterestOk returns a tuple with the OpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetOpenInterestOk() (*int64, bool) {
	if o == nil || o.OpenInterest == nil {
		return nil, false
	}
	return o.OpenInterest, true
}

// HasOpenInterest returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasOpenInterest() bool {
	if o != nil && o.OpenInterest != nil {
		return true
	}

	return false
}

// SetOpenInterest gets a reference to the given int64 and assigns it to the OpenInterest field.
func (o *MarketStatsPoint) SetOpenInterest(v int64) {
	o.OpenInterest = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetPrice() int64 {
	if o == nil || o.Price == nil {
		var ret int64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetPriceOk() (*int64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int64 and assigns it to the Price field.
func (o *MarketStatsPoint) SetPrice(v int64) {
	o.Price = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetTs() int64 {
	if o == nil || o.Ts == nil {
		var ret int64
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetTsOk() (*int64, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given int64 and assigns it to the Ts field.
func (o *MarketStatsPoint) SetTs(v int64) {
	o.Ts = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetVolume() int64 {
	if o == nil || o.Volume == nil {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetVolumeOk() (*int64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *MarketStatsPoint) SetVolume(v int64) {
	o.Volume = &v
}

// GetYesAsk returns the YesAsk field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetYesAsk() int64 {
	if o == nil || o.YesAsk == nil {
		var ret int64
		return ret
	}
	return *o.YesAsk
}

// GetYesAskOk returns a tuple with the YesAsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetYesAskOk() (*int64, bool) {
	if o == nil || o.YesAsk == nil {
		return nil, false
	}
	return o.YesAsk, true
}

// HasYesAsk returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasYesAsk() bool {
	if o != nil && o.YesAsk != nil {
		return true
	}

	return false
}

// SetYesAsk gets a reference to the given int64 and assigns it to the YesAsk field.
func (o *MarketStatsPoint) SetYesAsk(v int64) {
	o.YesAsk = &v
}

// GetYesBid returns the YesBid field value if set, zero value otherwise.
func (o *MarketStatsPoint) GetYesBid() int64 {
	if o == nil || o.YesBid == nil {
		var ret int64
		return ret
	}
	return *o.YesBid
}

// GetYesBidOk returns a tuple with the YesBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketStatsPoint) GetYesBidOk() (*int64, bool) {
	if o == nil || o.YesBid == nil {
		return nil, false
	}
	return o.YesBid, true
}

// HasYesBid returns a boolean if a field has been set.
func (o *MarketStatsPoint) HasYesBid() bool {
	if o != nil && o.YesBid != nil {
		return true
	}

	return false
}

// SetYesBid gets a reference to the given int64 and assigns it to the YesBid field.
func (o *MarketStatsPoint) SetYesBid(v int64) {
	o.YesBid = &v
}

func (o MarketStatsPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DollarOpenInterest != nil {
		toSerialize["dollar_open_interest"] = o.DollarOpenInterest
	}
	if o.DollarVolume != nil {
		toSerialize["dollar_volume"] = o.DollarVolume
	}
	if o.OpenInterest != nil {
		toSerialize["open_interest"] = o.OpenInterest
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.YesAsk != nil {
		toSerialize["yes_ask"] = o.YesAsk
	}
	if o.YesBid != nil {
		toSerialize["yes_bid"] = o.YesBid
	}
	return json.Marshal(toSerialize)
}

type NullableMarketStatsPoint struct {
	value *MarketStatsPoint
	isSet bool
}

func (v NullableMarketStatsPoint) Get() *MarketStatsPoint {
	return v.value
}

func (v *NullableMarketStatsPoint) Set(val *MarketStatsPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketStatsPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketStatsPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketStatsPoint(val *MarketStatsPoint) *NullableMarketStatsPoint {
	return &NullableMarketStatsPoint{value: val, isSet: true}
}

func (v NullableMarketStatsPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketStatsPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


