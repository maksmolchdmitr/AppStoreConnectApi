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

// checks if the AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations{}

// AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations struct for AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations
type AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks                           `json:"links,omitempty"`
	Meta  *PagingInformation                                                                    `json:"meta,omitempty"`
	Data  []AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner `json:"data,omitempty"`
}

// NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations instantiates a new AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations() *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations {
	this := AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations{}
	return &this
}

// NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsWithDefaults instantiates a new AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsWithDefaults() *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations {
	this := AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetData() []AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) GetDataOk() ([]AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner and assigns it to the Data field.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) SetData(v []AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) {
	o.Data = v
}

func (o AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) ToMap() (map[string]interface{}, error) {
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

type NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations struct {
	value *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations
	isSet bool
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) Get() *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations {
	return v.value
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) Set(val *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations(val *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations {
	return &NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
