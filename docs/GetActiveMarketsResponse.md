# GetActiveMarketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMarkets** | Pointer to [**[]MarketActivity**](MarketActivity.md) |  | [optional] 

## Methods

### NewGetActiveMarketsResponse

`func NewGetActiveMarketsResponse() *GetActiveMarketsResponse`

NewGetActiveMarketsResponse instantiates a new GetActiveMarketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActiveMarketsResponseWithDefaults

`func NewGetActiveMarketsResponseWithDefaults() *GetActiveMarketsResponse`

NewGetActiveMarketsResponseWithDefaults instantiates a new GetActiveMarketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveMarkets

`func (o *GetActiveMarketsResponse) GetActiveMarkets() []MarketActivity`

GetActiveMarkets returns the ActiveMarkets field if non-nil, zero value otherwise.

### GetActiveMarketsOk

`func (o *GetActiveMarketsResponse) GetActiveMarketsOk() (*[]MarketActivity, bool)`

GetActiveMarketsOk returns a tuple with the ActiveMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMarkets

`func (o *GetActiveMarketsResponse) SetActiveMarkets(v []MarketActivity)`

SetActiveMarkets sets ActiveMarkets field to given value.

### HasActiveMarkets

`func (o *GetActiveMarketsResponse) HasActiveMarkets() bool`

HasActiveMarkets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


