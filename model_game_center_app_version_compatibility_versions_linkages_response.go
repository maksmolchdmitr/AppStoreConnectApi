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

// checks if the GameCenterAppVersionCompatibilityVersionsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionCompatibilityVersionsLinkagesResponse{}

// GameCenterAppVersionCompatibilityVersionsLinkagesResponse struct for GameCenterAppVersionCompatibilityVersionsLinkagesResponse
type GameCenterAppVersionCompatibilityVersionsLinkagesResponse struct {
	Data  []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner `json:"data"`
	Links PagedDocumentLinks                                                `json:"links"`
	Meta  *PagingInformation                                                `json:"meta,omitempty"`
}

// NewGameCenterAppVersionCompatibilityVersionsLinkagesResponse instantiates a new GameCenterAppVersionCompatibilityVersionsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionCompatibilityVersionsLinkagesResponse(data []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner, links PagedDocumentLinks) *GameCenterAppVersionCompatibilityVersionsLinkagesResponse {
	this := GameCenterAppVersionCompatibilityVersionsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAppVersionCompatibilityVersionsLinkagesResponseWithDefaults instantiates a new GameCenterAppVersionCompatibilityVersionsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionCompatibilityVersionsLinkagesResponseWithDefaults() *GameCenterAppVersionCompatibilityVersionsLinkagesResponse {
	this := GameCenterAppVersionCompatibilityVersionsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetData() []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner {
	if o == nil {
		var ret []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetDataOk() ([]GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetData(v []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterAppVersionCompatibilityVersionsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionCompatibilityVersionsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse struct {
	value *GameCenterAppVersionCompatibilityVersionsLinkagesResponse
	isSet bool
}

func (v NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) Get() *GameCenterAppVersionCompatibilityVersionsLinkagesResponse {
	return v.value
}

func (v *NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) Set(val *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse(val *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) *NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse {
	return &NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionCompatibilityVersionsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
