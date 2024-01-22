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

// checks if the AppStoreVersionExperimentTreatmentLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationCreateRequestData{}

// AppStoreVersionExperimentTreatmentLocalizationCreateRequestData struct for AppStoreVersionExperimentTreatmentLocalizationCreateRequestData
type AppStoreVersionExperimentTreatmentLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships `json:"relationships"`
}

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestData instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestData(type_ string, attributes AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes, relationships AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships) *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetAttributes() AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetAttributesOk() (*AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) SetAttributes(v AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetRelationships() AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) GetRelationshipsOk() (*AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) SetRelationships(v AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData struct {
	value *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) Get() *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) Set(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequestData) *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


