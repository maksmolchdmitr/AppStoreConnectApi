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

// checks if the GameCenterLeaderboardSetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetsResponse{}

// GameCenterLeaderboardSetsResponse struct for GameCenterLeaderboardSetsResponse
type GameCenterLeaderboardSetsResponse struct {
	Data []GameCenterLeaderboardSet `json:"data"`
	Included []GameCenterLeaderboardSetsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewGameCenterLeaderboardSetsResponse instantiates a new GameCenterLeaderboardSetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetsResponse(data []GameCenterLeaderboardSet, links PagedDocumentLinks) *GameCenterLeaderboardSetsResponse {
	this := GameCenterLeaderboardSetsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetsResponseWithDefaults instantiates a new GameCenterLeaderboardSetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetsResponseWithDefaults() *GameCenterLeaderboardSetsResponse {
	this := GameCenterLeaderboardSetsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetsResponse) GetData() []GameCenterLeaderboardSet {
	if o == nil {
		var ret []GameCenterLeaderboardSet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetsResponse) GetDataOk() ([]GameCenterLeaderboardSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetsResponse) SetData(v []GameCenterLeaderboardSet) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetsResponse) GetIncluded() []GameCenterLeaderboardSetsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetsResponse) GetIncludedOk() ([]GameCenterLeaderboardSetsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardSetsResponse) SetIncluded(v []GameCenterLeaderboardSetsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterLeaderboardSetsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterLeaderboardSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterLeaderboardSetsResponse struct {
	value *GameCenterLeaderboardSetsResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetsResponse) Get() *GameCenterLeaderboardSetsResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetsResponse) Set(val *GameCenterLeaderboardSetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetsResponse(val *GameCenterLeaderboardSetsResponse) *NullableGameCenterLeaderboardSetsResponse {
	return &NullableGameCenterLeaderboardSetsResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


