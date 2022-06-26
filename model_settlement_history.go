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

// SettlementHistory struct for SettlementHistory
type SettlementHistory struct {
	DeterminedTime *time.Time `json:"determined_time,omitempty"`
	MarketId *string `json:"market_id,omitempty"`
	MarketResult *string `json:"market_result,omitempty"`
	MarketTitle *string `json:"market_title,omitempty"`
	NoCount *int64 `json:"no_count,omitempty"`
	NoTotalCost *int64 `json:"no_total_cost,omitempty"`
	Profit *int64 `json:"profit,omitempty"`
	SettledTime *time.Time `json:"settled_time,omitempty"`
	YesCount *int64 `json:"yes_count,omitempty"`
	YesTotalCost *int64 `json:"yes_total_cost,omitempty"`
}

// NewSettlementHistory instantiates a new SettlementHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettlementHistory() *SettlementHistory {
	this := SettlementHistory{}
	return &this
}

// NewSettlementHistoryWithDefaults instantiates a new SettlementHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettlementHistoryWithDefaults() *SettlementHistory {
	this := SettlementHistory{}
	return &this
}

// GetDeterminedTime returns the DeterminedTime field value if set, zero value otherwise.
func (o *SettlementHistory) GetDeterminedTime() time.Time {
	if o == nil || o.DeterminedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DeterminedTime
}

// GetDeterminedTimeOk returns a tuple with the DeterminedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetDeterminedTimeOk() (*time.Time, bool) {
	if o == nil || o.DeterminedTime == nil {
		return nil, false
	}
	return o.DeterminedTime, true
}

// HasDeterminedTime returns a boolean if a field has been set.
func (o *SettlementHistory) HasDeterminedTime() bool {
	if o != nil && o.DeterminedTime != nil {
		return true
	}

	return false
}

// SetDeterminedTime gets a reference to the given time.Time and assigns it to the DeterminedTime field.
func (o *SettlementHistory) SetDeterminedTime(v time.Time) {
	o.DeterminedTime = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *SettlementHistory) GetMarketId() string {
	if o == nil || o.MarketId == nil {
		var ret string
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetMarketIdOk() (*string, bool) {
	if o == nil || o.MarketId == nil {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *SettlementHistory) HasMarketId() bool {
	if o != nil && o.MarketId != nil {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given string and assigns it to the MarketId field.
func (o *SettlementHistory) SetMarketId(v string) {
	o.MarketId = &v
}

// GetMarketResult returns the MarketResult field value if set, zero value otherwise.
func (o *SettlementHistory) GetMarketResult() string {
	if o == nil || o.MarketResult == nil {
		var ret string
		return ret
	}
	return *o.MarketResult
}

// GetMarketResultOk returns a tuple with the MarketResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetMarketResultOk() (*string, bool) {
	if o == nil || o.MarketResult == nil {
		return nil, false
	}
	return o.MarketResult, true
}

// HasMarketResult returns a boolean if a field has been set.
func (o *SettlementHistory) HasMarketResult() bool {
	if o != nil && o.MarketResult != nil {
		return true
	}

	return false
}

// SetMarketResult gets a reference to the given string and assigns it to the MarketResult field.
func (o *SettlementHistory) SetMarketResult(v string) {
	o.MarketResult = &v
}

// GetMarketTitle returns the MarketTitle field value if set, zero value otherwise.
func (o *SettlementHistory) GetMarketTitle() string {
	if o == nil || o.MarketTitle == nil {
		var ret string
		return ret
	}
	return *o.MarketTitle
}

// GetMarketTitleOk returns a tuple with the MarketTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetMarketTitleOk() (*string, bool) {
	if o == nil || o.MarketTitle == nil {
		return nil, false
	}
	return o.MarketTitle, true
}

// HasMarketTitle returns a boolean if a field has been set.
func (o *SettlementHistory) HasMarketTitle() bool {
	if o != nil && o.MarketTitle != nil {
		return true
	}

	return false
}

// SetMarketTitle gets a reference to the given string and assigns it to the MarketTitle field.
func (o *SettlementHistory) SetMarketTitle(v string) {
	o.MarketTitle = &v
}

// GetNoCount returns the NoCount field value if set, zero value otherwise.
func (o *SettlementHistory) GetNoCount() int64 {
	if o == nil || o.NoCount == nil {
		var ret int64
		return ret
	}
	return *o.NoCount
}

// GetNoCountOk returns a tuple with the NoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetNoCountOk() (*int64, bool) {
	if o == nil || o.NoCount == nil {
		return nil, false
	}
	return o.NoCount, true
}

// HasNoCount returns a boolean if a field has been set.
func (o *SettlementHistory) HasNoCount() bool {
	if o != nil && o.NoCount != nil {
		return true
	}

	return false
}

// SetNoCount gets a reference to the given int64 and assigns it to the NoCount field.
func (o *SettlementHistory) SetNoCount(v int64) {
	o.NoCount = &v
}

// GetNoTotalCost returns the NoTotalCost field value if set, zero value otherwise.
func (o *SettlementHistory) GetNoTotalCost() int64 {
	if o == nil || o.NoTotalCost == nil {
		var ret int64
		return ret
	}
	return *o.NoTotalCost
}

// GetNoTotalCostOk returns a tuple with the NoTotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetNoTotalCostOk() (*int64, bool) {
	if o == nil || o.NoTotalCost == nil {
		return nil, false
	}
	return o.NoTotalCost, true
}

// HasNoTotalCost returns a boolean if a field has been set.
func (o *SettlementHistory) HasNoTotalCost() bool {
	if o != nil && o.NoTotalCost != nil {
		return true
	}

	return false
}

// SetNoTotalCost gets a reference to the given int64 and assigns it to the NoTotalCost field.
func (o *SettlementHistory) SetNoTotalCost(v int64) {
	o.NoTotalCost = &v
}

// GetProfit returns the Profit field value if set, zero value otherwise.
func (o *SettlementHistory) GetProfit() int64 {
	if o == nil || o.Profit == nil {
		var ret int64
		return ret
	}
	return *o.Profit
}

// GetProfitOk returns a tuple with the Profit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetProfitOk() (*int64, bool) {
	if o == nil || o.Profit == nil {
		return nil, false
	}
	return o.Profit, true
}

// HasProfit returns a boolean if a field has been set.
func (o *SettlementHistory) HasProfit() bool {
	if o != nil && o.Profit != nil {
		return true
	}

	return false
}

// SetProfit gets a reference to the given int64 and assigns it to the Profit field.
func (o *SettlementHistory) SetProfit(v int64) {
	o.Profit = &v
}

// GetSettledTime returns the SettledTime field value if set, zero value otherwise.
func (o *SettlementHistory) GetSettledTime() time.Time {
	if o == nil || o.SettledTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SettledTime
}

// GetSettledTimeOk returns a tuple with the SettledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetSettledTimeOk() (*time.Time, bool) {
	if o == nil || o.SettledTime == nil {
		return nil, false
	}
	return o.SettledTime, true
}

// HasSettledTime returns a boolean if a field has been set.
func (o *SettlementHistory) HasSettledTime() bool {
	if o != nil && o.SettledTime != nil {
		return true
	}

	return false
}

// SetSettledTime gets a reference to the given time.Time and assigns it to the SettledTime field.
func (o *SettlementHistory) SetSettledTime(v time.Time) {
	o.SettledTime = &v
}

// GetYesCount returns the YesCount field value if set, zero value otherwise.
func (o *SettlementHistory) GetYesCount() int64 {
	if o == nil || o.YesCount == nil {
		var ret int64
		return ret
	}
	return *o.YesCount
}

// GetYesCountOk returns a tuple with the YesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetYesCountOk() (*int64, bool) {
	if o == nil || o.YesCount == nil {
		return nil, false
	}
	return o.YesCount, true
}

// HasYesCount returns a boolean if a field has been set.
func (o *SettlementHistory) HasYesCount() bool {
	if o != nil && o.YesCount != nil {
		return true
	}

	return false
}

// SetYesCount gets a reference to the given int64 and assigns it to the YesCount field.
func (o *SettlementHistory) SetYesCount(v int64) {
	o.YesCount = &v
}

// GetYesTotalCost returns the YesTotalCost field value if set, zero value otherwise.
func (o *SettlementHistory) GetYesTotalCost() int64 {
	if o == nil || o.YesTotalCost == nil {
		var ret int64
		return ret
	}
	return *o.YesTotalCost
}

// GetYesTotalCostOk returns a tuple with the YesTotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementHistory) GetYesTotalCostOk() (*int64, bool) {
	if o == nil || o.YesTotalCost == nil {
		return nil, false
	}
	return o.YesTotalCost, true
}

// HasYesTotalCost returns a boolean if a field has been set.
func (o *SettlementHistory) HasYesTotalCost() bool {
	if o != nil && o.YesTotalCost != nil {
		return true
	}

	return false
}

// SetYesTotalCost gets a reference to the given int64 and assigns it to the YesTotalCost field.
func (o *SettlementHistory) SetYesTotalCost(v int64) {
	o.YesTotalCost = &v
}

func (o SettlementHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeterminedTime != nil {
		toSerialize["determined_time"] = o.DeterminedTime
	}
	if o.MarketId != nil {
		toSerialize["market_id"] = o.MarketId
	}
	if o.MarketResult != nil {
		toSerialize["market_result"] = o.MarketResult
	}
	if o.MarketTitle != nil {
		toSerialize["market_title"] = o.MarketTitle
	}
	if o.NoCount != nil {
		toSerialize["no_count"] = o.NoCount
	}
	if o.NoTotalCost != nil {
		toSerialize["no_total_cost"] = o.NoTotalCost
	}
	if o.Profit != nil {
		toSerialize["profit"] = o.Profit
	}
	if o.SettledTime != nil {
		toSerialize["settled_time"] = o.SettledTime
	}
	if o.YesCount != nil {
		toSerialize["yes_count"] = o.YesCount
	}
	if o.YesTotalCost != nil {
		toSerialize["yes_total_cost"] = o.YesTotalCost
	}
	return json.Marshal(toSerialize)
}

type NullableSettlementHistory struct {
	value *SettlementHistory
	isSet bool
}

func (v NullableSettlementHistory) Get() *SettlementHistory {
	return v.value
}

func (v *NullableSettlementHistory) Set(val *SettlementHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableSettlementHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableSettlementHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettlementHistory(val *SettlementHistory) *NullableSettlementHistory {
	return &NullableSettlementHistory{value: val, isSet: true}
}

func (v NullableSettlementHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettlementHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

