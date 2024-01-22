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

// checks if the TerritoryAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryAvailability{}

// TerritoryAvailability struct for TerritoryAvailability
type TerritoryAvailability struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *TerritoryAvailabilityAttributes `json:"attributes,omitempty"`
	Relationships *InAppPurchasePricePointRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewTerritoryAvailability instantiates a new TerritoryAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryAvailability(type_ string, id string) *TerritoryAvailability {
	this := TerritoryAvailability{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTerritoryAvailabilityWithDefaults instantiates a new TerritoryAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryAvailabilityWithDefaults() *TerritoryAvailability {
	this := TerritoryAvailability{}
	return &this
}

// GetType returns the Type field value
func (o *TerritoryAvailability) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerritoryAvailability) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerritoryAvailability) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TerritoryAvailability) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TerritoryAvailability) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TerritoryAvailability) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TerritoryAvailability) GetAttributes() TerritoryAvailabilityAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret TerritoryAvailabilityAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailability) GetAttributesOk() (*TerritoryAvailabilityAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TerritoryAvailability) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TerritoryAvailabilityAttributes and assigns it to the Attributes field.
func (o *TerritoryAvailability) SetAttributes(v TerritoryAvailabilityAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TerritoryAvailability) GetRelationships() InAppPurchasePricePointRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InAppPurchasePricePointRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailability) GetRelationshipsOk() (*InAppPurchasePricePointRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TerritoryAvailability) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InAppPurchasePricePointRelationships and assigns it to the Relationships field.
func (o *TerritoryAvailability) SetRelationships(v InAppPurchasePricePointRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TerritoryAvailability) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailability) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TerritoryAvailability) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *TerritoryAvailability) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o TerritoryAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTerritoryAvailability struct {
	value *TerritoryAvailability
	isSet bool
}

func (v NullableTerritoryAvailability) Get() *TerritoryAvailability {
	return v.value
}

func (v *NullableTerritoryAvailability) Set(val *TerritoryAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryAvailability(val *TerritoryAvailability) *NullableTerritoryAvailability {
	return &NullableTerritoryAvailability{value: val, isSet: true}
}

func (v NullableTerritoryAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

