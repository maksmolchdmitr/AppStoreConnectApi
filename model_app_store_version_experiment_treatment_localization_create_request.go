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

// checks if the AppStoreVersionExperimentTreatmentLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationCreateRequest{}

// AppStoreVersionExperimentTreatmentLocalizationCreateRequest struct for AppStoreVersionExperimentTreatmentLocalizationCreateRequest
type AppStoreVersionExperimentTreatmentLocalizationCreateRequest struct {
	Data AppStoreVersionExperimentTreatmentLocalizationCreateRequestData `json:"data"`
}

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequest instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequest(data AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) *AppStoreVersionExperimentTreatmentLocalizationCreateRequest {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationCreateRequest {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequest) GetData() AppStoreVersionExperimentTreatmentLocalizationCreateRequestData {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequest) GetDataOk() (*AppStoreVersionExperimentTreatmentLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequest) SetData(v AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest struct {
	value *AppStoreVersionExperimentTreatmentLocalizationCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) Get() *AppStoreVersionExperimentTreatmentLocalizationCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) Set(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequest) *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
