# UserBatchOrdersCreateSingleOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**JSONError**](JSONError.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the order submit operation | [optional] 

## Methods

### NewUserBatchOrdersCreateSingleOrderResponse

`func NewUserBatchOrdersCreateSingleOrderResponse() *UserBatchOrdersCreateSingleOrderResponse`

NewUserBatchOrdersCreateSingleOrderResponse instantiates a new UserBatchOrdersCreateSingleOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBatchOrdersCreateSingleOrderResponseWithDefaults

`func NewUserBatchOrdersCreateSingleOrderResponseWithDefaults() *UserBatchOrdersCreateSingleOrderResponse`

NewUserBatchOrdersCreateSingleOrderResponseWithDefaults instantiates a new UserBatchOrdersCreateSingleOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetError() JSONError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetErrorOk() (*JSONError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserBatchOrdersCreateSingleOrderResponse) SetError(v JSONError)`

SetError sets Error field to given value.

### HasError

`func (o *UserBatchOrdersCreateSingleOrderResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOrder

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UserBatchOrdersCreateSingleOrderResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UserBatchOrdersCreateSingleOrderResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStatus

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserBatchOrdersCreateSingleOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserBatchOrdersCreateSingleOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserBatchOrdersCreateSingleOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


