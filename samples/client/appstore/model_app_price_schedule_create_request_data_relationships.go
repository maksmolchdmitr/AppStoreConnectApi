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

// checks if the AppPriceScheduleCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceScheduleCreateRequestDataRelationships{}

// AppPriceScheduleCreateRequestDataRelationships struct for AppPriceScheduleCreateRequestDataRelationships
type AppPriceScheduleCreateRequestDataRelationships struct {
	App AppAvailabilityV2CreateRequestDataRelationshipsApp `json:"app"`
	BaseTerritory AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory `json:"baseTerritory"`
	ManualPrices AppPriceScheduleCreateRequestDataRelationshipsManualPrices `json:"manualPrices"`
}

// NewAppPriceScheduleCreateRequestDataRelationships instantiates a new AppPriceScheduleCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceScheduleCreateRequestDataRelationships(app AppAvailabilityV2CreateRequestDataRelationshipsApp, baseTerritory AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory, manualPrices AppPriceScheduleCreateRequestDataRelationshipsManualPrices) *AppPriceScheduleCreateRequestDataRelationships {
	this := AppPriceScheduleCreateRequestDataRelationships{}
	this.App = app
	this.BaseTerritory = baseTerritory
	this.ManualPrices = manualPrices
	return &this
}

// NewAppPriceScheduleCreateRequestDataRelationshipsWithDefaults instantiates a new AppPriceScheduleCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceScheduleCreateRequestDataRelationshipsWithDefaults() *AppPriceScheduleCreateRequestDataRelationships {
	this := AppPriceScheduleCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppPriceScheduleCreateRequestDataRelationships) GetApp() AppAvailabilityV2CreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityV2CreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppPriceScheduleCreateRequestDataRelationships) SetApp(v AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetBaseTerritory returns the BaseTerritory field value
func (o *AppPriceScheduleCreateRequestDataRelationships) GetBaseTerritory() AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory {
	if o == nil {
		var ret AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory
		return ret
	}

	return o.BaseTerritory
}

// GetBaseTerritoryOk returns a tuple with the BaseTerritory field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequestDataRelationships) GetBaseTerritoryOk() (*AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseTerritory, true
}

// SetBaseTerritory sets field value
func (o *AppPriceScheduleCreateRequestDataRelationships) SetBaseTerritory(v AppPriceScheduleCreateRequestDataRelationshipsBaseTerritory) {
	o.BaseTerritory = v
}

// GetManualPrices returns the ManualPrices field value
func (o *AppPriceScheduleCreateRequestDataRelationships) GetManualPrices() AppPriceScheduleCreateRequestDataRelationshipsManualPrices {
	if o == nil {
		var ret AppPriceScheduleCreateRequestDataRelationshipsManualPrices
		return ret
	}

	return o.ManualPrices
}

// GetManualPricesOk returns a tuple with the ManualPrices field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequestDataRelationships) GetManualPricesOk() (*AppPriceScheduleCreateRequestDataRelationshipsManualPrices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualPrices, true
}

// SetManualPrices sets field value
func (o *AppPriceScheduleCreateRequestDataRelationships) SetManualPrices(v AppPriceScheduleCreateRequestDataRelationshipsManualPrices) {
	o.ManualPrices = v
}

func (o AppPriceScheduleCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceScheduleCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["baseTerritory"] = o.BaseTerritory
	toSerialize["manualPrices"] = o.ManualPrices
	return toSerialize, nil
}

type NullableAppPriceScheduleCreateRequestDataRelationships struct {
	value *AppPriceScheduleCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppPriceScheduleCreateRequestDataRelationships) Get() *AppPriceScheduleCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationships) Set(val *AppPriceScheduleCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceScheduleCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceScheduleCreateRequestDataRelationships(val *AppPriceScheduleCreateRequestDataRelationships) *NullableAppPriceScheduleCreateRequestDataRelationships {
	return &NullableAppPriceScheduleCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppPriceScheduleCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


