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

// checks if the CiProductRelationshipsPrimaryRepositoriesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiProductRelationshipsPrimaryRepositoriesDataInner{}

// CiProductRelationshipsPrimaryRepositoriesDataInner struct for CiProductRelationshipsPrimaryRepositoriesDataInner
type CiProductRelationshipsPrimaryRepositoriesDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewCiProductRelationshipsPrimaryRepositoriesDataInner instantiates a new CiProductRelationshipsPrimaryRepositoriesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiProductRelationshipsPrimaryRepositoriesDataInner(type_ string, id string) *CiProductRelationshipsPrimaryRepositoriesDataInner {
	this := CiProductRelationshipsPrimaryRepositoriesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiProductRelationshipsPrimaryRepositoriesDataInnerWithDefaults instantiates a new CiProductRelationshipsPrimaryRepositoriesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiProductRelationshipsPrimaryRepositoriesDataInnerWithDefaults() *CiProductRelationshipsPrimaryRepositoriesDataInner {
	this := CiProductRelationshipsPrimaryRepositoriesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiProductRelationshipsPrimaryRepositoriesDataInner) SetId(v string) {
	o.Id = v
}

func (o CiProductRelationshipsPrimaryRepositoriesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiProductRelationshipsPrimaryRepositoriesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCiProductRelationshipsPrimaryRepositoriesDataInner struct {
	value *CiProductRelationshipsPrimaryRepositoriesDataInner
	isSet bool
}

func (v NullableCiProductRelationshipsPrimaryRepositoriesDataInner) Get() *CiProductRelationshipsPrimaryRepositoriesDataInner {
	return v.value
}

func (v *NullableCiProductRelationshipsPrimaryRepositoriesDataInner) Set(val *CiProductRelationshipsPrimaryRepositoriesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCiProductRelationshipsPrimaryRepositoriesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCiProductRelationshipsPrimaryRepositoriesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiProductRelationshipsPrimaryRepositoriesDataInner(val *CiProductRelationshipsPrimaryRepositoriesDataInner) *NullableCiProductRelationshipsPrimaryRepositoriesDataInner {
	return &NullableCiProductRelationshipsPrimaryRepositoriesDataInner{value: val, isSet: true}
}

func (v NullableCiProductRelationshipsPrimaryRepositoriesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiProductRelationshipsPrimaryRepositoriesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
