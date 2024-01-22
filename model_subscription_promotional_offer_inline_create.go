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

// checks if the SubscriptionPromotionalOfferInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferInlineCreate{}

// SubscriptionPromotionalOfferInlineCreate struct for SubscriptionPromotionalOfferInlineCreate
type SubscriptionPromotionalOfferInlineCreate struct {
	Type          string                                                 `json:"type"`
	Id            *string                                                `json:"id,omitempty"`
	Attributes    SubscriptionPromotionalOfferInlineCreateAttributes     `json:"attributes"`
	Relationships *SubscriptionPromotionalOfferInlineCreateRelationships `json:"relationships,omitempty"`
}

// NewSubscriptionPromotionalOfferInlineCreate instantiates a new SubscriptionPromotionalOfferInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferInlineCreate(type_ string, attributes SubscriptionPromotionalOfferInlineCreateAttributes) *SubscriptionPromotionalOfferInlineCreate {
	this := SubscriptionPromotionalOfferInlineCreate{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSubscriptionPromotionalOfferInlineCreateWithDefaults instantiates a new SubscriptionPromotionalOfferInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferInlineCreateWithDefaults() *SubscriptionPromotionalOfferInlineCreate {
	this := SubscriptionPromotionalOfferInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPromotionalOfferInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPromotionalOfferInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionPromotionalOfferInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionPromotionalOfferInlineCreate) GetAttributes() SubscriptionPromotionalOfferInlineCreateAttributes {
	if o == nil {
		var ret SubscriptionPromotionalOfferInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) GetAttributesOk() (*SubscriptionPromotionalOfferInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionPromotionalOfferInlineCreate) SetAttributes(v SubscriptionPromotionalOfferInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferInlineCreate) GetRelationships() SubscriptionPromotionalOfferInlineCreateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionPromotionalOfferInlineCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) GetRelationshipsOk() (*SubscriptionPromotionalOfferInlineCreateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionPromotionalOfferInlineCreateRelationships and assigns it to the Relationships field.
func (o *SubscriptionPromotionalOfferInlineCreate) SetRelationships(v SubscriptionPromotionalOfferInlineCreateRelationships) {
	o.Relationships = &v
}

func (o SubscriptionPromotionalOfferInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferInlineCreate struct {
	value *SubscriptionPromotionalOfferInlineCreate
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferInlineCreate) Get() *SubscriptionPromotionalOfferInlineCreate {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferInlineCreate) Set(val *SubscriptionPromotionalOfferInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferInlineCreate(val *SubscriptionPromotionalOfferInlineCreate) *NullableSubscriptionPromotionalOfferInlineCreate {
	return &NullableSubscriptionPromotionalOfferInlineCreate{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
