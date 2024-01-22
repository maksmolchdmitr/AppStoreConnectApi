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

// checks if the AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration{}

// AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration struct for AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration
type AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration struct {
	Data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData `json:"data"`
}

// NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration instantiates a new AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration(data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration {
	this := AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration{}
	this.Data = data
	return &this
}

// NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationWithDefaults instantiates a new AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationWithDefaults() *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration {
	this := AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration{}
	return &this
}

// GetData returns the Data field value
func (o *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) GetData() AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData {
	if o == nil {
		var ret AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) SetData(v AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) {
	o.Data = v
}

func (o AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration struct {
	value *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration
	isSet bool
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) Get() *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration {
	return v.value
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) Set(val *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration(val *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) *NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration {
	return &NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


