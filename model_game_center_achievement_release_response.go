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

// checks if the GameCenterAchievementReleaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementReleaseResponse{}

// GameCenterAchievementReleaseResponse struct for GameCenterAchievementReleaseResponse
type GameCenterAchievementReleaseResponse struct {
	Data     GameCenterAchievementRelease                         `json:"data"`
	Included []GameCenterAchievementReleasesResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                        `json:"links"`
}

// NewGameCenterAchievementReleaseResponse instantiates a new GameCenterAchievementReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementReleaseResponse(data GameCenterAchievementRelease, links DocumentLinks) *GameCenterAchievementReleaseResponse {
	this := GameCenterAchievementReleaseResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAchievementReleaseResponseWithDefaults instantiates a new GameCenterAchievementReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleaseResponseWithDefaults() *GameCenterAchievementReleaseResponse {
	this := GameCenterAchievementReleaseResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementReleaseResponse) GetData() GameCenterAchievementRelease {
	if o == nil {
		var ret GameCenterAchievementRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseResponse) GetDataOk() (*GameCenterAchievementRelease, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementReleaseResponse) SetData(v GameCenterAchievementRelease) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterAchievementReleaseResponse) GetIncluded() []GameCenterAchievementReleasesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterAchievementReleasesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseResponse) GetIncludedOk() ([]GameCenterAchievementReleasesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterAchievementReleaseResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterAchievementReleasesResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterAchievementReleaseResponse) SetIncluded(v []GameCenterAchievementReleasesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterAchievementReleaseResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAchievementReleaseResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterAchievementReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementReleaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterAchievementReleaseResponse struct {
	value *GameCenterAchievementReleaseResponse
	isSet bool
}

func (v NullableGameCenterAchievementReleaseResponse) Get() *GameCenterAchievementReleaseResponse {
	return v.value
}

func (v *NullableGameCenterAchievementReleaseResponse) Set(val *GameCenterAchievementReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementReleaseResponse(val *GameCenterAchievementReleaseResponse) *NullableGameCenterAchievementReleaseResponse {
	return &NullableGameCenterAchievementReleaseResponse{value: val, isSet: true}
}

func (v NullableGameCenterAchievementReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
