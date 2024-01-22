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

// checks if the AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData{}

// AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData struct for AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData
type AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData instantiates a new AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData(type_ string, id string) *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData {
	this := AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageDataWithDefaults instantiates a new AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageDataWithDefaults() *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData {
	this := AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) SetId(v string) {
	o.Id = v
}

func (o AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData struct {
	value *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) Get() *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) Set(val *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData(val *AppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) *NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData {
	return &NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationRelationshipsAppClipHeaderImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


