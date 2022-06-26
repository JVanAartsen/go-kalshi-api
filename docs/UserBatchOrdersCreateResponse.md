# UserBatchOrdersCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]UserBatchOrdersCreateSingleOrderResponse**](UserBatchOrdersCreateSingleOrderResponse.md) | An array of responses corresponding to orders in the request. | 

## Methods

### NewUserBatchOrdersCreateResponse

`func NewUserBatchOrdersCreateResponse(orders []UserBatchOrdersCreateSingleOrderResponse, ) *UserBatchOrdersCreateResponse`

NewUserBatchOrdersCreateResponse instantiates a new UserBatchOrdersCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBatchOrdersCreateResponseWithDefaults

`func NewUserBatchOrdersCreateResponseWithDefaults() *UserBatchOrdersCreateResponse`

NewUserBatchOrdersCreateResponseWithDefaults instantiates a new UserBatchOrdersCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *UserBatchOrdersCreateResponse) GetOrders() []UserBatchOrdersCreateSingleOrderResponse`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *UserBatchOrdersCreateResponse) GetOrdersOk() (*[]UserBatchOrdersCreateSingleOrderResponse, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *UserBatchOrdersCreateResponse) SetOrders(v []UserBatchOrdersCreateSingleOrderResponse)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


