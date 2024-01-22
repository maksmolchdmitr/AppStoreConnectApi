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

// checks if the AppPreviewRelationshipsAppPreviewSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewRelationshipsAppPreviewSet{}

// AppPreviewRelationshipsAppPreviewSet struct for AppPreviewRelationshipsAppPreviewSet
type AppPreviewRelationshipsAppPreviewSet struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data *AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner `json:"data,omitempty"`
}

// NewAppPreviewRelationshipsAppPreviewSet instantiates a new AppPreviewRelationshipsAppPreviewSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewRelationshipsAppPreviewSet() *AppPreviewRelationshipsAppPreviewSet {
	this := AppPreviewRelationshipsAppPreviewSet{}
	return &this
}

// NewAppPreviewRelationshipsAppPreviewSetWithDefaults instantiates a new AppPreviewRelationshipsAppPreviewSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewRelationshipsAppPreviewSetWithDefaults() *AppPreviewRelationshipsAppPreviewSet {
	this := AppPreviewRelationshipsAppPreviewSet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPreviewRelationshipsAppPreviewSet) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppPreviewRelationshipsAppPreviewSet) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewRelationshipsAppPreviewSet) GetData() AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) GetDataOk() (*AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner and assigns it to the Data field.
func (o *AppPreviewRelationshipsAppPreviewSet) SetData(v AppCustomProductPageLocalizationRelationshipsAppPreviewSetsDataInner) {
	o.Data = &v
}

func (o AppPreviewRelationshipsAppPreviewSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewRelationshipsAppPreviewSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPreviewRelationshipsAppPreviewSet struct {
	value *AppPreviewRelationshipsAppPreviewSet
	isSet bool
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) Get() *AppPreviewRelationshipsAppPreviewSet {
	return v.value
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) Set(val *AppPreviewRelationshipsAppPreviewSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewRelationshipsAppPreviewSet(val *AppPreviewRelationshipsAppPreviewSet) *NullableAppPreviewRelationshipsAppPreviewSet {
	return &NullableAppPreviewRelationshipsAppPreviewSet{value: val, isSet: true}
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

