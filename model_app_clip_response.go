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

// checks if the AppClipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipResponse{}

// AppClipResponse struct for AppClipResponse
type AppClipResponse struct {
	Data     AppClip                         `json:"data"`
	Included []AppClipsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                   `json:"links"`
}

// NewAppClipResponse instantiates a new AppClipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipResponse(data AppClip, links DocumentLinks) *AppClipResponse {
	this := AppClipResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipResponseWithDefaults instantiates a new AppClipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipResponseWithDefaults() *AppClipResponse {
	this := AppClipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipResponse) GetData() AppClip {
	if o == nil {
		var ret AppClip
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipResponse) GetDataOk() (*AppClip, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipResponse) SetData(v AppClip) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipResponse) GetIncluded() []AppClipsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipResponse) GetIncludedOk() ([]AppClipsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipsResponseIncludedInner and assigns it to the Included field.
func (o *AppClipResponse) SetIncluded(v []AppClipsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipResponse struct {
	value *AppClipResponse
	isSet bool
}

func (v NullableAppClipResponse) Get() *AppClipResponse {
	return v.value
}

func (v *NullableAppClipResponse) Set(val *AppClipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipResponse(val *AppClipResponse) *NullableAppClipResponse {
	return &NullableAppClipResponse{value: val, isSet: true}
}

func (v NullableAppClipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
