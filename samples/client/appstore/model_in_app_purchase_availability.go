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

// checks if the InAppPurchaseAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAvailability{}

// InAppPurchaseAvailability struct for InAppPurchaseAvailability
type InAppPurchaseAvailability struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppAvailabilityV2Attributes `json:"attributes,omitempty"`
	Relationships *InAppPurchaseAvailabilityRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewInAppPurchaseAvailability instantiates a new InAppPurchaseAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAvailability(type_ string, id string) *InAppPurchaseAvailability {
	this := InAppPurchaseAvailability{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchaseAvailabilityWithDefaults instantiates a new InAppPurchaseAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAvailabilityWithDefaults() *InAppPurchaseAvailability {
	this := InAppPurchaseAvailability{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseAvailability) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailability) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseAvailability) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchaseAvailability) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailability) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchaseAvailability) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InAppPurchaseAvailability) GetAttributes() AppAvailabilityV2Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppAvailabilityV2Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailability) GetAttributesOk() (*AppAvailabilityV2Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InAppPurchaseAvailability) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppAvailabilityV2Attributes and assigns it to the Attributes field.
func (o *InAppPurchaseAvailability) SetAttributes(v AppAvailabilityV2Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InAppPurchaseAvailability) GetRelationships() InAppPurchaseAvailabilityRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InAppPurchaseAvailabilityRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailability) GetRelationshipsOk() (*InAppPurchaseAvailabilityRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InAppPurchaseAvailability) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InAppPurchaseAvailabilityRelationships and assigns it to the Relationships field.
func (o *InAppPurchaseAvailability) SetRelationships(v InAppPurchaseAvailabilityRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InAppPurchaseAvailability) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailability) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InAppPurchaseAvailability) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *InAppPurchaseAvailability) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o InAppPurchaseAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAvailability) ToMap() (map[string]interface{}, error) {
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

type NullableInAppPurchaseAvailability struct {
	value *InAppPurchaseAvailability
	isSet bool
}

func (v NullableInAppPurchaseAvailability) Get() *InAppPurchaseAvailability {
	return v.value
}

func (v *NullableInAppPurchaseAvailability) Set(val *InAppPurchaseAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAvailability(val *InAppPurchaseAvailability) *NullableInAppPurchaseAvailability {
	return &NullableInAppPurchaseAvailability{value: val, isSet: true}
}

func (v NullableInAppPurchaseAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


