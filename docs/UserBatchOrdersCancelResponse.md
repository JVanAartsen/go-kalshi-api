# UserBatchOrdersCancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]UserBatchOrdersCancelSingleOrderResponse**](UserBatchOrdersCancelSingleOrderResponse.md) | An array of responses corresponding to the orders in the request. | 

## Methods

### NewUserBatchOrdersCancelResponse

`func NewUserBatchOrdersCancelResponse(orders []UserBatchOrdersCancelSingleOrderResponse, ) *UserBatchOrdersCancelResponse`

NewUserBatchOrdersCancelResponse instantiates a new UserBatchOrdersCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBatchOrdersCancelResponseWithDefaults

`func NewUserBatchOrdersCancelResponseWithDefaults() *UserBatchOrdersCancelResponse`

NewUserBatchOrdersCancelResponseWithDefaults instantiates a new UserBatchOrdersCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *UserBatchOrdersCancelResponse) GetOrders() []UserBatchOrdersCancelSingleOrderResponse`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *UserBatchOrdersCancelResponse) GetOrdersOk() (*[]UserBatchOrdersCancelSingleOrderResponse, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *UserBatchOrdersCancelResponse) SetOrders(v []UserBatchOrdersCancelSingleOrderResponse)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


