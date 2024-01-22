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

// checks if the SubscriptionIntroductoryOfferUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferUpdateRequestData{}

// SubscriptionIntroductoryOfferUpdateRequestData struct for SubscriptionIntroductoryOfferUpdateRequestData
type SubscriptionIntroductoryOfferUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionIntroductoryOfferUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewSubscriptionIntroductoryOfferUpdateRequestData instantiates a new SubscriptionIntroductoryOfferUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferUpdateRequestData(type_ string, id string) *SubscriptionIntroductoryOfferUpdateRequestData {
	this := SubscriptionIntroductoryOfferUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionIntroductoryOfferUpdateRequestDataWithDefaults instantiates a new SubscriptionIntroductoryOfferUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferUpdateRequestDataWithDefaults() *SubscriptionIntroductoryOfferUpdateRequestData {
	this := SubscriptionIntroductoryOfferUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionIntroductoryOfferUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionIntroductoryOfferUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetAttributes() SubscriptionIntroductoryOfferUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionIntroductoryOfferUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) GetAttributesOk() (*SubscriptionIntroductoryOfferUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionIntroductoryOfferUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionIntroductoryOfferUpdateRequestData) SetAttributes(v SubscriptionIntroductoryOfferUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o SubscriptionIntroductoryOfferUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferUpdateRequestData struct {
	value *SubscriptionIntroductoryOfferUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestData) Get() *SubscriptionIntroductoryOfferUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestData) Set(val *SubscriptionIntroductoryOfferUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferUpdateRequestData(val *SubscriptionIntroductoryOfferUpdateRequestData) *NullableSubscriptionIntroductoryOfferUpdateRequestData {
	return &NullableSubscriptionIntroductoryOfferUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


