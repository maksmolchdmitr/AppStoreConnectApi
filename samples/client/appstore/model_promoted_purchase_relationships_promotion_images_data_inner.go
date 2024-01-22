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

// checks if the PromotedPurchaseRelationshipsPromotionImagesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseRelationshipsPromotionImagesDataInner{}

// PromotedPurchaseRelationshipsPromotionImagesDataInner struct for PromotedPurchaseRelationshipsPromotionImagesDataInner
type PromotedPurchaseRelationshipsPromotionImagesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewPromotedPurchaseRelationshipsPromotionImagesDataInner instantiates a new PromotedPurchaseRelationshipsPromotionImagesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseRelationshipsPromotionImagesDataInner(type_ string, id string) *PromotedPurchaseRelationshipsPromotionImagesDataInner {
	this := PromotedPurchaseRelationshipsPromotionImagesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPromotedPurchaseRelationshipsPromotionImagesDataInnerWithDefaults instantiates a new PromotedPurchaseRelationshipsPromotionImagesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseRelationshipsPromotionImagesDataInnerWithDefaults() *PromotedPurchaseRelationshipsPromotionImagesDataInner {
	this := PromotedPurchaseRelationshipsPromotionImagesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PromotedPurchaseRelationshipsPromotionImagesDataInner) SetId(v string) {
	o.Id = v
}

func (o PromotedPurchaseRelationshipsPromotionImagesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseRelationshipsPromotionImagesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePromotedPurchaseRelationshipsPromotionImagesDataInner struct {
	value *PromotedPurchaseRelationshipsPromotionImagesDataInner
	isSet bool
}

func (v NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) Get() *PromotedPurchaseRelationshipsPromotionImagesDataInner {
	return v.value
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) Set(val *PromotedPurchaseRelationshipsPromotionImagesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseRelationshipsPromotionImagesDataInner(val *PromotedPurchaseRelationshipsPromotionImagesDataInner) *NullablePromotedPurchaseRelationshipsPromotionImagesDataInner {
	return &NullablePromotedPurchaseRelationshipsPromotionImagesDataInner{value: val, isSet: true}
}

func (v NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImagesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


