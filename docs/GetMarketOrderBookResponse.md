# GetMarketOrderBookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderBook** | [**OrderBook**](OrderBook.md) |  | 

## Methods

### NewGetMarketOrderBookResponse

`func NewGetMarketOrderBookResponse(orderBook OrderBook, ) *GetMarketOrderBookResponse`

NewGetMarketOrderBookResponse instantiates a new GetMarketOrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketOrderBookResponseWithDefaults

`func NewGetMarketOrderBookResponseWithDefaults() *GetMarketOrderBookResponse`

NewGetMarketOrderBookResponseWithDefaults instantiates a new GetMarketOrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderBook

`func (o *GetMarketOrderBookResponse) GetOrderBook() OrderBook`

GetOrderBook returns the OrderBook field if non-nil, zero value otherwise.

### GetOrderBookOk

`func (o *GetMarketOrderBookResponse) GetOrderBookOk() (*OrderBook, bool)`

GetOrderBookOk returns a tuple with the OrderBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBook

`func (o *GetMarketOrderBookResponse) SetOrderBook(v OrderBook)`

SetOrderBook sets OrderBook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


