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

// checks if the AppStoreVersionExperimentTreatmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentsResponse{}

// AppStoreVersionExperimentTreatmentsResponse struct for AppStoreVersionExperimentTreatmentsResponse
type AppStoreVersionExperimentTreatmentsResponse struct {
	Data []AppStoreVersionExperimentTreatment `json:"data"`
	Included []AppStoreVersionExperimentTreatmentsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentsResponse instantiates a new AppStoreVersionExperimentTreatmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentsResponse(data []AppStoreVersionExperimentTreatment, links PagedDocumentLinks) *AppStoreVersionExperimentTreatmentsResponse {
	this := AppStoreVersionExperimentTreatmentsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionExperimentTreatmentsResponseWithDefaults instantiates a new AppStoreVersionExperimentTreatmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentsResponseWithDefaults() *AppStoreVersionExperimentTreatmentsResponse {
	this := AppStoreVersionExperimentTreatmentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentsResponse) GetData() []AppStoreVersionExperimentTreatment {
	if o == nil {
		var ret []AppStoreVersionExperimentTreatment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetDataOk() ([]AppStoreVersionExperimentTreatment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentsResponse) SetData(v []AppStoreVersionExperimentTreatment) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetIncluded() []AppStoreVersionExperimentTreatmentsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionExperimentTreatmentsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetIncludedOk() ([]AppStoreVersionExperimentTreatmentsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionExperimentTreatmentsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionExperimentTreatmentsResponse) SetIncluded(v []AppStoreVersionExperimentTreatmentsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionExperimentTreatmentsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionExperimentTreatmentsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionExperimentTreatmentsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppStoreVersionExperimentTreatmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppStoreVersionExperimentTreatmentsResponse struct {
	value *AppStoreVersionExperimentTreatmentsResponse
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentsResponse) Get() *AppStoreVersionExperimentTreatmentsResponse {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponse) Set(val *AppStoreVersionExperimentTreatmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentsResponse(val *AppStoreVersionExperimentTreatmentsResponse) *NullableAppStoreVersionExperimentTreatmentsResponse {
	return &NullableAppStoreVersionExperimentTreatmentsResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


