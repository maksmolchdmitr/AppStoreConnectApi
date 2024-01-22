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

// checks if the AppStoreVersionExperimentV2RelationshipsControlVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentV2RelationshipsControlVersions{}

// AppStoreVersionExperimentV2RelationshipsControlVersions struct for AppStoreVersionExperimentV2RelationshipsControlVersions
type AppStoreVersionExperimentV2RelationshipsControlVersions struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData `json:"data,omitempty"`
}

// NewAppStoreVersionExperimentV2RelationshipsControlVersions instantiates a new AppStoreVersionExperimentV2RelationshipsControlVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentV2RelationshipsControlVersions() *AppStoreVersionExperimentV2RelationshipsControlVersions {
	this := AppStoreVersionExperimentV2RelationshipsControlVersions{}
	return &this
}

// NewAppStoreVersionExperimentV2RelationshipsControlVersionsWithDefaults instantiates a new AppStoreVersionExperimentV2RelationshipsControlVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentV2RelationshipsControlVersionsWithDefaults() *AppStoreVersionExperimentV2RelationshipsControlVersions {
	this := AppStoreVersionExperimentV2RelationshipsControlVersions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetData() []AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData {
	if o == nil || IsNil(o.Data) {
		var ret []AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) GetDataOk() ([]AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData and assigns it to the Data field.
func (o *AppStoreVersionExperimentV2RelationshipsControlVersions) SetData(v []AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData) {
	o.Data = v
}

func (o AppStoreVersionExperimentV2RelationshipsControlVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentV2RelationshipsControlVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentV2RelationshipsControlVersions struct {
	value *AppStoreVersionExperimentV2RelationshipsControlVersions
	isSet bool
}

func (v NullableAppStoreVersionExperimentV2RelationshipsControlVersions) Get() *AppStoreVersionExperimentV2RelationshipsControlVersions {
	return v.value
}

func (v *NullableAppStoreVersionExperimentV2RelationshipsControlVersions) Set(val *AppStoreVersionExperimentV2RelationshipsControlVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentV2RelationshipsControlVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentV2RelationshipsControlVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentV2RelationshipsControlVersions(val *AppStoreVersionExperimentV2RelationshipsControlVersions) *NullableAppStoreVersionExperimentV2RelationshipsControlVersions {
	return &NullableAppStoreVersionExperimentV2RelationshipsControlVersions{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentV2RelationshipsControlVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentV2RelationshipsControlVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

