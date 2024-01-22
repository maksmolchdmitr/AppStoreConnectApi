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

// checks if the GameCenterAchievementReleasesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementReleasesResponse{}

// GameCenterAchievementReleasesResponse struct for GameCenterAchievementReleasesResponse
type GameCenterAchievementReleasesResponse struct {
	Data     []GameCenterAchievementRelease                       `json:"data"`
	Included []GameCenterAchievementReleasesResponseIncludedInner `json:"included,omitempty"`
	Links    PagedDocumentLinks                                   `json:"links"`
	Meta     *PagingInformation                                   `json:"meta,omitempty"`
}

// NewGameCenterAchievementReleasesResponse instantiates a new GameCenterAchievementReleasesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementReleasesResponse(data []GameCenterAchievementRelease, links PagedDocumentLinks) *GameCenterAchievementReleasesResponse {
	this := GameCenterAchievementReleasesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAchievementReleasesResponseWithDefaults instantiates a new GameCenterAchievementReleasesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleasesResponseWithDefaults() *GameCenterAchievementReleasesResponse {
	this := GameCenterAchievementReleasesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementReleasesResponse) GetData() []GameCenterAchievementRelease {
	if o == nil {
		var ret []GameCenterAchievementRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleasesResponse) GetDataOk() ([]GameCenterAchievementRelease, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementReleasesResponse) SetData(v []GameCenterAchievementRelease) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterAchievementReleasesResponse) GetIncluded() []GameCenterAchievementReleasesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterAchievementReleasesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleasesResponse) GetIncludedOk() ([]GameCenterAchievementReleasesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterAchievementReleasesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterAchievementReleasesResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterAchievementReleasesResponse) SetIncluded(v []GameCenterAchievementReleasesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterAchievementReleasesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleasesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAchievementReleasesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterAchievementReleasesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleasesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterAchievementReleasesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterAchievementReleasesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterAchievementReleasesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementReleasesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementReleasesResponse struct {
	value *GameCenterAchievementReleasesResponse
	isSet bool
}

func (v NullableGameCenterAchievementReleasesResponse) Get() *GameCenterAchievementReleasesResponse {
	return v.value
}

func (v *NullableGameCenterAchievementReleasesResponse) Set(val *GameCenterAchievementReleasesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementReleasesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementReleasesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementReleasesResponse(val *GameCenterAchievementReleasesResponse) *NullableGameCenterAchievementReleasesResponse {
	return &NullableGameCenterAchievementReleasesResponse{value: val, isSet: true}
}

func (v NullableGameCenterAchievementReleasesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementReleasesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
