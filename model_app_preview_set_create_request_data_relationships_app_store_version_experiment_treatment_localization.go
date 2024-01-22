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

// checks if the AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization{}

// AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization struct for AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization
type AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization struct {
	Data *AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData `json:"data,omitempty"`
}

// NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization instantiates a new AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization {
	this := AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization{}
	return &this
}

// NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalizationWithDefaults instantiates a new AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalizationWithDefaults() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization {
	this := AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) GetData() AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData {
	if o == nil || IsNil(o.Data) {
		var ret AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) GetDataOk() (*AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData and assigns it to the Data field.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) SetData(v AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalizationData) {
	o.Data = &v
}

func (o AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization struct {
	value *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization
	isSet bool
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) Get() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization {
	return v.value
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) Set(val *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization(val *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization {
	return &NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization{value: val, isSet: true}
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionExperimentTreatmentLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}