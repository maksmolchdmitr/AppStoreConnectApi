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

// checks if the GameCenterGroupGameCenterLeaderboardsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupGameCenterLeaderboardsLinkagesResponse{}

// GameCenterGroupGameCenterLeaderboardsLinkagesResponse struct for GameCenterGroupGameCenterLeaderboardsLinkagesResponse
type GameCenterGroupGameCenterLeaderboardsLinkagesResponse struct {
	Data  []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
	Links PagedDocumentLinks                                             `json:"links"`
	Meta  *PagingInformation                                             `json:"meta,omitempty"`
}

// NewGameCenterGroupGameCenterLeaderboardsLinkagesResponse instantiates a new GameCenterGroupGameCenterLeaderboardsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupGameCenterLeaderboardsLinkagesResponse(data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, links PagedDocumentLinks) *GameCenterGroupGameCenterLeaderboardsLinkagesResponse {
	this := GameCenterGroupGameCenterLeaderboardsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterGroupGameCenterLeaderboardsLinkagesResponseWithDefaults instantiates a new GameCenterGroupGameCenterLeaderboardsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupGameCenterLeaderboardsLinkagesResponseWithDefaults() *GameCenterGroupGameCenterLeaderboardsLinkagesResponse {
	this := GameCenterGroupGameCenterLeaderboardsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterGroupGameCenterLeaderboardsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupGameCenterLeaderboardsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse struct {
	value *GameCenterGroupGameCenterLeaderboardsLinkagesResponse
	isSet bool
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) Get() *GameCenterGroupGameCenterLeaderboardsLinkagesResponse {
	return v.value
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) Set(val *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse(val *GameCenterGroupGameCenterLeaderboardsLinkagesResponse) *NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse {
	return &NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse{value: val, isSet: true}
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
