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

// checks if the SubscriptionOfferCodeRelationshipsPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeRelationshipsPrices{}

// SubscriptionOfferCodeRelationshipsPrices struct for SubscriptionOfferCodeRelationshipsPrices
type SubscriptionOfferCodeRelationshipsPrices struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []SubscriptionOfferCodeRelationshipsPricesDataInner `json:"data,omitempty"`
}

// NewSubscriptionOfferCodeRelationshipsPrices instantiates a new SubscriptionOfferCodeRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeRelationshipsPrices() *SubscriptionOfferCodeRelationshipsPrices {
	this := SubscriptionOfferCodeRelationshipsPrices{}
	return &this
}

// NewSubscriptionOfferCodeRelationshipsPricesWithDefaults instantiates a new SubscriptionOfferCodeRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeRelationshipsPricesWithDefaults() *SubscriptionOfferCodeRelationshipsPrices {
	this := SubscriptionOfferCodeRelationshipsPrices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *SubscriptionOfferCodeRelationshipsPrices) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionOfferCodeRelationshipsPrices) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetData() []SubscriptionOfferCodeRelationshipsPricesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SubscriptionOfferCodeRelationshipsPricesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) GetDataOk() ([]SubscriptionOfferCodeRelationshipsPricesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationshipsPrices) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubscriptionOfferCodeRelationshipsPricesDataInner and assigns it to the Data field.
func (o *SubscriptionOfferCodeRelationshipsPrices) SetData(v []SubscriptionOfferCodeRelationshipsPricesDataInner) {
	o.Data = v
}

func (o SubscriptionOfferCodeRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeRelationshipsPrices) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionOfferCodeRelationshipsPrices struct {
	value *SubscriptionOfferCodeRelationshipsPrices
	isSet bool
}

func (v NullableSubscriptionOfferCodeRelationshipsPrices) Get() *SubscriptionOfferCodeRelationshipsPrices {
	return v.value
}

func (v *NullableSubscriptionOfferCodeRelationshipsPrices) Set(val *SubscriptionOfferCodeRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeRelationshipsPrices(val *SubscriptionOfferCodeRelationshipsPrices) *NullableSubscriptionOfferCodeRelationshipsPrices {
	return &NullableSubscriptionOfferCodeRelationshipsPrices{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


