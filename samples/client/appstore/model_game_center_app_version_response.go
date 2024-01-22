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

// checks if the GameCenterAppVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionResponse{}

// GameCenterAppVersionResponse struct for GameCenterAppVersionResponse
type GameCenterAppVersionResponse struct {
	Data GameCenterAppVersion `json:"data"`
	Included []GameCenterAppVersionsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewGameCenterAppVersionResponse instantiates a new GameCenterAppVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionResponse(data GameCenterAppVersion, links DocumentLinks) *GameCenterAppVersionResponse {
	this := GameCenterAppVersionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterAppVersionResponseWithDefaults instantiates a new GameCenterAppVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionResponseWithDefaults() *GameCenterAppVersionResponse {
	this := GameCenterAppVersionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionResponse) GetData() GameCenterAppVersion {
	if o == nil {
		var ret GameCenterAppVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionResponse) GetDataOk() (*GameCenterAppVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionResponse) SetData(v GameCenterAppVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterAppVersionResponse) GetIncluded() []GameCenterAppVersionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterAppVersionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionResponse) GetIncludedOk() ([]GameCenterAppVersionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterAppVersionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterAppVersionsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterAppVersionResponse) SetIncluded(v []GameCenterAppVersionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterAppVersionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterAppVersionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterAppVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterAppVersionResponse struct {
	value *GameCenterAppVersionResponse
	isSet bool
}

func (v NullableGameCenterAppVersionResponse) Get() *GameCenterAppVersionResponse {
	return v.value
}

func (v *NullableGameCenterAppVersionResponse) Set(val *GameCenterAppVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionResponse(val *GameCenterAppVersionResponse) *NullableGameCenterAppVersionResponse {
	return &NullableGameCenterAppVersionResponse{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

