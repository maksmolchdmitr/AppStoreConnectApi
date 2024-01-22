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

// checks if the InAppPurchaseV2RelationshipsPricePoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsPricePoints{}

// InAppPurchaseV2RelationshipsPricePoints struct for InAppPurchaseV2RelationshipsPricePoints
type InAppPurchaseV2RelationshipsPricePoints struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks  `json:"links,omitempty"`
	Meta  *PagingInformation                                           `json:"meta,omitempty"`
	Data  []InAppPurchasePriceRelationshipsInAppPurchasePricePointData `json:"data,omitempty"`
}

// NewInAppPurchaseV2RelationshipsPricePoints instantiates a new InAppPurchaseV2RelationshipsPricePoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsPricePoints() *InAppPurchaseV2RelationshipsPricePoints {
	this := InAppPurchaseV2RelationshipsPricePoints{}
	return &this
}

// NewInAppPurchaseV2RelationshipsPricePointsWithDefaults instantiates a new InAppPurchaseV2RelationshipsPricePoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsPricePointsWithDefaults() *InAppPurchaseV2RelationshipsPricePoints {
	this := InAppPurchaseV2RelationshipsPricePoints{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *InAppPurchaseV2RelationshipsPricePoints) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *InAppPurchaseV2RelationshipsPricePoints) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetData() []InAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	if o == nil || IsNil(o.Data) {
		var ret []InAppPurchasePriceRelationshipsInAppPurchasePricePointData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) GetDataOk() ([]InAppPurchasePriceRelationshipsInAppPurchasePricePointData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsPricePoints) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InAppPurchasePriceRelationshipsInAppPurchasePricePointData and assigns it to the Data field.
func (o *InAppPurchaseV2RelationshipsPricePoints) SetData(v []InAppPurchasePriceRelationshipsInAppPurchasePricePointData) {
	o.Data = v
}

func (o InAppPurchaseV2RelationshipsPricePoints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsPricePoints) ToMap() (map[string]interface{}, error) {
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

type NullableInAppPurchaseV2RelationshipsPricePoints struct {
	value *InAppPurchaseV2RelationshipsPricePoints
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsPricePoints) Get() *InAppPurchaseV2RelationshipsPricePoints {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsPricePoints) Set(val *InAppPurchaseV2RelationshipsPricePoints) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsPricePoints) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsPricePoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsPricePoints(val *InAppPurchaseV2RelationshipsPricePoints) *NullableInAppPurchaseV2RelationshipsPricePoints {
	return &NullableInAppPurchaseV2RelationshipsPricePoints{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsPricePoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsPricePoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
