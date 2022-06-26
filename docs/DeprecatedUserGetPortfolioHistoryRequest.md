# DeprecatedUserGetPortfolioHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDate** | Pointer to **time.Time** | Restricts the response to orders before a timestamp in: query | [optional] 
**MaxTs** | Pointer to **int64** | Restricts the response to orders before a timestamp in unix seconds, overrides max_date, defaults to now. in: query | [optional] 
**MinDate** | Pointer to **time.Time** | Restricts the response to orders after a timestamp in: query | [optional] 
**MinTs** | Pointer to **int64** | Restricts the response to orders after a timestamp in unix seconds, overrides min_date, defaults to one hour before now. in: query | [optional] 
**NumBuckets** | Pointer to **int32** | Determines the number of buckets to average over when performing subsampling, defaults to 1440. in: query | [optional] 

## Methods

### NewDeprecatedUserGetPortfolioHistoryRequest

`func NewDeprecatedUserGetPortfolioHistoryRequest() *DeprecatedUserGetPortfolioHistoryRequest`

NewDeprecatedUserGetPortfolioHistoryRequest instantiates a new DeprecatedUserGetPortfolioHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedUserGetPortfolioHistoryRequestWithDefaults

`func NewDeprecatedUserGetPortfolioHistoryRequestWithDefaults() *DeprecatedUserGetPortfolioHistoryRequest`

NewDeprecatedUserGetPortfolioHistoryRequestWithDefaults instantiates a new DeprecatedUserGetPortfolioHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMaxDate() time.Time`

GetMaxDate returns the MaxDate field if non-nil, zero value otherwise.

### GetMaxDateOk

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMaxDateOk() (*time.Time, bool)`

GetMaxDateOk returns a tuple with the MaxDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) SetMaxDate(v time.Time)`

SetMaxDate sets MaxDate field to given value.

### HasMaxDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) HasMaxDate() bool`

HasMaxDate returns a boolean if a field has been set.

### GetMaxTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMaxTs() int64`

GetMaxTs returns the MaxTs field if non-nil, zero value otherwise.

### GetMaxTsOk

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMaxTsOk() (*int64, bool)`

GetMaxTsOk returns a tuple with the MaxTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) SetMaxTs(v int64)`

SetMaxTs sets MaxTs field to given value.

### HasMaxTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) HasMaxTs() bool`

HasMaxTs returns a boolean if a field has been set.

### GetMinDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMinDate() time.Time`

GetMinDate returns the MinDate field if non-nil, zero value otherwise.

### GetMinDateOk

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMinDateOk() (*time.Time, bool)`

GetMinDateOk returns a tuple with the MinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) SetMinDate(v time.Time)`

SetMinDate sets MinDate field to given value.

### HasMinDate

`func (o *DeprecatedUserGetPortfolioHistoryRequest) HasMinDate() bool`

HasMinDate returns a boolean if a field has been set.

### GetMinTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMinTs() int64`

GetMinTs returns the MinTs field if non-nil, zero value otherwise.

### GetMinTsOk

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetMinTsOk() (*int64, bool)`

GetMinTsOk returns a tuple with the MinTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) SetMinTs(v int64)`

SetMinTs sets MinTs field to given value.

### HasMinTs

`func (o *DeprecatedUserGetPortfolioHistoryRequest) HasMinTs() bool`

HasMinTs returns a boolean if a field has been set.

### GetNumBuckets

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetNumBuckets() int32`

GetNumBuckets returns the NumBuckets field if non-nil, zero value otherwise.

### GetNumBucketsOk

`func (o *DeprecatedUserGetPortfolioHistoryRequest) GetNumBucketsOk() (*int32, bool)`

GetNumBucketsOk returns a tuple with the NumBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBuckets

`func (o *DeprecatedUserGetPortfolioHistoryRequest) SetNumBuckets(v int32)`

SetNumBuckets sets NumBuckets field to given value.

### HasNumBuckets

`func (o *DeprecatedUserGetPortfolioHistoryRequest) HasNumBuckets() bool`

HasNumBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


