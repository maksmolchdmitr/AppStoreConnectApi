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

// checks if the AppInfoRelationshipsAgeRatingDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoRelationshipsAgeRatingDeclaration{}

// AppInfoRelationshipsAgeRatingDeclaration struct for AppInfoRelationshipsAgeRatingDeclaration
type AppInfoRelationshipsAgeRatingDeclaration struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *AppInfoRelationshipsAgeRatingDeclarationData               `json:"data,omitempty"`
}

// NewAppInfoRelationshipsAgeRatingDeclaration instantiates a new AppInfoRelationshipsAgeRatingDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoRelationshipsAgeRatingDeclaration() *AppInfoRelationshipsAgeRatingDeclaration {
	this := AppInfoRelationshipsAgeRatingDeclaration{}
	return &this
}

// NewAppInfoRelationshipsAgeRatingDeclarationWithDefaults instantiates a new AppInfoRelationshipsAgeRatingDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoRelationshipsAgeRatingDeclarationWithDefaults() *AppInfoRelationshipsAgeRatingDeclaration {
	this := AppInfoRelationshipsAgeRatingDeclaration{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppInfoRelationshipsAgeRatingDeclaration) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoRelationshipsAgeRatingDeclaration) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppInfoRelationshipsAgeRatingDeclaration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppInfoRelationshipsAgeRatingDeclaration) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppInfoRelationshipsAgeRatingDeclaration) GetData() AppInfoRelationshipsAgeRatingDeclarationData {
	if o == nil || IsNil(o.Data) {
		var ret AppInfoRelationshipsAgeRatingDeclarationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoRelationshipsAgeRatingDeclaration) GetDataOk() (*AppInfoRelationshipsAgeRatingDeclarationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppInfoRelationshipsAgeRatingDeclaration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppInfoRelationshipsAgeRatingDeclarationData and assigns it to the Data field.
func (o *AppInfoRelationshipsAgeRatingDeclaration) SetData(v AppInfoRelationshipsAgeRatingDeclarationData) {
	o.Data = &v
}

func (o AppInfoRelationshipsAgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoRelationshipsAgeRatingDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppInfoRelationshipsAgeRatingDeclaration struct {
	value *AppInfoRelationshipsAgeRatingDeclaration
	isSet bool
}

func (v NullableAppInfoRelationshipsAgeRatingDeclaration) Get() *AppInfoRelationshipsAgeRatingDeclaration {
	return v.value
}

func (v *NullableAppInfoRelationshipsAgeRatingDeclaration) Set(val *AppInfoRelationshipsAgeRatingDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoRelationshipsAgeRatingDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoRelationshipsAgeRatingDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoRelationshipsAgeRatingDeclaration(val *AppInfoRelationshipsAgeRatingDeclaration) *NullableAppInfoRelationshipsAgeRatingDeclaration {
	return &NullableAppInfoRelationshipsAgeRatingDeclaration{value: val, isSet: true}
}

func (v NullableAppInfoRelationshipsAgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoRelationshipsAgeRatingDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
