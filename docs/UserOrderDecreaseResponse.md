# UserOrderDecreaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**Order**](Order.md) |  | 
**ReducedBy** | **int32** | Result of the decrease operation | 

## Methods

### NewUserOrderDecreaseResponse

`func NewUserOrderDecreaseResponse(order Order, reducedBy int32, ) *UserOrderDecreaseResponse`

NewUserOrderDecreaseResponse instantiates a new UserOrderDecreaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrderDecreaseResponseWithDefaults

`func NewUserOrderDecreaseResponseWithDefaults() *UserOrderDecreaseResponse`

NewUserOrderDecreaseResponseWithDefaults instantiates a new UserOrderDecreaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *UserOrderDecreaseResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UserOrderDecreaseResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UserOrderDecreaseResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.


### GetReducedBy

`func (o *UserOrderDecreaseResponse) GetReducedBy() int32`

GetReducedBy returns the ReducedBy field if non-nil, zero value otherwise.

### GetReducedByOk

`func (o *UserOrderDecreaseResponse) GetReducedByOk() (*int32, bool)`

GetReducedByOk returns a tuple with the ReducedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedBy

`func (o *UserOrderDecreaseResponse) SetReducedBy(v int32)`

SetReducedBy sets ReducedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


