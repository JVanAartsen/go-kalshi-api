# DeprecatedPortfolioMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | **int64** |  | 
**BalanceChange** | Pointer to **int64** |  | [optional] 
**CumulativeDeposits** | **int64** |  | 
**CumulativeNumberSettlements** | **int64** | Count of settlements member has had from account creation to timestamp (inclusive) | 
**CumulativeWithdrawals** | **int64** |  | 
**Reason** | Pointer to **string** | Reasons for the portfolio value change, if applicable | [optional] 
**Ts** | **int64** | Timestamp of the read in UNIX timestamp. (https://www.unixtimestamp.com/) | 
**V** | **int64** |  | 

## Methods

### NewDeprecatedPortfolioMeasurement

`func NewDeprecatedPortfolioMeasurement(a int64, cumulativeDeposits int64, cumulativeNumberSettlements int64, cumulativeWithdrawals int64, ts int64, v int64, ) *DeprecatedPortfolioMeasurement`

NewDeprecatedPortfolioMeasurement instantiates a new DeprecatedPortfolioMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedPortfolioMeasurementWithDefaults

`func NewDeprecatedPortfolioMeasurementWithDefaults() *DeprecatedPortfolioMeasurement`

NewDeprecatedPortfolioMeasurementWithDefaults instantiates a new DeprecatedPortfolioMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *DeprecatedPortfolioMeasurement) GetA() int64`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *DeprecatedPortfolioMeasurement) GetAOk() (*int64, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *DeprecatedPortfolioMeasurement) SetA(v int64)`

SetA sets A field to given value.


### GetBalanceChange

`func (o *DeprecatedPortfolioMeasurement) GetBalanceChange() int64`

GetBalanceChange returns the BalanceChange field if non-nil, zero value otherwise.

### GetBalanceChangeOk

`func (o *DeprecatedPortfolioMeasurement) GetBalanceChangeOk() (*int64, bool)`

GetBalanceChangeOk returns a tuple with the BalanceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChange

`func (o *DeprecatedPortfolioMeasurement) SetBalanceChange(v int64)`

SetBalanceChange sets BalanceChange field to given value.

### HasBalanceChange

`func (o *DeprecatedPortfolioMeasurement) HasBalanceChange() bool`

HasBalanceChange returns a boolean if a field has been set.

### GetCumulativeDeposits

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeDeposits() int64`

GetCumulativeDeposits returns the CumulativeDeposits field if non-nil, zero value otherwise.

### GetCumulativeDepositsOk

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeDepositsOk() (*int64, bool)`

GetCumulativeDepositsOk returns a tuple with the CumulativeDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeDeposits

`func (o *DeprecatedPortfolioMeasurement) SetCumulativeDeposits(v int64)`

SetCumulativeDeposits sets CumulativeDeposits field to given value.


### GetCumulativeNumberSettlements

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeNumberSettlements() int64`

GetCumulativeNumberSettlements returns the CumulativeNumberSettlements field if non-nil, zero value otherwise.

### GetCumulativeNumberSettlementsOk

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeNumberSettlementsOk() (*int64, bool)`

GetCumulativeNumberSettlementsOk returns a tuple with the CumulativeNumberSettlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeNumberSettlements

`func (o *DeprecatedPortfolioMeasurement) SetCumulativeNumberSettlements(v int64)`

SetCumulativeNumberSettlements sets CumulativeNumberSettlements field to given value.


### GetCumulativeWithdrawals

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeWithdrawals() int64`

GetCumulativeWithdrawals returns the CumulativeWithdrawals field if non-nil, zero value otherwise.

### GetCumulativeWithdrawalsOk

`func (o *DeprecatedPortfolioMeasurement) GetCumulativeWithdrawalsOk() (*int64, bool)`

GetCumulativeWithdrawalsOk returns a tuple with the CumulativeWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeWithdrawals

`func (o *DeprecatedPortfolioMeasurement) SetCumulativeWithdrawals(v int64)`

SetCumulativeWithdrawals sets CumulativeWithdrawals field to given value.


### GetReason

`func (o *DeprecatedPortfolioMeasurement) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeprecatedPortfolioMeasurement) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeprecatedPortfolioMeasurement) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeprecatedPortfolioMeasurement) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTs

`func (o *DeprecatedPortfolioMeasurement) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *DeprecatedPortfolioMeasurement) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *DeprecatedPortfolioMeasurement) SetTs(v int64)`

SetTs sets Ts field to given value.


### GetV

`func (o *DeprecatedPortfolioMeasurement) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *DeprecatedPortfolioMeasurement) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *DeprecatedPortfolioMeasurement) SetV(v int64)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


