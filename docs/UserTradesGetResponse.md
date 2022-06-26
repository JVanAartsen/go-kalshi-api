# UserTradesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trades** | [**[]UserTrade**](UserTrade.md) |  | 

## Methods

### NewUserTradesGetResponse

`func NewUserTradesGetResponse(trades []UserTrade, ) *UserTradesGetResponse`

NewUserTradesGetResponse instantiates a new UserTradesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTradesGetResponseWithDefaults

`func NewUserTradesGetResponseWithDefaults() *UserTradesGetResponse`

NewUserTradesGetResponseWithDefaults instantiates a new UserTradesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrades

`func (o *UserTradesGetResponse) GetTrades() []UserTrade`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *UserTradesGetResponse) GetTradesOk() (*[]UserTrade, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *UserTradesGetResponse) SetTrades(v []UserTrade)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


