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

// checks if the BetaLicenseAgreementUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreementUpdateRequestData{}

// BetaLicenseAgreementUpdateRequestData struct for BetaLicenseAgreementUpdateRequestData
type BetaLicenseAgreementUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaLicenseAgreementAttributes `json:"attributes,omitempty"`
}

// NewBetaLicenseAgreementUpdateRequestData instantiates a new BetaLicenseAgreementUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreementUpdateRequestData(type_ string, id string) *BetaLicenseAgreementUpdateRequestData {
	this := BetaLicenseAgreementUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaLicenseAgreementUpdateRequestDataWithDefaults instantiates a new BetaLicenseAgreementUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementUpdateRequestDataWithDefaults() *BetaLicenseAgreementUpdateRequestData {
	this := BetaLicenseAgreementUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaLicenseAgreementUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaLicenseAgreementUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaLicenseAgreementUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaLicenseAgreementUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaLicenseAgreementUpdateRequestData) GetAttributes() BetaLicenseAgreementAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaLicenseAgreementAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementUpdateRequestData) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaLicenseAgreementUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaLicenseAgreementAttributes and assigns it to the Attributes field.
func (o *BetaLicenseAgreementUpdateRequestData) SetAttributes(v BetaLicenseAgreementAttributes) {
	o.Attributes = &v
}

func (o BetaLicenseAgreementUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreementUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBetaLicenseAgreementUpdateRequestData struct {
	value *BetaLicenseAgreementUpdateRequestData
	isSet bool
}

func (v NullableBetaLicenseAgreementUpdateRequestData) Get() *BetaLicenseAgreementUpdateRequestData {
	return v.value
}

func (v *NullableBetaLicenseAgreementUpdateRequestData) Set(val *BetaLicenseAgreementUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreementUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreementUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreementUpdateRequestData(val *BetaLicenseAgreementUpdateRequestData) *NullableBetaLicenseAgreementUpdateRequestData {
	return &NullableBetaLicenseAgreementUpdateRequestData{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreementUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreementUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


