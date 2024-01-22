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

// checks if the AppPriceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceRelationships{}

// AppPriceRelationships struct for AppPriceRelationships
type AppPriceRelationships struct {
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	PriceTier *AppPricePointV2RelationshipsPriceTier `json:"priceTier,omitempty"`
}

// NewAppPriceRelationships instantiates a new AppPriceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceRelationships() *AppPriceRelationships {
	this := AppPriceRelationships{}
	return &this
}

// NewAppPriceRelationshipsWithDefaults instantiates a new AppPriceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceRelationshipsWithDefaults() *AppPriceRelationships {
	this := AppPriceRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppPriceRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppPriceRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *AppPriceRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetPriceTier returns the PriceTier field value if set, zero value otherwise.
func (o *AppPriceRelationships) GetPriceTier() AppPricePointV2RelationshipsPriceTier {
	if o == nil || IsNil(o.PriceTier) {
		var ret AppPricePointV2RelationshipsPriceTier
		return ret
	}
	return *o.PriceTier
}

// GetPriceTierOk returns a tuple with the PriceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceRelationships) GetPriceTierOk() (*AppPricePointV2RelationshipsPriceTier, bool) {
	if o == nil || IsNil(o.PriceTier) {
		return nil, false
	}
	return o.PriceTier, true
}

// HasPriceTier returns a boolean if a field has been set.
func (o *AppPriceRelationships) HasPriceTier() bool {
	if o != nil && !IsNil(o.PriceTier) {
		return true
	}

	return false
}

// SetPriceTier gets a reference to the given AppPricePointV2RelationshipsPriceTier and assigns it to the PriceTier field.
func (o *AppPriceRelationships) SetPriceTier(v AppPricePointV2RelationshipsPriceTier) {
	o.PriceTier = &v
}

func (o AppPriceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.PriceTier) {
		toSerialize["priceTier"] = o.PriceTier
	}
	return toSerialize, nil
}

type NullableAppPriceRelationships struct {
	value *AppPriceRelationships
	isSet bool
}

func (v NullableAppPriceRelationships) Get() *AppPriceRelationships {
	return v.value
}

func (v *NullableAppPriceRelationships) Set(val *AppPriceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceRelationships(val *AppPriceRelationships) *NullableAppPriceRelationships {
	return &NullableAppPriceRelationships{value: val, isSet: true}
}

func (v NullableAppPriceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


