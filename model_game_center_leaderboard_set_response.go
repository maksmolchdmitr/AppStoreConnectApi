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

// checks if the GameCenterLeaderboardSetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetResponse{}

// GameCenterLeaderboardSetResponse struct for GameCenterLeaderboardSetResponse
type GameCenterLeaderboardSetResponse struct {
	Data     GameCenterLeaderboardSet                         `json:"data"`
	Included []GameCenterLeaderboardSetsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                    `json:"links"`
}

// NewGameCenterLeaderboardSetResponse instantiates a new GameCenterLeaderboardSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetResponse(data GameCenterLeaderboardSet, links DocumentLinks) *GameCenterLeaderboardSetResponse {
	this := GameCenterLeaderboardSetResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetResponseWithDefaults instantiates a new GameCenterLeaderboardSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetResponseWithDefaults() *GameCenterLeaderboardSetResponse {
	this := GameCenterLeaderboardSetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetResponse) GetData() GameCenterLeaderboardSet {
	if o == nil {
		var ret GameCenterLeaderboardSet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetResponse) GetDataOk() (*GameCenterLeaderboardSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetResponse) SetData(v GameCenterLeaderboardSet) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetResponse) GetIncluded() []GameCenterLeaderboardSetsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetResponse) GetIncludedOk() ([]GameCenterLeaderboardSetsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardSetResponse) SetIncluded(v []GameCenterLeaderboardSetsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardSetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetResponse struct {
	value *GameCenterLeaderboardSetResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetResponse) Get() *GameCenterLeaderboardSetResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetResponse) Set(val *GameCenterLeaderboardSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetResponse(val *GameCenterLeaderboardSetResponse) *NullableGameCenterLeaderboardSetResponse {
	return &NullableGameCenterLeaderboardSetResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
