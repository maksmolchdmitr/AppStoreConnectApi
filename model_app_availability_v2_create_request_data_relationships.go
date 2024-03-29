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

// checks if the AppAvailabilityV2CreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2CreateRequestDataRelationships{}

// AppAvailabilityV2CreateRequestDataRelationships struct for AppAvailabilityV2CreateRequestDataRelationships
type AppAvailabilityV2CreateRequestDataRelationships struct {
	App                     AppAvailabilityV2CreateRequestDataRelationshipsApp                     `json:"app"`
	TerritoryAvailabilities AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities `json:"territoryAvailabilities"`
}

// NewAppAvailabilityV2CreateRequestDataRelationships instantiates a new AppAvailabilityV2CreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2CreateRequestDataRelationships(app AppAvailabilityV2CreateRequestDataRelationshipsApp, territoryAvailabilities AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities) *AppAvailabilityV2CreateRequestDataRelationships {
	this := AppAvailabilityV2CreateRequestDataRelationships{}
	this.App = app
	this.TerritoryAvailabilities = territoryAvailabilities
	return &this
}

// NewAppAvailabilityV2CreateRequestDataRelationshipsWithDefaults instantiates a new AppAvailabilityV2CreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2CreateRequestDataRelationshipsWithDefaults() *AppAvailabilityV2CreateRequestDataRelationships {
	this := AppAvailabilityV2CreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppAvailabilityV2CreateRequestDataRelationships) GetApp() AppAvailabilityV2CreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequestDataRelationships) GetAppOk() (*AppAvailabilityV2CreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppAvailabilityV2CreateRequestDataRelationships) SetApp(v AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetTerritoryAvailabilities returns the TerritoryAvailabilities field value
func (o *AppAvailabilityV2CreateRequestDataRelationships) GetTerritoryAvailabilities() AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities
		return ret
	}

	return o.TerritoryAvailabilities
}

// GetTerritoryAvailabilitiesOk returns a tuple with the TerritoryAvailabilities field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequestDataRelationships) GetTerritoryAvailabilitiesOk() (*AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerritoryAvailabilities, true
}

// SetTerritoryAvailabilities sets field value
func (o *AppAvailabilityV2CreateRequestDataRelationships) SetTerritoryAvailabilities(v AppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities) {
	o.TerritoryAvailabilities = v
}

func (o AppAvailabilityV2CreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2CreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["territoryAvailabilities"] = o.TerritoryAvailabilities
	return toSerialize, nil
}

type NullableAppAvailabilityV2CreateRequestDataRelationships struct {
	value *AppAvailabilityV2CreateRequestDataRelationships
	isSet bool
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationships) Get() *AppAvailabilityV2CreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationships) Set(val *AppAvailabilityV2CreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2CreateRequestDataRelationships(val *AppAvailabilityV2CreateRequestDataRelationships) *NullableAppAvailabilityV2CreateRequestDataRelationships {
	return &NullableAppAvailabilityV2CreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
