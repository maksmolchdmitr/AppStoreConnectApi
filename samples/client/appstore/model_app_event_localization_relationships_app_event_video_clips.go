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

// checks if the AppEventLocalizationRelationshipsAppEventVideoClips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationshipsAppEventVideoClips{}

// AppEventLocalizationRelationshipsAppEventVideoClips struct for AppEventLocalizationRelationshipsAppEventVideoClips
type AppEventLocalizationRelationshipsAppEventVideoClips struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppEventLocalizationRelationshipsAppEventVideoClipsDataInner `json:"data,omitempty"`
}

// NewAppEventLocalizationRelationshipsAppEventVideoClips instantiates a new AppEventLocalizationRelationshipsAppEventVideoClips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationshipsAppEventVideoClips() *AppEventLocalizationRelationshipsAppEventVideoClips {
	this := AppEventLocalizationRelationshipsAppEventVideoClips{}
	return &this
}

// NewAppEventLocalizationRelationshipsAppEventVideoClipsWithDefaults instantiates a new AppEventLocalizationRelationshipsAppEventVideoClips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsAppEventVideoClipsWithDefaults() *AppEventLocalizationRelationshipsAppEventVideoClips {
	this := AppEventLocalizationRelationshipsAppEventVideoClips{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetData() []AppEventLocalizationRelationshipsAppEventVideoClipsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppEventLocalizationRelationshipsAppEventVideoClipsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) GetDataOk() ([]AppEventLocalizationRelationshipsAppEventVideoClipsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppEventLocalizationRelationshipsAppEventVideoClipsDataInner and assigns it to the Data field.
func (o *AppEventLocalizationRelationshipsAppEventVideoClips) SetData(v []AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) {
	o.Data = v
}

func (o AppEventLocalizationRelationshipsAppEventVideoClips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationshipsAppEventVideoClips) ToMap() (map[string]interface{}, error) {
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

type NullableAppEventLocalizationRelationshipsAppEventVideoClips struct {
	value *AppEventLocalizationRelationshipsAppEventVideoClips
	isSet bool
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClips) Get() *AppEventLocalizationRelationshipsAppEventVideoClips {
	return v.value
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClips) Set(val *AppEventLocalizationRelationshipsAppEventVideoClips) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClips) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationshipsAppEventVideoClips(val *AppEventLocalizationRelationshipsAppEventVideoClips) *NullableAppEventLocalizationRelationshipsAppEventVideoClips {
	return &NullableAppEventLocalizationRelationshipsAppEventVideoClips{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

