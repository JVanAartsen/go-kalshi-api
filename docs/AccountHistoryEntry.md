# AccountHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AccountHistoryEntryData**](AccountHistoryEntryData.md) |  | [optional] 
**Type** | **string** | Type of entry, one of Deposit, Withdrawal, Order, or Settlement | 

## Methods

### NewAccountHistoryEntry

`func NewAccountHistoryEntry(type_ string, ) *AccountHistoryEntry`

NewAccountHistoryEntry instantiates a new AccountHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHistoryEntryWithDefaults

`func NewAccountHistoryEntryWithDefaults() *AccountHistoryEntry`

NewAccountHistoryEntryWithDefaults instantiates a new AccountHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccountHistoryEntry) GetData() AccountHistoryEntryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountHistoryEntry) GetDataOk() (*AccountHistoryEntryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountHistoryEntry) SetData(v AccountHistoryEntryData)`

SetData sets Data field to given value.

### HasData

`func (o *AccountHistoryEntry) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *AccountHistoryEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountHistoryEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountHistoryEntry) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


