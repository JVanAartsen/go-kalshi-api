# UserBatchOrdersCancelSingleOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**JSONError**](JSONError.md) |  | [optional] 
**Id** | **string** | ID of the order | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**ReducedBy** | Pointer to **int32** | Result of the decrease operation | [optional] 

## Methods

### NewUserBatchOrdersCancelSingleOrderResponse

`func NewUserBatchOrdersCancelSingleOrderResponse(id string, ) *UserBatchOrdersCancelSingleOrderResponse`

NewUserBatchOrdersCancelSingleOrderResponse instantiates a new UserBatchOrdersCancelSingleOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBatchOrdersCancelSingleOrderResponseWithDefaults

`func NewUserBatchOrdersCancelSingleOrderResponseWithDefaults() *UserBatchOrdersCancelSingleOrderResponse`

NewUserBatchOrdersCancelSingleOrderResponseWithDefaults instantiates a new UserBatchOrdersCancelSingleOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetError() JSONError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetErrorOk() (*JSONError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserBatchOrdersCancelSingleOrderResponse) SetError(v JSONError)`

SetError sets Error field to given value.

### HasError

`func (o *UserBatchOrdersCancelSingleOrderResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBatchOrdersCancelSingleOrderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOrder

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UserBatchOrdersCancelSingleOrderResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UserBatchOrdersCancelSingleOrderResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReducedBy

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetReducedBy() int32`

GetReducedBy returns the ReducedBy field if non-nil, zero value otherwise.

### GetReducedByOk

`func (o *UserBatchOrdersCancelSingleOrderResponse) GetReducedByOk() (*int32, bool)`

GetReducedByOk returns a tuple with the ReducedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedBy

`func (o *UserBatchOrdersCancelSingleOrderResponse) SetReducedBy(v int32)`

SetReducedBy sets ReducedBy field to given value.

### HasReducedBy

`func (o *UserBatchOrdersCancelSingleOrderResponse) HasReducedBy() bool`

HasReducedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


