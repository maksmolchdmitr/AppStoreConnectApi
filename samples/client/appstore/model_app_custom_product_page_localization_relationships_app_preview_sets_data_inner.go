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

// checks if the AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner{}

// AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner struct for AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner
type AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner instantiates a new AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner(type_ string, id string) *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner {
	this := AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInnerWithDefaults instantiates a new AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInnerWithDefaults() *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner {
	this := AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner struct {
	value *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) Get() *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) Set(val *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner(val *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) *NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner {
	return &NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


