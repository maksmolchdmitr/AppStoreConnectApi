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

// checks if the AgeRatingDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRatingDeclaration{}

// AgeRatingDeclaration struct for AgeRatingDeclaration
type AgeRatingDeclaration struct {
	Type       string                          `json:"type"`
	Id         string                          `json:"id"`
	Attributes *AgeRatingDeclarationAttributes `json:"attributes,omitempty"`
	Links      *ResourceLinks                  `json:"links,omitempty"`
}

// NewAgeRatingDeclaration instantiates a new AgeRatingDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRatingDeclaration(type_ string, id string) *AgeRatingDeclaration {
	this := AgeRatingDeclaration{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAgeRatingDeclarationWithDefaults instantiates a new AgeRatingDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRatingDeclarationWithDefaults() *AgeRatingDeclaration {
	this := AgeRatingDeclaration{}
	return &this
}

// GetType returns the Type field value
func (o *AgeRatingDeclaration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclaration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgeRatingDeclaration) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AgeRatingDeclaration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclaration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgeRatingDeclaration) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AgeRatingDeclaration) GetAttributes() AgeRatingDeclarationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AgeRatingDeclarationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclaration) GetAttributesOk() (*AgeRatingDeclarationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AgeRatingDeclaration) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AgeRatingDeclarationAttributes and assigns it to the Attributes field.
func (o *AgeRatingDeclaration) SetAttributes(v AgeRatingDeclarationAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgeRatingDeclaration) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclaration) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgeRatingDeclaration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AgeRatingDeclaration) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeRatingDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAgeRatingDeclaration struct {
	value *AgeRatingDeclaration
	isSet bool
}

func (v NullableAgeRatingDeclaration) Get() *AgeRatingDeclaration {
	return v.value
}

func (v *NullableAgeRatingDeclaration) Set(val *AgeRatingDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRatingDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRatingDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRatingDeclaration(val *AgeRatingDeclaration) *NullableAgeRatingDeclaration {
	return &NullableAgeRatingDeclaration{value: val, isSet: true}
}

func (v NullableAgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeRatingDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
