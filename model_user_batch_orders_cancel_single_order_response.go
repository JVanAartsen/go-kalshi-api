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

// UserBatchOrdersCancelSingleOrderResponse Response for cancelling a batch of orders
type UserBatchOrdersCancelSingleOrderResponse struct {
	Error *JSONError `json:"error,omitempty"`
	// ID of the order
	Id string `json:"id"`
	Order *Order `json:"order,omitempty"`
	// Result of the decrease operation
	ReducedBy *int32 `json:"reduced_by,omitempty"`
}

// NewUserBatchOrdersCancelSingleOrderResponse instantiates a new UserBatchOrdersCancelSingleOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBatchOrdersCancelSingleOrderResponse(id string) *UserBatchOrdersCancelSingleOrderResponse {
	this := UserBatchOrdersCancelSingleOrderResponse{}
	this.Id = id
	return &this
}

// NewUserBatchOrdersCancelSingleOrderResponseWithDefaults instantiates a new UserBatchOrdersCancelSingleOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBatchOrdersCancelSingleOrderResponseWithDefaults() *UserBatchOrdersCancelSingleOrderResponse {
	this := UserBatchOrdersCancelSingleOrderResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetError() JSONError {
	if o == nil || o.Error == nil {
		var ret JSONError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetErrorOk() (*JSONError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given JSONError and assigns it to the Error field.
func (o *UserBatchOrdersCancelSingleOrderResponse) SetError(v JSONError) {
	o.Error = &v
}

// GetId returns the Id field value
func (o *UserBatchOrdersCancelSingleOrderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserBatchOrdersCancelSingleOrderResponse) SetId(v string) {
	o.Id = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetOrder() Order {
	if o == nil || o.Order == nil {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetOrderOk() (*Order, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *UserBatchOrdersCancelSingleOrderResponse) SetOrder(v Order) {
	o.Order = &v
}

// GetReducedBy returns the ReducedBy field value if set, zero value otherwise.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetReducedBy() int32 {
	if o == nil || o.ReducedBy == nil {
		var ret int32
		return ret
	}
	return *o.ReducedBy
}

// GetReducedByOk returns a tuple with the ReducedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) GetReducedByOk() (*int32, bool) {
	if o == nil || o.ReducedBy == nil {
		return nil, false
	}
	return o.ReducedBy, true
}

// HasReducedBy returns a boolean if a field has been set.
func (o *UserBatchOrdersCancelSingleOrderResponse) HasReducedBy() bool {
	if o != nil && o.ReducedBy != nil {
		return true
	}

	return false
}

// SetReducedBy gets a reference to the given int32 and assigns it to the ReducedBy field.
func (o *UserBatchOrdersCancelSingleOrderResponse) SetReducedBy(v int32) {
	o.ReducedBy = &v
}

func (o UserBatchOrdersCancelSingleOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ReducedBy != nil {
		toSerialize["reduced_by"] = o.ReducedBy
	}
	return json.Marshal(toSerialize)
}

type NullableUserBatchOrdersCancelSingleOrderResponse struct {
	value *UserBatchOrdersCancelSingleOrderResponse
	isSet bool
}

func (v NullableUserBatchOrdersCancelSingleOrderResponse) Get() *UserBatchOrdersCancelSingleOrderResponse {
	return v.value
}

func (v *NullableUserBatchOrdersCancelSingleOrderResponse) Set(val *UserBatchOrdersCancelSingleOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBatchOrdersCancelSingleOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBatchOrdersCancelSingleOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBatchOrdersCancelSingleOrderResponse(val *UserBatchOrdersCancelSingleOrderResponse) *NullableUserBatchOrdersCancelSingleOrderResponse {
	return &NullableUserBatchOrdersCancelSingleOrderResponse{value: val, isSet: true}
}

func (v NullableUserBatchOrdersCancelSingleOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBatchOrdersCancelSingleOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

