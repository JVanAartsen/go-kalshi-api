# TradesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trades** | [**[]PublicTrade**](PublicTrade.md) |  | 

## Methods

### NewTradesGetResponse

`func NewTradesGetResponse(trades []PublicTrade, ) *TradesGetResponse`

NewTradesGetResponse instantiates a new TradesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradesGetResponseWithDefaults

`func NewTradesGetResponseWithDefaults() *TradesGetResponse`

NewTradesGetResponseWithDefaults instantiates a new TradesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrades

`func (o *TradesGetResponse) GetTrades() []PublicTrade`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *TradesGetResponse) GetTradesOk() (*[]PublicTrade, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *TradesGetResponse) SetTrades(v []PublicTrade)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


