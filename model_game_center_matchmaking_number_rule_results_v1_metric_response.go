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

// checks if the GameCenterMatchmakingNumberRuleResultsV1MetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingNumberRuleResultsV1MetricResponse{}

// GameCenterMatchmakingNumberRuleResultsV1MetricResponse struct for GameCenterMatchmakingNumberRuleResultsV1MetricResponse
type GameCenterMatchmakingNumberRuleResultsV1MetricResponse struct {
	Data  []GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner `json:"data"`
	Links PagedDocumentLinks                                                `json:"links"`
	Meta  *PagingInformation                                                `json:"meta,omitempty"`
}

// NewGameCenterMatchmakingNumberRuleResultsV1MetricResponse instantiates a new GameCenterMatchmakingNumberRuleResultsV1MetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingNumberRuleResultsV1MetricResponse(data []GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner, links PagedDocumentLinks) *GameCenterMatchmakingNumberRuleResultsV1MetricResponse {
	this := GameCenterMatchmakingNumberRuleResultsV1MetricResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseWithDefaults instantiates a new GameCenterMatchmakingNumberRuleResultsV1MetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseWithDefaults() *GameCenterMatchmakingNumberRuleResultsV1MetricResponse {
	this := GameCenterMatchmakingNumberRuleResultsV1MetricResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetData() []GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner {
	if o == nil {
		var ret []GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetDataOk() ([]GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) SetData(v []GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterMatchmakingNumberRuleResultsV1MetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingNumberRuleResultsV1MetricResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse struct {
	value *GameCenterMatchmakingNumberRuleResultsV1MetricResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) Get() *GameCenterMatchmakingNumberRuleResultsV1MetricResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) Set(val *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse(val *GameCenterMatchmakingNumberRuleResultsV1MetricResponse) *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse {
	return &NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}