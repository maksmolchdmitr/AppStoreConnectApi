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

// checks if the AppRelationshipsEndUserLicenseAgreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsEndUserLicenseAgreement{}

// AppRelationshipsEndUserLicenseAgreement struct for AppRelationshipsEndUserLicenseAgreement
type AppRelationshipsEndUserLicenseAgreement struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *AppRelationshipsEndUserLicenseAgreementData                `json:"data,omitempty"`
}

// NewAppRelationshipsEndUserLicenseAgreement instantiates a new AppRelationshipsEndUserLicenseAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsEndUserLicenseAgreement() *AppRelationshipsEndUserLicenseAgreement {
	this := AppRelationshipsEndUserLicenseAgreement{}
	return &this
}

// NewAppRelationshipsEndUserLicenseAgreementWithDefaults instantiates a new AppRelationshipsEndUserLicenseAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsEndUserLicenseAgreementWithDefaults() *AppRelationshipsEndUserLicenseAgreement {
	this := AppRelationshipsEndUserLicenseAgreement{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsEndUserLicenseAgreement) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsEndUserLicenseAgreement) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsEndUserLicenseAgreement) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppRelationshipsEndUserLicenseAgreement) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsEndUserLicenseAgreement) GetData() AppRelationshipsEndUserLicenseAgreementData {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsEndUserLicenseAgreementData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsEndUserLicenseAgreement) GetDataOk() (*AppRelationshipsEndUserLicenseAgreementData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsEndUserLicenseAgreement) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsEndUserLicenseAgreementData and assigns it to the Data field.
func (o *AppRelationshipsEndUserLicenseAgreement) SetData(v AppRelationshipsEndUserLicenseAgreementData) {
	o.Data = &v
}

func (o AppRelationshipsEndUserLicenseAgreement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsEndUserLicenseAgreement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsEndUserLicenseAgreement struct {
	value *AppRelationshipsEndUserLicenseAgreement
	isSet bool
}

func (v NullableAppRelationshipsEndUserLicenseAgreement) Get() *AppRelationshipsEndUserLicenseAgreement {
	return v.value
}

func (v *NullableAppRelationshipsEndUserLicenseAgreement) Set(val *AppRelationshipsEndUserLicenseAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsEndUserLicenseAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsEndUserLicenseAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsEndUserLicenseAgreement(val *AppRelationshipsEndUserLicenseAgreement) *NullableAppRelationshipsEndUserLicenseAgreement {
	return &NullableAppRelationshipsEndUserLicenseAgreement{value: val, isSet: true}
}

func (v NullableAppRelationshipsEndUserLicenseAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsEndUserLicenseAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
