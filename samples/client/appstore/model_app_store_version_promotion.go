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

// checks if the AppStoreVersionPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPromotion{}

// AppStoreVersionPromotion struct for AppStoreVersionPromotion
type AppStoreVersionPromotion struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewAppStoreVersionPromotion instantiates a new AppStoreVersionPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPromotion(type_ string, id string) *AppStoreVersionPromotion {
	this := AppStoreVersionPromotion{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionPromotionWithDefaults instantiates a new AppStoreVersionPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPromotionWithDefaults() *AppStoreVersionPromotion {
	this := AppStoreVersionPromotion{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionPromotion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionPromotion) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionPromotion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionPromotion) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionPromotion) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotion) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionPromotion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppStoreVersionPromotion) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppStoreVersionPromotion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAppStoreVersionPromotion struct {
	value *AppStoreVersionPromotion
	isSet bool
}

func (v NullableAppStoreVersionPromotion) Get() *AppStoreVersionPromotion {
	return v.value
}

func (v *NullableAppStoreVersionPromotion) Set(val *AppStoreVersionPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPromotion(val *AppStoreVersionPromotion) *NullableAppStoreVersionPromotion {
	return &NullableAppStoreVersionPromotion{value: val, isSet: true}
}

func (v NullableAppStoreVersionPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

