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

// checks if the GameCenterAppVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionsResponse{}

// GameCenterAppVersionsResponse struct for GameCenterAppVersionsResponse
type GameCenterAppVersionsResponse struct {
	Data     []GameCenterAppVersion                       `json:"data"`
	Included []GameCenterAppVersionsResponseIncludedInner `json:"included,omitempty"`
	Links    PagedDocumentLinks                           `json:"links"`
	Meta     *PagingInformation                           `json:"meta,omitempty"`
}

// NewGameCenterAppVersionsResponse instantiates a new GameCenterAppVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionsResponse(data []GameCenterAppVersion, links PagedDocumentLinks) *GameCenterAppVersionsResponse {
	this := GameCenterAppVersionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAppVersionsResponseWithDefaults instantiates a new GameCenterAppVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionsResponseWithDefaults() *GameCenterAppVersionsResponse {
	this := GameCenterAppVersionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionsResponse) GetData() []GameCenterAppVersion {
	if o == nil {
		var ret []GameCenterAppVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionsResponse) GetDataOk() ([]GameCenterAppVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionsResponse) SetData(v []GameCenterAppVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterAppVersionsResponse) GetIncluded() []GameCenterAppVersionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterAppVersionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionsResponse) GetIncludedOk() ([]GameCenterAppVersionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterAppVersionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterAppVersionsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterAppVersionsResponse) SetIncluded(v []GameCenterAppVersionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterAppVersionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAppVersionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterAppVersionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterAppVersionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterAppVersionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterAppVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterAppVersionsResponse struct {
	value *GameCenterAppVersionsResponse
	isSet bool
}

func (v NullableGameCenterAppVersionsResponse) Get() *GameCenterAppVersionsResponse {
	return v.value
}

func (v *NullableGameCenterAppVersionsResponse) Set(val *GameCenterAppVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionsResponse(val *GameCenterAppVersionsResponse) *NullableGameCenterAppVersionsResponse {
	return &NullableGameCenterAppVersionsResponse{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}