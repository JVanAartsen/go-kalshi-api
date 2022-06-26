# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**No** | **[][]int32** |  | 
**Yes** | **[][]int32** |  | 

## Methods

### NewOrderBook

`func NewOrderBook(no [][]int32, yes [][]int32, ) *OrderBook`

NewOrderBook instantiates a new OrderBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookWithDefaults

`func NewOrderBookWithDefaults() *OrderBook`

NewOrderBookWithDefaults instantiates a new OrderBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNo

`func (o *OrderBook) GetNo() [][]int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *OrderBook) GetNoOk() (*[][]int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *OrderBook) SetNo(v [][]int32)`

SetNo sets No field to given value.


### GetYes

`func (o *OrderBook) GetYes() [][]int32`

GetYes returns the Yes field if non-nil, zero value otherwise.

### GetYesOk

`func (o *OrderBook) GetYesOk() (*[][]int32, bool)`

GetYesOk returns a tuple with the Yes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYes

`func (o *OrderBook) SetYes(v [][]int32)`

SetYes sets Yes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


