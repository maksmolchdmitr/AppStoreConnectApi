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

// checks if the AppStoreVersionSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionSubmissionResponse{}

// AppStoreVersionSubmissionResponse struct for AppStoreVersionSubmissionResponse
type AppStoreVersionSubmissionResponse struct {
	// Deprecated
	Data AppStoreVersionSubmission `json:"data"`
	Included []AppStoreVersion `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppStoreVersionSubmissionResponse instantiates a new AppStoreVersionSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionSubmissionResponse(data AppStoreVersionSubmission, links DocumentLinks) *AppStoreVersionSubmissionResponse {
	this := AppStoreVersionSubmissionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionSubmissionResponseWithDefaults instantiates a new AppStoreVersionSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionSubmissionResponseWithDefaults() *AppStoreVersionSubmissionResponse {
	this := AppStoreVersionSubmissionResponse{}
	return &this
}

// GetData returns the Data field value
// Deprecated
func (o *AppStoreVersionSubmissionResponse) GetData() AppStoreVersionSubmission {
	if o == nil {
		var ret AppStoreVersionSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppStoreVersionSubmissionResponse) GetDataOk() (*AppStoreVersionSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
// Deprecated
func (o *AppStoreVersionSubmissionResponse) SetData(v AppStoreVersionSubmission) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionSubmissionResponse) GetIncluded() []AppStoreVersion {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionResponse) GetIncludedOk() ([]AppStoreVersion, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionSubmissionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersion and assigns it to the Included field.
func (o *AppStoreVersionSubmissionResponse) SetIncluded(v []AppStoreVersion) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionSubmissionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionSubmissionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionSubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppStoreVersionSubmissionResponse struct {
	value *AppStoreVersionSubmissionResponse
	isSet bool
}

func (v NullableAppStoreVersionSubmissionResponse) Get() *AppStoreVersionSubmissionResponse {
	return v.value
}

func (v *NullableAppStoreVersionSubmissionResponse) Set(val *AppStoreVersionSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionSubmissionResponse(val *AppStoreVersionSubmissionResponse) *NullableAppStoreVersionSubmissionResponse {
	return &NullableAppStoreVersionSubmissionResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


