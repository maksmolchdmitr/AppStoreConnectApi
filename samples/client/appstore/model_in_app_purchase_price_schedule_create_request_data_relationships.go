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

// checks if the InAppPurchasePriceScheduleCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceScheduleCreateRequestDataRelationships{}

// InAppPurchasePriceScheduleCreateRequestDataRelationships struct for InAppPurchasePriceScheduleCreateRequestDataRelationships
type InAppPurchasePriceScheduleCreateRequestDataRelationships struct {
	InAppPurchase InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2 `json:"inAppPurchase"`
	BaseTerritory *InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory `json:"baseTerritory,omitempty"`
	ManualPrices InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices `json:"manualPrices"`
}

// NewInAppPurchasePriceScheduleCreateRequestDataRelationships instantiates a new InAppPurchasePriceScheduleCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceScheduleCreateRequestDataRelationships(inAppPurchase InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2, manualPrices InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices) *InAppPurchasePriceScheduleCreateRequestDataRelationships {
	this := InAppPurchasePriceScheduleCreateRequestDataRelationships{}
	this.InAppPurchase = inAppPurchase
	this.ManualPrices = manualPrices
	return &this
}

// NewInAppPurchasePriceScheduleCreateRequestDataRelationshipsWithDefaults instantiates a new InAppPurchasePriceScheduleCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceScheduleCreateRequestDataRelationshipsWithDefaults() *InAppPurchasePriceScheduleCreateRequestDataRelationships {
	this := InAppPurchasePriceScheduleCreateRequestDataRelationships{}
	return &this
}

// GetInAppPurchase returns the InAppPurchase field value
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetInAppPurchase() InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2 {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2
		return ret
	}

	return o.InAppPurchase
}

// GetInAppPurchaseOk returns a tuple with the InAppPurchase field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetInAppPurchaseOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAppPurchase, true
}

// SetInAppPurchase sets field value
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) SetInAppPurchase(v InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2) {
	o.InAppPurchase = v
}

// GetBaseTerritory returns the BaseTerritory field value if set, zero value otherwise.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetBaseTerritory() InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory {
	if o == nil || IsNil(o.BaseTerritory) {
		var ret InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory
		return ret
	}
	return *o.BaseTerritory
}

// GetBaseTerritoryOk returns a tuple with the BaseTerritory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetBaseTerritoryOk() (*InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory, bool) {
	if o == nil || IsNil(o.BaseTerritory) {
		return nil, false
	}
	return o.BaseTerritory, true
}

// HasBaseTerritory returns a boolean if a field has been set.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) HasBaseTerritory() bool {
	if o != nil && !IsNil(o.BaseTerritory) {
		return true
	}

	return false
}

// SetBaseTerritory gets a reference to the given InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory and assigns it to the BaseTerritory field.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) SetBaseTerritory(v InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory) {
	o.BaseTerritory = &v
}

// GetManualPrices returns the ManualPrices field value
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetManualPrices() InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices {
	if o == nil {
		var ret InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices
		return ret
	}

	return o.ManualPrices
}

// GetManualPricesOk returns a tuple with the ManualPrices field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) GetManualPricesOk() (*InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualPrices, true
}

// SetManualPrices sets field value
func (o *InAppPurchasePriceScheduleCreateRequestDataRelationships) SetManualPrices(v InAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices) {
	o.ManualPrices = v
}

func (o InAppPurchasePriceScheduleCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceScheduleCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inAppPurchase"] = o.InAppPurchase
	if !IsNil(o.BaseTerritory) {
		toSerialize["baseTerritory"] = o.BaseTerritory
	}
	toSerialize["manualPrices"] = o.ManualPrices
	return toSerialize, nil
}

type NullableInAppPurchasePriceScheduleCreateRequestDataRelationships struct {
	value *InAppPurchasePriceScheduleCreateRequestDataRelationships
	isSet bool
}

func (v NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) Get() *InAppPurchasePriceScheduleCreateRequestDataRelationships {
	return v.value
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) Set(val *InAppPurchasePriceScheduleCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceScheduleCreateRequestDataRelationships(val *InAppPurchasePriceScheduleCreateRequestDataRelationships) *NullableInAppPurchasePriceScheduleCreateRequestDataRelationships {
	return &NullableInAppPurchasePriceScheduleCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


