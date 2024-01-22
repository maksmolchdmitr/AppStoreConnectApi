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

// checks if the GameCenterLeaderboardReleasesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardReleasesResponse{}

// GameCenterLeaderboardReleasesResponse struct for GameCenterLeaderboardReleasesResponse
type GameCenterLeaderboardReleasesResponse struct {
	Data []GameCenterLeaderboardRelease `json:"data"`
	Included []GameCenterLeaderboardReleasesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewGameCenterLeaderboardReleasesResponse instantiates a new GameCenterLeaderboardReleasesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardReleasesResponse(data []GameCenterLeaderboardRelease, links PagedDocumentLinks) *GameCenterLeaderboardReleasesResponse {
	this := GameCenterLeaderboardReleasesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardReleasesResponseWithDefaults instantiates a new GameCenterLeaderboardReleasesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardReleasesResponseWithDefaults() *GameCenterLeaderboardReleasesResponse {
	this := GameCenterLeaderboardReleasesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardReleasesResponse) GetData() []GameCenterLeaderboardRelease {
	if o == nil {
		var ret []GameCenterLeaderboardRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleasesResponse) GetDataOk() ([]GameCenterLeaderboardRelease, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardReleasesResponse) SetData(v []GameCenterLeaderboardRelease) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardReleasesResponse) GetIncluded() []GameCenterLeaderboardReleasesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardReleasesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleasesResponse) GetIncludedOk() ([]GameCenterLeaderboardReleasesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardReleasesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardReleasesResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardReleasesResponse) SetIncluded(v []GameCenterLeaderboardReleasesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardReleasesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleasesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardReleasesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterLeaderboardReleasesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleasesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterLeaderboardReleasesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterLeaderboardReleasesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterLeaderboardReleasesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardReleasesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterLeaderboardReleasesResponse struct {
	value *GameCenterLeaderboardReleasesResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardReleasesResponse) Get() *GameCenterLeaderboardReleasesResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardReleasesResponse) Set(val *GameCenterLeaderboardReleasesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardReleasesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardReleasesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardReleasesResponse(val *GameCenterLeaderboardReleasesResponse) *NullableGameCenterLeaderboardReleasesResponse {
	return &NullableGameCenterLeaderboardReleasesResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardReleasesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardReleasesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

