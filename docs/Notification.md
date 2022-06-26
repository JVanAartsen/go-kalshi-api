# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CreatedTs** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsDelivered** | Pointer to **bool** |  | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Notification) GetContent() map[string]map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Notification) GetContentOk() (*map[string]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Notification) SetContent(v map[string]map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *Notification) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedTs

`func (o *Notification) GetCreatedTs() time.Time`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *Notification) GetCreatedTsOk() (*time.Time, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *Notification) SetCreatedTs(v time.Time)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *Notification) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetId

`func (o *Notification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDelivered

`func (o *Notification) GetIsDelivered() bool`

GetIsDelivered returns the IsDelivered field if non-nil, zero value otherwise.

### GetIsDeliveredOk

`func (o *Notification) GetIsDeliveredOk() (*bool, bool)`

GetIsDeliveredOk returns a tuple with the IsDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelivered

`func (o *Notification) SetIsDelivered(v bool)`

SetIsDelivered sets IsDelivered field to given value.

### HasIsDelivered

`func (o *Notification) HasIsDelivered() bool`

HasIsDelivered returns a boolean if a field has been set.

### GetIsRead

`func (o *Notification) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Notification) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Notification) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *Notification) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetLink

`func (o *Notification) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Notification) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Notification) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Notification) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Notification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *Notification) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Notification) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Notification) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Notification) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


