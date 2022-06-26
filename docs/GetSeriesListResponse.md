# GetSeriesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]Series**](Series.md) |  | [optional] 

## Methods

### NewGetSeriesListResponse

`func NewGetSeriesListResponse() *GetSeriesListResponse`

NewGetSeriesListResponse instantiates a new GetSeriesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSeriesListResponseWithDefaults

`func NewGetSeriesListResponseWithDefaults() *GetSeriesListResponse`

NewGetSeriesListResponseWithDefaults instantiates a new GetSeriesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *GetSeriesListResponse) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *GetSeriesListResponse) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *GetSeriesListResponse) SetSeries(v []Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *GetSeriesListResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


