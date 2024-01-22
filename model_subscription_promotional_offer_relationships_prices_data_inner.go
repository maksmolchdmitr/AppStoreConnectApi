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

// checks if the SubscriptionPromotionalOfferRelationshipsPricesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferRelationshipsPricesDataInner{}

// SubscriptionPromotionalOfferRelationshipsPricesDataInner struct for SubscriptionPromotionalOfferRelationshipsPricesDataInner
type SubscriptionPromotionalOfferRelationshipsPricesDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewSubscriptionPromotionalOfferRelationshipsPricesDataInner instantiates a new SubscriptionPromotionalOfferRelationshipsPricesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferRelationshipsPricesDataInner(type_ string, id string) *SubscriptionPromotionalOfferRelationshipsPricesDataInner {
	this := SubscriptionPromotionalOfferRelationshipsPricesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionPromotionalOfferRelationshipsPricesDataInnerWithDefaults instantiates a new SubscriptionPromotionalOfferRelationshipsPricesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferRelationshipsPricesDataInnerWithDefaults() *SubscriptionPromotionalOfferRelationshipsPricesDataInner {
	this := SubscriptionPromotionalOfferRelationshipsPricesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionPromotionalOfferRelationshipsPricesDataInner) SetId(v string) {
	o.Id = v
}

func (o SubscriptionPromotionalOfferRelationshipsPricesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferRelationshipsPricesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner struct {
	value *SubscriptionPromotionalOfferRelationshipsPricesDataInner
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) Get() *SubscriptionPromotionalOfferRelationshipsPricesDataInner {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) Set(val *SubscriptionPromotionalOfferRelationshipsPricesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferRelationshipsPricesDataInner(val *SubscriptionPromotionalOfferRelationshipsPricesDataInner) *NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner {
	return &NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPricesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
