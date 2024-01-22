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

// checks if the InAppPurchaseV2RelationshipsAppStoreReviewScreenshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsAppStoreReviewScreenshot{}

// InAppPurchaseV2RelationshipsAppStoreReviewScreenshot struct for InAppPurchaseV2RelationshipsAppStoreReviewScreenshot
type InAppPurchaseV2RelationshipsAppStoreReviewScreenshot struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData `json:"data,omitempty"`
}

// NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshot instantiates a new InAppPurchaseV2RelationshipsAppStoreReviewScreenshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshot() *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot {
	this := InAppPurchaseV2RelationshipsAppStoreReviewScreenshot{}
	return &this
}

// NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotWithDefaults instantiates a new InAppPurchaseV2RelationshipsAppStoreReviewScreenshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotWithDefaults() *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot {
	this := InAppPurchaseV2RelationshipsAppStoreReviewScreenshot{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) GetData() InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData {
	if o == nil || IsNil(o.Data) {
		var ret InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) GetDataOk() (*InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData and assigns it to the Data field.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) SetData(v InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) {
	o.Data = &v
}

func (o InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot struct {
	value *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) Get() *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) Set(val *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot(val *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot {
	return &NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


