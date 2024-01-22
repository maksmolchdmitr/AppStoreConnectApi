/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions{}

// GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions struct for GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions
type GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions struct {
	Result                     *BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds        `json:"result,omitempty"`
	GameCenterMatchmakingQueue *AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters `json:"gameCenterMatchmakingQueue,omitempty"`
}

// NewGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions instantiates a new GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions() *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions {
	this := GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions{}
	return &this
}

// NewGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensionsWithDefaults instantiates a new GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensionsWithDefaults() *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions {
	this := GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) GetResult() BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds {
	if o == nil || IsNil(o.Result) {
		var ret BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) GetResultOk() (*BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds and assigns it to the Result field.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) SetResult(v BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds) {
	o.Result = &v
}

// GetGameCenterMatchmakingQueue returns the GameCenterMatchmakingQueue field value if set, zero value otherwise.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) GetGameCenterMatchmakingQueue() AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters {
	if o == nil || IsNil(o.GameCenterMatchmakingQueue) {
		var ret AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters
		return ret
	}
	return *o.GameCenterMatchmakingQueue
}

// GetGameCenterMatchmakingQueueOk returns a tuple with the GameCenterMatchmakingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) GetGameCenterMatchmakingQueueOk() (*AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters, bool) {
	if o == nil || IsNil(o.GameCenterMatchmakingQueue) {
		return nil, false
	}
	return o.GameCenterMatchmakingQueue, true
}

// HasGameCenterMatchmakingQueue returns a boolean if a field has been set.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) HasGameCenterMatchmakingQueue() bool {
	if o != nil && !IsNil(o.GameCenterMatchmakingQueue) {
		return true
	}

	return false
}

// SetGameCenterMatchmakingQueue gets a reference to the given AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters and assigns it to the GameCenterMatchmakingQueue field.
func (o *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) SetGameCenterMatchmakingQueue(v AppsBetaTesterUsagesV1MetricResponseDataInnerDimensionsBetaTesters) {
	o.GameCenterMatchmakingQueue = &v
}

func (o GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.GameCenterMatchmakingQueue) {
		toSerialize["gameCenterMatchmakingQueue"] = o.GameCenterMatchmakingQueue
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions struct {
	value *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions
	isSet bool
}

func (v NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) Get() *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions {
	return v.value
}

func (v *NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) Set(val *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions(val *GameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) *NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions {
	return &NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingBooleanRuleResultsV1MetricResponseDataInnerDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
