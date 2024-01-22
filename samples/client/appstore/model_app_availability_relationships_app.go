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

// checks if the AppAvailabilityRelationshipsApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityRelationshipsApp{}

// AppAvailabilityRelationshipsApp struct for AppAvailabilityRelationshipsApp
type AppAvailabilityRelationshipsApp struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data *AppAvailabilityV2CreateRequestDataRelationshipsAppData `json:"data,omitempty"`
}

// NewAppAvailabilityRelationshipsApp instantiates a new AppAvailabilityRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityRelationshipsApp() *AppAvailabilityRelationshipsApp {
	this := AppAvailabilityRelationshipsApp{}
	return &this
}

// NewAppAvailabilityRelationshipsAppWithDefaults instantiates a new AppAvailabilityRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityRelationshipsAppWithDefaults() *AppAvailabilityRelationshipsApp {
	this := AppAvailabilityRelationshipsApp{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppAvailabilityRelationshipsApp) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationshipsApp) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppAvailabilityRelationshipsApp) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppAvailabilityRelationshipsApp) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppAvailabilityRelationshipsApp) GetData() AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	if o == nil || IsNil(o.Data) {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsAppData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationshipsApp) GetDataOk() (*AppAvailabilityV2CreateRequestDataRelationshipsAppData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppAvailabilityRelationshipsApp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppAvailabilityV2CreateRequestDataRelationshipsAppData and assigns it to the Data field.
func (o *AppAvailabilityRelationshipsApp) SetData(v AppAvailabilityV2CreateRequestDataRelationshipsAppData) {
	o.Data = &v
}

func (o AppAvailabilityRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityRelationshipsApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppAvailabilityRelationshipsApp struct {
	value *AppAvailabilityRelationshipsApp
	isSet bool
}

func (v NullableAppAvailabilityRelationshipsApp) Get() *AppAvailabilityRelationshipsApp {
	return v.value
}

func (v *NullableAppAvailabilityRelationshipsApp) Set(val *AppAvailabilityRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityRelationshipsApp(val *AppAvailabilityRelationshipsApp) *NullableAppAvailabilityRelationshipsApp {
	return &NullableAppAvailabilityRelationshipsApp{value: val, isSet: true}
}

func (v NullableAppAvailabilityRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


