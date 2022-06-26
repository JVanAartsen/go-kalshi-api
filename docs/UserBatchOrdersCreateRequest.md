# UserBatchOrdersCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]UserOrderCreateRequest**](UserOrderCreateRequest.md) | An array of individual orders to place | 

## Methods

### NewUserBatchOrdersCreateRequest

`func NewUserBatchOrdersCreateRequest(orders []UserOrderCreateRequest, ) *UserBatchOrdersCreateRequest`

NewUserBatchOrdersCreateRequest instantiates a new UserBatchOrdersCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBatchOrdersCreateRequestWithDefaults

`func NewUserBatchOrdersCreateRequestWithDefaults() *UserBatchOrdersCreateRequest`

NewUserBatchOrdersCreateRequestWithDefaults instantiates a new UserBatchOrdersCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *UserBatchOrdersCreateRequest) GetOrders() []UserOrderCreateRequest`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *UserBatchOrdersCreateRequest) GetOrdersOk() (*[]UserOrderCreateRequest, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *UserBatchOrdersCreateRequest) SetOrders(v []UserOrderCreateRequest)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


