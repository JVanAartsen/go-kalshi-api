# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstOpenDate** | Pointer to **time.Time** |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewSeries

`func NewSeries() *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstOpenDate

`func (o *Series) GetFirstOpenDate() time.Time`

GetFirstOpenDate returns the FirstOpenDate field if non-nil, zero value otherwise.

### GetFirstOpenDateOk

`func (o *Series) GetFirstOpenDateOk() (*time.Time, bool)`

GetFirstOpenDateOk returns a tuple with the FirstOpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOpenDate

`func (o *Series) SetFirstOpenDate(v time.Time)`

SetFirstOpenDate sets FirstOpenDate field to given value.

### HasFirstOpenDate

`func (o *Series) HasFirstOpenDate() bool`

HasFirstOpenDate returns a boolean if a field has been set.

### GetFrequency

`func (o *Series) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Series) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Series) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Series) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTicker

`func (o *Series) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Series) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Series) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *Series) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTitle

`func (o *Series) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Series) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Series) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Series) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


