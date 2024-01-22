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

// checks if the SubscriptionOfferCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCode{}

// SubscriptionOfferCode struct for SubscriptionOfferCode
type SubscriptionOfferCode struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionOfferCodeAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionOfferCodeRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewSubscriptionOfferCode instantiates a new SubscriptionOfferCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCode(type_ string, id string) *SubscriptionOfferCode {
	this := SubscriptionOfferCode{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeWithDefaults instantiates a new SubscriptionOfferCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeWithDefaults() *SubscriptionOfferCode {
	this := SubscriptionOfferCode{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCode) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCode) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionOfferCode) GetAttributes() SubscriptionOfferCodeAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionOfferCodeAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCode) GetAttributesOk() (*SubscriptionOfferCodeAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionOfferCode) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionOfferCodeAttributes and assigns it to the Attributes field.
func (o *SubscriptionOfferCode) SetAttributes(v SubscriptionOfferCodeAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionOfferCode) GetRelationships() SubscriptionOfferCodeRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionOfferCodeRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCode) GetRelationshipsOk() (*SubscriptionOfferCodeRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionOfferCode) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionOfferCodeRelationships and assigns it to the Relationships field.
func (o *SubscriptionOfferCode) SetRelationships(v SubscriptionOfferCodeRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionOfferCode) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCode) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionOfferCode) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *SubscriptionOfferCode) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o SubscriptionOfferCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCode) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionOfferCode struct {
	value *SubscriptionOfferCode
	isSet bool
}

func (v NullableSubscriptionOfferCode) Get() *SubscriptionOfferCode {
	return v.value
}

func (v *NullableSubscriptionOfferCode) Set(val *SubscriptionOfferCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCode(val *SubscriptionOfferCode) *NullableSubscriptionOfferCode {
	return &NullableSubscriptionOfferCode{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


