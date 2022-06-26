# PortfolioMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | **int64** |  | 
**BalanceChanges** | Pointer to **[]int64** | Amount of the underlying balance change, if applicable, in cents | [optional] 
**CumulativeDeposits** | **int64** |  | 
**CumulativeNumberSettlements** | **int64** | Count of settlements member has had from account creation to timestamp (inclusive) | 
**CumulativeWithdrawals** | **int64** |  | 
**Reasons** | Pointer to **[]string** | Reasons for the portfolio value change, if applicable | [optional] 
**Ts** | **int64** | Timestamp of the read in UNIX timestamp. (https://www.unixtimestamp.com/) | 
**V** | **int64** |  | 

## Methods

### NewPortfolioMeasurement

`func NewPortfolioMeasurement(a int64, cumulativeDeposits int64, cumulativeNumberSettlements int64, cumulativeWithdrawals int64, ts int64, v int64, ) *PortfolioMeasurement`

NewPortfolioMeasurement instantiates a new PortfolioMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioMeasurementWithDefaults

`func NewPortfolioMeasurementWithDefaults() *PortfolioMeasurement`

NewPortfolioMeasurementWithDefaults instantiates a new PortfolioMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *PortfolioMeasurement) GetA() int64`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *PortfolioMeasurement) GetAOk() (*int64, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *PortfolioMeasurement) SetA(v int64)`

SetA sets A field to given value.


### GetBalanceChanges

`func (o *PortfolioMeasurement) GetBalanceChanges() []int64`

GetBalanceChanges returns the BalanceChanges field if non-nil, zero value otherwise.

### GetBalanceChangesOk

`func (o *PortfolioMeasurement) GetBalanceChangesOk() (*[]int64, bool)`

GetBalanceChangesOk returns a tuple with the BalanceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChanges

`func (o *PortfolioMeasurement) SetBalanceChanges(v []int64)`

SetBalanceChanges sets BalanceChanges field to given value.

### HasBalanceChanges

`func (o *PortfolioMeasurement) HasBalanceChanges() bool`

HasBalanceChanges returns a boolean if a field has been set.

### GetCumulativeDeposits

`func (o *PortfolioMeasurement) GetCumulativeDeposits() int64`

GetCumulativeDeposits returns the CumulativeDeposits field if non-nil, zero value otherwise.

### GetCumulativeDepositsOk

`func (o *PortfolioMeasurement) GetCumulativeDepositsOk() (*int64, bool)`

GetCumulativeDepositsOk returns a tuple with the CumulativeDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeDeposits

`func (o *PortfolioMeasurement) SetCumulativeDeposits(v int64)`

SetCumulativeDeposits sets CumulativeDeposits field to given value.


### GetCumulativeNumberSettlements

`func (o *PortfolioMeasurement) GetCumulativeNumberSettlements() int64`

GetCumulativeNumberSettlements returns the CumulativeNumberSettlements field if non-nil, zero value otherwise.

### GetCumulativeNumberSettlementsOk

`func (o *PortfolioMeasurement) GetCumulativeNumberSettlementsOk() (*int64, bool)`

GetCumulativeNumberSettlementsOk returns a tuple with the CumulativeNumberSettlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeNumberSettlements

`func (o *PortfolioMeasurement) SetCumulativeNumberSettlements(v int64)`

SetCumulativeNumberSettlements sets CumulativeNumberSettlements field to given value.


### GetCumulativeWithdrawals

`func (o *PortfolioMeasurement) GetCumulativeWithdrawals() int64`

GetCumulativeWithdrawals returns the CumulativeWithdrawals field if non-nil, zero value otherwise.

### GetCumulativeWithdrawalsOk

`func (o *PortfolioMeasurement) GetCumulativeWithdrawalsOk() (*int64, bool)`

GetCumulativeWithdrawalsOk returns a tuple with the CumulativeWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeWithdrawals

`func (o *PortfolioMeasurement) SetCumulativeWithdrawals(v int64)`

SetCumulativeWithdrawals sets CumulativeWithdrawals field to given value.


### GetReasons

`func (o *PortfolioMeasurement) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *PortfolioMeasurement) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *PortfolioMeasurement) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *PortfolioMeasurement) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetTs

`func (o *PortfolioMeasurement) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *PortfolioMeasurement) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *PortfolioMeasurement) SetTs(v int64)`

SetTs sets Ts field to given value.


### GetV

`func (o *PortfolioMeasurement) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *PortfolioMeasurement) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *PortfolioMeasurement) SetV(v int64)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


