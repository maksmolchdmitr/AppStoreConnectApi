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

// checks if the GameCenterAchievementGroupAchievementLinkageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementGroupAchievementLinkageResponse{}

// GameCenterAchievementGroupAchievementLinkageResponse struct for GameCenterAchievementGroupAchievementLinkageResponse
type GameCenterAchievementGroupAchievementLinkageResponse struct {
	Data GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewGameCenterAchievementGroupAchievementLinkageResponse instantiates a new GameCenterAchievementGroupAchievementLinkageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementGroupAchievementLinkageResponse(data GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, links DocumentLinks) *GameCenterAchievementGroupAchievementLinkageResponse {
	this := GameCenterAchievementGroupAchievementLinkageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAchievementGroupAchievementLinkageResponseWithDefaults instantiates a new GameCenterAchievementGroupAchievementLinkageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementGroupAchievementLinkageResponseWithDefaults() *GameCenterAchievementGroupAchievementLinkageResponse {
	this := GameCenterAchievementGroupAchievementLinkageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementGroupAchievementLinkageResponse) GetData() GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData {
	if o == nil {
		var ret GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementGroupAchievementLinkageResponse) GetDataOk() (*GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementGroupAchievementLinkageResponse) SetData(v GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterAchievementGroupAchievementLinkageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementGroupAchievementLinkageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAchievementGroupAchievementLinkageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterAchievementGroupAchievementLinkageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementGroupAchievementLinkageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterAchievementGroupAchievementLinkageResponse struct {
	value *GameCenterAchievementGroupAchievementLinkageResponse
	isSet bool
}

func (v NullableGameCenterAchievementGroupAchievementLinkageResponse) Get() *GameCenterAchievementGroupAchievementLinkageResponse {
	return v.value
}

func (v *NullableGameCenterAchievementGroupAchievementLinkageResponse) Set(val *GameCenterAchievementGroupAchievementLinkageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementGroupAchievementLinkageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementGroupAchievementLinkageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementGroupAchievementLinkageResponse(val *GameCenterAchievementGroupAchievementLinkageResponse) *NullableGameCenterAchievementGroupAchievementLinkageResponse {
	return &NullableGameCenterAchievementGroupAchievementLinkageResponse{value: val, isSet: true}
}

func (v NullableGameCenterAchievementGroupAchievementLinkageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementGroupAchievementLinkageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


