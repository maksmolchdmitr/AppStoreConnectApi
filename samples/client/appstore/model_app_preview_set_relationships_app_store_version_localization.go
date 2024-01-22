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

// checks if the AppPreviewSetRelationshipsAppStoreVersionLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetRelationshipsAppStoreVersionLocalization{}

// AppPreviewSetRelationshipsAppStoreVersionLocalization struct for AppPreviewSetRelationshipsAppStoreVersionLocalization
type AppPreviewSetRelationshipsAppStoreVersionLocalization struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data *AppPreviewSetRelationshipsAppStoreVersionLocalizationData `json:"data,omitempty"`
}

// NewAppPreviewSetRelationshipsAppStoreVersionLocalization instantiates a new AppPreviewSetRelationshipsAppStoreVersionLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetRelationshipsAppStoreVersionLocalization() *AppPreviewSetRelationshipsAppStoreVersionLocalization {
	this := AppPreviewSetRelationshipsAppStoreVersionLocalization{}
	return &this
}

// NewAppPreviewSetRelationshipsAppStoreVersionLocalizationWithDefaults instantiates a new AppPreviewSetRelationshipsAppStoreVersionLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetRelationshipsAppStoreVersionLocalizationWithDefaults() *AppPreviewSetRelationshipsAppStoreVersionLocalization {
	this := AppPreviewSetRelationshipsAppStoreVersionLocalization{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) GetData() AppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	if o == nil || IsNil(o.Data) {
		var ret AppPreviewSetRelationshipsAppStoreVersionLocalizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) GetDataOk() (*AppPreviewSetRelationshipsAppStoreVersionLocalizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionLocalizationData and assigns it to the Data field.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalization) SetData(v AppPreviewSetRelationshipsAppStoreVersionLocalizationData) {
	o.Data = &v
}

func (o AppPreviewSetRelationshipsAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetRelationshipsAppStoreVersionLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPreviewSetRelationshipsAppStoreVersionLocalization struct {
	value *AppPreviewSetRelationshipsAppStoreVersionLocalization
	isSet bool
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) Get() *AppPreviewSetRelationshipsAppStoreVersionLocalization {
	return v.value
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) Set(val *AppPreviewSetRelationshipsAppStoreVersionLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetRelationshipsAppStoreVersionLocalization(val *AppPreviewSetRelationshipsAppStoreVersionLocalization) *NullableAppPreviewSetRelationshipsAppStoreVersionLocalization {
	return &NullableAppPreviewSetRelationshipsAppStoreVersionLocalization{value: val, isSet: true}
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


