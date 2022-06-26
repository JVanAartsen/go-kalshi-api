# GetEventByTickerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**EventData**](EventData.md) |  | 

## Methods

### NewGetEventByTickerResponse

`func NewGetEventByTickerResponse(event EventData, ) *GetEventByTickerResponse`

NewGetEventByTickerResponse instantiates a new GetEventByTickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventByTickerResponseWithDefaults

`func NewGetEventByTickerResponseWithDefaults() *GetEventByTickerResponse`

NewGetEventByTickerResponseWithDefaults instantiates a new GetEventByTickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GetEventByTickerResponse) GetEvent() EventData`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetEventByTickerResponse) GetEventOk() (*EventData, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetEventByTickerResponse) SetEvent(v EventData)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


