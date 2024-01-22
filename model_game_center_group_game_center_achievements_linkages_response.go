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

// checks if the GameCenterGroupGameCenterAchievementsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupGameCenterAchievementsLinkagesResponse{}

// GameCenterGroupGameCenterAchievementsLinkagesResponse struct for GameCenterGroupGameCenterAchievementsLinkagesResponse
type GameCenterGroupGameCenterAchievementsLinkagesResponse struct {
	Data  []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData `json:"data"`
	Links PagedDocumentLinks                                                        `json:"links"`
	Meta  *PagingInformation                                                        `json:"meta,omitempty"`
}

// NewGameCenterGroupGameCenterAchievementsLinkagesResponse instantiates a new GameCenterGroupGameCenterAchievementsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupGameCenterAchievementsLinkagesResponse(data []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, links PagedDocumentLinks) *GameCenterGroupGameCenterAchievementsLinkagesResponse {
	this := GameCenterGroupGameCenterAchievementsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterGroupGameCenterAchievementsLinkagesResponseWithDefaults instantiates a new GameCenterGroupGameCenterAchievementsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupGameCenterAchievementsLinkagesResponseWithDefaults() *GameCenterGroupGameCenterAchievementsLinkagesResponse {
	this := GameCenterGroupGameCenterAchievementsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetData() []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData {
	if o == nil {
		var ret []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetDataOk() ([]GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetData(v []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterGroupGameCenterAchievementsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupGameCenterAchievementsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGameCenterGroupGameCenterAchievementsLinkagesResponse struct {
	value *GameCenterGroupGameCenterAchievementsLinkagesResponse
	isSet bool
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) Get() *GameCenterGroupGameCenterAchievementsLinkagesResponse {
	return v.value
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) Set(val *GameCenterGroupGameCenterAchievementsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupGameCenterAchievementsLinkagesResponse(val *GameCenterGroupGameCenterAchievementsLinkagesResponse) *NullableGameCenterGroupGameCenterAchievementsLinkagesResponse {
	return &NullableGameCenterGroupGameCenterAchievementsLinkagesResponse{value: val, isSet: true}
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
