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

// checks if the GameCenterLeaderboardSetReleaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetReleaseResponse{}

// GameCenterLeaderboardSetReleaseResponse struct for GameCenterLeaderboardSetReleaseResponse
type GameCenterLeaderboardSetReleaseResponse struct {
	Data GameCenterLeaderboardSetRelease `json:"data"`
	Included []GameCenterLeaderboardSetReleasesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewGameCenterLeaderboardSetReleaseResponse instantiates a new GameCenterLeaderboardSetReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetReleaseResponse(data GameCenterLeaderboardSetRelease, links DocumentLinks) *GameCenterLeaderboardSetReleaseResponse {
	this := GameCenterLeaderboardSetReleaseResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetReleaseResponseWithDefaults instantiates a new GameCenterLeaderboardSetReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetReleaseResponseWithDefaults() *GameCenterLeaderboardSetReleaseResponse {
	this := GameCenterLeaderboardSetReleaseResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetReleaseResponse) GetData() GameCenterLeaderboardSetRelease {
	if o == nil {
		var ret GameCenterLeaderboardSetRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetReleaseResponse) GetDataOk() (*GameCenterLeaderboardSetRelease, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetReleaseResponse) SetData(v GameCenterLeaderboardSetRelease) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetReleaseResponse) GetIncluded() []GameCenterLeaderboardSetReleasesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetReleasesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetReleaseResponse) GetIncludedOk() ([]GameCenterLeaderboardSetReleasesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetReleaseResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetReleasesResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardSetReleaseResponse) SetIncluded(v []GameCenterLeaderboardSetReleasesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetReleaseResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetReleaseResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetReleaseResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardSetReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetReleaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetReleaseResponse struct {
	value *GameCenterLeaderboardSetReleaseResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetReleaseResponse) Get() *GameCenterLeaderboardSetReleaseResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetReleaseResponse) Set(val *GameCenterLeaderboardSetReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetReleaseResponse(val *GameCenterLeaderboardSetReleaseResponse) *NullableGameCenterLeaderboardSetReleaseResponse {
	return &NullableGameCenterLeaderboardSetReleaseResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


