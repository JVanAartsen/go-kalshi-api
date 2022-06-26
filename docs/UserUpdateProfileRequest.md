# UserUpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaCode** | **string** | User&#39;s phone area code. | 
**BirthDate** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | **string** | User&#39;s country 2 digits code | 
**CountryCode** | **string** | User&#39;s phone country code. Should be 1 for now because only USA accounts are accepted. | 
**FinishedFre** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | **string** | User&#39;s phone number. | 
**PostalCode** | **string** | User&#39;s address postal code | 
**State** | **string** | User&#39;s state 2 digits code | 
**Street1** | Pointer to **string** |  | [optional] 
**Street2** | Pointer to **string** |  | [optional] 
**UseBidAsk** | Pointer to **bool** |  | [optional] 
**Watchlist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserUpdateProfileRequest

`func NewUserUpdateProfileRequest(areaCode string, country string, countryCode string, phoneNumber string, postalCode string, state string, ) *UserUpdateProfileRequest`

NewUserUpdateProfileRequest instantiates a new UserUpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateProfileRequestWithDefaults

`func NewUserUpdateProfileRequestWithDefaults() *UserUpdateProfileRequest`

NewUserUpdateProfileRequestWithDefaults instantiates a new UserUpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaCode

`func (o *UserUpdateProfileRequest) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *UserUpdateProfileRequest) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *UserUpdateProfileRequest) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.


### GetBirthDate

`func (o *UserUpdateProfileRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UserUpdateProfileRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UserUpdateProfileRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *UserUpdateProfileRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetCity

`func (o *UserUpdateProfileRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserUpdateProfileRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserUpdateProfileRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserUpdateProfileRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *UserUpdateProfileRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserUpdateProfileRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserUpdateProfileRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCountryCode

`func (o *UserUpdateProfileRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserUpdateProfileRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserUpdateProfileRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetFinishedFre

`func (o *UserUpdateProfileRequest) GetFinishedFre() bool`

GetFinishedFre returns the FinishedFre field if non-nil, zero value otherwise.

### GetFinishedFreOk

`func (o *UserUpdateProfileRequest) GetFinishedFreOk() (*bool, bool)`

GetFinishedFreOk returns a tuple with the FinishedFre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedFre

`func (o *UserUpdateProfileRequest) SetFinishedFre(v bool)`

SetFinishedFre sets FinishedFre field to given value.

### HasFinishedFre

`func (o *UserUpdateProfileRequest) HasFinishedFre() bool`

HasFinishedFre returns a boolean if a field has been set.

### GetFirstName

`func (o *UserUpdateProfileRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserUpdateProfileRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserUpdateProfileRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserUpdateProfileRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserUpdateProfileRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserUpdateProfileRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserUpdateProfileRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserUpdateProfileRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserUpdateProfileRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserUpdateProfileRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserUpdateProfileRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetPostalCode

`func (o *UserUpdateProfileRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserUpdateProfileRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserUpdateProfileRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetState

`func (o *UserUpdateProfileRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserUpdateProfileRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserUpdateProfileRequest) SetState(v string)`

SetState sets State field to given value.


### GetStreet1

`func (o *UserUpdateProfileRequest) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *UserUpdateProfileRequest) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *UserUpdateProfileRequest) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *UserUpdateProfileRequest) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### GetStreet2

`func (o *UserUpdateProfileRequest) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *UserUpdateProfileRequest) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *UserUpdateProfileRequest) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *UserUpdateProfileRequest) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetUseBidAsk

`func (o *UserUpdateProfileRequest) GetUseBidAsk() bool`

GetUseBidAsk returns the UseBidAsk field if non-nil, zero value otherwise.

### GetUseBidAskOk

`func (o *UserUpdateProfileRequest) GetUseBidAskOk() (*bool, bool)`

GetUseBidAskOk returns a tuple with the UseBidAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBidAsk

`func (o *UserUpdateProfileRequest) SetUseBidAsk(v bool)`

SetUseBidAsk sets UseBidAsk field to given value.

### HasUseBidAsk

`func (o *UserUpdateProfileRequest) HasUseBidAsk() bool`

HasUseBidAsk returns a boolean if a field has been set.

### GetWatchlist

`func (o *UserUpdateProfileRequest) GetWatchlist() []string`

GetWatchlist returns the Watchlist field if non-nil, zero value otherwise.

### GetWatchlistOk

`func (o *UserUpdateProfileRequest) GetWatchlistOk() (*[]string, bool)`

GetWatchlistOk returns a tuple with the Watchlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlist

`func (o *UserUpdateProfileRequest) SetWatchlist(v []string)`

SetWatchlist sets Watchlist field to given value.

### HasWatchlist

`func (o *UserUpdateProfileRequest) HasWatchlist() bool`

HasWatchlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


