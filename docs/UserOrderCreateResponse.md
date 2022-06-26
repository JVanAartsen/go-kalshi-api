# UserOrderCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**Order**](Order.md) |  | 
**Status** | **string** | Status of the order submit operation | 

## Methods

### NewUserOrderCreateResponse

`func NewUserOrderCreateResponse(order Order, status string, ) *UserOrderCreateResponse`

NewUserOrderCreateResponse instantiates a new UserOrderCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrderCreateResponseWithDefaults

`func NewUserOrderCreateResponseWithDefaults() *UserOrderCreateResponse`

NewUserOrderCreateResponseWithDefaults instantiates a new UserOrderCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *UserOrderCreateResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UserOrderCreateResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UserOrderCreateResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.


### GetStatus

`func (o *UserOrderCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserOrderCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserOrderCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


