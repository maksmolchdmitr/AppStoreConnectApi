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

// checks if the PrereleaseVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrereleaseVersionResponse{}

// PrereleaseVersionResponse struct for PrereleaseVersionResponse
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion                         `json:"data"`
	Included []PreReleaseVersionsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                             `json:"links"`
}

// NewPrereleaseVersionResponse instantiates a new PrereleaseVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrereleaseVersionResponse(data PrereleaseVersion, links DocumentLinks) *PrereleaseVersionResponse {
	this := PrereleaseVersionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPrereleaseVersionResponseWithDefaults instantiates a new PrereleaseVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrereleaseVersionResponseWithDefaults() *PrereleaseVersionResponse {
	this := PrereleaseVersionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PrereleaseVersionResponse) GetData() PrereleaseVersion {
	if o == nil {
		var ret PrereleaseVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionResponse) GetDataOk() (*PrereleaseVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PrereleaseVersionResponse) SetData(v PrereleaseVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *PrereleaseVersionResponse) GetIncluded() []PreReleaseVersionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []PreReleaseVersionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionResponse) GetIncludedOk() ([]PreReleaseVersionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *PrereleaseVersionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []PreReleaseVersionsResponseIncludedInner and assigns it to the Included field.
func (o *PrereleaseVersionResponse) SetIncluded(v []PreReleaseVersionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *PrereleaseVersionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PrereleaseVersionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o PrereleaseVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrereleaseVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullablePrereleaseVersionResponse struct {
	value *PrereleaseVersionResponse
	isSet bool
}

func (v NullablePrereleaseVersionResponse) Get() *PrereleaseVersionResponse {
	return v.value
}

func (v *NullablePrereleaseVersionResponse) Set(val *PrereleaseVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrereleaseVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrereleaseVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrereleaseVersionResponse(val *PrereleaseVersionResponse) *NullablePrereleaseVersionResponse {
	return &NullablePrereleaseVersionResponse{value: val, isSet: true}
}

func (v NullablePrereleaseVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrereleaseVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
