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

// checks if the InAppPurchaseAppStoreReviewScreenshotUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAppStoreReviewScreenshotUpdateRequestData{}

// InAppPurchaseAppStoreReviewScreenshotUpdateRequestData struct for InAppPurchaseAppStoreReviewScreenshotUpdateRequestData
type InAppPurchaseAppStoreReviewScreenshotUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceImageUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewInAppPurchaseAppStoreReviewScreenshotUpdateRequestData instantiates a new InAppPurchaseAppStoreReviewScreenshotUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAppStoreReviewScreenshotUpdateRequestData(type_ string, id string) *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData {
	this := InAppPurchaseAppStoreReviewScreenshotUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchaseAppStoreReviewScreenshotUpdateRequestDataWithDefaults instantiates a new InAppPurchaseAppStoreReviewScreenshotUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAppStoreReviewScreenshotUpdateRequestDataWithDefaults() *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData {
	this := InAppPurchaseAppStoreReviewScreenshotUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetAttributes() AppClipAdvancedExperienceImageUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) SetAttributes(v AppClipAdvancedExperienceImageUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData struct {
	value *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData
	isSet bool
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) Get() *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData {
	return v.value
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) Set(val *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData(val *InAppPurchaseAppStoreReviewScreenshotUpdateRequestData) *NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData {
	return &NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData{value: val, isSet: true}
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


