# JSONError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewJSONError

`func NewJSONError() *JSONError`

NewJSONError instantiates a new JSONError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONErrorWithDefaults

`func NewJSONErrorWithDefaults() *JSONError`

NewJSONErrorWithDefaults instantiates a new JSONError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JSONError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JSONError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JSONError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *JSONError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *JSONError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *JSONError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *JSONError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *JSONError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *JSONError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JSONError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JSONError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JSONError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetService

`func (o *JSONError) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *JSONError) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *JSONError) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *JSONError) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


